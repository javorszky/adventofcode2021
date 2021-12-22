package day18

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_parse(t *testing.T) {
	type args struct {
		in string
	}

	tests := []struct {
		name string
		args args
		want *node
	}{
		{
			name: "parses a single node",
			args: args{in: "4"},
			want: &node{
				value: 4,
			},
		},
		{
			name: "parses a simple pair",
			args: args{in: "[1,2]"},
			want: &node{
				value: 0,
				left:  &node{value: 1},
				right: &node{value: 2},
			},
		},
		{
			name: "parses a pair where the left is a pair itself",
			args: args{in: "[[5,9],2]"},
			want: &node{
				value: 0,
				left: &node{
					value: 0,
					left: &node{
						value: 5,
					},
					right: &node{
						value: 9,
					},
				},
				right: &node{value: 2},
			},
		},
		{
			name: "parses a complex string",
			args: args{in: "[[[[4,3],4],4],[7,[[8,4],9]]]"},
			want: &node{
				left: &node{
					left: &node{
						left: &node{
							left:  &node{value: 4},
							right: &node{value: 3},
						},
						right: &node{value: 4},
					},
					right: &node{value: 4},
				},
				right: &node{
					left: &node{value: 7},
					right: &node{
						left: &node{
							left:  &node{value: 8},
							right: &node{value: 4},
						},
						right: &node{value: 9},
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, parse(tt.args.in))
		})
	}
}

func Test_isLeaf(t *testing.T) {
	type args struct {
		in *node
	}

	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "node is leaf with val 0",
			args: args{in: &node{
				value: 0,
				left:  nil,
				right: nil,
			}},
			want: true,
		},
		{
			name: "node is leaf with val 1-9",
			args: args{in: &node{
				value: 4,
				left:  nil,
				right: nil,
			}},
			want: true,
		},
		{
			name: "node is leaf with val 32",
			args: args{in: &node{
				value: 32,
				left:  nil,
				right: nil,
			}},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, isLeaf(tt.args.in), "isLeaf(%v)", tt.args.in)
		})
	}
}

func Test_isPair(t *testing.T) {
	type args struct {
		in *node
	}

	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "returns yes for pair",
			args: args{in: &node{
				left:  &node{value: 1},
				right: &node{value: 2},
			}},
			want: true,
		},
		{
			name: "returns no for a leaf",
			args: args{in: &node{value: 3}},
			want: false,
		},
		{
			name: "returns no for a complex node",
			args: args{in: &node{
				left: &node{value: 4},
				right: &node{
					left:  &node{value: 9},
					right: &node{value: 7},
				},
			}},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, isPair(tt.args.in), "isPair(%v)", tt.args.in)
		})
	}
}

func Test_addNodes(t *testing.T) {
	type args struct {
		left  *node
		right *node
	}

	tests := []struct {
		name string
		args args
		want *node
	}{
		{
			name: "adds two nodes together",
			args: args{
				left: &node{
					left: &node{
						left: &node{value: 1},
						right: &node{
							left:  &node{value: 8},
							right: &node{value: 7},
						},
					},
					right: &node{value: 6},
				},
				right: &node{value: 5},
			},
			want: &node{
				left: &node{
					left: &node{
						left: &node{value: 1},
						right: &node{
							left:  &node{value: 8},
							right: &node{value: 7},
						},
					},
					right: &node{value: 6},
				},
				right: &node{value: 5},
			},
		},
		{
			name: "adds same two nodes together, but other way",
			args: args{
				right: &node{
					left: &node{
						left: &node{value: 1},
						right: &node{
							left:  &node{value: 8},
							right: &node{value: 7},
						},
					},
					right: &node{value: 6},
				},
				left: &node{value: 5},
			},
			want: &node{
				left: &node{value: 5},
				right: &node{
					left: &node{
						left: &node{value: 1},
						right: &node{
							left:  &node{value: 8},
							right: &node{value: 7},
						},
					},
					right: &node{value: 6},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, addNodes(tt.args.left, tt.args.right), "addNodes(%v, %v)", tt.args.left, tt.args.right)
		})
	}
}

func Test_gatherNodesFromLeft(t *testing.T) {
	type args struct {
		root *node
	}

	tests := []struct {
		name string
		args args
		want []*node
	}{
		{
			name: "gathers small tree",
			args: args{
				root: &node{
					left: &node{
						left: &node{value: 1},
						right: &node{
							left:  &node{value: 2},
							right: &node{value: 3},
						},
					},
					right: &node{
						left: &node{value: 4},
						right: &node{
							left:  &node{value: 5},
							right: &node{value: 6},
						},
					},
				},
			},
			want: []*node{
				{value: 1},
				{value: 2},
				{value: 3},
				{value: 4},
				{value: 5},
				{value: 6},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, gatherNodesFromLeft(tt.args.root), "gatherNodesFromLeft(%v)", tt.args.root)
		})
	}
}

func Test_gatherPairsBelowFour(t *testing.T) {
	type args struct {
		root *node
	}

	tests := []struct {
		name string
		args args
		want map[int][]*node
	}{
		{
			name: "gathers nodes in tiers",
			args: args{
				root: &node{ // 0
					left: &node{ // 1
						left: &node{ // 2
							left: &node{value: 9}, // 3
							right: &node{ // 3
								left:  &node{value: 1}, // 4
								right: &node{value: 8}, // 4
							},
						},
						right: &node{value: 2}, // 2
					},
					right: &node{ // 1
						left: &node{ // 2
							left: &node{value: 7}, // 3
							right: &node{ // 3
								left:  &node{value: 3}, // 4
								right: &node{value: 6}, // 4
							},
						},
						right: &node{value: 4}, // 2
					},
				},
			},
			want: map[int][]*node{
				0: { // 0
					&node{
						left: &node{
							left: &node{
								left: &node{value: 9},
								right: &node{
									left:  &node{value: 1},
									right: &node{value: 8},
								},
							},
							right: &node{value: 2},
						},
						right: &node{
							left: &node{
								left: &node{value: 7},
								right: &node{
									left:  &node{value: 3},
									right: &node{value: 6},
								},
							},
							right: &node{value: 4},
						},
					},
				}, // 0
				1: {
					&node{
						left: &node{
							left: &node{value: 9},
							right: &node{
								left:  &node{value: 1},
								right: &node{value: 8},
							},
						},
						right: &node{value: 2},
					},
					&node{
						left: &node{
							left: &node{value: 7},
							right: &node{
								left:  &node{value: 3},
								right: &node{value: 6},
							},
						},
						right: &node{value: 4},
					},
				}, // 1
				2: {
					&node{
						left: &node{value: 9},
						right: &node{
							left:  &node{value: 1},
							right: &node{value: 8},
						},
					},
					&node{value: 2},
					&node{ // 2
						left: &node{value: 7}, // 3
						right: &node{ // 3
							left:  &node{value: 3}, // 4
							right: &node{value: 6}, // 4
						},
					},
					&node{value: 4},
				}, // 2
				3: {
					&node{value: 9},
					&node{ // 3
						left:  &node{value: 1}, // 4
						right: &node{value: 8}, // 4
					},
					&node{value: 7},
					&node{ // 3
						left:  &node{value: 3}, // 4
						right: &node{value: 6}, // 4
					},
				}, // 3
				4: {
					&node{value: 1},
					&node{value: 8},
					&node{value: 3},
					&node{value: 6},
				}, // 4
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, gatherNodesAtTiers(tt.args.root), "gatherNodesAtTiers(%v)", tt.args.root)
		})
	}
}
