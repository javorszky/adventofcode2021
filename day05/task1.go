package day05

const bitSizeForCoordinates = 10

func task1(input []string) int {
	tuples := getTuples(input)
	euclideanTuples := filterToEuclideanLines(tuples)
	lines := mapLinesTuples(euclideanTuples)
	twos := 0

	for _, v := range lines {
		if v > 1 {
			twos++
		}
	}

	return twos
}

func task1TuplesReverse(input []string) int {
	tuples := getTuplesReversed(input)
	euclideanTuples := filterToEuclideanLines(tuples)
	lines := mapLinesTuples(euclideanTuples)
	twos := 0

	for _, v := range lines {
		if v > 1 {
			twos++
		}
	}

	return twos
}

func task1TuplesStrings(input []string) int {
	tuples := getTuplesString(input)
	euclideanTuples := filterToEuclideanLines(tuples)
	lines := mapLinesTuples(euclideanTuples)
	twos := 0

	for _, v := range lines {
		if v > 1 {
			twos++
		}
	}

	return twos
}

func task1SlicyReverse(input []string) int {
	coords := getCoordinateSliceReverse(input)
	euclideanCoords := filterToEuclideanLinesSlice(coords)
	lines := mapLinesSlice(euclideanCoords)
	twos := 0

	for _, v := range lines {
		if v > 1 {
			twos++
		}
	}

	return twos
}

func task1SlicyRegex(input []string) int {
	coords := getCoordinateSliceRegex(input)
	euclideanCoords := filterToEuclideanLinesSlice(coords)
	lines := mapLinesSlice(euclideanCoords)
	twos := 0

	for _, v := range lines {
		if v > 1 {
			twos++
		}
	}

	return twos
}

func task1SlicyStrings(input []string) int {
	coords := getCoordinateSliceStrings(input)
	euclideanCoords := filterToEuclideanLinesSlice(coords)
	lines := mapLinesSlice(euclideanCoords)
	twos := 0

	for _, v := range lines {
		if v > 1 {
			twos++
		}
	}

	return twos
}

func filterToEuclideanLines(tuples []tuple) []tuple {
	out := make([]tuple, 0)

	for _, t := range tuples {
		if t[0][0] != t[1][0] && t[0][1] != t[1][1] {
			continue
		}

		out = append(out, t)
	}

	return out
}

func filterToEuclideanLinesSlice(coords []uint) []uint {
	out := make([]uint, 0)

	for i := 0; i < len(coords); i += 4 {
		if coords[i] != coords[i+2] && coords[i+1] != coords[i+3] {
			continue
		}

		out = append(out, coords[i:i+4]...)
	}

	return out
}
