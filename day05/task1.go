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
