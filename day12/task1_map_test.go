package day12

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_parseIntoNodeMap(t *testing.T) {
	type args struct {
		input []string
	}

	tests := []struct {
		name string
		args args
		want map[string][]string
	}{
		{
			name: "parses super simple graph into nodemap",
			args: args{
				input: []string{
					"start-ak",
					"end-ak",
					"start-end",
				},
			},
			want: map[string][]string{
				"start": {"ak", "end"},
				"end":   {"ak", "start"},
				"ak":    {"start", "end"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, parseIntoNodeMap(tt.args.input), "parseIntoNodeMap(%v)", tt.args.input)
		})
	}
}

func Test_walkNodeMap(t *testing.T) {
	type args struct {
		currentNode string
		allNodes    func() map[string][]string
		currentPath []string
		cntns       func([]string, string) bool
	}

	tests := []struct {
		name string
		args args
		want [][]string
	}{
		{
			name: "walks nodemap super simple",
			args: args{
				currentNode: startName,
				allNodes: func() map[string][]string {
					return parseIntoNodeMap([]string{
						"start-ak",
						"ak-end",
						"start-end",
					})
				},
				currentPath: []string{},
				cntns:       contains,
			},
			want: [][]string{
				{"start", "end"},
				{"start", "ak", "end"},
			},
		},
		{
			name: "walks nodemap small",
			args: args{
				currentNode: startName,
				allNodes: func() map[string][]string {
					return parseIntoNodeMap([]string{
						"start-A",
						"start-b",
						"A-c",
						"A-b",
						"b-d",
						"A-end",
						"b-end",
					})
				},
				currentPath: []string{},
				cntns:       contains,
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
			name: "walks nodemap slightly larger",
			args: args{
				currentNode: startName,
				allNodes: func() map[string][]string {
					return parseIntoNodeMap(getInputs("example_input.txt"))
				},
				currentPath: []string{},
				cntns:       contains,
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
			assert.ElementsMatchf(t, tt.want, walkNodeMap(
				tt.args.currentNode,
				tt.args.allNodes(),
				tt.args.currentPath,
				tt.args.cntns,
			),
				"walkNodeMap(%v, %v, %v, %v)",
				tt.args.currentNode,
				tt.args.allNodes(),
				tt.args.currentPath,
				tt.args.cntns)
		})
	}
}
