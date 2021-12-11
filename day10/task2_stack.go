package day10

import (
	"log"
	"sort"
)

func task2Stack(input []string) int {
	scores := make([]int, 0)
	stack := make([]int32, 0, 40)

LineLoop:
	for _, line := range input {
		stack = stack[:0]

		for _, ch := range line {
			switch ch {
			// opening characters
			case 0x28: // (
				stack = append(stack, 0x29)
			case 0x5b: // [
				stack = append(stack, 0x5d)
			case 0x7b: // {
				stack = append(stack, 0x7d)
			case 0x3c: // <
				stack = append(stack, 0x3e)
			default:
				if stack[len(stack)-1] != ch {
					continue LineLoop
				}

				stack = stack[:len(stack)-1]
			}
		}

		l := len(stack) - 1
		acc := 0

		for i := 0; i <= l; i++ {
			acc *= 5

			switch stack[l-i] {
			case 0x29: // (
				acc++
			case 0x5d: // [
				acc += 2
			case 0x7d: // {
				acc += 3
			case 0x3e: // <
				acc += 4
			default:
				log.Fatalf("this should not have happened. Encountered a non opening character: [%s]", string(stack[l-i]))
			}
		}

		scores = append(scores, acc)
	}

	sort.Ints(scores)

	return scores[(len(scores) / 2)]
}
