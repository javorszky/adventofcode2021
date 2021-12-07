package day06

var next = map[int][]int{
	8: {7},
	7: {6},
	6: {5},
	5: {4},
	4: {3},
	3: {2},
	2: {1},
	1: {0},
	0: {6, 8},
}

func tick(in []int) []int {
	out := make([]int, 0)

	for _, i := range in {
		out = append(out, next[i]...)
	}

	return out
}

func task1_tick(input string) int {
	return calcState(parseFishSplitAtoi(input), targetDay)
}

func calcState(thing []int, ticks int) int {
	for i := 0; i < ticks; i++ {
		thing = tick(thing)
	}

	return len(thing)
}
