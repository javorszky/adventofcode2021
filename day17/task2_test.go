package day17

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_firstNNumbers(t *testing.T) {
	type args struct {
		max int
	}

	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "returns slice of sum of first n numbers, max 10",
			args: args{
				max: 10,
			},
			want: []int{
				0,  // 0
				1,  // 0 + 1
				3,  // 0 + 1 + 2
				6,  // 0 + 1 + 2 + 3
				10, // 0 + 1 + 2 + 3 + 4
				15, // 0 + 1 + 2 + 3 + 4 + 5
				21, // 0 + 1 + 2 + 3 + 4 + 5 + 6
				28, // 0 + 1 + 2 + 3 + 4 + 5 + 6 + 7
				36, // 0 + 1 + 2 + 3 + 4 + 5 + 6 + 7 + 8
				45, // 0 + 1 + 2 + 3 + 4 + 5 + 6 + 7 + 8 + 9
				55, // 0 + 1 + 2 + 3 + 4 + 5 + 6 + 7 + 8 + 9 + 10
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, firstNNumbers(tt.args.max), "firstNNumbers(%v)", tt.args.max)
		})
	}
}

func Test_getDirectShots(t *testing.T) {
	type args struct {
		xMin int
		xMax int
		yMin int
		yMax int
	}

	tests := []struct {
		name string
		args args
		want [][2]int
	}{
		{
			name: "gets all direct shots",
			args: args{
				xMin: 26,
				xMax: 29,
				yMin: -81,
				yMax: -75,
			},
			want: [][2]int{
				{26, -81},
				{26, -80},
				{26, -79},
				{26, -78},
				{26, -77},
				{26, -76},
				{26, -75},
				{27, -81},
				{27, -80},
				{27, -79},
				{27, -78},
				{27, -77},
				{27, -76},
				{27, -75},
				{28, -81},
				{28, -80},
				{28, -79},
				{28, -78},
				{28, -77},
				{28, -76},
				{28, -75},
				{29, -81},
				{29, -80},
				{29, -79},
				{29, -78},
				{29, -77},
				{29, -76},
				{29, -75},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.ElementsMatchf(t, tt.want, getDirectShots(tt.args.xMin, tt.args.xMax, tt.args.yMin, tt.args.yMax),
				"getDirectShots(%v, %v, %v, %v)", tt.args.xMin, tt.args.xMax, tt.args.yMin, tt.args.yMax)
		})
	}
}

func Test_getXTargetSpeeds(t *testing.T) {
	type args struct {
		xMin int
		xMax int
	}

	tests := []struct {
		name string
		args args
		want map[int][]int
	}{
		{
			name: "gets map",
			args: args{
				xMin: 20,
				xMax: 30,
			},
			want: map[int][]int{
				1: {20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30},
				2: {11, 12, 13, 14, 15},
				3: {8, 9, 10, 11},
				4: {7, 8, 9},
				5: {6, 7, 8},
				6: {6, 7},
				7: {7},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, getXTargetSpeeds(tt.args.xMin, tt.args.xMax),
				"getXTargetSpeeds(%v, %v)", tt.args.xMin, tt.args.xMax)
		})
	}
}

func Test_task2(t *testing.T) {
	type args struct {
		coords []int
	}

	tests := []struct {
		name string
		args args
		f    func([]int) int
		want int
	}{
		{
			name: "runs task 2 on test input",
			args: args{coords: []int{20, 30, -10, -5}},
			f:    task2,
			want: 112,
		},
		{
			name: "runs task 2 on test input",
			args: args{coords: []int{20, 30, -10, -5}},
			f:    task2Functions,
			want: 112,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, tt.f(tt.args.coords), "task2(%v)", tt.args.coords)
		})
	}
}

func Test_getFiringSolutions(t *testing.T) {
	type args struct {
		xMin int
		xMax int
		yMin int
		yMax int
	}

	tests := []struct {
		name string
		args args
		want [][2]int
	}{
		{
			name: "example firing solutions",
			args: args{
				xMin: 20,
				xMax: 30,
				yMin: -10,
				yMax: -5,
			},
			want: [][2]int{
				{6, 0},
				{6, 1},
				{6, 2},
				{6, 3},
				{6, 4},
				{6, 5},
				{6, 6},
				{6, 7},
				{6, 8},
				{6, 9},
				{7, -1},
				{7, 0},
				{7, 1},
				{7, 2},
				{7, 3},
				{7, 4},
				{7, 5},
				{7, 6},
				{7, 7},
				{7, 8},
				{7, 9},
				{8, -2},
				{8, -1},
				{8, 0},
				{8, 1},
				{9, -2},
				{9, -1},
				{9, 0},
				{10, -2},
				{10, -1},
				{11, -4},
				{11, -3},
				{11, -2},
				{11, -1},
				{12, -4},
				{12, -3},
				{12, -2},
				{13, -4},
				{13, -3},
				{13, -2},
				{14, -4},
				{14, -3},
				{14, -2},
				{15, -4},
				{15, -3},
				{15, -2},
				{20, -10},
				{20, -9},
				{20, -8},
				{20, -7},
				{20, -6},
				{20, -5},
				{21, -10},
				{21, -9},
				{21, -7},
				{21, -8},
				{21, -6},
				{21, -5},
				{22, -10},
				{22, -9},
				{22, -8},
				{22, -7},
				{22, -6},
				{22, -5},
				{23, -10},
				{23, -9},
				{23, -8},
				{23, -7},
				{23, -6},
				{23, -5},
				{24, -10},
				{24, -9},
				{24, -8},
				{24, -7},
				{24, -6},
				{24, -5},
				{25, -10},
				{25, -9},
				{25, -8},
				{25, -7},
				{25, -6},
				{25, -5},
				{26, -10},
				{26, -9},
				{26, -8},
				{26, -7},
				{26, -6},
				{26, -5},
				{27, -10},
				{27, -9},
				{27, -8},
				{27, -7},
				{27, -6},
				{27, -5},
				{28, -10},
				{28, -9},
				{28, -8},
				{28, -7},
				{28, -6},
				{28, -5},
				{29, -10},
				{29, -9},
				{29, -8},
				{29, -7},
				{29, -6},
				{29, -5},
				{30, -10},
				{30, -9},
				{30, -8},
				{30, -7},
				{30, -6},
				{30, -5},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.ElementsMatchf(t, tt.want, getFiringSolutions(tt.args.xMin, tt.args.xMax, tt.args.yMin, tt.args.yMax),
				"getFiringSolutions(%v, %v, %v, %v)", tt.args.xMin, tt.args.xMax, tt.args.yMin, tt.args.yMax)
		})
	}
}

func Test_getYTargetSpeeds(t *testing.T) {
	type args struct {
		yMin int
		yMax int
	}

	tests := []struct {
		name string
		args args
		want map[int][]int
	}{
		{
			name: "gets y coordinate starting positions",
			args: args{
				yMin: -20,
				yMax: -10,
			},
			want: map[int][]int{
				1: {-10, -11, -12, -13, -14, -15, -16, -17, -18, -19, -20},
				2: {-5, -6, -7, -8, -9},
				3: {-3, -4, -5},
				4: {-1, -2, -3},
				5: {-0, -1, -2},
				6: {-0},
			},
		},
		{
			name: "gets y coordinate starting positions",
			args: args{
				yMin: -10,
				yMax: -5,
			},
			want: map[int][]int{
				1: {-5, -6, -7, -8, -9, -10},
				2: {-2, -3, -4},
				3: {-1, -2},
				4: {0, -1},
				5: {0},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, getYTargetSpeeds(tt.args.yMin, tt.args.yMax),
				"getYTargetSpeeds(%v, %v)", tt.args.yMin, tt.args.yMax)
		})
	}
}
