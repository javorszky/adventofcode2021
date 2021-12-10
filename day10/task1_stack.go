package day10

var pairs = map[int32]int32{
	0x28: 0x29, // (:)
	0x5b: 0x5d, // [:]
	0x7b: 0x7d, // {:}
	0x3c: 0x3e, // <:>
}

// ): 3 points.
// ]: 57 points.
// }: 1197 points.
// >: 25137 points.
var invalidScoreStack = map[int32]int{
	0x29: 3,     // )
	0x5d: 57,    // ]
	0x7d: 1197,  // }
	0x3e: 25137, // >
}

func task1Stack(input []string) int {
	acc := 0
	stack := make([]int32, 0, 30)

	for _, line := range input {
		stack = stack[:0]

	Loop:
		for _, ch := range line {
			switch closing := pairs[ch]; closing {
			case 0:
				//if len(stack) == 0 || stack[len(stack)-1] != ch {
				if stack[len(stack)-1] != ch {
					acc += invalidScoreStack[ch]

					break Loop
				}

				stack = stack[:len(stack)-1]
			default:
				// we found the character, which means it was an opening character, so let's stick a closing one here.
				stack = append(stack, closing)
			}
		}
	}

	return acc
}
