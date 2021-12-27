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
	xFrom, xTo, yFrom, yTo, zFrom, zTo int
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
		xFrom: numbers[0],
		xTo:   numbers[1],
		yFrom: numbers[2],
		yTo:   numbers[3],
		zFrom: numbers[4],
		zTo:   numbers[5],
		flip:  ins,
	}
}
