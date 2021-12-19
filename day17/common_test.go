package day17

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_getInputs(t *testing.T) {
	type args struct {
		fn string
	}

	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "reads input correctly",
			args: args{fn: "input.txt"},
			want: []int{135, 155, -102, -78},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.ElementsMatchf(t, tt.want, getInputs(tt.args.fn), "getInputs(%s)", tt.args.fn)
		})
	}
}
