package day22

import (
	"errors"
)

const (
	cubeLowerLimit = -50
	cubeUpperLimit = 50
)

func task1(input []instruction) int {
	limited := make([]instruction, 0)

	for _, ins := range input {
		limitedInstruction, err := limitInstructionTo50(ins)
		if err != nil {
			continue
		}

		limited = append(limited, limitedInstruction)
	}

	kuqe := make(cubespace)

	for _, inst := range limited {
		kuqe.applyInstructions(inst)
	}

	kuqe.Collapse()

	return kuqe.Lights()
}

func limitInstructionTo50(in instruction) (instruction, error) {
	xmin, ymin, zmin := cubeLowerLimit, cubeLowerLimit, cubeLowerLimit
	xmax, ymax, zmax := cubeUpperLimit, cubeUpperLimit, cubeUpperLimit

	if in.XFrom > cubeUpperLimit ||
		in.YFrom > cubeUpperLimit ||
		in.ZFrom > cubeUpperLimit ||
		in.XTo < cubeLowerLimit ||
		in.YTo < cubeLowerLimit ||
		in.ZTo < cubeLowerLimit {
		// If any part of the cube is outside our bounds, then return an error.
		return instruction{}, errors.New("out of bounds")
	}

	if in.XFrom > xmin {
		xmin = in.XFrom
	}

	if in.XTo < xmax {
		xmax = in.XTo
	}

	if in.YFrom > ymin {
		ymin = in.YFrom
	}

	if in.YTo < ymax {
		ymax = in.YTo
	}

	if in.ZFrom > zmin {
		zmin = in.ZFrom
	}

	if in.ZTo < zmax {
		zmax = in.ZTo
	}

	return instruction{
		XFrom: xmin,
		XTo:   xmax,
		YFrom: ymin,
		YTo:   ymax,
		ZFrom: zmin,
		ZTo:   zmax,
		Flip:  in.Flip,
	}, nil
}
