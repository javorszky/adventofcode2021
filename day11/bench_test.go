package day11

import "testing"

func Benchmark_Tasks(b *testing.B) {
	benchmarks := []struct {
		name string
		fn   func([]string) int
	}{
		{
			name: "day 11 task 1",
			fn:   task1,
		},
		{
			name: "day 11 task 2",
			fn:   task2,
		},
	}

	input := benchInput(b, "input.txt")

	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				bm.fn(input)
			}
		})
	}
}

func benchInput(b testing.TB, fn string) []string {
	b.Helper()

	return getInputs(fn)
}
