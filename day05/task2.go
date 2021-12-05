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

func task2SlicyReversed(input []string) int {
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

func task2SlicyRegex(input []string) int {
	coords := getCoordinateSliceRegex(input)
	lines := mapLinesSlice(coords)
	twos := 0

	for _, v := range lines {
		if v > 1 {
			twos++
		}
	}

	return twos
}

func task2SlicyStrings(input []string) int {
	coords := getCoordinateSliceStrings(input)
	lines := mapLinesSlice(coords)
	twos := 0

	for _, v := range lines {
		if v > 1 {
			twos++
		}
	}

	return twos
}
