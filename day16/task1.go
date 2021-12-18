package day16

import "strings"

type lengthType int

const (
	unknownLengthType lengthType = iota
	subPacketLength
	subPacketNumber
)

func task1(input string) interface{} {
	reader := strings.NewReader(input)

	built := newBuilder(reader).build()

	return built.AllVersions()
}

type packet interface {
	Version() int
	Type() int
	SubPackets() []packet
	AllVersions() int
	LengthType() lengthType
}

type literal struct {
	packetVersion int
	packetType    int
	value         int
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
