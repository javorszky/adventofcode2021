package day08

import (
	"log"
	"math"
	"strings"

	"github.com/javorszky/adventofcode2021/util"
)

func task2(input []string) int {
	acc := 0

	for _, i := range input {
		acc += deduceMatch(i)
	}

	return acc
}

func deduce(input string) int {
	set := newDisplay()
	parts := strings.Split(input, " | ")
	leftNums := strings.Split(parts[0], " ")
	rightNums := strings.Split(parts[0], " ")
	allNums := append(leftNums, rightNums...)

	set.parse(allNums)

	if !set.IsSolved() {
		log.Fatalf("could not solve set for line%s[ %s ]%s"+
			"current state:%s"+
			"%#v%s",
			util.NewLine, input, util.NewLine,
			util.NewLine,
			set.State(), util.NewLine)
	}

	a := 0
	l := len(rightNums)

	for i, rightnum := range rightNums {
		a += set.GetNumber(rightnum) * int(math.Pow10(l-i))
	}

	return a
}
