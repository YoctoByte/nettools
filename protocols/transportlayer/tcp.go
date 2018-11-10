package transportlayer

import (
	"encoding/binary"
	"errors"
	"fmt"
	"github.com/yoctobyte/nettools/protocols"
)

type TCP struct {
	SourcePort      uint16
	DestinationPort uint16
	SequenceNumber  uint32
	AckNumber       uint32
	DataOffset      uint8
	FlagNS          bool
	FlagCWR         bool
	FlagECE         bool
	FlagURG         bool
	FlagACK         bool
	FlagPSH         bool
	FlagRST         bool
	FlagSYN         bool
	FlagFin         bool
	WindowSize      uint16
	Checksum        uint16
	UrgentPointer   uint16
	Payload         protocols.Payload
}

func (segment TCP) FromBytes(bytes []byte) (*TCP, error) {
	if len(bytes) < 20 {
		return nil, errors.New("not enough bytes for tcp segment")
	}

	// Parse header
	segment.SourcePort = binary.BigEndian.Uint16(bytes[:2])
	segment.DestinationPort = binary.BigEndian.Uint16(bytes[2:4])
	segment.SequenceNumber = binary.BigEndian.Uint32(bytes[4:8])
	segment.AckNumber = binary.BigEndian.Uint32(bytes[8:12])
	segment.DataOffset = uint8(bytes[12]) >> 4
	segment.FlagNS = uint8(bytes[12])%2 == 1
	segment.FlagCWR = uint8(bytes[13])>>7%2 == 1
	segment.FlagECE = uint8(bytes[13])>>6%2 == 1
	segment.FlagURG = uint8(bytes[13])>>5%2 == 1
	segment.FlagACK = uint8(bytes[13])>>4%2 == 1
	segment.FlagPSH = uint8(bytes[13])>>3%2 == 1
	segment.FlagRST = uint8(bytes[13])>>2%2 == 1
	segment.FlagSYN = uint8(bytes[13])>>1%2 == 1
	segment.FlagFin = uint8(bytes[13])>>0%2 == 1
	segment.WindowSize = binary.BigEndian.Uint16(bytes[14:16])
	segment.Checksum = binary.BigEndian.Uint16(bytes[16:18])
	segment.UrgentPointer = binary.BigEndian.Uint16(bytes[18:20])
	if segment.DataOffset < 5 {
		return nil, errors.New("data offset < 5")
	}

	// Parse payload
	rawPayload := bytes[segment.DataOffset*4:]
	segment.Payload = &protocols.RawPayload{Bytes: rawPayload}

	return &segment, nil
}

func (segment *TCP) ToBytes() ([]byte, error) {
	// TODO
	return nil, nil
}

func (segment *TCP) String() string {
	header := fmt.Sprintf("TCP\nsource port: %v\ndestination port: %v\n", segment.SourcePort, segment.DestinationPort)
	return header + protocols.Indent(segment.Payload.String())
}
