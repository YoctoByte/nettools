package linklayer

type EtherType uint16

var (
	EtherTypeIPv4 EtherType = 0x0800
	EtherTypeARP  EtherType = 0x0806
	EtherTypeWOL  EtherType = 0x0842
	EtherTypeIPv6 EtherType = 0x86dd
)

var EtherTypeString = map[EtherType]string{
	EtherTypeIPv4: "Ipv4",
	EtherTypeARP:  "ARP",
	EtherTypeWOL:  "Wake-On-Lan",
	EtherTypeIPv6: "IPv6",
}
