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
	XFrom int  `json:"x_from,omitempty"`
	XTo   int  `json:"x_to,omitempty"`
	YFrom int  `json:"y_from,omitempty"`
	YTo   int  `json:"y_to,omitempty"`
	ZFrom int  `json:"z_from,omitempty"`
	ZTo   int  `json:"z_to,omitempty"`
	Flip  flip `json:"Flip,omitempty"`
}

func (i instruction) Lights() int {
	return (i.XTo - i.XFrom + 1) * (i.YTo - i.YFrom + 1) * (i.ZTo - i.ZFrom + 1)
}

func (i instruction) String() string {
	return fmt.Sprintf("%d/%d/%d/%d/%d/%d/%s",
		i.XFrom, i.XTo,
		i.YFrom, i.YTo,
		i.ZFrom, i.ZTo,
		i.Flip,
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
		XFrom: numbers[0],
		XTo:   numbers[1],
		YFrom: numbers[2],
		YTo:   numbers[3],
		ZFrom: numbers[4],
		ZTo:   numbers[5],
		Flip:  ins,
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
	if box.XFrom > otherBox.XTo || box.YFrom > otherBox.YTo || box.ZFrom > otherBox.ZTo {
		return []instruction{box, otherBox}
	}

	// they do not overlap, because box doesn't start until after otherBox ends.
	if otherBox.XFrom > box.XTo || otherBox.YFrom > box.YTo || otherBox.ZFrom > box.ZTo {
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
	xmin := box.XFrom
	if otherBox.XFrom > xmin {
		xmin = otherBox.XFrom
	}

	xmax := box.XTo
	if otherBox.XTo < xmax {
		xmax = otherBox.XTo
	}

	ymin := box.YFrom
	if otherBox.YFrom > ymin {
		ymin = otherBox.YFrom
	}

	ymax := box.YTo
	if otherBox.YTo < ymax {
		ymax = otherBox.YTo
	}

	zmin := box.ZFrom
	if otherBox.ZFrom > zmin {
		zmin = otherBox.ZFrom
	}

	zmax := box.ZTo
	if otherBox.ZTo < zmax {
		zmax = otherBox.ZTo
	}

	if xmin > xmax || ymin > ymax || zmin > zmax {
		return instruction{}, errors.New("out of bounds")
	}

	return instruction{
		XFrom: xmin,
		XTo:   xmax,
		YFrom: ymin,
		YTo:   ymax,
		ZFrom: zmin,
		ZTo:   zmax,
		Flip:  otherBox.Flip,
	}, nil
}

// find faces, 6 of them.
func findTopFace(box, overlapBox instruction) []instruction {
	// there is no face here, the overlap box is at the edge
	if box.ZTo == overlapBox.ZTo {
		return nil
	}

	return []instruction{
		{
			XFrom: overlapBox.XFrom,
			XTo:   overlapBox.XTo,
			YFrom: overlapBox.YFrom,
			YTo:   overlapBox.YTo,
			ZFrom: overlapBox.ZTo + 1,
			ZTo:   box.ZTo,
			Flip:  box.Flip,
		},
	}
}

func findBottomFace(box, overlapBox instruction) []instruction {
	// there is no face here, the overlap box is at the edge
	if box.ZFrom == overlapBox.ZFrom {
		return nil
	}

	return []instruction{
		{
			XFrom: overlapBox.XFrom,
			XTo:   overlapBox.XTo,
			YFrom: overlapBox.YFrom,
			YTo:   overlapBox.YTo,
			ZFrom: box.ZFrom,
			ZTo:   overlapBox.ZFrom - 1,
			Flip:  box.Flip,
		},
	}
}

func findFrontFace(box, overlapBox instruction) []instruction {
	// there is no face here, the overlap box is at the front edge
	if box.XTo == overlapBox.XTo {
		return nil
	}

	return []instruction{
		{
			XFrom: overlapBox.XTo + 1,
			XTo:   box.XTo,
			YFrom: overlapBox.YFrom,
			YTo:   overlapBox.YTo,
			ZFrom: overlapBox.ZFrom,
			ZTo:   overlapBox.ZTo,
			Flip:  box.Flip,
		},
	}
}

func findBackFace(box, overlapBox instruction) []instruction {
	// there is no face here, the overlap box is at the back edge
	if box.XFrom == overlapBox.XFrom {
		return nil
	}

	return []instruction{
		{
			XFrom: box.XFrom,
			XTo:   overlapBox.XFrom - 1,
			YFrom: overlapBox.YFrom,
			YTo:   overlapBox.YTo,
			ZFrom: overlapBox.ZFrom,
			ZTo:   overlapBox.ZTo,
			Flip:  box.Flip,
		},
	}
}

func findLeftFace(box, overlapBox instruction) []instruction {
	// there is no face here, the overlap box is at the left edge
	if box.YFrom == overlapBox.YFrom {
		return nil
	}

	return []instruction{
		{
			XFrom: overlapBox.XFrom,
			XTo:   overlapBox.XTo,
			YFrom: box.YFrom,
			YTo:   overlapBox.YFrom - 1,
			ZFrom: overlapBox.ZFrom,
			ZTo:   overlapBox.ZTo,
			Flip:  box.Flip,
		},
	}
}

func findRightFace(box, overlapBox instruction) []instruction {
	// there is no face here, the overlap box is at the right edge
	if box.YTo == overlapBox.YTo {
		return nil
	}

	return []instruction{
		{
			XFrom: overlapBox.XFrom,
			XTo:   overlapBox.XTo,
			YFrom: overlapBox.YTo + 1,
			YTo:   box.YTo,
			ZFrom: overlapBox.ZFrom,
			ZTo:   overlapBox.ZTo,
			Flip:  box.Flip,
		},
	}
}

// find edges, 12 of them.
func findTopLeftEdge(box, overlapBox instruction) []instruction {
	// there is no edge here, the overlap box is at either the top or the left edge.
	if box.ZTo == overlapBox.ZTo || box.YFrom == overlapBox.YFrom {
		return nil
	}

	return []instruction{
		{
			XFrom: overlapBox.XFrom,
			XTo:   overlapBox.XTo,
			YFrom: box.YFrom,
			YTo:   overlapBox.YFrom - 1,
			ZFrom: overlapBox.ZTo + 1,
			ZTo:   box.ZTo,
			Flip:  box.Flip,
		},
	}
}

func findTopBackEdge(box, overlapBox instruction) []instruction {
	// there is no edge here, the overlap box is at either the top or the back edge.
	if box.ZTo == overlapBox.ZTo || box.XFrom == overlapBox.XFrom {
		return nil
	}

	return []instruction{
		{
			XFrom: box.XFrom,
			XTo:   overlapBox.XFrom - 1,
			YFrom: overlapBox.YFrom,
			YTo:   overlapBox.YTo,
			ZFrom: overlapBox.ZTo + 1,
			ZTo:   box.ZTo,
			Flip:  box.Flip,
		},
	}
}

func findTopRightEdge(box, overlapBox instruction) []instruction {
	// there is no edge here, the overlap box is at either the top or the left edge.
	if box.ZTo == overlapBox.ZTo || box.YTo == overlapBox.YTo {
		return nil
	}

	return []instruction{
		{
			XFrom: overlapBox.XFrom,
			XTo:   overlapBox.XTo,
			YFrom: overlapBox.YTo + 1,
			YTo:   box.YTo,
			ZFrom: overlapBox.ZTo + 1,
			ZTo:   box.ZTo,
			Flip:  box.Flip,
		},
	}
}

func findTopFrontEdge(box, overlapBox instruction) []instruction {
	// there is no edge here, the overlap box is at either the top or the back edge.
	if box.ZTo == overlapBox.ZTo || box.XTo == overlapBox.XTo {
		return nil
	}

	return []instruction{
		{
			XFrom: overlapBox.XTo + 1,
			XTo:   box.XTo,
			YFrom: overlapBox.YFrom,
			YTo:   overlapBox.YTo,
			ZFrom: overlapBox.ZTo + 1,
			ZTo:   box.ZTo,
			Flip:  box.Flip,
		},
	}
}

func findBottomLeftEdge(box, overlapBox instruction) []instruction {
	// there is no edge here, the overlap box is at either the bottom or the left edge.
	if box.ZFrom == overlapBox.ZFrom || box.YFrom == overlapBox.YFrom {
		return nil
	}

	return []instruction{
		{
			XFrom: overlapBox.XFrom,
			XTo:   overlapBox.XTo,
			YFrom: box.YFrom,
			YTo:   overlapBox.YFrom - 1,
			ZFrom: box.ZFrom,
			ZTo:   overlapBox.ZFrom - 1,
			Flip:  box.Flip,
		},
	}
}

func findBottomBackEdge(box, overlapBox instruction) []instruction {
	// there is no edge here, the overlap box is at either the bottom or the back edge.
	if box.ZFrom == overlapBox.ZFrom || box.XFrom == overlapBox.XFrom {
		return nil
	}

	return []instruction{
		{
			XFrom: box.XFrom,
			XTo:   overlapBox.XFrom - 1,
			YFrom: overlapBox.YFrom,
			YTo:   overlapBox.YTo,
			ZFrom: box.ZFrom,
			ZTo:   overlapBox.ZFrom - 1,
			Flip:  box.Flip,
		},
	}
}

func findBottomRightEdge(box, overlapBox instruction) []instruction {
	// there is no edge here, the overlap box is at either the bottom or the left edge.
	if box.ZFrom == overlapBox.ZFrom || box.YTo == overlapBox.YTo {
		return nil
	}

	return []instruction{
		{
			XFrom: overlapBox.XFrom,
			XTo:   overlapBox.XTo,
			YFrom: overlapBox.YTo + 1,
			YTo:   box.YTo,
			ZFrom: box.ZFrom,
			ZTo:   overlapBox.ZFrom - 1,
			Flip:  box.Flip,
		},
	}
}

func findBottomFrontEdge(box, overlapBox instruction) []instruction {
	// there is no edge here, the overlap box is at either the bottom or the back edge.
	if box.ZFrom == overlapBox.ZFrom || box.XTo == overlapBox.XTo {
		return nil
	}

	return []instruction{
		{
			XFrom: overlapBox.XTo + 1,
			XTo:   box.XTo,
			YFrom: overlapBox.YFrom,
			YTo:   overlapBox.YTo,
			ZFrom: box.ZFrom,
			ZTo:   overlapBox.ZFrom - 1,
			Flip:  box.Flip,
		},
	}
}

func findFrontRightEdge(box, overlapBox instruction) []instruction {
	// there is no edge here, the overlap box is at either the front or the left edge.
	if box.XTo == overlapBox.XTo || box.YTo == overlapBox.YTo {
		return nil
	}

	return []instruction{
		{
			XFrom: overlapBox.XTo + 1,
			XTo:   box.XTo,
			YFrom: overlapBox.YTo + 1,
			YTo:   box.YTo,
			ZFrom: overlapBox.ZFrom,
			ZTo:   overlapBox.ZTo,
			Flip:  box.Flip,
		},
	}
}

func findFrontLeftEdge(box, overlapBox instruction) []instruction {
	// there is no edge here, the overlap box is at either the front or the left edge.
	if box.XTo == overlapBox.XTo || box.YFrom == overlapBox.YFrom {
		return nil
	}

	return []instruction{
		{
			XFrom: overlapBox.XTo + 1,
			XTo:   box.XTo,
			YFrom: box.YFrom,
			YTo:   overlapBox.YFrom - 1,
			ZFrom: overlapBox.ZFrom,
			ZTo:   overlapBox.ZTo,
			Flip:  box.Flip,
		},
	}
}

func findBackRightEdge(box, overlapBox instruction) []instruction {
	// there is no edge here, the overlap box is at either the back or the left edge.
	if box.XFrom == overlapBox.XFrom || box.YTo == overlapBox.YTo {
		return nil
	}

	return []instruction{
		{
			XFrom: box.XFrom,
			XTo:   overlapBox.XFrom - 1,
			YFrom: overlapBox.YTo + 1,
			YTo:   box.YTo,
			ZFrom: overlapBox.ZFrom,
			ZTo:   overlapBox.ZTo,
			Flip:  box.Flip,
		},
	}
}

func findBackLeftEdge(box, overlapBox instruction) []instruction {
	// there is no edge here, the overlap box is at either the back or the left edge.
	if box.XFrom == overlapBox.XFrom || box.YFrom == overlapBox.YFrom {
		return nil
	}

	return []instruction{
		{
			XFrom: box.XFrom,
			XTo:   overlapBox.XFrom - 1,
			YFrom: box.YFrom,
			YTo:   overlapBox.YFrom - 1,
			ZFrom: overlapBox.ZFrom,
			ZTo:   overlapBox.ZTo,
			Flip:  box.Flip,
		},
	}
}

// find corners, 8 of them.
func findTopBackLeftCorner(box, overlapBox instruction) []instruction {
	// there is no corner here, the overlap box is at either the back or the left or the top edge.
	if box.ZTo == overlapBox.ZTo || box.XFrom == overlapBox.XFrom || box.YFrom == overlapBox.YFrom {
		return nil
	}

	return []instruction{
		{
			XFrom: box.XFrom,
			XTo:   overlapBox.XFrom - 1,
			YFrom: box.YFrom,
			YTo:   overlapBox.YFrom - 1,
			ZFrom: overlapBox.ZTo + 1,
			ZTo:   box.ZTo,
			Flip:  box.Flip,
		},
	}
}

func findTopBackRightCorner(box, overlapBox instruction) []instruction {
	// there is no corner here, the overlap box is at either the back or the right or the top edge.
	if box.ZTo == overlapBox.ZTo || box.XFrom == overlapBox.XFrom || box.YTo == overlapBox.YTo {
		return nil
	}

	return []instruction{
		{
			XFrom: box.XFrom,
			XTo:   overlapBox.XFrom - 1,
			YFrom: overlapBox.YTo + 1,
			YTo:   box.YTo,
			ZFrom: overlapBox.ZTo + 1,
			ZTo:   box.ZTo,
			Flip:  box.Flip,
		},
	}
}

func findTopFrontLeftCorner(box, overlapBox instruction) []instruction {
	// there is no corner here, the overlap box is at either the front or the left or the top edge.
	if box.ZTo == overlapBox.ZTo || box.XTo == overlapBox.XTo || box.YFrom == overlapBox.YFrom {
		return nil
	}

	return []instruction{
		{
			XFrom: overlapBox.XTo + 1,
			XTo:   box.XTo,
			YFrom: box.YFrom,
			YTo:   overlapBox.YFrom - 1,
			ZFrom: overlapBox.ZTo + 1,
			ZTo:   box.ZTo,
			Flip:  box.Flip,
		},
	}
}

func findTopFrontRightCorner(box, overlapBox instruction) []instruction {
	// there is no corner here, the overlap box is at either the front or the right or the top edge.
	if box.ZTo == overlapBox.ZTo || box.XTo == overlapBox.XTo || box.YTo == overlapBox.YTo {
		return nil
	}

	return []instruction{
		{
			XFrom: overlapBox.XTo + 1,
			XTo:   box.XTo,
			YFrom: overlapBox.YTo + 1,
			YTo:   box.YTo,
			ZFrom: overlapBox.ZTo + 1,
			ZTo:   box.ZTo,
			Flip:  box.Flip,
		},
	}
}

func findBottomBackLeftCorner(box, overlapBox instruction) []instruction {
	// there is no corner here, the overlap box is at either the back or the left or the bottom edge.
	if box.ZFrom == overlapBox.ZFrom || box.XFrom == overlapBox.XFrom || box.YFrom == overlapBox.YFrom {
		return nil
	}

	return []instruction{
		{
			XFrom: box.XFrom,
			XTo:   overlapBox.XFrom - 1,
			YFrom: box.YFrom,
			YTo:   overlapBox.YFrom - 1,
			ZFrom: box.ZFrom,
			ZTo:   overlapBox.ZFrom - 1,
			Flip:  box.Flip,
		},
	}
}

func findBottomBackRightCorner(box, overlapBox instruction) []instruction {
	// there is no corner here, the overlap box is at either the back or the right or the bottom edge.
	if box.ZFrom == overlapBox.ZFrom || box.XFrom == overlapBox.XFrom || box.YTo == overlapBox.YTo {
		return nil
	}

	return []instruction{
		{
			XFrom: box.XFrom,
			XTo:   overlapBox.XFrom - 1,
			YFrom: overlapBox.YTo + 1,
			YTo:   box.YTo,
			ZFrom: box.ZFrom,
			ZTo:   overlapBox.ZFrom - 1,
			Flip:  box.Flip,
		},
	}
}

func findBottomFrontLeftCorner(box, overlapBox instruction) []instruction {
	// there is no corner here, the overlap box is at either the front or the left or the bottom edge.
	if box.ZFrom == overlapBox.ZFrom || box.XTo == overlapBox.XTo || box.YFrom == overlapBox.YFrom {
		return nil
	}

	return []instruction{
		{
			XFrom: overlapBox.XTo + 1,
			XTo:   box.XTo,
			YFrom: box.YFrom,
			YTo:   overlapBox.YFrom - 1,
			ZFrom: box.ZFrom,
			ZTo:   overlapBox.ZFrom - 1,
			Flip:  box.Flip,
		},
	}
}

func findBottomFrontRightCorner(box, overlapBox instruction) []instruction {
	// there is no corner here, the overlap box is at either the front or the right or the bottom edge.
	if box.ZFrom == overlapBox.ZFrom || box.XTo == overlapBox.XTo || box.YTo == overlapBox.YTo {
		return nil
	}

	return []instruction{
		{
			XFrom: overlapBox.XTo + 1,
			XTo:   box.XTo,
			YFrom: overlapBox.YTo + 1,
			YTo:   box.YTo,
			ZFrom: box.ZFrom,
			ZTo:   overlapBox.ZFrom - 1,
			Flip:  box.Flip,
		},
	}
}

func mergeBoxes(box, otherBox instruction) []instruction {
	// different flips will never be merged
	if box.Flip != otherBox.Flip {
		return []instruction{box, otherBox}
	}

	// try to merge along the x axis. yfrom, yto, zfrom, zto need to match, xfrom == xto
	if box.ZFrom == otherBox.ZFrom && box.ZTo == otherBox.ZTo && box.YFrom == otherBox.YFrom && box.YTo == otherBox.YTo {
		smallerXFrom := box.XFrom
		if otherBox.XFrom < smallerXFrom {
			smallerXFrom = otherBox.XFrom
		}

		largerXTo := box.XTo
		if otherBox.XTo > largerXTo {
			largerXTo = otherBox.XTo
		}

		if (box.XFrom <= otherBox.XFrom && box.XTo >= otherBox.XFrom) ||
			(otherBox.XFrom <= box.XFrom && otherBox.XTo >= box.XFrom) ||
			(box.XTo+1 == otherBox.XFrom || otherBox.XTo+1 == box.XFrom) {
			// If they don't touch.
			return []instruction{
				{
					XFrom: smallerXFrom,
					XTo:   largerXTo,
					YFrom: box.YFrom,
					YTo:   box.YTo,
					ZFrom: box.ZFrom,
					ZTo:   box.ZTo,
					Flip:  box.Flip,
				},
			}
		}
	}

	// try to merge along the y axis. xfrom, xto, zfrom, zto need to match, yfrom == yto
	if box.ZFrom == otherBox.ZFrom && box.ZTo == otherBox.ZTo && box.XFrom == otherBox.XFrom && box.XTo == otherBox.XTo {
		smallerYFrom := box.YFrom
		if otherBox.YFrom < smallerYFrom {
			smallerYFrom = otherBox.YFrom
		}

		largerYTo := box.YTo
		if otherBox.YTo > largerYTo {
			largerYTo = otherBox.YTo
		}

		if (box.YFrom <= otherBox.YFrom && box.YTo >= otherBox.YFrom) ||
			(otherBox.YFrom <= box.YFrom && otherBox.YTo >= box.YFrom) ||
			(box.YTo+1 == otherBox.YFrom || otherBox.YTo+1 == box.YFrom) {
			// If they don't touch.
			return []instruction{
				{
					XFrom: box.XFrom,
					XTo:   box.XTo,
					YFrom: smallerYFrom,
					YTo:   largerYTo,
					ZFrom: box.ZFrom,
					ZTo:   box.ZTo,
					Flip:  box.Flip,
				},
			}
		}
	}

	// try to merge along the z axis. xfrom, xto, yfrom, yto need to match.
	if box.YFrom == otherBox.YFrom && box.YTo == otherBox.YTo && box.XFrom == otherBox.XFrom && box.XTo == otherBox.XTo {
		smallerZFrom := box.ZFrom
		if otherBox.ZFrom < smallerZFrom {
			smallerZFrom = otherBox.ZFrom
		}

		largerZTo := box.ZTo
		if otherBox.ZTo > largerZTo {
			largerZTo = otherBox.ZTo
		}

		if (box.ZFrom <= otherBox.ZFrom && box.ZTo >= otherBox.ZFrom) ||
			(otherBox.ZFrom <= box.ZFrom && otherBox.ZTo >= box.ZFrom) ||
			(box.ZTo+1 == otherBox.ZFrom || otherBox.ZTo+1 == box.ZFrom) {
			// If there's no gap between the two boxes in z.
			return []instruction{
				{
					XFrom: box.XFrom,
					XTo:   box.XTo,
					YFrom: box.YFrom,
					YTo:   box.YTo,
					ZFrom: smallerZFrom,
					ZTo:   largerZTo,
					Flip:  box.Flip,
				},
			}
		}
	}

	return []instruction{box, otherBox}
}

func filterOffs(instructions map[string]instruction) map[string]instruction {
	onlyOns := make(map[string]instruction)

	for k, i := range instructions {
		if i.Flip == on {
			onlyOns[k] = i
		}
	}

	return onlyOns
}
