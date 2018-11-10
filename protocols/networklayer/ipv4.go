package networklayer

import (
	"encoding/binary"
	"errors"
	"fmt"
	"github.com/yoctobyte/nettools/addresses"
	"github.com/yoctobyte/nettools/protocols"
	"math"
)

type Ipv4 struct {
	Version        uint8
	Ihl            uint8
	Dscp           uint8
	Ecn            uint8
	TotalLength    uint16
	Identification uint16
	Flag0          bool
	FlagDF         bool
	FlagMF         bool
	Offset         uint16
	Ttl            uint8
	Protocol       IPProtocol
	HeaderChecksum uint16
	SourceIP       addresses.IPv4Address
	DestinationIP  addresses.IPv4Address
	Options        []uint32
	Payload        protocols.Payload
}

func (packet Ipv4) FromBytes(bytes []byte) (*Ipv4, error) {
	if len(bytes) < 20 {
		return nil, errors.New("ipv4 package not long enough")
	}

	// Parse header
	version := bytes[0] >> 4
	ihl := uint8(bytes[0]) % 16
	if ihl < 5 {
		return nil, errors.New("ihl < 5")
	}
	dscp := bytes[1] >> 2
	ecn := bytes[1] % 4
	totalLength := binary.BigEndian.Uint16(bytes[2:4])
	identification := binary.BigEndian.Uint16(bytes[4:6])
	offset := binary.BigEndian.Uint16(bytes[2:4]) % uint16(math.Pow(2, 13))
	ttl := uint8(bytes[8])
	protocol := IPProtocol(bytes[9])
	checksum := binary.BigEndian.Uint16(bytes[10:12])
	sourceIP, err := addresses.IPv4Address{}.FromBytes(bytes[12:16])
	if err != nil {
		panic(err)
	}
	destinationIP, err := addresses.IPv4Address{}.FromBytes(bytes[16:20])
	if err != nil {
		panic(err)
	}
	if totalLength != uint16(len(bytes)) {
		// TODO: warning
		// panic(fmt.Sprintf("totalLength < len(bytes). %v < %v", totalLength, len(bytes)))
	}
	if totalLength < 4*uint16(ihl) {
		return nil, errors.New("total length field value too low")
	}

	// Parse payload.
	rawPayload := bytes[ihl*4 : totalLength]
	payload, err := ParseProtocolPayload(rawPayload, protocol)
	if err != nil {
		payload = protocols.RawPayload{Bytes: rawPayload}
	}

	packet.Version = version
	packet.Ihl = ihl
	packet.Dscp = dscp
	packet.Ecn = ecn
	packet.TotalLength = totalLength
	packet.Identification = identification
	packet.Offset = offset
	packet.Ttl = ttl
	packet.Protocol = protocol
	packet.HeaderChecksum = checksum
	packet.SourceIP = *sourceIP
	packet.DestinationIP = *destinationIP
	packet.Payload = payload
	return &packet, nil
}

func (packet Ipv4) ToBytes() ([]byte, error) {
	// TODO
	return nil, nil
}

func (packet Ipv4) String() string {
	header := fmt.Sprintf("IPv4\nsource: %v\ndestination: %v\nprotocol: %v (%v)\n",
		packet.SourceIP.String(), packet.DestinationIP.String(), packet.Protocol, ProtocolString[packet.Protocol])
	return header + protocols.Indent(packet.Payload.String())
}
