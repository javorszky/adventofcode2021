package day06

import (
	"fmt"
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
			var out []int
			copy(out, tt.in)

			for i := 0; i < tt.tick; i++ {
				out = tick(out)
				fmt.Printf("ran tick: %v\n", out)
			}

			assert.Equalf(t, tt.want, out, "tick(%v)", tt.in)
		})
	}
}

func Test_tick8(t *testing.T) {
	tests := []struct {
		name string
		in   []int
		want []int
	}{
		{
			name: "tick8 is the same as 8x tick",
			in:   []int{2, 3, 2, 0, 1},
			want: []int{2, 3, 2, 0, 1, 2, 3, 4, 4, 5},
			//want: []int{1, 2, 1, 6, 0, 1, 2, 3, 3, 4, 8},

		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var out []int
			copy(out, tt.in)

			assert.Equalf(t, tt.want, tick8(tt.in), "tick8(%v)", tt.in)

			for i := 0; i < 8; i++ {
				out = tick(out)
			}

			assert.Equal(t, tt.want, out)
		})
	}
}
