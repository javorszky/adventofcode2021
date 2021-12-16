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
