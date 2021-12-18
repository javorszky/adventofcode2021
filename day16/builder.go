package day16

import (
	"io"
	"log"
	"strconv"
	"strings"
)

type state int

const (
	_ state = iota
	headerWork
	subPacketsLen
	subPacketsCount
	lengthIDWork
	doneWork
)

type builder struct {
	state     state
	charCount int
	reader    io.Reader
}

func (b *builder) build() packet {
	var p packet

loop:
	for {
		switch b.state {
		case headerWork:
			pV, pT, read := b.parseHeader()
			b.charCount += read
			switch pT {
			case 4:
				p = literal{
					packetVersion: pV,
					packetType:    pT,
					value:         b.parseLiteral(),
				}
				b.state = doneWork
			default:
				lengthID, read := b.parseLengthID()
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

func (b *builder) parseHeader() (int, int, int) {
	v := make([]byte, 3)
	read := 0
	// read version.
	vRead, err := b.reader.Read(v)
	if err != nil || vRead != 3 {
		log.Fatalf("reading reader encountered an issue reading version: read %d bytes, got error: %s", vRead, err)
	}

	read += vRead

	pv, err := strconv.ParseInt(string(v), 2, 8)
	if err != nil {
		log.Fatalf("could not parse %s into base10 integer", string(v))
	}

	tRead, err := b.reader.Read(v)
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

func (b *builder) parseLengthID() (lengthType, int) {
	read := 0
	id := make([]byte, 1)

	idRead, err := b.reader.Read(id)
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

		return unknownLengthType, read
	}
}

func (b *builder) parseLiteral() int {
	var value strings.Builder

	goOn := true
	read := 0
	check := make([]byte, 1)
	val := make([]byte, 4)

	for {
		cR, err := b.reader.Read(check)
		if err != nil {
			log.Fatalf("could not read check bit: %s", err)
		}

		read += cR

		if string(check) == "0" {
			goOn = false
		}

		vR, err := b.reader.Read(val)
		if err != nil {
			log.Fatalf("builder.parseLiteral: could not read val bit: %s", err)
		}

		read += vR

		value.WriteString(string(val))

		if !goOn {
			break
		}
	}

	b.charCount += read
	valueString := value.String()

	parsedInt, err := strconv.ParseInt(valueString, 2, len(valueString)+1)
	if err != nil {
		log.Fatalf("could not parse binary string [%s] into int: %s", valueString, err)
	}

	return int(parsedInt)
}

func newBuilder(reader io.Reader) *builder {
	return &builder{
		state:     headerWork,
		reader:    reader,
		charCount: 0,
	}
}
