package main

import (
	"fmt"
	"github.com/yoctobyte/nettools/protocols/linklayer"
	"github.com/yoctobyte/nettools/protocols/networklayer"
	"github.com/yoctobyte/nettools/protocols/transportlayer"
	"github.com/yoctobyte/nettools/rawsocket"
)

func main() {
	rawSock := rawsocket.RawSocket{}.New(make(chan *linklayer.Ethernet, 10), make(chan *linklayer.Ethernet, 10))
	rawSock.Start()
	for {
		frame := rawSock.ReceiveFrame()
		if frame.EtherType == linklayer.EtherTypeIPv4 {
			packet, ok := frame.Payload.(*networklayer.Ipv4)
			if !ok {
				continue
			}
			if packet.Protocol == networklayer.ProtocolUDP {
				datagram, ok := packet.Payload.(*transportlayer.UDP)
				if !ok {
					fmt.Println("not ok")
					continue
				}
				if datagram.DestinationPort == 67 {
					fmt.Println(frame)
				}
			}
		}
	}
}
