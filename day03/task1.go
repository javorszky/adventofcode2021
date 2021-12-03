package day03

import (
	"log"
)

const (
	zeroByte rune = 0x30
	oneByte  rune = 0x31
)

func task1(input []uint, width int) uint {
	var gamma uint

	local := make([]uint, len(input))
	copy(local, input)

	var mask uint = (1 << width) - 1

	for i := width - 1; i >= 0; i-- {
		ones := 0
		zeroes := 0

		var currentCheck uint = 1 << i

		for j, v := range local {
			if v < currentCheck {
				zeroes++
			} else {
				ones++
				local[j] -= 1 << i
			}
		}

		if ones > zeroes {
			gamma += 1 << i
		}
	}

	return gamma * (^gamma - ^mask)
}

func task1Strings(input []string) uint {
	width := len(input[0])
	vertical := make([]int, width)

	var mask uint = (1 << width) - 1

	for _, line := range input {
		for j, char := range line {
			switch char {
			case zeroByte:
				vertical[j]--
			case oneByte:
				vertical[j]++
			default:
				log.Fatalf("we've encountered a char [%v] that should not have happened in line %s", char, line)
			}
		}
	}

	var gamma uint

	for k, val := range vertical {
		if val < 0 {
			continue
		}

		gamma += 1 << (width - 1 - k)
	}

	return gamma * (^gamma - ^mask)
}
