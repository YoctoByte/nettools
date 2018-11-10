package applicationlayer

type DNS struct {
	ID            uint16
	FlagQR        bool
	Opcode        uint8
	FlagAA        bool
	FlagTC        bool
	FlagRD        bool
	FlagRA        bool
	FlagZ         bool
	FlagAD        bool
	FlagCD        bool
	RCode         uint8
	Questions     []DNSQuestion
	Answers       []DNSRecord
	Authoritative []DNSRecord
	Additional    []DNSRecord
}

type DNSQuestion []byte
type DNSRecord []byte

func (message DNS) FromBytes([]byte) (*DNS, error) {
	// TODO
	return nil, nil
}

func (message DNS) ToBytes() ([]byte, error) {
	// TODO
	return nil, nil
}

func (message DNS) String() string {
	// TODO
	return ""
}
