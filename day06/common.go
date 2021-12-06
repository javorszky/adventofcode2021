package day06

import (
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

const filename = "day06/input.txt"

var byteToInt = map[uint8]int{
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
}

// getInputs reads the input.txt file and returns them as a slice of strings for each row.
func getInputs(fn string) string {
	data, err := ioutil.ReadFile(fn)
	if err != nil {
		panic(err)
	}

	return strings.TrimRight(string(data), "\n")
}

func parseFishSplitAtoi(in string) []int {
	numStrings := strings.Split(in, ",")
	out := make([]int, len(numStrings))

	for i, ns := range numStrings {
		n, err := strconv.Atoi(ns)
		if err != nil {
			log.Fatalf("failed to convert string [%s] into int: %s\n", ns, err)
		}

		out[i] = n
	}

	return out
}

func parseFishWalkAtoi(in string) []int {
	out := make([]int, 0)

	for _, ns := range in {
		switch ns {
		case 0xa, 0x2c:
			continue
		default:
			n, err := strconv.Atoi(string(ns))
			if err != nil {
				log.Fatalf("failed to convert string [%s] into int: %s\n", string(ns), err)
			}

			out = append(out, n)
		}
	}

	return out
}

func parseFishSplitMap(in string) []int {
	numStrings := strings.Split(in, ",")
	out := make([]int, len(numStrings))

	for i, ns := range numStrings {
		out[i] = byteToInt[[]byte(ns)[0]]
	}

	return out
}

func parseFishWalkMap(in string) []int {
	out := make([]int, 0)

	for _, ns := range in {
		switch ns {
		case 0xa, 0x2c:
			continue
		default:
			out = append(out, byteToInt[uint8(ns)])
		}
	}

	return out
}

func parseFishForAtoi(in string) []int {
	out := make([]int, 0)

	for i := 0; i < len(in); i++ {
		switch in[i] {
		case 0xa, 0x2c:
			continue
		default:
			n, err := strconv.Atoi(string(in[i]))
			if err != nil {
				log.Fatalf("failed to convert string [%s] into int: %s\n", string(in[i]), err)
			}

			out = append(out, n)
		}
	}

	return out
}

func parseFishForMap(in string) []int {
	out := make([]int, 0)

	for i := 0; i < len(in); i++ {
		switch in[i] {
		case 0xa, 0x2c:
			continue
		default:
			out = append(out, byteToInt[in[i]])
		}
	}

	return out
}
