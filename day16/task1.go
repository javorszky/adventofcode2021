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

	return reader
}

type packet interface {
	Version() int
	Type() int
}

type literal struct {
	packetVersion int
	packetType    int
	value         int
}

func (l literal) Version() int {
	return l.packetVersion
}

func (l literal) Type() int {
	return l.packetType
}

type operator struct {
	packetVersion int
	packetType    int
	lengthTypeID  lengthType
	subPackets    []packet
}

func (o operator) Version() int {
	return o.packetVersion
}

func (o operator) Type() int {
	return o.packetType
}
