package day20

import (
	"log"
)

type Enhancement struct {
	s string
}

type pixel string

const (
	unknownPx pixel = "?"
	darkPx    pixel = "."
	lightPx   pixel = "#"
)

func NewEnhancement(s string) Enhancement {
	return Enhancement{s: s}
}

func (e Enhancement) subPixel(in []byte) pixel {
	idx := 0

	if len(in) != 9 {
		log.Fatalf("input subpixel is not 9 long")
	}

	for i, sp := range in {
		switch sp {
		case 0x2e:
			// dot, do nothing
		case 0x23:
			// hash
			idx += 1 << ((len(in) - 1) - i)
		default:
			log.Fatalf("yo, this should not have happened, unknown byte: %s", string(sp))
		}
	}

	return parseChar(e.s[idx])
}

func parseChar(c byte) pixel {
	switch c {
	case 0x23:
		return lightPx
	case 0x2e:
		return darkPx
	default:
		log.Fatalf("Enhancement has weird character in it: [%s]", string(c))

		return unknownPx
	}
}
