package day05

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_filterToEuclideanLines(t *testing.T) {
	tests := []struct {
		name   string
		tuples []tuple
		want   []tuple
	}{
		{
			name: "keeps all elements in",
			tuples: []tuple{
				{
					{1, 2},
					{1, 4},
				},
				{
					{1, 2},
					{4, 2},
				},
			},
			want: []tuple{
				{
					{1, 2},
					{1, 4},
				},
				{
					{1, 2},
					{4, 2},
				},
			},
		},
		{
			name:   "handles empty list correctly",
			tuples: []tuple{},
			want:   []tuple{},
		},
		{
			name:   "handles nil list correctly",
			tuples: nil,
			want:   []tuple{},
		},
		{
			name: "removes element because neither dimensions match",
			tuples: []tuple{
				{
					{1, 2},
					{1, 4},
				},
				{
					{1, 4},
					{5, 8},
				},
				{
					{1, 2},
					{4, 2},
				},
			},
			want: []tuple{
				{
					{1, 2},
					{1, 4},
				},
				{
					{1, 2},
					{4, 2},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := filterToEuclideanLines(tt.tuples)
			assert.Equal(t, tt.want, got)
		})
	}
}

func Test_filterToEuclideanLinesSlice(t *testing.T) {
	tests := []struct {
		name   string
		coords []uint
		want   []uint
	}{
		{
			name:   "keeps all elements in",
			coords: []uint{1, 2, 1, 4, 1, 2, 4, 2},
			want:   []uint{1, 2, 1, 4, 1, 2, 4, 2},
		},
		{
			name:   "handles empty list correctly",
			coords: []uint{},
			want:   []uint{},
		},
		{
			name:   "handles nil list correctly",
			coords: nil,
			want:   []uint{},
		},
		{
			name:   "removes element because neither dimensions match",
			coords: []uint{1, 2, 1, 4, 1, 4, 5, 8, 1, 2, 4, 2},
			want:   []uint{1, 2, 1, 4, 1, 2, 4, 2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, filterToEuclideanLinesSlice(tt.coords), "filterToEuclideanLinesSlice(%v)", tt.coords)
		})
	}
}

func Test_task1(t *testing.T) {
	input := []string{
		"0,9 -> 5,9",
		"8,0 -> 0,8",
		"9,4 -> 3,4",
		"2,2 -> 2,1",
		"7,0 -> 7,4",
		"6,4 -> 2,0",
		"0,9 -> 2,9",
		"3,4 -> 1,4",
		"0,0 -> 8,8",
		"5,5 -> 8,2",
	}

	tests := []struct {
		name  string
		input []string
		fn    func([]string) int
		want  int
	}{
		{
			name:  "task 1 with example input works correctly",
			input: input,
			fn:    task1,
			want:  5,
		},
		{
			name:  "task 1 slicy with example input works correcly",
			input: input,
			fn:    task1Slicy,
			want:  5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, task1(tt.input), "task1(%v)", tt.input)
		})
	}
}
