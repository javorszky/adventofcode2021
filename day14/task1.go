package day14

import (
	"strings"
)

func task1(template string, rules []string) int {
	betterRules := parseBetterRules(rules)
	polymer := parsePolymer(template)

	for j := 0; j < 10; j++ {
		polymer = work(polymer, betterRules)
	}

	counts := map[uint]int{}
	for _, ch := range polymer {
		counts[ch]++
	}

	least := len(polymer)
	most := 0

	for _, v := range counts {
		if v > most {
			most = v
		}

		if v < least {
			least = v
		}
	}

	return most - least
}

func work(polymer []uint, betterRules map[uint]uint) []uint {
	grown := make([]uint, len(polymer)*2-1)
	grown[0] = polymer[0]

	for k := 0; k < len(polymer)-1; k++ {
		grown[2*k+1] = betterRules[polymer[k]<<8|polymer[k+1]]
		grown[2*k+2] = polymer[k+1]
	}

	return grown
}

func parsePolymer(template string) []uint {
	polymer := make([]uint, len(template))

	for i, ch := range template {
		polymer[i] = uint(ch)
	}

	return polymer
}

func parseBetterRules(rules []string) map[uint]uint {
	br := make(map[uint]uint)

	for _, line := range rules {
		parts := strings.Split(line, " -> ")
		br[uint(parts[0][0])<<8|uint(parts[0][1])] = uint(parts[1][0])
	}

	return br
}
