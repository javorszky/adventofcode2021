package day10

import "log"

var pairs = map[uint8]uint8{
	0x28: 0x29, // (:)
	0x5b: 0x5d, // [:]
	0x7b: 0x7d, // {:}
	0x3c: 0x3e, // <:>
}

func task1Stack(input []string) int {
	acc := 0

	for _, line := range input {
		stack := make([]uint8, 0)

		for i := 0; i < len(line); i++ {
			closing, ok := pairs[line[i]]
			if ok {
				stack = append(stack, closing)

				continue
			}

			if len(stack) == 0 || stack[len(stack)-1] != line[i] {
				acc += invalidScore[line[i]]

				break
			}

			if stack[len(stack)-1] == line[i] {
				stack = stack[:len(stack)-1]

				continue
			}

			log.Fatalf("entirely unknown character [%s]. This should not have happened", string(line[i]))
		}
	}

	return acc
}
