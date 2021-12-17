package day16

import (
	"io"
)

type state int

const (
	blockSize = 4

	unknown state = iota
	headerWork
	literalWork
	subPacketsLen
	subPacketsCount
	lengthIDWork
)

type builder struct {
	state     state
	next      state
	charCount int
}

func build(reader io.Reader) packet {
	for {

		break
	}

	return literal{
		packetVersion: 1,
		packetType:    2,
		value:         3,
	}
}
