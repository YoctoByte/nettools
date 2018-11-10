package linklayer

import (
	"encoding/binary"
	"errors"
	"fmt"
	"github.com/yoctobyte/nettools/addresses"
	"github.com/yoctobyte/nettools/protocols"
	"github.com/yoctobyte/nettools/protocols/networklayer"
)

type Ethernet struct {
	DestinationAddress *addresses.MacAddress
	SourceAddress      *addresses.MacAddress
	EtherType          EtherType
	QTag               uint32
	Payload            protocols.Payload
}

func (frame Ethernet) FromBytes(bytes []byte) (*Ethernet, error) {
	if len(bytes) < 18 || (bytes[12] == 0x80 && bytes[13] == 0x00 && len(bytes) < 22) {
		return nil, errors.New("ethernet fragment not long enough")
	}

	// Parse header:
	frame.DestinationAddress, _ = addresses.MacAddress{}.FromBytes(bytes[:6])
	frame.SourceAddress, _ = addresses.MacAddress{}.FromBytes(bytes[6:12])
	var rawPayload []byte
	if bytes[12] == 0x80 && bytes[13] == 0x00 { // QTag is set
		frame.QTag = binary.BigEndian.Uint32(bytes[12:16])
		frame.EtherType = EtherType(binary.BigEndian.Uint16(bytes[16:18]))
		rawPayload = bytes[18:]
	} else {
		frame.EtherType = EtherType(binary.BigEndian.Uint16(bytes[12:14]))
		rawPayload = bytes[14:]
	}

	// Parse Payload:
	var err error
	switch frame.EtherType {
	case EtherTypeIPv4:
		frame.Payload, err = networklayer.Ipv4{}.FromBytes(rawPayload)
		if err != nil {
			frame.Payload = &protocols.RawPayload{Bytes: rawPayload}
		}
	case EtherTypeARP:
		frame.Payload = &protocols.RawPayload{Bytes: rawPayload}
	case EtherTypeWOL:
		frame.Payload = &protocols.RawPayload{Bytes: rawPayload}
	case EtherTypeIPv6:
		frame.Payload = &protocols.RawPayload{Bytes: rawPayload}
	default:
		frame.Payload = &protocols.RawPayload{Bytes: rawPayload}
	}
	return &frame, nil
}

func (frame *Ethernet) ToBytes() ([]byte, error) {
	// Set destination mac address:
	var bytes []byte
	if frame.DestinationAddress == nil {
		return nil, errors.New("destination address not set")
	}
	destinationBytes, err := frame.DestinationAddress.ToBytes()
	if err != nil {
		return nil, errors.New("could not convert destination address to bytes: " + err.Error())
	}
	bytes = append(bytes, destinationBytes...)

	// Set source mac address:
	if frame.SourceAddress == nil {
		return nil, errors.New("source address not set")
	}
	sourceBytes, err := frame.SourceAddress.ToBytes()
	if err != nil {
		return nil, errors.New(" could not convert source address to bytes: " + err.Error())
	}
	bytes = append(bytes, sourceBytes...)

	// TODO

	return bytes, nil
}

func (frame *Ethernet) String() string {
	var header string
	if frame.QTag != 0 {
		header = fmt.Sprintf("ETHERNET\ndestination: %v\nsource: %v\nqtag: %v\nethertype: %v (%v)\n",
			frame.DestinationAddress, frame.SourceAddress, frame.QTag, frame.EtherType, EtherTypeString[frame.EtherType])
	} else {
		header = fmt.Sprintf("ETHERNET\ndestination: %v\nsource: %v\nethertype: %v (%v)\n",
			frame.DestinationAddress, frame.SourceAddress, frame.EtherType, EtherTypeString[frame.EtherType])
	}
	return header + protocols.Indent(frame.Payload.String())
}
