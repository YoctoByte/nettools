package networklayer

import (
	"github.com/yoctobyte/nettools/protocols"
	"github.com/yoctobyte/nettools/protocols/transportlayer"
)

type IPProtocol uint8

var (
	ProtocolICMP  IPProtocol = 0x01
	ProtocolIGMP  IPProtocol = 0x02
	ProtocolTCP   IPProtocol = 0x06
	ProtocolUDP   IPProtocol = 0x11
	ProtocolENCAP IPProtocol = 0x29
	ProtocolOSPF  IPProtocol = 0x59
	ProtocolSCTP  IPProtocol = 0x84
)

var ProtocolString = map[IPProtocol]string{
	ProtocolICMP:  "ICMP",
	ProtocolIGMP:  "IGMP",
	ProtocolTCP:   "TCP",
	ProtocolUDP:   "UDP",
	ProtocolENCAP: "IPv6 encapsulation",
	ProtocolOSPF:  "OSPF",
	ProtocolSCTP:  "SCTP",
}

func ParseProtocolPayload(rawPayload []byte, protocol IPProtocol) (protocols.Payload, error) {
	var payload protocols.Payload
	var err error
	switch protocol {
	case ProtocolICMP: // Internet control message protocol
		payload = &protocols.RawPayload{Bytes: rawPayload}
	case ProtocolIGMP: // Internet group management protocol
		payload = &protocols.RawPayload{Bytes: rawPayload}
	case ProtocolTCP: // Transmission control protocol
		payload, err = transportlayer.TCP{}.FromBytes(rawPayload)
		if err != nil {
			payload = &protocols.RawPayload{Bytes: rawPayload}
		}
	case ProtocolUDP: // User datagram protocol
		payload, err = transportlayer.UDP{}.FromBytes(rawPayload)
		if err != nil {
			payload = &protocols.RawPayload{Bytes: rawPayload}
		}
	case ProtocolENCAP: // IPv6 encapsulation
		payload = &protocols.RawPayload{Bytes: rawPayload}
	case ProtocolOSPF: // Open shortest path first
		payload = &protocols.RawPayload{Bytes: rawPayload}
	case ProtocolSCTP: // Stream control transmission protocol
		payload = &protocols.RawPayload{Bytes: rawPayload}
	default:
		payload = &protocols.RawPayload{Bytes: rawPayload}
	}
	return payload, nil
}
