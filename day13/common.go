package day13

import (
	"io/ioutil"
	"math"
	"strings"
)

//var folds = regexp.MustCompile(`^fold along (.)=(\d+)$`)

// getInputs reads the input.txt file and returns them as a slice of strings for each row.
func getInputs(fn string) ([]string, []string) {
	data, err := ioutil.ReadFile(fn)
	if err != nil {
		panic(err)
	}

	dotsAndFolds := strings.Split(string(data), "\n\n")

	return strings.Split(dotsAndFolds[0], "\n"), strings.Split(dotsAndFolds[1], "\n")
}

func makePaper(dots []string) map[uint]uint {
	paper := make(map[uint]uint)
	x := make([]uint, 0)
	y := make([]uint, 0)
	sawComma := true

	for _, line := range dots {
		// reset counters and accumulators.
		sawComma = false
		x = x[:0]
		y = y[:0]

		for _, char := range line {
			num := charToInt[char]
			// if it's a comma, set the comma
			if num > 9 {
				sawComma = true

				continue
			}
			// if we've seen a comma, start accumulating into other slice.
			if !sawComma {
				x = append(x, num)
			} else {
				y = append(y, num)
			}
		}

		xacc := uint(0)
		xlen := len(x)

		for xi, xn := range x {
			xacc += xn * uint(math.Pow10(xlen-xi-1))
		}

		yacc := uint(0)
		ylen := len(y)

		for yi, yn := range y {
			yacc += yn * uint(math.Pow10(ylen-yi-1))
		}

		paper[xacc<<11|yacc] = 1
	}

	return paper
}

var charToInt = map[int32]uint{
	0x30: 0,
	0x31: 1,
	0x32: 2,
	0x33: 3,
	0x34: 4,
	0x35: 5,
	0x36: 6,
	0x37: 7,
	0x38: 8,
	0x39: 9,
	0x2c: 10,
}

//
//1
//2
//4
//8
//16 - 5
//32
//64
//128
//256
//512 - 10
//1024
//2048
