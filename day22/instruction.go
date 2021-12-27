package day22

import (
	"log"
	"regexp"
	"strconv"
)

type flip int

const (
	unknownInstruction flip = iota
	on
	off
)

type instruction struct {
	xfrom, xto, yfrom, yto, zfrom, zto int
	flip                               flip
}

var reInstruction = regexp.MustCompile(`^(on|off) x=(-?\d+)\.\.(-?\d+),y=(-?\d+)\.\.(-?\d+),z=(-?\d+)\.\.(-?\d+)$`)

func parseInstruction(s string) instruction {
	parts := reInstruction.FindStringSubmatch(s)
	numbers := make([]int, 6)

	for i, p := range parts[2:] {
		n, err := strconv.Atoi(p)
		if err != nil {
			log.Fatalf("failed to convert string [%s] to int while parsing [%s]: %s", p, s, err)
		}

		numbers[i] = n
	}

	var ins flip

	switch parts[1] {
	case "on":
		ins = on
	case "off":
		ins = off
	default:
		ins = unknownInstruction

		log.Fatalf("unknown ins state: %s", parts[1])
	}

	return instruction{
		xfrom: numbers[0],
		xto:   numbers[1],
		yfrom: numbers[2],
		yto:   numbers[3],
		zfrom: numbers[4],
		zto:   numbers[5],
		flip:  ins,
	}
}
