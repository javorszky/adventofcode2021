package day10

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_hasSyntaxError(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want bool
	}{
		{
			name: "finds closing chars",
			s:    "((([[((<(([<((})>])([(<(",
			want: true,
		},
		{
			name: "finds closing chars 2",
			s:    "[{{<({<{{<({[{)]})<{<",
			want: true,
		},
		{
			name: "does not find closing chars",
			s:    "[[([{[<<(<<[{[",
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, hasSyntaxError(tt.s), "hasSyntaxError(%v)", tt.s)
		})
	}
}

func Test_gatherClosingScore(t *testing.T) {
	tests := []struct {
		name     string
		openings string
		want     int
	}{
		{
			name:     "tests example given",
			openings: "<{([", // we need to use the reverse here from the example: ])}>
			want:     294,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, gatherClosingScore(tt.openings), "gatherClosingScore(%v)", tt.openings)
		})
	}
}

func Test_task2(t *testing.T) {
	tests := []struct {
		name     string
		filename string
		want     int
	}{
		{
			name:     "solves the example",
			filename: "example_input.txt",
			want:     288957,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			input := getInputs(tt.filename)
			assert.Equalf(t, tt.want, task2(input), "task2(%v)", input)
		})
	}
}
