package day03

import (
	"log"
)

const (
	mask     uint = 0b111111111111
	zeroByte rune = 0x30
	oneByte       = 0x31
)

func task1(input []string) uint {
	width := len(input[0])
	vertical := make([]int, width)

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
