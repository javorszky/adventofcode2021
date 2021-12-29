package day22

import (
	"errors"
)

const (
	cubeLowerLimit = -50
	cubeUpperLimit = 50
)

func task1(input []instruction) interface{} {
	limited := make([]instruction, 0)

	for _, ins := range input {
		limitedInstruction, err := limitInstructionTo50(ins)
		if err != nil {
			continue
		}

		limited = append(limited, limitedInstruction)
	}

	kuqe := make(cubespace)

	for i, inst := range limited {
		kuqe.applyInstructions(inst)

		if i > 1 {
			break
		}
	}

	return len(limited)
}

func limitInstructionTo50(in instruction) (instruction, error) {
	xmin, ymin, zmin := cubeLowerLimit, cubeLowerLimit, cubeLowerLimit
	xmax, ymax, zmax := cubeUpperLimit, cubeUpperLimit, cubeUpperLimit

	if in.xFrom > cubeUpperLimit ||
		in.yFrom > cubeUpperLimit ||
		in.zFrom > cubeUpperLimit ||
		in.xTo < cubeLowerLimit ||
		in.yTo < cubeLowerLimit ||
		in.zTo < cubeLowerLimit {
		// If any part of the cube is outside our bounds, then return an error.
		return instruction{}, errors.New("out of bounds")
	}

	if in.xFrom > xmin {
		xmin = in.xFrom
	}

	if in.xTo < xmax {
		xmax = in.xTo
	}

	if in.yFrom > ymin {
		ymin = in.yFrom
	}

	if in.yTo < ymax {
		ymax = in.yTo
	}

	if in.zFrom > zmin {
		zmin = in.zFrom
	}

	if in.zTo < zmax {
		zmax = in.zTo
	}

	return instruction{
		xFrom: xmin,
		xTo:   xmax,
		yFrom: ymin,
		yTo:   ymax,
		zFrom: zmin,
		zTo:   zmax,
		flip:  in.flip,
	}, nil
}
