package day05

func task2(input []string) int {
	tuples := getTuples(input)
	lines := mapLinesTuples(tuples)
	twos := 0

	for _, v := range lines {
		if v > 1 {
			twos++
		}
	}

	return twos
}

func task2TupleStrings(input []string) int {
	tuples := getTuplesString(input)
	lines := mapLinesTuples(tuples)
	twos := 0

	for _, v := range lines {
		if v > 1 {
			twos++
		}
	}

	return twos
}

func task2TupleReverse(input []string) int {
	tuples := getTuplesReversed(input)
	lines := mapLinesTuples(tuples)
	twos := 0

	for _, v := range lines {
		if v > 1 {
			twos++
		}
	}

	return twos
}

func task2Slicy(input []string) int {
	coords := getCoordinateSliceReverse(input)
	lines := mapLinesSlice(coords)
	twos := 0

	for _, v := range lines {
		if v > 1 {
			twos++
		}
	}

	return twos
}
