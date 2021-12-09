package day08

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_display_parse(t *testing.T) {
	tests := []struct {
		name string
		in   []string
		want map[uint]uint
	}{
		{
			name: "parses a one",
			in:   []string{"gf"},
			want: map[uint]uint{
				segmentA: 0b0011111,
				segmentB: 0b0011111,
				segmentC: 0b1100000,
				segmentD: 0b0011111,
				segmentE: 0b0011111,
				segmentF: 0b1100000,
				segmentG: 0b0011111,
			},
		},
		{
			name: "parses a one and a seven",
			in:   []string{"gf", "agf"},
			want: map[uint]uint{
				segmentA: 0b0000001,
				segmentB: 0b0011110,
				segmentC: 0b1100000,
				segmentD: 0b0011110,
				segmentE: 0b0011110,
				segmentF: 0b1100000,
				segmentG: 0b0011110,
			},
		},
		{
			name: "parses a one and a seven and a four",
			in:   []string{"gf", "agf", "cfge"},
			want: map[uint]uint{
				segmentA: 0b0000001, // top
				segmentB: 0b0010100, // top left
				segmentC: 0b1100000, // top right
				segmentD: 0b0010100, // belt
				segmentE: 0b0001010, // bottom left
				segmentF: 0b1100000, // bottom right
				segmentG: 0b0001010, // bottom
			},
		},
		{
			name: "parses a six",
			in:   []string{"abcdef"},
			want: map[uint]uint{
				segmentA: 0b0111111, // top
				segmentB: 0b0111111, // top left
				segmentC: 0b1111111, // top right
				segmentD: 0b1111111, // belt
				segmentE: 0b1111111, // bottom left
				segmentF: 0b0111111, // bottom right
				segmentG: 0b0111111, // bottom
			},
		},
		{
			name: "parses a five",
			in:   []string{"abefg"}, // missing: c, d
			want: map[uint]uint{
				segmentA: 0b1110011, // top
				segmentB: 0b1111111, // top left
				segmentC: 0b1111111, // top right
				segmentD: 0b1110011, // belt
				segmentE: 0b1111111, // bottom left
				segmentF: 0b1111111, // bottom right
				segmentG: 0b1110011, // bottom
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := NewDisplay()
			d.parse(tt.in)
			assert.Equalf(t, tt.want, d.State(), "parse(%v)", tt.in)
		})
	}
}
