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
			name: "extrapolates two intersecting coords to lines",
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
			assert.Equalf(t, tt.want, mapLinesTuples(tt.tuples), "mapLinesTuples(%v)", tt.tuples)
		})
	}
}

func Test_getTuples(t *testing.T) {
	fileData := []string{
		"242,601 -> 242,18",
		"938,357 -> 938,128",
		"920,574 -> 750,574",
		"804,978 -> 804,813",
		"955,932 -> 68,45",
		"232,604 -> 232,843",
	}

	ts := []tuple{
		{
			{242, 601},
			{242, 18},
		},
		{
			{938, 357},
			{938, 128},
		},
		{
			{920, 574},
			{750, 574},
		},
		{
			{804, 978},

			{804, 813},
		},
		{
			{955, 932},
			{68, 45},
		},
		{
			{232, 604},
			{232, 843},
		},
	}

	tests := []struct {
		name string
		fn   func([]string) []tuple
		want []tuple
	}{
		{
			name: "gettuples works with regex",
			fn:   getTuples,
			want: ts,
		},
		{
			name: "gettuples works with string split",
			fn:   getTuplesString,
			want: ts,
		},
		{
			name: "gettuples works with reversed char by char parsing",
			fn:   getTuplesReversed,
			want: ts,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, tt.fn(fileData), "getTuples(%v)", fileData)
		})
	}
}

func Test_getCoordinateSliceReverse(t *testing.T) {
	tests := []struct {
		name     string
		fileData []string
		want     []uint
	}{
		{
			name: "parses one line correctly",
			fileData: []string{
				"242,601 -> 242,18",
			},
			want: []uint{242, 601, 242, 18},
		},
		{
			name: "parses three lines correctly",
			fileData: []string{
				"242,601 -> 242,18",
				"938,357 -> 938,128",
				"920,574 -> 750,574",
			},
			want: []uint{
				242, 601, 242, 18,
				938, 357, 938, 128,
				920, 574, 750, 574,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, getCoordinateSliceReverse(tt.fileData), "getCoordinateSliceReverse(%v)", tt.fileData)
		})
	}
}

func Test_mapLinesSlice(t *testing.T) {
	tests := []struct {
		name   string
		coords []uint
		want   map[uint]uint
	}{
		{
			name:   "extrapolates one vertical tuple to lines, first bigger",
			coords: []uint{1, 2, 1, 4},
			want: map[uint]uint{
				0b00000000010000000010: 1,
				0b00000000010000000011: 1,
				0b00000000010000000100: 1,
			},
		},
		{
			name:   "extrapolates one vertical tuple to lines, second bigger",
			coords: []uint{1, 4, 1, 2},
			want: map[uint]uint{
				0b00000000010000000010: 1,
				0b00000000010000000011: 1,
				0b00000000010000000100: 1,
			},
		},
		{
			name:   "extrapolates one horizontal tuple to lines, first bigger",
			coords: []uint{1, 6, 5, 6},
			want: map[uint]uint{
				0b00000000010000000110: 1,
				0b00000000100000000110: 1,
				0b00000000110000000110: 1,
				0b00000001000000000110: 1,
				0b00000001010000000110: 1,
			},
		},
		{
			name:   "extrapolates one horizontal tuple to lines, second bigger",
			coords: []uint{5, 6, 1, 6},
			want: map[uint]uint{
				0b00000000010000000110: 1,
				0b00000000100000000110: 1,
				0b00000000110000000110: 1,
				0b00000001000000000110: 1,
				0b00000001010000000110: 1,
			},
		},
		{
			name:   "extrapolates two intersecting coords to lines",
			coords: []uint{5, 6, 1, 6, 3, 2, 3, 8},
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
			name:   "diagonal, top left to bottom right",
			coords: []uint{1, 4, 4, 7},
			want: map[uint]uint{
				0b00000000010000000100: 1, // 1,4
				0b00000000100000000101: 1, // 2,5
				0b00000000110000000110: 1, // 3,6
				0b00000001000000000111: 1, // 4,7
			},
		},
		{
			name:   "diagonal, bottom right to top left",
			coords: []uint{4, 7, 1, 4},
			want: map[uint]uint{
				0b00000000010000000100: 1, // 1,4
				0b00000000100000000101: 1, // 2,5
				0b00000000110000000110: 1, // 3,6
				0b00000001000000000111: 1, // 4,7
			},
		},
		{
			name: "diagonal, bottom left to top right",
			coords: []uint{1, 7,
				4, 4},
			want: map[uint]uint{
				0b00000001000000000100: 1, // 4,4
				0b00000000110000000101: 1, // 3,5
				0b00000000100000000110: 1, // 2,6
				0b00000000010000000111: 1, // 1,7
			},
		},
		{
			name:   "diagonal, top right to bottom left",
			coords: []uint{4, 4, 1, 7},
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
			assert.Equalf(t, tt.want, mapLinesSlice(tt.coords), "mapLinesSlice(%v)", tt.coords)
		})
	}
}
