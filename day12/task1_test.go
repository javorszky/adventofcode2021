package day12

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_contains(t *testing.T) {
	type args struct {
		path    []string
		element string
	}

	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "contains returns true for exact match",
			args: args{
				path:    []string{"ab", "bc", "DE", "de", "fg"},
				element: "de",
			},
			want: true,
		},
		{
			name: "contains returns true for exact match",
			args: args{
				path:    []string{"ab", "bc", "DE", "de", "fg"},
				element: "xo",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, contains(tt.args.path, tt.args.element))
		})
	}
}

func Test_walkNodes(t *testing.T) {
	type args struct {
		start       func() *node
		currentPath []string
	}

	tests := []struct {
		name string
		args args
		want [][]string
	}{
		{
			name: "walks nodes super simple",
			args: args{
				start: func() *node {
					nodes := parseIntoNodes([]string{
						"start-ak",
						"ak-end",
						"start-end",
					})

					return nodes["start"]
				},
				currentPath: []string{},
			},
			want: [][]string{
				{"start", "end"},
				{"start", "ak", "end"},
			},
		},
		{
			name: "walks nodes super simple",
			args: args{
				start: func() *node {
					nodes := parseIntoNodes([]string{
						"start-A",
						"start-b",
						"A-c",
						"A-b",
						"b-d",
						"A-end",
						"b-end",
					})

					return nodes["start"]
				},
				currentPath: []string{},
			},
			want: [][]string{
				{"start", "A", "b", "A", "c", "A", "end"},
				{"start", "A", "b", "A", "end"},
				{"start", "A", "b", "end"},
				{"start", "A", "c", "A", "b", "A", "end"},
				{"start", "A", "c", "A", "b", "end"},
				{"start", "A", "c", "A", "end"},
				{"start", "A", "end"},
				{"start", "b", "A", "c", "A", "end"},
				{"start", "b", "A", "end"},
				{"start", "b", "end"},
			},
		},
		{
			name: "walks nodes example slightly larger",
			args: args{
				start: func() *node {
					nodes := parseIntoNodes(getInputs("example_input.txt"))

					return nodes["start"]
				},
				currentPath: []string{},
			},
			want: [][]string{
				{"start", "HN", "dc", "HN", "end"},
				{"start", "HN", "dc", "HN", "kj", "HN", "end"},
				{"start", "HN", "dc", "end"},
				{"start", "HN", "dc", "kj", "HN", "end"},
				{"start", "HN", "end"},
				{"start", "HN", "kj", "HN", "dc", "HN", "end"},
				{"start", "HN", "kj", "HN", "dc", "end"},
				{"start", "HN", "kj", "HN", "end"},
				{"start", "HN", "kj", "dc", "HN", "end"},
				{"start", "HN", "kj", "dc", "end"},
				{"start", "dc", "HN", "end"},
				{"start", "dc", "HN", "kj", "HN", "end"},
				{"start", "dc", "end"},
				{"start", "dc", "kj", "HN", "end"},
				{"start", "kj", "HN", "dc", "HN", "end"},
				{"start", "kj", "HN", "dc", "end"},
				{"start", "kj", "HN", "end"},
				{"start", "kj", "dc", "HN", "end"},
				{"start", "kj", "dc", "end"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.ElementsMatchf(t, tt.want, walkNodes(tt.args.start(), tt.args.currentPath),
				"walkNodes(%v, %v)", tt.args.start, tt.args.currentPath)
		})
	}
}
