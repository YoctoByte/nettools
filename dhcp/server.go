package dhcp

import (
	"github.com/yoctobyte/nettools/protocols/linklayer"
	"github.com/yoctobyte/nettools/protocols/networklayer"
	"github.com/yoctobyte/nettools/protocols/transportlayer"
	"github.com/yoctobyte/nettools/rawsocket"
)

type Server struct {
	socket *rawsocket.RawSocket
	config Config
}

type Config struct {
}

func (s Server) New(config Config) Server {
	s.socket = rawsocket.RawSocket{}.New(make(chan *linklayer.Ethernet, 10), make(chan *linklayer.Ethernet, 10))
	s.config = config
	return s
}

func (s *Server) Listen(dhcpChan chan *linklayer.Ethernet) {
	s.socket.Start()
	for {
		frame := s.socket.ReceiveFrame()
		dhcpChan <- frame
		if frame.EtherType == linklayer.EtherTypeIPv4 {
			packet, ok := frame.Payload.(networklayer.Ipv4)
			if !ok {
				continue
			}
			if packet.Protocol == networklayer.ProtocolUDP {
				datagram, ok := packet.Payload.(transportlayer.UDP)
				if !ok {
					continue
				}
				if datagram.DestinationPort == 67 {
					dhcpChan <- frame
				}
			}
		}
	}
}
