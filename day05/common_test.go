package day05

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_mapLines(t *testing.T) {
	tests := []struct {
		name   string
		tuples []tuple
		want   map[uint]uint
	}{
		{
			name: "extrapolates one vertical tuple to lines, first bigger",
			tuples: []tuple{
				{
					{1, 2},
					{1, 4},
				},
			},
			want: map[uint]uint{
				0b00000000010000000010: 1,
				0b00000000010000000011: 1,
				0b00000000010000000100: 1,
			},
		},
		{
			name: "extrapolates one vertical tuple to lines, second bigger",
			tuples: []tuple{
				{
					{1, 4},
					{1, 2},
				},
			},
			want: map[uint]uint{
				0b00000000010000000010: 1,
				0b00000000010000000011: 1,
				0b00000000010000000100: 1,
			},
		},
		{
			name: "extrapolates one horizontal tuple to lines, first bigger",
			tuples: []tuple{
				{
					{1, 6},
					{5, 6},
				},
			},
			want: map[uint]uint{
				0b00000000010000000110: 1,
				0b00000000100000000110: 1,
				0b00000000110000000110: 1,
				0b00000001000000000110: 1,
				0b00000001010000000110: 1,
			},
		},
		{
			name: "extrapolates one horizontal tuple to lines, second bigger",
			tuples: []tuple{
				{
					{5, 6},
					{1, 6},
				},
			},
			want: map[uint]uint{
				0b00000000010000000110: 1,
				0b00000000100000000110: 1,
				0b00000000110000000110: 1,
				0b00000001000000000110: 1,
				0b00000001010000000110: 1,
			},
		},
		{
			name: "extrapolates two intersecting tuples to lines",
			tuples: []tuple{
				{
					{5, 6},
					{1, 6},
				},
				{
					{3, 2},
					{3, 8},
				},
			},
			want: map[uint]uint{
				0b00000000110000000010: 1, // 3,2
				0b00000000110000000011: 1, // 3,3
				0b00000000110000000100: 1, // 3,4
				0b00000000110000000101: 1, // 3,5
				0b00000000110000000111: 1, // 3,7
				0b00000000110000001000: 1, // 3,8

				0b00000000010000000110: 1, // 1,6
				0b00000000100000000110: 1, // 2,6
				0b00000000110000000110: 2, // 3,6
				0b00000001000000000110: 1, // 4,6
				0b00000001010000000110: 1, // 5,6
			},
		},
		{
			name: "diagonal, top left to bottom right",
			tuples: []tuple{
				{
					{1, 4},
					{4, 7},
				},
			},
			want: map[uint]uint{
				0b00000000010000000100: 1, // 1,4
				0b00000000100000000101: 1, // 2,5
				0b00000000110000000110: 1, // 3,6
				0b00000001000000000111: 1, // 4,7
			},
		},
		{
			name: "diagonal, bottom right to top left",
			tuples: []tuple{
				{
					{4, 7},
					{1, 4},
				},
			},
			want: map[uint]uint{
				0b00000000010000000100: 1, // 1,4
				0b00000000100000000101: 1, // 2,5
				0b00000000110000000110: 1, // 3,6
				0b00000001000000000111: 1, // 4,7
			},
		},
		{
			name: "diagonal, bottom left to top right",
			tuples: []tuple{
				{
					{1, 7},
					{4, 4},
				},
			},
			want: map[uint]uint{
				0b00000001000000000100: 1, // 4,4
				0b00000000110000000101: 1, // 3,5
				0b00000000100000000110: 1, // 2,6
				0b00000000010000000111: 1, // 1,7
			},
		},
		{
			name: "diagonal, top right to bottom left",
			tuples: []tuple{
				{
					{4, 4},
					{1, 7},
				},
			},
			want: map[uint]uint{
				0b00000001000000000100: 1, // 4,4
				0b00000000110000000101: 1, // 3,5
				0b00000000100000000110: 1, // 2,6
				0b00000000010000000111: 1, // 1,7
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, mapLines(tt.tuples), "mapLines(%v)", tt.tuples)
		})
	}
}
