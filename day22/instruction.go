package day22

import (
	"errors"
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

func (i instruction) Lights() int {
	return (i.xTo - i.xFrom + 1) * (i.yTo - i.yFrom + 1) * (i.zTo - i.zFrom + 1)
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

func overlap(box, otherBox instruction) []instruction {
	// they do not overlap, because box ends before otherBox begins.
	if box.xFrom > otherBox.xTo || box.yFrom > otherBox.yTo || box.zFrom > box.zTo {
		return []instruction{box, otherBox}
	}

	// they do not overlap, because box doesn't start until after otherBox ends.
	if box.xTo < otherBox.xFrom || box.yTo < otherBox.yFrom || box.zTo < otherBox.zFrom {
		return []instruction{box, otherBox}
	}

	instructions := make([]instruction, 0)

	overlapBox, err := findOverlapBox(box, otherBox)
	if err != nil {
		log.Fatalf("despite checking for overlaps, we couldn't find the box. This should not have happened\n"+
			"box:      %v\n"+
			"otherbox: %v", box, otherBox)
	}

	instructions = append(instructions, overlapBox)

	for _, f := range []func(instruction, instruction) []instruction{
		findTopFace,
		findBottomFace,
		findLeftFace,
		findRightFace,
		findFrontFace,
		findBackFace,
		findTopLeftEdge,
		findTopBackEdge,
		findTopRightEdge,
		findTopFrontEdge,
		findBottomLeftEdge,
		findBottomBackEdge,
		findBottomRightEdge,
		findBottomFrontEdge,
		findFrontLeftEdge,
		findFrontRightEdge,
		findBackLeftEdge,
		findBackRightEdge,
		findTopBackLeftCorner,
		findTopBackRightCorner,
		findTopFrontLeftCorner,
		findTopFrontRightCorner,
		findBottomBackLeftCorner,
		findBottomBackRightCorner,
		findBottomFrontLeftCorner,
		findBottomFrontRightCorner,
	} {
		instructions = append(instructions, f(box, overlapBox)...)
	}

	/*
		xxx  xxx  xxx
		xxx  xox  xxx
		xxx  xxx  xxx

		ends up being

		aaa  bbb  ccc
		aaa  doe  ccc
		aaa  fff  ccc
	*/
	return instructions
}

func findOverlapBox(box, otherBox instruction) (instruction, error) {
	// Generate overlap box.
	xmin := box.xFrom
	if otherBox.xFrom > xmin {
		xmin = otherBox.xFrom
	}

	xmax := box.xTo
	if otherBox.xTo < xmax {
		xmax = otherBox.xTo
	}

	ymin := box.yFrom
	if otherBox.yFrom > ymin {
		ymin = otherBox.yFrom
	}

	ymax := box.yTo
	if otherBox.yTo < ymax {
		ymax = otherBox.yTo
	}

	zmin := box.zFrom
	if otherBox.zFrom < zmin {
		zmin = otherBox.zFrom
	}

	zmax := box.zTo
	if otherBox.zTo < zmax {
		zmax = otherBox.zTo
	}

	if xmin >= xmax || ymin >= ymax || zmin >= zmax {
		return instruction{}, errors.New("out of bounds")
	}

	return instruction{
		xFrom: xmin,
		xTo:   xmax,
		yFrom: ymin,
		yTo:   ymax,
		zFrom: zmin,
		zTo:   zmax,
		flip:  otherBox.flip,
	}, nil
}

// find faces, 6 of them.
func findTopFace(box, overlapBox instruction) []instruction {
	// there is no face here, the overlap box is at the edge
	if box.zTo == overlapBox.zTo {
		return nil
	}

	return []instruction{
		{
			xFrom: overlapBox.xFrom,
			xTo:   overlapBox.xTo,
			yFrom: overlapBox.yFrom,
			yTo:   overlapBox.yTo,
			zFrom: overlapBox.zTo,
			zTo:   box.zTo,
			flip:  box.flip,
		},
	}
}

func findBottomFace(box, overlapBox instruction) []instruction {
	// there is no face here, the overlap box is at the edge
	if box.zFrom == overlapBox.zFrom {
		return nil
	}

	return []instruction{
		{
			xFrom: overlapBox.xFrom,
			xTo:   overlapBox.xTo,
			yFrom: overlapBox.yFrom,
			yTo:   overlapBox.yTo,
			zFrom: box.zFrom,
			zTo:   overlapBox.zFrom,
			flip:  box.flip,
		},
	}
}

func findFrontFace(box, overlapBox instruction) []instruction {
	// there is no face here, the overlap box is at the front edge
	if box.xTo == overlapBox.xTo {
		return nil
	}

	return []instruction{
		{
			xFrom: overlapBox.xTo,
			xTo:   box.xTo,
			yFrom: overlapBox.yFrom,
			yTo:   overlapBox.yTo,
			zFrom: overlapBox.zFrom,
			zTo:   overlapBox.zTo,
			flip:  box.flip,
		},
	}
}

func findBackFace(box, overlapBox instruction) []instruction {
	// there is no face here, the overlap box is at the back edge
	if box.xFrom == overlapBox.xFrom {
		return nil
	}

	return []instruction{
		{
			xFrom: box.xFrom,
			xTo:   overlapBox.xFrom,
			yFrom: overlapBox.yFrom,
			yTo:   overlapBox.yTo,
			zFrom: overlapBox.zFrom,
			zTo:   overlapBox.zTo,
			flip:  box.flip,
		},
	}
}

func findLeftFace(box, overlapBox instruction) []instruction {
	// there is no face here, the overlap box is at the left edge
	if box.yFrom == overlapBox.yFrom {
		return nil
	}

	return []instruction{
		{
			xFrom: overlapBox.xFrom,
			xTo:   overlapBox.xTo,
			yFrom: box.yFrom,
			yTo:   overlapBox.yFrom,
			zFrom: overlapBox.zFrom,
			zTo:   overlapBox.zTo,
			flip:  box.flip,
		},
	}
}

func findRightFace(box, overlapBox instruction) []instruction {
	// there is no face here, the overlap box is at the right edge
	if box.yTo == overlapBox.yTo {
		return nil
	}

	return []instruction{
		{
			xFrom: overlapBox.xFrom,
			xTo:   overlapBox.xTo,
			yFrom: overlapBox.yTo,
			yTo:   box.yTo,
			zFrom: overlapBox.zFrom,
			zTo:   overlapBox.zTo,
			flip:  box.flip,
		},
	}
}

// find edges, 12 of them.
func findTopLeftEdge(box, overlapBox instruction) []instruction {
	return nil
}

func findTopBackEdge(box, overlapBox instruction) []instruction {
	return nil
}

func findTopRightEdge(box, overlapBox instruction) []instruction {
	return nil
}

func findTopFrontEdge(box, overlapBox instruction) []instruction {
	return nil
}

func findBottomLeftEdge(box, overlapBox instruction) []instruction {
	return nil
}

func findBottomBackEdge(box, overlapBox instruction) []instruction {
	return nil
}

func findBottomRightEdge(box, overlapBox instruction) []instruction {
	return nil
}

func findBottomFrontEdge(box, overlapBox instruction) []instruction {
	return nil
}

func findFrontLeftEdge(box, overlapBox instruction) []instruction {
	return nil
}

func findFrontRightEdge(box, overlapBox instruction) []instruction {
	return nil
}

func findBackLeftEdge(box, overlapBox instruction) []instruction {
	return nil
}

func findBackRightEdge(box, overlapBox instruction) []instruction {
	return nil
}

// find corners, 8 of them.
func findTopBackLeftCorner(box, overlapBox instruction) []instruction {
	return nil
}

func findTopBackRightCorner(box, overlapBox instruction) []instruction {
	return nil
}

func findTopFrontLeftCorner(box, overlapBox instruction) []instruction {
	return nil
}

func findTopFrontRightCorner(box, overlapBox instruction) []instruction {
	return nil
}

func findBottomBackLeftCorner(box, overlapBox instruction) []instruction {
	return nil
}

func findBottomBackRightCorner(box, overlapBox instruction) []instruction {
	return nil
}

func findBottomFrontLeftCorner(box, overlapBox instruction) []instruction {
	return nil
}

func findBottomFrontRightCorner(box, overlapBox instruction) []instruction {
	return nil
}
