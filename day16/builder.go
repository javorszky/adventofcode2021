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
	doneWork
)

type builder struct {
	state     state
	charCount int
}

func (b *builder) build(reader io.Reader) packet {
	var p packet

loop:
	for {
		switch b.state {
		case headerWork:
			pV, pT, read := b.parseHeader(reader)
			b.charCount += read
			switch pT {
			case 4:
				b.state = literalWork
				p = literal{
					packetVersion: pV,
					packetType:    pT,
					value:         0,
				}
			default:
				lengthID, read := b.parseLengthID(reader)
				b.charCount += read
				p = operator{
					packetVersion: pV,
					packetType:    pT,
					lengthTypeID:  lengthID,
					subPackets:    nil,
				}

				switch lengthID {
				case subPacketLength:
					b.state = subPacketsLen
				default:
					b.state = subPacketsCount
				}
			}
		case literalWork:
		case subPacketsLen:
		case subPacketsCount:
		case lengthIDWork:

		case doneWork:
			break loop
		default:
			log.Fatalf("builder encountered an unknown state: %v", b.state)
		}
	}

	return p
}

func (b *builder) parseHeader(reader io.Reader) (int, int, int) {
	v := make([]byte, 3)
	read := 0
	// read version.
	vRead, err := reader.Read(v)
	if err != nil || vRead != 3 {
		log.Fatalf("reading reader encountered an issue reading version: read %d bytes, got error: %s", vRead, err)
	}

	read += vRead

	pv, err := strconv.ParseInt(string(v), 2, 8)
	if err != nil {
		log.Fatalf("could not parse %s into base10 integer", string(v))
	}

	tRead, err := reader.Read(v)
	if err != nil || tRead != 3 {
		log.Fatalf("reading reader encountered an issue reading type: read %d bytes, got error: %s", tRead, err)
	}

	read += tRead

	pt, err := strconv.ParseInt(string(v), 2, 8)
	if err != nil {
		log.Fatalf("could not parse %s into base10 integer", string(v))
	}

	return int(pv), int(pt), read
}

func (b *builder) parseLengthID(reader io.Reader) (lengthType, int) {
	read := 0
	id := make([]byte, 1)

	idRead, err := reader.Read(id)
	if err != nil || idRead != 1 {
		log.Fatalf("reading reader encountered an issue reading length type id: read %d bytes, got error: %s", idRead, err)
	}

	read += idRead

	pid, err := strconv.ParseInt(string(id), 2, 1)
	if err != nil {
		log.Fatalf("could not parse %s into base10 integer", string(id))
	}

	switch pid {
	case 0:
		return subPacketLength, read
	case 1:
		return subPacketNumber, read
	default:
		log.Fatalf("unknown length type id, can't continue")
	}
}

func newBuilder() *builder {
	return &builder{
		state:     headerWork,
		charCount: 0,
	}
}
