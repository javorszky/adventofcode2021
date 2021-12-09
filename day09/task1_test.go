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
