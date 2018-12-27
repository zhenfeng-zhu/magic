package udp

import (
	"context"
	"log"

	"v2ray.com/core/common/buf"
	"v2ray.com/core/common/net"
	"v2ray.com/core/transport/internet"
)

// Payload represents a single UDP payload.
type Payload struct {
	Content             *buf.Buffer
	Source              net.Destination
	OriginalDestination net.Destination
}

type HubOption func(h *Hub)

func HubCapacity(capacity int) HubOption {
	return func(h *Hub) {
		h.capacity = capacity
	}
}

func HubReceiveOriginalDestination(r bool) HubOption {
	return func(h *Hub) {
		h.recvOrigDest = r
	}
}

type Hub struct {
	conn         *net.UDPConn
	cache        chan *Payload
	capacity     int
	recvOrigDest bool
}

func ListenUDP(ctx context.Context, address net.Address, port net.Port, streamSettings *internet.MemoryStreamConfig, options ...HubOption) (*Hub, error) {
	hub := &Hub{
		capacity:     256,
		recvOrigDest: false,
	}
	for _, opt := range options {
		opt(hub)
	}

	var sockopt *internet.SocketConfig
	if streamSettings != nil {
		sockopt = streamSettings.SocketSettings
	}
	if sockopt != nil && sockopt.ReceiveOriginalDestAddress {
		hub.recvOrigDest = true
	}

	udpConn, err := internet.ListenSystemPacket(ctx, &net.UDPAddr{
		IP:   address.IP(),
		Port: int(port),
	}, sockopt)
	if err != nil {
		return nil, err
	}
	log.Println("listening udp on ", address, ":", port)
	newError("listening UDP on ", address, ":", port).WriteToLog()
	hub.conn = udpConn.(*net.UDPConn)
	hub.cache = make(chan *Payload, hub.capacity)

	go hub.start()
	return hub, nil
}

// Close implements net.Listener.
func (h *Hub) Close() error {
	h.conn.Close()
	return nil
}

func (h *Hub) WriteTo(payload []byte, dest net.Destination) (int, error) {
	return h.conn.WriteToUDP(payload, &net.UDPAddr{
		IP:   dest.Address.IP(),
		Port: int(dest.Port),
	})
}

func (h *Hub) start() {
	c := h.cache
	defer close(c)

	oobBytes := make([]byte, 256)

	for {
		buffer := buf.New()
		var noob int
		var addr *net.UDPAddr
		rawBytes := buffer.Extend(buf.Size)

		n, noob, _, addr, err := ReadUDPMsg(h.conn, rawBytes, oobBytes)
		if err != nil {
			newError("failed to read UDP msg").Base(err).WriteToLog()
			buffer.Release()
			break
		}
		buffer.Resize(0, int32(n))

		if buffer.IsEmpty() {
			buffer.Release()
			continue
		}

		payload := &Payload{
			Content: buffer,
			Source:  net.UDPDestination(net.IPAddress(addr.IP), net.Port(addr.Port)),
		}
		if h.recvOrigDest && noob > 0 {
			payload.OriginalDestination = RetrieveOriginalDest(oobBytes[:noob])
			if payload.OriginalDestination.IsValid() {
				newError("UDP original destination: ", payload.OriginalDestination).AtDebug().WriteToLog()
			} else {
				newError("failed to read UDP original destination").WriteToLog()
			}
		}

		select {
		case c <- payload:
		default:
			buffer.Release()
			payload.Content = nil
		}

	}
}

// Addr implements net.Listener.
func (h *Hub) Addr() net.Addr {
	return h.conn.LocalAddr()
}

func (h *Hub) Receive() <-chan *Payload {
	return h.cache
}
