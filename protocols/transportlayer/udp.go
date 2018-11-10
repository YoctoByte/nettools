package transportlayer

import (
	"encoding/binary"
	"errors"
	"fmt"
	"github.com/yoctobyte/nettools/protocols"
)

type UDP struct {
	SourcePort      uint16
	DestinationPort uint16
	Length          uint16
	Checksum        uint16
	Payload         protocols.Payload
}

func (datagram UDP) FromBytes(bytes []byte) (*UDP, error) {
	if len(bytes) < 8 {
		return nil, errors.New("not enough bytes for udp datagram")
	}
	datagram.SourcePort = binary.BigEndian.Uint16(bytes[:2])
	datagram.DestinationPort = binary.BigEndian.Uint16(bytes[2:4])
	datagram.Length = binary.BigEndian.Uint16(bytes[4:6])
	datagram.Checksum = binary.BigEndian.Uint16(bytes[6:8])
	if datagram.Length < 8 {
		return nil, errors.New("length < 8")
	}
	if datagram.Length != uint16(len(bytes)) {
		panic(fmt.Sprintf("length != len(bytes). %v, %v", datagram.Length, len(bytes)))
	}
	rawPayload := bytes[8:datagram.Length]
	datagram.Payload = protocols.RawPayload{Bytes: rawPayload}
	return &datagram, nil
}

func (datagram UDP) ToBytes() ([]byte, error) {
	// TODO
	return nil, nil
}

func (datagram UDP) String() string {
	header := fmt.Sprintf("UDP\nsource port: %v\ndestination port: %v\n", datagram.SourcePort, datagram.DestinationPort)
	return header + protocols.Indent(datagram.Payload.String())
}
