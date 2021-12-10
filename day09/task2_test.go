package day09

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_binToCoords(t *testing.T) {
	tests := []struct {
		name  string
		in    uint
		want  int
		want1 int
	}{
		{
			name:  "parses bin for 97,34",
			in:    0b11000010100010,
			want:  97,
			want1: 34,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := binToCoords(tt.in)
			assert.Equalf(t, tt.want, got, "binToCoords(%v)", tt.in)
			assert.Equalf(t, tt.want1, got1, "binToCoords(%v)", tt.in)
			bin2 := coordsToBin(got, got1)
			assert.Equalf(t, tt.in, bin2, "coordsToBin(%d, %d)", got, got1)
		})
	}
}

func Test_adjacentCoord(t *testing.T) {
	type args struct {
		thisOne   uint
		direction string
	}

	tests := []struct {
		name string
		args args
		want uint
	}{
		{
			name: "gets left adjacent for 50, 50",
			args: args{
				thisOne:   0b01100100110010,
				direction: "left",
			},
			want: 0b01100100110001,
		},
		{
			name: "gets right adjacent for 50, 50",
			args: args{
				thisOne:   0b01100100110010,
				direction: "right",
			},
			want: 0b01100100110011,
		},
		{
			name: "gets up adjacent for 50, 50",
			args: args{
				thisOne:   0b01100100110010,
				direction: "up",
			},
			want: 0b01100010110010,
		},
		{
			name: "gets down adjacent for 50, 50",
			args: args{
				thisOne:   0b01100100110010,
				direction: "down",
			},
			want: 0b01100110110010,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tx, ty := binToCoords(tt.args.thisOne)
			assert.Equalf(t, tt.want, adjacentCoord(tt.args.thisOne, tt.args.direction),
				"adjacentCoord([%d, %d], %s)", tx, ty, tt.args.direction)
		})
	}
}

func Test_grabBasin(t *testing.T) {
	field := [][]int{
		{2, 1, 9, 9, 9, 4, 3, 2, 1, 0},
		{3, 9, 8, 7, 8, 9, 4, 9, 2, 1},
		{9, 8, 5, 6, 7, 8, 9, 8, 9, 2},
		{8, 7, 6, 7, 8, 9, 6, 7, 8, 9},
		{9, 8, 9, 9, 9, 6, 5, 6, 7, 8},
	}

	type args struct {
		coord    uint
		fullGrid [][]int
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "finds basin top left at 0-1",
			args: args{
				coord:    0b00000000000001,
				fullGrid: field,
			},
			want: 3,
		},
		{
			name: "finds basin top right at 0-9",
			args: args{
				coord:    0b00000000001001,
				fullGrid: field,
			},
			want: 9,
		},
		{
			name: "finds basin middle at 2-2",
			args: args{
				coord:    0b00000100000010,
				fullGrid: field,
			},
			want: 14,
		},
		{
			name: "finds basin bottom at 4-6",
			args: args{
				coord:    0b00001000000110,
				fullGrid: field,
			},
			want: 9,
		},
		{
			name: "finds basin from a point that has 0 value",
			args: args{
				coord: 0b00000100000100,
				fullGrid: [][]int{
					{9, 9, 9, 9, 9, 9, 9},
					{9, 8, 8, 9, 7, 3, 9},
					{9, 5, 6, 4, 0, 2, 9},
					{9, 9, 9, 9, 9, 9, 9},
				},
			},
			want: 9,
		},
		{
			name: "finds basin from a point that has 0 value",
			args: args{
				coord: 0b00000100000100, // 2,4 from top left 2 down, 4 right
				fullGrid: [][]int{
					{9, 9, 9, 9, 9, 9, 9},
					{9, 8, 8, 9, 9, 3, 9},
					{9, 5, 6, 9, 0, 9, 9},
					{9, 9, 9, 9, 9, 9, 9},
				},
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, grabBasin(tt.args.coord, fullGridBin(tt.args.fullGrid)),
				"grabBasin(%v, %v)", tt.args.coord, tt.args.fullGrid)
		})
	}
}

func Test_task2(t *testing.T) {
	field := [][]int{
		{2, 1, 9, 9, 9, 4, 3, 2, 1, 0},
		{3, 9, 8, 7, 8, 9, 4, 9, 2, 1},
		{9, 8, 5, 6, 7, 8, 9, 8, 9, 2},
		{8, 7, 6, 7, 8, 9, 6, 7, 8, 9},
		{9, 8, 9, 9, 9, 6, 5, 6, 7, 8},
	}

	lowestPoints := map[uint]int{
		0b00000000000001: 1,
		0b00000000001001: 0,
		0b00000100000010: 5,
		0b00001000000110: 5,
	}

	type args struct {
		fullGrid     [][]int
		lowestPoints map[uint]int
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "solves example for task 2",
			args: args{
				fullGrid:     field,
				lowestPoints: lowestPoints,
			},
			want: 1134,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, task2(tt.args.fullGrid, tt.args.lowestPoints),
				"task2(%v, %v)", tt.args.fullGrid, tt.args.lowestPoints)
		})
	}
}

func Test_fullGridBin(t *testing.T) {
	tests := []struct {
		name   string
		inGrid [][]int
		want   map[uint]int
	}{
		{
			name: "parses a 3x3 fullgrid",
			inGrid: [][]int{
				{1, 2, 3},
				{4, 5, 6},
				{7, 8, 9},
			},
			want: map[uint]int{
				0b00000000000000: 1,
				0b00000000000001: 2,
				0b00000000000010: 3,
				0b00000010000000: 4,
				0b00000010000001: 5,
				0b00000010000010: 6,
				0b00000100000000: 7,
				0b00000100000001: 8,
				0b00000100000010: 9,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fGridBin := fullGridBin(tt.inGrid)
			assert.Equalf(t, tt.want, fGridBin, "fullGridBin(%v)", tt.inGrid)
			for i := 0; i < 3; i++ {
				for j := 0; j < 3; j++ {
					assert.Equalf(t, tt.inGrid[i][j], fGridBin[coordsToBin(i, j)], "comparing value for [][]int grid and bin grid")
				}
			}
		})
	}
}
