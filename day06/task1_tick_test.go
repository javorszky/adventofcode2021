package day06

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_tick(t *testing.T) {
	tests := []struct {
		name string
		in   []int
		tick int
		want []int
	}{
		{
			name: "example step 1",
			in:   []int{3, 4, 3, 1, 2},
			tick: 1,
			want: []int{2, 3, 2, 0, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			out := make([]int, len(tt.in))
			copy(out, tt.in)

			for i := 0; i < tt.tick; i++ {
				out = tick(out)
			}

			assert.Equalf(t, tt.want, out, "tick(%v)", tt.in)
		})
	}
}
