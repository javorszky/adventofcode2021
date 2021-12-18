package day16

import (
	"io"
	"log"
	"strconv"
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
	charCount int
}

func (b *builder) build(reader io.Reader) packet {
	//var p packet
	for {
		switch b.state {
		case headerWork:
			//pV, pT, err := b.parseHeader(reader)
		case literalWork:
		case subPacketsLen:
		case subPacketsCount:
		case lengthIDWork:
		default:
			log.Fatalf("builder encountered an unknown state: %v", b.state)
		}
		break
	}

	return literal{
		packetVersion: 1,
		packetType:    2,
		value:         3,
	}
}

func (b *builder) parseHeader(reader io.Reader) (int, int, error) {
	v := make([]byte, 3)

	// read version.
	read, err := reader.Read(v)
	if err != nil || read != 3 {
		return 0, 0, err
	}

	pv, err := strconv.ParseInt(string(v), 2, 8)
	if err != nil {
		log.Fatalf("could not parse %s into base10 integer", string(v))
	}

	read, err = reader.Read(v)
	if err != nil || read != 3 {
		return 0, 0, err
	}

	pt, err := strconv.ParseInt(string(v), 2, 8)
	if err != nil {
		log.Fatalf("could not parse %s into base10 integer", string(v))
	}

	return int(pv), int(pt), nil
}

func newBuilder() *builder {
	return &builder{
		state:     headerWork,
		charCount: 0,
	}
}
