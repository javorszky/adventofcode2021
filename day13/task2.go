package day13

import (
	"log"
	"math"
	"strings"

	"github.com/javorszky/adventofcode2021/util"
)

func task2(dots, folds []string) string {
	paper := makePaper(dots)
	n := uint(0)
	nParts := make([]uint, 0)

	for _, f := range folds {
		parts := strings.Split(f, "=")
		nParts = nParts[:0]

		for _, char := range parts[1] {
			nParts = append(nParts, charToInt[char])
		}

		nLen := len(nParts)
		n = 0

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
	}

	return visualize(paper)
}

func visualize(paper map[uint]uint) string {
	largestRow := uint(0)
	largestCol := uint(0)

	for k := range paper {
		col := k >> 11
		row := k &^ (col << 11)

		if col > largestCol {
			largestCol = col
		}

		if row > largestRow {
			largestRow = row
		}
	}

	var sb strings.Builder

	for r := uint(0); r <= largestRow; r++ {
		for c := uint(0); c <= largestCol; c++ {
			if _, ok := paper[c<<11|r]; ok {
				sb.WriteString("#")

				continue
			}

			sb.WriteString(" ")
		}

		sb.WriteString(util.NewLine)
	}

	return sb.String()
}
