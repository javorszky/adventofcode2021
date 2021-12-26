package day21

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

func task2(p1, p2 int) int {
	return simulateMultiverses(p1, p2)
}

func calculatePossibilities() map[int]int {
	poss := make(map[int]int)

	for i := 1; i <= 3; i++ {
		for j := 1; j <= 3; j++ {
			for k := 1; k <= 3; k++ {
				poss[i+j+k]++
			}
		}
	}

	return poss
}

type Universe struct {
	p1Step  int
	p1Score int
	p2Step  int
	p2Score int
}

func (u Universe) String() string {
	return fmt.Sprintf("%d/%d/%d/%d", u.p1Step, u.p1Score, u.p2Step, u.p2Score)
}

func (u *Universe) Marshal(s string) {
	parts := strings.Split(s, "/")
	if len(parts) != 4 {
		log.Fatalf("can not marshal universe from string [%s]", s)
	}

	numbers := make([]int, len(parts))

	for i, subString := range parts {
		n, err := strconv.Atoi(subString)
		if err != nil {
			log.Fatalf("can not marshal universe from string: strconv failed for [%s]: %s", subString, err)
		}

		numbers[i] = n
	}

	u.p1Step = numbers[0]
	u.p1Score = numbers[1]
	u.p2Step = numbers[2]
	u.p2Score = numbers[3]
}

func simulateMultiverses(p1, p2 int) int {
	rollPossibilities := calculatePossibilities()

	return len(rollPossibilities)
}
