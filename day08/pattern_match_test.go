package day08

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_parseString(t *testing.T) {
	tests := []struct {
		name string
		in   string
		want uint
	}{
		{
			name: "parses ef",
			in:   "ef",
			want: 0b0110000,
		},
		{
			name: "parses fe",
			in:   "fe",
			want: 0b0110000,
		},
		{
			name: "parses abc",
			in:   "abc",
			want: 0b0000111,
		},
		{
			name: "parses cgdb",
			in:   "cgdb",
			want: 0b1001110,
		},
		{
			name: "parses efb",
			in:   "efb",
			want: 0b0110010,
		},
		{
			name: "parses cefabd",
			in:   "cefabd",
			want: 0b0111111,
		},
		{
			name: "parses cdfgeb",
			in:   "cdfgeb",
			want: 0b1111110,
		},
		{
			name: "parses cagedb",
			in:   "cagedb",
			want: 0b1011111,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, parseNumberString(tt.in), "parseString(%v)", tt.in)
		})
	}
}

func Test_deduceMatch(t *testing.T) {
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
			assert.Equalf(t, tt.want, deduceMatch(tt.in), "deduceMatch(%v)", tt.in)
		})
	}
}
