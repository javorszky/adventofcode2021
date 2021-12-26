package day20

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_task1(t *testing.T) {
	tests := []struct {
		name  string
		input func() (string, string)
		want  int
	}{
		{
			name: "solves example input",
			input: func() (string, string) {
				return getInputs("example_input.txt")
			},
			want: 35,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mask, img := tt.input()

			assert.Equalf(t, tt.want, task1(mask, img), "task1(%v, %v)", mask, img)
		})
	}
}
