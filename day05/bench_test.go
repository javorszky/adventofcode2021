package day05

import "testing"

func Benchmark(b *testing.B) {
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

func benchInput(b testing.TB, filename string) []string {
	b.Helper()

	return getInputs(filename)
}
