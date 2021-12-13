package day13

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_makePaper(t *testing.T) {
	type args struct {
		dots []string
	}

	tests := []struct {
		name string
		args args
		want map[uint]uint
	}{
		{
			name: "parses a single coordinate",
			args: args{dots: []string{"533,911"}},
			want: map[uint]uint{
				0b0100001010101110001111: 1,
			},
		},
		{
			name: "parses small input into paper",
			args: args{dots: []string{
				"543,332",
				"1300,1300",
				"0,0",
				"392,66",
			}},
			want: map[uint]uint{
				0b0100001111100101001100: 1, // 543,332
				0b1010001010010100010100: 1, // 1300,1300
				0b0000000000000000000000: 1, // 0,0
				0b0011000100000001000010: 1, // 392,66
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, makePaper(tt.args.dots))
		})
	}
}

func Test_getVerticalFolded(t *testing.T) {
	type args struct {
		coord uint
		fold  uint
	}

	tests := []struct {
		name string
		args args
		want uint
	}{
		{
			name: "fold is below coordinate, ignores",
			args: args{
				coord: 234<<11 | 111,
				fold:  199,
			},
			want: 234<<11 | 111,
		},
		{
			name: "folds 2,14 at 7",
			args: args{
				coord: 2<<11 | 14,
				fold:  7,
			},
			want: 2 << 11,
		},
		{
			name: "folds 234,111 at 199 -> 164,111",
			args: args{
				// 234 to the right, 111 down
				coord: 234<<11 | 111,
				fold:  76,
			},
			want: 234<<11 | 41,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, getVerticalFolded(tt.args.coord, tt.args.fold),
				"getVerticalFolded(%v, %v)", tt.args.coord, tt.args.fold)
		})
	}
}

func Test_foldUp(t *testing.T) {
	type args struct {
		paper func() map[uint]uint
		x     uint
	}

	tests := []struct {
		name string
		args args
		want map[uint]uint
	}{
		{
			name: "folds example up at 7",
			args: args{
				paper: func() map[uint]uint {
					d, _ := getInputs("example_input.txt")

					return makePaper(d)
				},
				x: 7,
			},
			want: map[uint]uint{
				0 << 11:    1,
				2 << 11:    1,
				3 << 11:    1,
				6 << 11:    1,
				9 << 11:    1,
				0<<11 | 1:  1,
				4<<11 | 1:  1,
				6<<11 | 2:  1,
				10<<11 | 2: 1,
				0<<11 | 3:  1,
				4<<11 | 3:  1,
				1<<11 | 4:  1,
				3<<11 | 4:  1,
				6<<11 | 4:  1,
				8<<11 | 4:  1,
				9<<11 | 4:  1,
				10<<11 | 4: 1,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, foldUp(tt.args.paper(), tt.args.x), "foldUp(%v, %v)", tt.args.paper, tt.args.x)
		})
	}
}

func Test_getHorizontalFolded(t *testing.T) {
	type args struct {
		coord uint
		fold  uint
	}

	tests := []struct {
		name string
		args args
		want uint
	}{
		{
			name: "fold is to the right of coordinate, ignores",
			args: args{
				coord: 188<<11 | 111,
				fold:  199,
			},
			want: 188<<11 | 111,
		},
		{
			name: "folds 8,2 at 5",
			args: args{
				coord: 8<<11 | 2,
				fold:  5,
			},
			want: 2<<11 | 2,
		},
		{
			name: "folds 234,111 at 199 -> 164,111",
			args: args{
				// 234 to the right, 111 down
				coord: 234<<11 | 111,
				fold:  199,
			},
			want: 164<<11 | 111,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, getHorizontalFolded(tt.args.coord, tt.args.fold),
				"getHorizontalFolded(%v, %v)", tt.args.coord, tt.args.fold)
		})
	}
}

func Test_foldLeft(t *testing.T) {
	type args struct {
		paper func() map[uint]uint
		y     uint
	}

	tests := []struct {
		name string
		args args
		want map[uint]uint
	}{
		{
			name: "folds example left at 5",
			args: args{
				paper: func() map[uint]uint {
					d, _ := getInputs("example_input.txt")

					return makePaper(d)
				},
				y: 5,
			},
			want: map[uint]uint{
				1 << 11:    1,
				3 << 11:    1,
				4 << 11:    1,
				4<<11 | 1:  1,
				0<<11 | 3:  1,
				0<<11 | 4:  1,
				2<<11 | 4:  1,
				3<<11 | 4:  1,
				1<<11 | 10: 1,
				2<<11 | 10: 1,
				4<<11 | 10: 1,
				4<<11 | 11: 1,
				0<<11 | 12: 1,
				4<<11 | 12: 1,
				0<<11 | 13: 1,
				0<<11 | 14: 1,
				2<<11 | 14: 1,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, foldLeft(tt.args.paper(), tt.args.y), "foldLeft(%v, %v)", tt.args.paper(), tt.args.y)
		})
	}
}
