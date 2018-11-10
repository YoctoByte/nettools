package rawsocket

import (
	"github.com/yoctobyte/nettools/protocols/linklayer"
	"math"
	"syscall"
)

type RawSocket struct {
	sock          int
	addr          syscall.Sockaddr
	receiveBuffer chan *linklayer.Ethernet
	sendBuffer    chan *linklayer.Ethernet
}

func (s RawSocket) New(receiveBuffer chan *linklayer.Ethernet, sendBuffer chan *linklayer.Ethernet) *RawSocket {
	sock, err := syscall.Socket(syscall.AF_PACKET, syscall.SOCK_RAW, int(htons(3)))
	if err != nil {
		panic(err)
	}
	return &RawSocket{
		sock:          sock,
		receiveBuffer: receiveBuffer,
		sendBuffer:    sendBuffer,
	}
}

// htons converts a short (uint16) from host-to-network byte order.
func htons(i uint16) uint16 {
	return (i<<8)&0xff00 | i>>8
}

func (s *RawSocket) Start() {
	// Start listening
	go func(receiveBuffer chan *linklayer.Ethernet) {
		recvBuf := make([]byte, int(math.Pow(2, 16)))
		for {
			read, _, err := syscall.Recvfrom(s.sock, recvBuf, 0)
			if err != nil {
				continue
			}
			ethernet, err := linklayer.Ethernet{}.FromBytes(recvBuf[:read])
			if err != nil {
				continue
			}
			receiveBuffer <- ethernet
		}
	}(s.receiveBuffer)

	// Start sending
	go func(sendBuffer chan *linklayer.Ethernet) {
		for {
			frame := <-s.sendBuffer
			bytes, err := frame.ToBytes()
			if err != nil {
				// TODO: warn
				continue
			}
			err = syscall.Sendto(s.sock, bytes, 0, nil)
			if err != nil {
				// TODO: warn
				continue
			}
		}
	}(s.sendBuffer)
}

func (s RawSocket) SendFrame(frame *linklayer.Ethernet) {
	s.sendBuffer <- frame
}

func (s RawSocket) ReceiveFrame() *linklayer.Ethernet {
	frame := <-s.receiveBuffer
	return frame
}
