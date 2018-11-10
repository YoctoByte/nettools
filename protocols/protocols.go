package protocols

import (
	"fmt"
	"strings"
)

func Indent(str string) string {
	indent := "    "
	var parts []string
	for _, part := range strings.Split(str, "\n") {
		parts = append(parts, indent+part)
	}
	return strings.Join(parts, "\n")
}

type Payload interface {
	ToBytes() ([]byte, error)
	String() string
}

type RawPayload struct {
	Bytes []byte
}

func (p RawPayload) ToBytes() ([]byte, error) {
	return p.Bytes, nil
}

func (p RawPayload) String() string {
	return fmt.Sprintf("%q\n", p.Bytes)
	//var stringBytes []string
	//for _, b := range p.Bytes {
	//	stringBytes = append(stringBytes, fmt.Sprintf("%02x", b))
	//}
	//return strings.Join(stringBytes, " ")
}
