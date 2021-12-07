package day06

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_tick(t *testing.T) {
	tests := []struct {
		name string
		in   []int
		want []int
	}{
		{
			name: "example step 1",
			in:   []int{3, 4, 3, 1, 2},
			want: []int{2, 3, 2, 0, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, tick(tt.in), "tick(%v)", tt.in)
		})
	}
}
