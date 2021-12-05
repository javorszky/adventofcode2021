package day05

import "testing"

func Benchmark_Tasks(b *testing.B) {
	benchmarks := []struct {
		name string
		fn   func([]string) int
	}{
		{
			name: "task 1 using full input",
			fn:   task1,
		},
		{
			name: "task 2 using full input",
			fn:   task2,
		},
	}
	for _, bm := range benchmarks {
		inputs := benchInput(b, "input.txt")
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
			name: "getTuples strings.Split",
			fn:   getTuplesString,
		},
		{
			name: "gettuples reversed",
			fn:   getTuplesReversed,
		},
	}
	for _, bm := range benchmarks {
		inputs := benchInput(b, "input.txt")
		b.Run(bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = bm.fn(inputs)
			}
		})
	}
}

func benchInput(b testing.TB, filename string) []string {
	b.Helper()

	return getInputs(filename)
}
