package day05

import "testing"

func Benchmark_Tasks(b *testing.B) {
	benchmarks := []struct {
		name string
		fn   func([]string) int
	}{
		{
			name: "task 1 tuple regex",
			fn:   task1,
		},
		{
			name: "task 1 tuple string split",
			fn:   task1TuplesStrings,
		},
		{
			name: "task 1 tuple reverse",
			fn:   task1TuplesReverse,
		},
		{
			name: "task 1 slicy regex",
			fn:   task1SlicyRegex,
		},
		{
			name: "task 1 slicy string split",
			fn:   task1SlicyStrings,
		},
		{
			name: "task 1 slicy reverse",
			fn:   task1SlicyReverse,
		},
		{
			name: "task 1 concurrent slicy reverse",
			fn:   task1SlicyReverseConcurrent,
		},
		{
			name: "task 2 tuple regex",
			fn:   task2,
		},
		{
			name: "task 2 tuple string split",
			fn:   task2TupleStrings,
		},
		{
			name: "task 2 tuple reverse",
			fn:   task2TupleReverse,
		},
		{
			name: "task 2 slicy regex",
			fn:   task2SlicyRegex,
		},
		{
			name: "task 2 slicy string split",
			fn:   task2SlicyStrings,
		},
		{
			name: "task 2 slicy reverse",
			fn:   task2SlicyReversed,
		},
	}

	inputs := benchInput(b, "input.txt")

	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = bm.fn(inputs)
			}
		})
	}
}

func Benchmark_GetTuples(b *testing.B) {
	benchmarks := []struct {
		name string
		fn   func([]string) []tuple
	}{
		{
			name: "getTuples regex",
			fn:   getTuples,
		},
		{
			name: "getTuples string split",
			fn:   getTuplesString,
		},
		{
			name: "getTuples reverse",
			fn:   getTuplesReversed,
		},
	}

	inputs := benchInput(b, "input.txt")

	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = bm.fn(inputs)
			}
		})
	}
}

func Benchmark_GetCoords(b *testing.B) {
	benchmarks := []struct {
		name string
		fn   func([]string) []uint
	}{
		{
			name: "getCoords regex",
			fn:   getCoordinateSliceRegex,
		},
		{
			name: "getCoords string split",
			fn:   getCoordinateSliceStrings,
		},
		{
			name: "getCoords reverse",
			fn:   getCoordinateSliceReverse,
		},
	}

	inputs := benchInput(b, "input.txt")

	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = bm.fn(inputs)
			}
		})
	}
}

func Benchmark_MapLinesTuples(b *testing.B) {
	benchmarks := []struct {
		name string
		fn   func([]tuple) map[uint]uint
	}{
		{
			name: "mapLinesTuples",
			fn:   mapLinesTuples,
		},
	}

	ts := benchTuples(b, benchInput(b, "input.txt"))

	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = bm.fn(ts)
			}
		})
	}
}

func Benchmark_MapLinesCoords(b *testing.B) {
	benchmarks := []struct {
		name string
		fn   func([]uint) map[uint]uint
	}{
		{
			name: "mapLinesCoords",
			fn:   mapLinesSlice,
		},
		{
			name: "mapLinesCoordsConcurrent",
			fn:   mapLinesSliceConcurrent,
		},
	}

	coords := benchCoords(b, benchInput(b, "input.txt"))

	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				bm.fn(coords)
			}
		})
	}
}

func benchInput(b testing.TB, filename string) []string {
	b.Helper()

	return getInputs(filename)
}

func benchTuples(b testing.TB, input []string) []tuple {
	b.Helper()

	return getTuples(input)
}

func benchCoords(b testing.TB, input []string) []uint {
	b.Helper()

	return getCoordinateSliceRegex(input)
}
