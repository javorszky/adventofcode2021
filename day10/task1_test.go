package day10

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_task1(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  int
	}{
		{
			name:  "solves example",
			input: "example_input.txt",
			want:  26397,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			input := getInputs(tt.input)
			assert.Equal(t, tt.want, task1(input))
		})
	}
}
