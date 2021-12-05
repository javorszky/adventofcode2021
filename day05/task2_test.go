package day05

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_task2(t *testing.T) {
	input := []string{
		"0,9 -> 5,9",
		"8,0 -> 0,8",
		"9,4 -> 3,4",
		"2,2 -> 2,1",
		"7,0 -> 7,4",
		"6,4 -> 2,0",
		"0,9 -> 2,9",
		"3,4 -> 1,4",
		"0,0 -> 8,8",
		"5,5 -> 8,2",
	}

	tests := []struct {
		name  string
		input []string
		fn    func([]string) int
		want  int
	}{
		{
			name:  "task 2 tuple example input",
			input: input,
			fn:    task2,
			want:  12,
		},
		{
			name:  "task 2 tuple strings example input",
			input: input,
			fn:    task2TupleStrings,
			want:  12,
		},
		{
			name:  "task 2 tuple reverse example input",
			input: input,
			fn:    task2TupleReverse,
			want:  12,
		},
		{
			name:  "task 2 slicy example input",
			input: input,
			fn:    task2Slicy,
			want:  12,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, task2(tt.input), "task2(%v)", tt.input)
		})
	}
}
