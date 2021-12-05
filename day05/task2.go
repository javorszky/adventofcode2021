package day05

func task2(input []string) int {
	tuples := getTuples(input)
	lines := mapLines(tuples)
	twos := 0

	for _, v := range lines {
		if v > 1 {
			twos++
		}
	}

	return twos
}
