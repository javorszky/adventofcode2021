package day13

import (
	"log"
	"math"
	"strings"

	"github.com/javorszky/adventofcode2021/util"
)

func task1(dots, folds []string) int {
	paper := makePaper(dots)
	n := uint(0)
	nParts := make([]uint, 0)

	for _, f := range folds {
		parts := strings.Split(f, "=")
		for _, char := range parts[1] {
			nParts = append(nParts, charToInt[char])
		}

		nLen := len(nParts)

		for ri, rn := range nParts {
			n += rn * uint(math.Pow10(nLen-ri-1))
		}

		switch parts[0] {
		case "fold along x":
			paper = foldLeft(paper, n)
		case "fold along y":
			paper = foldUp(paper, n)
		default:
			log.Fatalf("something went wrong, bad part: %s%s", parts[0], util.NewLine)
		}

		return len(paper)
	}

	log.Fatalf("there were no fold instructions, this should not have happened")

	return 0
}
