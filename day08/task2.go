package day08

import (
	"log"
	"math"
	"strings"
)

func task2(input []string) int {
	for _, i := range input {
		deduce(i)
	}

	return 0
}

func deduce(input string) int {
	set := newDisplay()
	parts := strings.Split(input, " | ")
	leftNums := strings.Split(parts[0], " ")
	rightNums := strings.Split(parts[0], " ")
	allNums := append(leftNums, rightNums...)

	set.parse(allNums)

	if !set.IsSolved() {
		log.Fatalf("could not solve set for line\n[ %s ]\n"+
			"current state:\n"+
			"%#v\n", input, set.State())
	}

	a := 0
	l := len(rightNums)

	for i, rightnum := range rightNums {
		a += set.GetNumber(rightnum) * int(math.Pow10(l-i))
	}

	return a
}
