package day16

import "strings"

type lengthType int

const (
	unknownLengthType lengthType = iota
	subPacketLength
	subPacketNumber
)

func task1(input string) int {
	reader := strings.NewReader(input)

	built := newBuilder(reader).build()

	return built.AllVersions()
}
