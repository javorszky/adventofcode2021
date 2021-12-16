package day15

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_makeMapMap(t *testing.T) {
	type args struct {
		input []string
	}

	tests := []struct {
		name string
		args args
		want map[int]map[int]int
	}{
		{
			name: "makes a map out of string slice",
			args: args{
				input: []string{
					"1234",
					"5678",
					"9012",
					"3456",
				},
			},
			want: map[int]map[int]int{
				0: {
					0: 1,
					1: 2,
					2: 3,
					3: 4,
				},
				1: {
					0: 5,
					1: 6,
					2: 7,
					3: 8,
				},
				2: {
					0: 9,
					1: 0,
					2: 1,
					3: 2,
				},
				3: {
					0: 3,
					1: 4,
					2: 5,
					3: 6,
				},
			},
		},
		{
			name: "makes a map out of string slice bigger",
			args: args{
				input: []string{
					"199999",
					"119999",
					"919999",
					"119999",
					"199111",
					"111191",
				},
			},
			want: map[int]map[int]int{
				0: {
					0: 1,
					1: 9,
					2: 9,
					3: 9,
					4: 9,
					5: 9,
				},
				1: {
					0: 1,
					1: 1,
					2: 9,
					3: 9,
					4: 9,
					5: 9,
				},
				2: {
					0: 9,
					1: 1,
					2: 9,
					3: 9,
					4: 9,
					5: 9,
				},
				3: {
					0: 1,
					1: 1,
					2: 9,
					3: 9,
					4: 9,
					5: 9,
				},
				4: {
					0: 1,
					1: 9,
					2: 9,
					3: 1,
					4: 1,
					5: 1,
				},
				5: {
					0: 1,
					1: 1,
					2: 1,
					3: 1,
					4: 9,
					5: 1,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, makeMapMap(tt.args.input), "makeMapMap(%v)", tt.args.input)
		})
	}
}

func Test_makeWalkOrderMap(t *testing.T) {
	type args struct {
		in map[int]map[int]int
	}

	tests := []struct {
		name string
		args args
		want [][2]int
	}{
		{
			name: "creates walk order for field",
			args: args{
				in: map[int]map[int]int{
					0: {
						0: 1,
						1: 2,
						2: 3,
						3: 4,
					},
					1: {
						0: 5,
						1: 6,
						2: 7,
						3: 8,
					},
					2: {
						0: 9,
						1: 0,
						2: 1,
						3: 2,
					},
					3: {
						0: 3,
						1: 4,
						2: 5,
						3: 6,
					},
				},
			},
			want: [][2]int{
				{0, 0},
				{0, 1}, {1, 0},
				{0, 2}, {1, 1}, {2, 0},
				{0, 3}, {1, 2}, {2, 1}, {3, 0},
				{1, 3}, {2, 2}, {3, 1},
				{2, 3}, {3, 2},
				{3, 3},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, makeWalkOrderMap(tt.args.in), "makeWalkOrderMap(%v)", tt.args.in)
		})
	}
}

func Test_makeRiskMapMap(t *testing.T) {
	startingField := map[int]map[int]int{
		0: {
			0: 1,
			1: 9,
			2: 9,
			3: 9,
			4: 9,
			5: 9,
		},
		1: {
			0: 1,
			1: 1,
			2: 9,
			3: 9,
			4: 9,
			5: 9,
		},
		2: {
			0: 9,
			1: 1,
			2: 9,
			3: 9,
			4: 9,
			5: 9,
		},
		3: {
			0: 1,
			1: 1,
			2: 9,
			3: 9,
			4: 9,
			5: 9,
		},
		4: {
			0: 1,
			1: 9,
			2: 9,
			3: 1,
			4: 1,
			5: 1,
		},
		5: {
			0: 1,
			1: 1,
			2: 1,
			3: 1,
			4: 9,
			5: 1,
		},
	}

	type args struct {
		field map[int]map[int]int
		order [][2]int
	}

	tests := []struct {
		name   string
		args   args
		f      func(map[int]map[int]int, [][2]int) map[int]map[int]int
		fAgain func(map[int]map[int]int, map[int]map[int]int) map[int]map[int]int
		want   map[int]map[int]int
	}{
		{
			name: "makes a risk map for an artificial input",
			args: args{
				field: startingField,
				order: makeWalkOrderMap(startingField),
			},
			want: map[int]map[int]int{
				0: {
					0: 0,
					1: 9,
					2: 18,
					3: 27,
					4: 36,
					5: 45,
				},
				1: {
					0: 1,
					1: 2,
					2: 11,
					3: 20,
					4: 29,
					5: 38,
				},
				2: {
					0: 10,
					1: 3,
					2: 12,
					3: 21,
					4: 30,
					5: 39,
				},
				3: {
					0: 11,
					1: 4,
					2: 13,
					3: 22,
					4: 31,
					5: 40,
				},
				4: {
					0: 12,
					1: 13,
					2: 22,
					3: 23,
					4: 24,
					5: 25,
				},
				5: {
					0: 13,
					1: 14,
					2: 15,
					3: 16,
					4: 25,
					5: 26,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, makeRiskMapMap(tt.args.field, tt.args.order),
				"makeRiskMapMap(%v, %v)", tt.args.field, tt.args.order)
		})
	}
}
