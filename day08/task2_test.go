package day08

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_deduce(t *testing.T) {
	tests := []struct {
		name string
		in   string
		want int
	}{
		{
			name: "deduces the example",
			in:   "acedgfb cdfbe gcdfa fbcad dab cefabd cdfgeb eafb cagedb ab | cdfeb fcadb cdfeb cdbaf",
			want: 5353,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, deduce(tt.in), "deduce(%v)", tt.in)
		})
	}
}
