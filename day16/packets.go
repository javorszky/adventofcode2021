package day16

import (
	"log"
	"sort"
)

type packet interface {
	Version() int
	Type() int
	SubPackets() []packet
	AllVersions() int
	LengthType() lengthType
	Value() int
}

type literal struct {
	packetVersion int
	packetType    int
	value         int
}

func (l *literal) Value() int {
	return l.value
}

func (l *literal) LengthType() lengthType {
	return unknownLengthType
}

func (l *literal) Version() int {
	return l.packetVersion
}

func (l *literal) Type() int {
	return l.packetType
}

func (l *literal) SubPackets() []packet {
	return nil
}

func (l *literal) AllVersions() int {
	return l.Version()
}

type operator struct {
	packetVersion int
	packetType    int
	lengthTypeID  lengthType
	subPackets    []packet
}

func (o *operator) Value() int {
	v := 0

	switch o.Type() {
	case 0:
		// sum packets
		for _, sp := range o.SubPackets() {
			v += sp.Value()
		}
	case 1:
		// product packets
		v = 1
		for _, sp := range o.SubPackets() {
			v *= sp.Value()
		}
	case 2:
		// minimum packets
		vs := make([]int, len(o.SubPackets()))
		for i, sp := range o.SubPackets() {
			vs[i] = sp.Value()
		}

		sort.Ints(vs)
		v = vs[0]
	case 3:
		// maximum packets
		vs := make([]int, len(o.SubPackets()))
		for i, sp := range o.SubPackets() {
			vs[i] = sp.Value()
		}

		sort.Ints(vs)
		v = vs[len(vs)-1]
	case 5:
		// greater than packets (exactly two subpackets)
		// value is 1 if the first subpacket is greater
		sps := o.SubPackets()
		if sps[0].Value() > sps[1].Value() {
			v = 1
		}
	case 6:
		// less than packates (exactly two subpackets)
		// value is 1 if the first subpacket is less
		sps := o.SubPackets()
		if sps[0].Value() < sps[1].Value() {
			v = 1
		}
	case 7:
		// equal to packets (exactly two subpackets)
		// value is 1 if the two subpackets are equal
		sps := o.SubPackets()
		if sps[0].Value() == sps[1].Value() {
			v = 1
		}
	default:
		log.Fatalf("unknown packet type %d", o.Type())
	}

	return v
}

func (o *operator) LengthType() lengthType {
	return o.lengthTypeID
}

func (o *operator) AllVersions() int {
	av := o.Version()
	for _, p := range o.subPackets {
		av += p.AllVersions()
	}

	return av
}

func (o *operator) SubPackets() []packet {
	return o.subPackets
}

func (o *operator) Version() int {
	return o.packetVersion
}

func (o *operator) Type() int {
	return o.packetType
}

func (o *operator) SetSubPackets(packets []packet) {
	o.subPackets = packets
}
