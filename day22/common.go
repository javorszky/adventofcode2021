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
		if i.XFrom < xmin {
			xmin = i.XFrom
		}

		if i.YFrom < ymin {
			ymin = i.YFrom
		}

		if i.ZFrom < zmin {
			zmin = i.ZFrom
		}

		if i.XTo > xmax {
			xmax = i.XTo
		}

		if i.YTo > ymax {
			ymax = i.YTo
		}

		if i.ZTo > zmax {
			zmax = i.ZTo
		}
	}

	return instruction{
		XFrom: xmin,
		XTo:   xmax,
		YFrom: ymin,
		YTo:   ymax,
		ZFrom: zmin,
		ZTo:   zmax,
		Flip:  off,
	}
}

func mergeMap(mergeThis, intoThis map[string]instruction) map[string]instruction {
	for k, v := range mergeThis {
		intoThis[k] = v
	}

	return intoThis
}
