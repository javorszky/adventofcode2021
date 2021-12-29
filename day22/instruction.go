package day22

import (
	"errors"
	"fmt"
	"log"
	"regexp"
	"strconv"
)

const (
	unknownInstruction flip = iota
	on
	off
)

var reInstruction = regexp.MustCompile(`^(on|off) x=(-?\d+)\.\.(-?\d+),y=(-?\d+)\.\.(-?\d+),z=(-?\d+)\.\.(-?\d+)$`)

type flip int

func (f flip) String() string {
	switch f {
	case on:
		return "on"
	case off:
		return "off"
	default:
		return "unknown"
	}
}

type instruction struct {
	xFrom, xTo, yFrom, yTo, zFrom, zTo int
	flip                               flip
}

func (i instruction) Lights() int {
	return (i.xTo - i.xFrom + 1) * (i.yTo - i.yFrom + 1) * (i.zTo - i.zFrom + 1)
}

func (i instruction) String() string {
	return fmt.Sprintf("%d/%d/%d/%d/%d/%d/%s",
		i.xFrom, i.xTo,
		i.yFrom, i.yTo,
		i.zFrom, i.zTo,
		i.flip,
	)
}

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

func overlapAndMerge(box, otherBox instruction) map[string]instruction {
	overlaps := overlap(box, otherBox)

	overlaps = mergeInstructionSlice(overlaps)

	newOverlapsMap := make(map[string]instruction)
	for _, v := range overlaps {
		newOverlapsMap[v.String()] = v
	}

	return newOverlapsMap
}

func overlap(box, otherBox instruction) []instruction {
	// they do not overlap, because box ends before otherBox begins.
	if box.xFrom >= otherBox.xTo || box.yFrom >= otherBox.yTo || box.zFrom >= otherBox.zTo {
		return []instruction{box, otherBox}
	}

	// they do not overlap, because box doesn't start until after otherBox ends.
	if otherBox.xFrom >= box.xTo || otherBox.yFrom >= box.yTo || otherBox.zFrom >= box.zTo {
		return []instruction{box, otherBox}
	}

	instructions := make([]instruction, 0)

	overlapBox, err := findOverlapBox(box, otherBox)
	if err != nil {
		log.Fatalf("despite checking for overlaps, we couldn't find the box. This should not have happened\n"+
			"box:      %v\n"+
			"otherBox: %v", box, otherBox)
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
		instructions = append(instructions, f(otherBox, overlapBox)...)
	}

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
	if otherBox.zFrom > zmin {
		zmin = otherBox.zFrom
	}

	zmax := box.zTo
	if otherBox.zTo < zmax {
		zmax = otherBox.zTo
	}

	if xmin > xmax || ymin > ymax || zmin > zmax {
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
			zFrom: overlapBox.zTo + 1,
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
			zTo:   overlapBox.zFrom - 1,
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
			xFrom: overlapBox.xTo + 1 + 1,
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
			xTo:   overlapBox.xFrom - 1 - 1,
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
			yTo:   overlapBox.yFrom - 1,
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
			yFrom: overlapBox.yTo + 1,
			yTo:   box.yTo,
			zFrom: overlapBox.zFrom,
			zTo:   overlapBox.zTo,
			flip:  box.flip,
		},
	}
}

// find edges, 12 of them.
func findTopLeftEdge(box, overlapBox instruction) []instruction {
	// there is no edge here, the overlap box is at either the top or the left edge.
	if box.zTo == overlapBox.zTo || box.yFrom == overlapBox.yFrom {
		return nil
	}

	return []instruction{
		{
			xFrom: overlapBox.xFrom,
			xTo:   overlapBox.xTo,
			yFrom: box.yFrom,
			yTo:   overlapBox.yFrom - 1,
			zFrom: overlapBox.zTo + 1,
			zTo:   box.zTo,
			flip:  box.flip,
		},
	}
}

func findTopBackEdge(box, overlapBox instruction) []instruction {
	// there is no edge here, the overlap box is at either the top or the back edge.
	if box.zTo == overlapBox.zTo || box.xFrom == overlapBox.xFrom {
		return nil
	}

	return []instruction{
		{
			xFrom: box.xFrom,
			xTo:   overlapBox.xFrom - 1,
			yFrom: overlapBox.yFrom,
			yTo:   overlapBox.yTo,
			zFrom: overlapBox.zTo + 1,
			zTo:   box.zTo,
			flip:  box.flip,
		},
	}
}

func findTopRightEdge(box, overlapBox instruction) []instruction {
	// there is no edge here, the overlap box is at either the top or the left edge.
	if box.zTo == overlapBox.zTo || box.yTo == overlapBox.yTo {
		return nil
	}

	return []instruction{
		{
			xFrom: overlapBox.xFrom,
			xTo:   overlapBox.xTo,
			yFrom: overlapBox.yTo + 1,
			yTo:   box.yTo,
			zFrom: overlapBox.zTo + 1,
			zTo:   box.zTo,
			flip:  box.flip,
		},
	}
}

func findTopFrontEdge(box, overlapBox instruction) []instruction {
	// there is no edge here, the overlap box is at either the top or the back edge.
	if box.zTo == overlapBox.zTo || box.xTo == overlapBox.xTo {
		return nil
	}

	return []instruction{
		{
			xFrom: overlapBox.xTo + 1,
			xTo:   box.xTo,
			yFrom: overlapBox.yFrom,
			yTo:   overlapBox.yTo,
			zFrom: overlapBox.zTo + 1,
			zTo:   box.zTo,
			flip:  box.flip,
		},
	}
}

func findBottomLeftEdge(box, overlapBox instruction) []instruction {
	// there is no edge here, the overlap box is at either the bottom or the left edge.
	if box.zFrom == overlapBox.zFrom || box.yFrom == overlapBox.yFrom {
		return nil
	}

	return []instruction{
		{
			xFrom: overlapBox.xFrom,
			xTo:   overlapBox.xTo,
			yFrom: box.yFrom,
			yTo:   overlapBox.yFrom - 1,
			zFrom: box.zFrom,
			zTo:   overlapBox.zFrom - 1,
			flip:  box.flip,
		},
	}
}

func findBottomBackEdge(box, overlapBox instruction) []instruction {
	// there is no edge here, the overlap box is at either the bottom or the back edge.
	if box.zFrom == overlapBox.zFrom || box.xFrom == overlapBox.xFrom {
		return nil
	}

	return []instruction{
		{
			xFrom: box.xFrom,
			xTo:   overlapBox.xFrom - 1,
			yFrom: overlapBox.yFrom,
			yTo:   overlapBox.yTo,
			zFrom: box.zFrom,
			zTo:   overlapBox.zFrom - 1,
			flip:  box.flip,
		},
	}
}

func findBottomRightEdge(box, overlapBox instruction) []instruction {
	// there is no edge here, the overlap box is at either the bottom or the left edge.
	if box.zFrom == overlapBox.zFrom || box.yTo == overlapBox.yTo {
		return nil
	}

	return []instruction{
		{
			xFrom: overlapBox.xFrom,
			xTo:   overlapBox.xTo,
			yFrom: overlapBox.yTo + 1,
			yTo:   box.yTo,
			zFrom: box.zFrom,
			zTo:   overlapBox.zFrom - 1,
			flip:  box.flip,
		},
	}
}

func findBottomFrontEdge(box, overlapBox instruction) []instruction {
	// there is no edge here, the overlap box is at either the bottom or the back edge.
	if box.zFrom == overlapBox.zFrom || box.xTo == overlapBox.xTo {
		return nil
	}

	return []instruction{
		{
			xFrom: overlapBox.xTo + 1,
			xTo:   box.xTo,
			yFrom: overlapBox.yFrom,
			yTo:   overlapBox.yTo,
			zFrom: box.zFrom,
			zTo:   overlapBox.zFrom - 1,
			flip:  box.flip,
		},
	}
}

func findFrontRightEdge(box, overlapBox instruction) []instruction {
	// there is no edge here, the overlap box is at either the front or the left edge.
	if box.xTo == overlapBox.xTo || box.yTo == overlapBox.yTo {
		return nil
	}

	return []instruction{
		{
			xFrom: overlapBox.xTo + 1,
			xTo:   box.xTo,
			yFrom: overlapBox.yTo + 1,
			yTo:   box.yTo,
			zFrom: overlapBox.zFrom,
			zTo:   overlapBox.zTo,
			flip:  box.flip,
		},
	}
}

func findFrontLeftEdge(box, overlapBox instruction) []instruction {
	// there is no edge here, the overlap box is at either the front or the left edge.
	if box.xTo == overlapBox.xTo || box.yFrom == overlapBox.yFrom {
		return nil
	}

	return []instruction{
		{
			xFrom: overlapBox.xTo + 1,
			xTo:   box.xTo,
			yFrom: box.yFrom,
			yTo:   overlapBox.yFrom - 1,
			zFrom: overlapBox.zFrom,
			zTo:   overlapBox.zTo,
			flip:  box.flip,
		},
	}
}

func findBackRightEdge(box, overlapBox instruction) []instruction {
	// there is no edge here, the overlap box is at either the back or the left edge.
	if box.xFrom == overlapBox.xFrom || box.yTo == overlapBox.yTo {
		return nil
	}

	return []instruction{
		{
			xFrom: box.xFrom,
			xTo:   overlapBox.xFrom - 1,
			yFrom: overlapBox.yTo + 1,
			yTo:   box.yTo,
			zFrom: overlapBox.zFrom,
			zTo:   overlapBox.zTo,
			flip:  box.flip,
		},
	}
}

func findBackLeftEdge(box, overlapBox instruction) []instruction {
	// there is no edge here, the overlap box is at either the back or the left edge.
	if box.xFrom == overlapBox.xFrom || box.yFrom == overlapBox.yFrom {
		return nil
	}

	return []instruction{
		{
			xFrom: box.xFrom,
			xTo:   overlapBox.xFrom - 1,
			yFrom: box.yFrom,
			yTo:   overlapBox.yFrom - 1,
			zFrom: overlapBox.zFrom,
			zTo:   overlapBox.zTo,
			flip:  box.flip,
		},
	}
}

// find corners, 8 of them.
func findTopBackLeftCorner(box, overlapBox instruction) []instruction {
	// there is no corner here, the overlap box is at either the back or the left or the top edge.
	if box.zTo == overlapBox.zTo || box.xFrom == overlapBox.xFrom || box.yFrom == overlapBox.yFrom {
		return nil
	}

	return []instruction{
		{
			xFrom: box.xFrom,
			xTo:   overlapBox.xFrom - 1,
			yFrom: box.yFrom,
			yTo:   overlapBox.yFrom - 1,
			zFrom: overlapBox.zTo + 1,
			zTo:   box.zTo,
			flip:  box.flip,
		},
	}
}

func findTopBackRightCorner(box, overlapBox instruction) []instruction {
	// there is no corner here, the overlap box is at either the back or the right or the top edge.
	if box.zTo == overlapBox.zTo || box.xFrom == overlapBox.xFrom || box.yTo == overlapBox.yTo {
		return nil
	}

	return []instruction{
		{
			xFrom: box.xFrom,
			xTo:   overlapBox.xFrom - 1,
			yFrom: overlapBox.yTo + 1,
			yTo:   box.yTo,
			zFrom: overlapBox.zTo + 1,
			zTo:   box.zTo,
			flip:  box.flip,
		},
	}
}

func findTopFrontLeftCorner(box, overlapBox instruction) []instruction {
	// there is no corner here, the overlap box is at either the front or the left or the top edge.
	if box.zTo == overlapBox.zTo || box.xTo == overlapBox.xTo || box.yFrom == overlapBox.yFrom {
		return nil
	}

	return []instruction{
		{
			xFrom: overlapBox.xTo + 1,
			xTo:   box.xTo,
			yFrom: box.yFrom,
			yTo:   overlapBox.yFrom - 1,
			zFrom: overlapBox.zTo + 1,
			zTo:   box.zTo,
			flip:  box.flip,
		},
	}
}

func findTopFrontRightCorner(box, overlapBox instruction) []instruction {
	// there is no corner here, the overlap box is at either the front or the right or the top edge.
	if box.zTo == overlapBox.zTo || box.xTo == overlapBox.xTo || box.yTo == overlapBox.yTo {
		return nil
	}

	return []instruction{
		{
			xFrom: overlapBox.xTo + 1,
			xTo:   box.xTo,
			yFrom: overlapBox.yTo + 1,
			yTo:   box.yTo,
			zFrom: overlapBox.zTo + 1,
			zTo:   box.zTo,
			flip:  box.flip,
		},
	}
}

func findBottomBackLeftCorner(box, overlapBox instruction) []instruction {
	// there is no corner here, the overlap box is at either the back or the left or the bottom edge.
	if box.zFrom == overlapBox.zFrom || box.xFrom == overlapBox.xFrom || box.yFrom == overlapBox.yFrom {
		return nil
	}

	return []instruction{
		{
			xFrom: box.xFrom,
			xTo:   overlapBox.xFrom - 1,
			yFrom: box.yFrom,
			yTo:   overlapBox.yFrom - 1,
			zFrom: box.zFrom,
			zTo:   overlapBox.zFrom - 1,
			flip:  box.flip,
		},
	}
}

func findBottomBackRightCorner(box, overlapBox instruction) []instruction {
	// there is no corner here, the overlap box is at either the back or the right or the bottom edge.
	if box.zFrom == overlapBox.zFrom || box.xFrom == overlapBox.xFrom || box.yTo == overlapBox.yTo {
		return nil
	}

	return []instruction{
		{
			xFrom: box.xFrom,
			xTo:   overlapBox.xFrom - 1,
			yFrom: overlapBox.yTo + 1,
			yTo:   box.yTo,
			zFrom: box.zFrom,
			zTo:   overlapBox.zFrom - 1,
			flip:  box.flip,
		},
	}
}

func findBottomFrontLeftCorner(box, overlapBox instruction) []instruction {
	// there is no corner here, the overlap box is at either the front or the left or the bottom edge.
	if box.zFrom == overlapBox.zFrom || box.xTo == overlapBox.xTo || box.yFrom == overlapBox.yFrom {
		return nil
	}

	return []instruction{
		{
			xFrom: overlapBox.xTo + 1,
			xTo:   box.xTo,
			yFrom: box.yFrom,
			yTo:   overlapBox.yFrom - 1,
			zFrom: box.zFrom,
			zTo:   overlapBox.zFrom - 1,
			flip:  box.flip,
		},
	}
}

func findBottomFrontRightCorner(box, overlapBox instruction) []instruction {
	// there is no corner here, the overlap box is at either the front or the right or the bottom edge.
	if box.zFrom == overlapBox.zFrom || box.xTo == overlapBox.xTo || box.yTo == overlapBox.yTo {
		return nil
	}

	return []instruction{
		{
			xFrom: overlapBox.xTo + 1,
			xTo:   box.xTo,
			yFrom: overlapBox.yTo + 1,
			yTo:   box.yTo,
			zFrom: box.zFrom,
			zTo:   overlapBox.zFrom - 1,
			flip:  box.flip,
		},
	}
}

func mergeBoxes(box, otherBox instruction) []instruction {
	// different flips will never be merged
	if box.flip != otherBox.flip {
		return []instruction{box, otherBox}
	}

	// try to merge along the x axis. yfrom, yto, zfrom, zto need to match, xfrom == xto
	if box.zFrom == otherBox.zFrom && box.yFrom == otherBox.yFrom && box.zTo == otherBox.zTo && box.yTo == otherBox.yTo {
		smallerXFrom := box.xFrom
		if otherBox.xFrom < smallerXFrom {
			smallerXFrom = otherBox.xFrom
		}

		largerXTo := box.xTo
		if otherBox.xTo > largerXTo {
			largerXTo = otherBox.xTo
		}

		if (box.xFrom <= otherBox.xFrom && box.xTo >= otherBox.xFrom) ||
			(otherBox.xFrom <= box.xFrom && otherBox.xTo >= box.xFrom) {
			// If they don't touch.
			return []instruction{
				{
					xFrom: smallerXFrom,
					xTo:   largerXTo,
					yFrom: box.yFrom,
					yTo:   box.yTo,
					zFrom: box.zFrom,
					zTo:   box.zTo,
					flip:  box.flip,
				},
			}
		}
	}

	// try to merge along the y axis. xfrom, xto, zfrom, zto need to match, yfrom == yto
	if box.zFrom == otherBox.zFrom && box.zTo == otherBox.zTo && box.xFrom == otherBox.xFrom && box.xTo == otherBox.xTo {
		smallerYFrom := box.yFrom
		if otherBox.yFrom < smallerYFrom {
			smallerYFrom = otherBox.yFrom
		}

		largerYTo := box.yTo
		if otherBox.yTo > largerYTo {
			largerYTo = otherBox.yTo
		}

		if (box.yFrom <= otherBox.yFrom && box.yTo >= otherBox.yFrom) ||
			(otherBox.yFrom <= box.yFrom && otherBox.yTo >= box.yFrom) {
			// If they don't touch.
			return []instruction{
				{
					xFrom: box.xFrom,
					xTo:   box.xTo,
					yFrom: smallerYFrom,
					yTo:   largerYTo,
					zFrom: box.zFrom,
					zTo:   box.zTo,
					flip:  box.flip,
				},
			}
		}
	}

	// try to merge along the z axis. xfrom, xto, yfrom, yto need to match.
	if box.yFrom == otherBox.yFrom && box.yTo == otherBox.yTo && box.xFrom == otherBox.xFrom && box.xTo == otherBox.xTo {
		smallerZFrom := box.zFrom
		if otherBox.zFrom < smallerZFrom {
			smallerZFrom = otherBox.zFrom
		}

		largerZTo := box.zTo
		if otherBox.zTo > largerZTo {
			largerZTo = otherBox.zTo
		}

		if (box.zFrom <= otherBox.zFrom && box.zTo >= otherBox.zFrom) ||
			(otherBox.zFrom <= box.zFrom && otherBox.zTo >= box.zFrom) {
			// If there's no gap between the two boxes in z.
			return []instruction{
				{
					xFrom: box.xFrom,
					xTo:   box.xTo,
					yFrom: box.yFrom,
					yTo:   box.yTo,
					zFrom: smallerZFrom,
					zTo:   largerZTo,
					flip:  box.flip,
				},
			}
		}
	}

	return []instruction{box, otherBox}
}

func filterOffs(instructions map[string]instruction) map[string]instruction {
	onlyOns := make(map[string]instruction)

	for k, i := range instructions {
		if i.flip == on {
			onlyOns[k] = i
		}
	}

	return onlyOns
}
