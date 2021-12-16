package day15

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_task1(t *testing.T) {
	type args struct {
		input []string
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "solves example input",
			args: args{input: getInputs("example_input.txt")},
			want: 40,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, task1(tt.args.input), "task1(%v)", tt.args.input)
		})
	}
}
