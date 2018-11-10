package addresses

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

// Mac addresses:
type MacAddress struct {
	bytes []byte
}

func (a MacAddress) FromBytes(bytes []byte) (*MacAddress, error) {
	if len(bytes) != 6 {
		return nil, errors.New("mac address should have a length of 6 bytes")
	}
	a.bytes = bytes
	return &a, nil
}

func (a *MacAddress) ToBytes() ([]byte, error) {
	if len(a.bytes) != 6 {
		return nil, errors.New("mac address should have a length of 6 bytes")
	}
	return a.bytes, nil
}

func (a *MacAddress) String() string {
	var stringBytes []string
	for _, b := range a.bytes {
		stringBytes = append(stringBytes, fmt.Sprintf("%02x", b))
	}
	return strings.Join(stringBytes, ":")
}

// IPv4 addresses:
type IPv4Address struct {
	bytes []byte
}

func (a IPv4Address) FromBytes(bytes []byte) (*IPv4Address, error) {
	if len(bytes) != 4 {
		return nil, errors.New("ipv4 address should have a length of 4 bytes")
	}
	a.bytes = bytes
	return &a, nil
}

func (a *IPv4Address) ToBytes() ([]byte, error) {
	if len(a.bytes) != 4 {
		return nil, errors.New("ipv4 address should have a length of 4 bytes")
	}
	return a.bytes, nil
}

func (a *IPv4Address) String() string {
	var stringParts []string
	for _, b := range a.bytes {
		stringParts = append(stringParts, strconv.Itoa(int(b)))
	}
	return strings.Join(stringParts, ".")
}

// IPv6 addresses:
type IPv6Address struct {
	bytes []byte
}

// TODO: same as for IPv4
