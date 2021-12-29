package day22

import (
	"io/ioutil"
	"strings"

	"github.com/javorszky/adventofcode2021/util"
)

// getInputs reads the input.txt file and returns them as a slice of strings for each row.
func getInputs(fn string) []string {
	data, err := ioutil.ReadFile(fn)
	if err != nil {
		panic(err)
	}

	return strings.Split(strings.TrimSpace(string(data)), util.NewLine)
}

func findBoundaries(in []instruction) instruction {
	xmin, ymin, zmin := 1<<32, 1<<32, 1<<32
	xmax, ymax, zmax := -1<<32, -1<<32, -1<<32

	for _, i := range in {
		if i.xFrom < xmin {
			xmin = i.xFrom
		}

		if i.yFrom < ymin {
			ymin = i.yFrom
		}

		if i.zFrom < zmin {
			zmin = i.zFrom
		}

		if i.xTo > xmax {
			xmax = i.xTo
		}

		if i.yTo > ymax {
			ymax = i.yTo
		}

		if i.zTo > zmax {
			zmax = i.zTo
		}
	}

	return instruction{
		xFrom: xmin,
		xTo:   xmax,
		yFrom: ymin,
		yTo:   ymax,
		zFrom: zmin,
		zTo:   zmax,
		flip:  off,
	}
}

func mergeMap(mergeThis, intoThis map[string]instruction) map[string]instruction {
	for k, v := range mergeThis {
		intoThis[k] = v
	}

	return intoThis
}
