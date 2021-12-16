package day15

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_tasks(t *testing.T) {
	type args struct {
		input []string
	}

	tests := []struct {
		name string
		args args
		f    func([]string) int
		want int
	}{
		{
			name: "solves example input for task 1",
			args: args{input: getInputs("example_input.txt")},
			f:    task1,
			want: 40,
		},
		{
			name: "solves example input for task 1 map",
			args: args{input: getInputs("example_input.txt")},
			f:    task1Map,
			want: 40,
		},
		{
			name: "solves actual input for task 1",
			args: args{input: getInputs("input.txt")},
			f:    task1,
			want: 373,
		},
		{
			name: "solves actual input for task 1 map",
			args: args{input: getInputs("input.txt")},
			f:    task1Map,
			want: 373,
		},
		{
			name: "solves example input for task 2",
			args: args{input: getInputs("example_input.txt")},
			f:    task2,
			want: 315,
		},
		//{
		//	name: "solves actual input for task 1",
		//	args: args{input: getInputs("input.txt")},
		//	want: 373,
		//},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.f(tt.args.input))
		})
	}
}
