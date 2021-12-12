package day11

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_parseIntoGrid(t *testing.T) {
	tests := []struct {
		name  string
		input []string
		want  map[uint]uint
	}{
		{
			name: "parses horizontally",
			input: []string{
				"9123443232",
			},
			want: map[uint]uint{
				0b00000000: 9,
				0b00000001: 1,
				0b00000010: 2,
				0b00000011: 3,
				0b00000100: 4,
				0b00000101: 4,
				0b00000110: 3,
				0b00000111: 2,
				0b00001000: 3,
				0b00001001: 2,
			},
		},
		{
			name: "parses vertically",
			input: []string{
				"9",
				"1",
				"2",
				"3",
				"4",
				"4",
				"3",
				"2",
				"3",
				"2",
			},
			want: map[uint]uint{
				0b00000000: 9,
				0b00010000: 1,
				0b00100000: 2,
				0b00110000: 3,
				0b01000000: 4,
				0b01010000: 4,
				0b01100000: 3,
				0b01110000: 2,
				0b10000000: 3,
				0b10010000: 2,
			},
		},
		{
			name: "parses a square",
			input: []string{
				"912",
				"344",
				"321",
			},
			want: map[uint]uint{
				0b00000000: 9,
				0b00000001: 1,
				0b00000010: 2,
				0b00010000: 3,
				0b00010001: 4,
				0b00010010: 4,
				0b00100000: 3,
				0b00100001: 2,
				0b00100010: 1,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, parseIntoGrid(tt.input))
		})
	}
}
