package day10

import (
	"regexp"
	"sort"
)

var reSyntaxError = regexp.MustCompile(`\)|]|}|>`)

func task2(input []string) int {
	r := getChunkReplacer()
	scores := make([]int, 0)

	for _, line := range input {
		for {
			l := len(line)
			line = r.Replace(line)

			if l == len(line) {
				break
			}
		}

		if line == `` || hasSyntaxError(line) {
			continue
		}

		scores = append(scores, gatherClosingScore(line))
	}

	sort.Ints(scores)

	return scores[len(scores)/2]
}

func hasSyntaxError(s string) bool {
	return reSyntaxError.MatchString(s)
}

// ): 1 point.
// ]: 2 points.
// }: 3 points.
// >: 4 points.
func gatherClosingScore(openings string) int {
	acc := 0

	for i := len(openings) - 1; i >= 0; i-- {
		acc = acc * 5
		acc += closingScores[openings[i]]
	}

	return acc
}

var closingScores = map[uint8]int{
	0x28: 1, // (
	0x5b: 2, // [
	0x7b: 3, // {
	0x3c: 4, // <
}
