package day10

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_task1(t *testing.T) {
	tests := []struct {
		name  string
		input string
		fn    func([]string) int
		want  int
	}{
		{
			name:  "solves example",
			input: "example_input.txt",
			fn:    task1,
			want:  26397,
		},
		{
			name:  "solves example",
			input: "example_input.txt",
			fn:    task1Stack,
			want:  26397,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			input := getInputs(tt.input)
			assert.Equal(t, tt.want, tt.fn(input))
		})
	}
}

func Benchmark(b *testing.B) {
	benchmarks := []struct {
		name     string
		filename string
		fn       func([]string) int
	}{
		{
			name:     "task 1 example",
			filename: "example_input.txt",
			fn:       task1,
		},
		{
			name:     "task 1 stack example",
			filename: "example_input.txt",
			fn:       task1Stack,
		},
		{
			name:     "task 1 nekkid stack example",
			filename: "example_input.txt",
			fn:       task1NakedStack,
		},
		{
			name:     "task 2 example",
			filename: "example_input.txt",
			fn:       task2,
		},
		{
			name:     "task 1 full input",
			filename: "input.txt",
			fn:       task1,
		},
		{
			name:     "task 1 stack full input",
			filename: "input.txt",
			fn:       task1Stack,
		},
		{
			name:     "task 1 nekkid stack full input",
			filename: "input.txt",
			fn:       task1NakedStack,
		},
		{
			name:     "task 2 full input",
			filename: "input.txt",
			fn:       task2,
		},
	}
	for _, bm := range benchmarks {
		input := benchInput(b, bm.filename)
		b.Run(bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				bm.fn(input)
			}
		})
	}
}

func benchInput(b testing.TB, filename string) []string {
	b.Helper()

	return getInputs(filename)
}
