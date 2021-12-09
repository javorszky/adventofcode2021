package day09

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_makeGrid(t *testing.T) {
	tests := []struct {
		name  string
		in    []string
		want  [][]int
		want1 [][]int
	}{
		{
			name: "parses a 3x3 grid correctly",
			in: []string{
				"123",
				"456",
				"789",
			},
			want: [][]int{
				{1, 2, 3},
				{4, 5, 6},
				{7, 8, 9},
			},
			want1: [][]int{
				{1, 4, 7},
				{2, 5, 8},
				{3, 6, 9},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := makeGrid(tt.in)
			assert.Equal(t, tt.want, got)
			assert.Equal(t, tt.want1, got1)
		})
	}
}

func Test_getValleys(t *testing.T) {
	tests := []struct {
		name string
		in   []int
		want []int
	}{
		{
			name: "finds dips in longish array",
			//          o                    o     o
			in:   []int{1, 2, 3, 4, 3, 3, 5, 2, 5, 3},
			want: []int{0, 7, 9},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, getValleys(tt.in), "getValleys(%v)", tt.in)
		})
	}
}

func Test_task1(t *testing.T) {
	tests := []struct {
		name  string
		input []string
		want  int
	}{
		{
			name: "parses example correctly",
			input: []string{
				"2199943210",
				"3987894921",
				"9856789892",
				"8767896789",
				"9899965678",
			},
			want: 15,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, task1(tt.input), "task1(%v)", tt.input)
		})
	}
}
