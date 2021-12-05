package day05

import (
	"io/ioutil"
	"log"
	"regexp"
	"strconv"
	"strings"
	"sync"
)

const (
	lineMatchLength = 5
)

var extractData = regexp.MustCompile(`^(\d+),(\d+) -> (\d+),(\d+)$`)
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

type tuple [2][2]uint

// getInputs reads the input.txt file and returns them as a slice of strings for each row.
func getInputs(fn string) []string {
	data, err := ioutil.ReadFile(fn)
	if err != nil {
		panic(err)
	}

	return strings.Split(strings.TrimRight(string(data), "\n"), "\n")
}

func getTuples(fileData []string) []tuple {
	tuples := make([]tuple, len(fileData))

	for i, line := range fileData {
		matches := extractData.FindStringSubmatch(line)
		if len(matches) != lineMatchLength {
			log.Fatalf("regex extracting data from line [%s] at index [%d] failed. Matches: %v, len should be 5, it's %d",
				line,
				i,
				matches,
				len(matches))
		}

		reInts := convertToUints(matches[1:])

		tuples[i] = tuple{
			{
				reInts[0],
				reInts[1],
			},
			{
				reInts[2],
				reInts[3],
			},
		}
	}

	return tuples
}

func getTuplesReversed(fileData []string) []tuple {
	tuples := make([]tuple, len(fileData))
	lineArray := [4]uint{}
	power := 1
	n := 0
	usedNonNumber := false
	acc := 0

	for j, line := range fileData {
		power = 1
		n = 0
		usedNonNumber = false
		acc = 0

		for i := len(line) - 1; i >= 0; i-- {
			switch line[i] {
			case 0x20, 0x2c, 0x2d, 0x3e:
				if usedNonNumber {
					continue
				}

				usedNonNumber = true
				lineArray[n] = uint(acc)
				acc = 0
				power = 1
				n++
			default:
				usedNonNumber = false
				acc += byteToInt[line[i]] * power

				power = power * 10
			}
		}

		lineArray[n] = uint(acc)

		tuples[j] = tuple{
			{lineArray[3], lineArray[2]},
			{lineArray[1], lineArray[0]},
		}
	}

	return tuples
}

func getTuplesString(fileData []string) []tuple {
	tuples := make([]tuple, len(fileData))

	for i, line := range fileData {
		pairs := strings.Split(line, " -> ")
		t1 := strings.Split(pairs[0], ",")
		t2 := strings.Split(pairs[1], ",")

		t1Uints := convertToUints(t1)
		t2Uints := convertToUints(t2)

		tuples[i] = tuple{
			{
				t1Uints[0],
				t1Uints[1],
			},
			{
				t2Uints[0],
				t2Uints[1],
			},
		}
	}

	return tuples
}

func getCoordinateSliceRegex(fileData []string) []uint {
	coords := make([]uint, len(fileData)*4)

	for i, line := range fileData {
		matches := extractData.FindStringSubmatch(line)
		if len(matches) != lineMatchLength {
			log.Fatalf("regex extracting data from line [%s] at index [%d] failed. Matches: %v, len should be 5, it's %d",
				line,
				i,
				matches,
				len(matches))
		}

		reInts := convertToUints(matches[1:])

		coords[i*4] = reInts[0]
		coords[i*4+1] = reInts[1]
		coords[i*4+2] = reInts[2]
		coords[i*4+3] = reInts[3]
	}

	return coords
}

func getCoordinateSliceStrings(fileData []string) []uint {
	coords := make([]uint, len(fileData)*4)

	for i, line := range fileData {
		pairs := strings.Split(line, " -> ")
		t1 := strings.Split(pairs[0], ",")
		t2 := strings.Split(pairs[1], ",")

		t1Uints := convertToUints(t1)
		t2Uints := convertToUints(t2)

		coords[i*4] = t1Uints[0]
		coords[i*4+1] = t1Uints[1]
		coords[i*4+2] = t2Uints[0]
		coords[i*4+3] = t2Uints[1]
	}

	return coords
}

func getCoordinateSliceReverse(fileData []string) []uint {
	coords := make([]uint, len(fileData)*4)
	power := 1
	n := 3
	usedNonNumber := false
	acc := 0

	for j, line := range fileData {
		power = 1
		n = 3
		usedNonNumber = false
		acc = 0

		for i := len(line) - 1; i >= 0; i-- {
			switch line[i] {
			case 0x20, 0x2c, 0x2d, 0x3e:
				if usedNonNumber {
					continue
				}

				usedNonNumber = true
				coords[j*4+n] = uint(acc)
				acc = 0
				power = 1
				n--
			default:
				usedNonNumber = false
				acc += byteToInt[line[i]] * power
				power = power * 10
			}
		}

		coords[j*4+n] = uint(acc)
	}

	return coords
}

func convertToUints(matches []string) []uint {
	reInts := make([]uint, len(matches))

	for i, m := range matches {
		num, err := strconv.Atoi(m)
		if err != nil {
			log.Fatalf("converting string [%s] to int: %s", m, err)
		}

		reInts[i] = uint(num)
	}

	return reInts
}

func mapLinesTuples(tuples []tuple) map[uint]uint {
	m := make(map[uint]uint)

	for _, t := range tuples {
		switch {
		case t[0][0] == t[1][0] && t[0][1] < t[1][1]:
			// first coördinate is the same and first tuple is smaller
			x := t[0][0] << bitSizeForCoordinates

			for i := t[0][1]; i <= t[1][1]; i++ {
				m[x|i]++
			}
		case t[0][0] == t[1][0] && t[0][1] > t[1][1]:
			// first coördinate is the same and second tuple is smaller
			x := t[0][0] << bitSizeForCoordinates

			for i := t[1][1]; i <= t[0][1]; i++ {
				m[x|i]++
			}
		case t[0][0] < t[1][0] && t[0][1] == t[1][1]:
			// second coördinate is the same, first tuple is smaller
			y := t[0][1]

			for i := t[0][0]; i <= t[1][0]; i++ {
				m[(i<<bitSizeForCoordinates)|y]++
			}
		case t[0][0] > t[1][0] && t[0][1] == t[1][1]:
			// second coördinate is the same, second tuple is smaller
			y := t[0][1]

			for i := t[1][0]; i <= t[0][0]; i++ {
				m[(i<<bitSizeForCoordinates)|y]++
			}
		case t[0][0] < t[1][0] && t[0][1] < t[1][1]:
			// top left to bottom right \
			for i := uint(0); i <= (t[1][0] - t[0][0]); i++ {
				x := t[0][0] + i
				y := t[0][1] + i
				m[x<<10|y]++
			}
		case t[0][0] < t[1][0] && t[0][1] > t[1][1]:
			// bottom left to top right /
			for i := uint(0); i <= (t[1][0] - t[0][0]); i++ {
				x := t[0][0] + i
				y := t[0][1] - i
				m[x<<10|y]++
			}
		case t[0][0] > t[1][0] && t[0][1] < t[1][1]:
			// top right to bottom left /
			for i := uint(0); i <= (t[0][0] - t[1][0]); i++ {
				x := t[0][0] - i
				y := t[0][1] + i
				m[x<<10|y]++
			}
		case t[0][0] > t[1][0] && t[0][1] > t[1][1]:
			// bottom right to top left \
			for i := uint(0); i <= (t[0][0] - t[1][0]); i++ {
				x := t[1][0] + i
				y := t[1][1] + i
				m[x<<10|y]++
			}
		default:
			log.Fatalf("well something went wrong, neither the first, nor the second coordinates were equal")
		}
	}

	return m
}

func mapLinesSlice(coords []uint) map[uint]uint {
	m := make(map[uint]uint)

	for h := 0; h < len(coords); h += 4 {
		switch {
		case coords[h] == coords[h+2] && coords[h+1] < coords[h+3]:
			// first coördinate is the same and first tuple is smaller
			x := coords[h] << bitSizeForCoordinates

			for i := coords[h+1]; i <= coords[h+3]; i++ {
				m[x|i]++
			}
		case coords[h] == coords[h+2] && coords[h+1] > coords[h+3]:
			// first coördinate is the same and second tuple is smaller
			x := coords[h] << bitSizeForCoordinates

			for i := coords[h+3]; i <= coords[h+1]; i++ {
				m[x|i]++
			}
		case coords[h] < coords[h+2] && coords[h+1] == coords[h+3]:
			// second coördinate is the same, first tuple is smaller
			y := coords[h+1]

			for i := coords[h]; i <= coords[h+2]; i++ {
				m[(i<<bitSizeForCoordinates)|y]++
			}
		case coords[h] > coords[h+2] && coords[h+1] == coords[h+3]:
			// second coördinate is the same, second tuple is smaller
			y := coords[h+1]

			for i := coords[h+2]; i <= coords[h]; i++ {
				m[(i<<bitSizeForCoordinates)|y]++
			}
		case coords[h] < coords[h+2] && coords[h+1] < coords[h+3]:
			// top left to bottom right \
			for i := uint(0); i <= (coords[h+2] - coords[h]); i++ {
				x := coords[h] + i
				y := coords[h+1] + i
				m[x<<10|y]++
			}
		case coords[h] < coords[h+2] && coords[h+1] > coords[h+3]:
			// bottom left to top right /
			for i := uint(0); i <= (coords[h+2] - coords[h]); i++ {
				x := coords[h] + i
				y := coords[h+1] - i
				m[x<<10|y]++
			}
		case coords[h] > coords[h+2] && coords[h+1] < coords[h+3]:
			// top right to bottom left /
			for i := uint(0); i <= (coords[h] - coords[h+2]); i++ {
				x := coords[h] - i
				y := coords[h+1] + i
				m[x<<10|y]++
			}
		case coords[h] > coords[h+2] && coords[h+1] > coords[h+3]:
			// bottom right to top left \
			for i := uint(0); i <= (coords[h] - coords[h+2]); i++ {
				x := coords[h+2] + i
				y := coords[h+3] + i
				m[x<<10|y]++
			}
		default:
			log.Fatalf("well something went wrong, neither the first, nor the second coordinates were equal")
		}
	}

	return m
}

func mapLinesSliceConcurrent(coords []uint) map[uint]uint {
	m := make(map[uint]uint)
	mut := &sync.RWMutex{}
	wg := sync.WaitGroup{}

	for q := 0; q < len(coords); q += 4 {
		wg.Add(1)

		go func(h int) {
			defer wg.Done()

			switch {
			case coords[h] == coords[h+2] && coords[h+1] < coords[h+3]:
				// first coördinate is the same and first tuple is smaller
				x := coords[h] << bitSizeForCoordinates

				for i := coords[h+1]; i <= coords[h+3]; i++ {
					incSyncMap(mut, &m, x|i)
				}
			case coords[h] == coords[h+2] && coords[h+1] > coords[h+3]:
				// first coördinate is the same and second tuple is smaller
				x := coords[h] << bitSizeForCoordinates

				for i := coords[h+3]; i <= coords[h+1]; i++ {
					incSyncMap(mut, &m, x|i)
				}
			case coords[h] < coords[h+2] && coords[h+1] == coords[h+3]:
				// second coördinate is the same, first tuple is smaller
				y := coords[h+1]

				for i := coords[h]; i <= coords[h+2]; i++ {
					incSyncMap(mut, &m, (i<<bitSizeForCoordinates)|y)
				}
			case coords[h] > coords[h+2] && coords[h+1] == coords[h+3]:
				// second coördinate is the same, second tuple is smaller
				y := coords[h+1]

				for i := coords[h+2]; i <= coords[h]; i++ {
					incSyncMap(mut, &m, (i<<bitSizeForCoordinates)|y)
				}
			case coords[h] < coords[h+2] && coords[h+1] < coords[h+3]:
				// top left to bottom right \
				for i := uint(0); i <= (coords[h+2] - coords[h]); i++ {
					x := coords[h] + i
					y := coords[h+1] + i

					incSyncMap(mut, &m, x<<10|y)
				}
			case coords[h] < coords[h+2] && coords[h+1] > coords[h+3]:
				// bottom left to top right /
				for i := uint(0); i <= (coords[h+2] - coords[h]); i++ {
					x := coords[h] + i
					y := coords[h+1] - i

					incSyncMap(mut, &m, x<<10|y)
				}
			case coords[h] > coords[h+2] && coords[h+1] < coords[h+3]:
				// top right to bottom left /
				for i := uint(0); i <= (coords[h] - coords[h+2]); i++ {
					x := coords[h] - i
					y := coords[h+1] + i

					incSyncMap(mut, &m, x<<10|y)
				}
			case coords[h] > coords[h+2] && coords[h+1] > coords[h+3]:
				// bottom right to top left \
				for i := uint(0); i <= (coords[h] - coords[h+2]); i++ {
					x := coords[h+2] + i
					y := coords[h+3] + i

					incSyncMap(mut, &m, x<<10|y)
				}
			default:
				log.Fatalf("well something went wrong, neither the first, nor the second coordinates were equal")
			}
		}(q)
	}

	wg.Wait()

	return m
}

func incSyncMap(mut *sync.RWMutex, m *map[uint]uint, key uint) {
	mut.Lock()
	(*m)[key]++
	mut.Unlock()
}
