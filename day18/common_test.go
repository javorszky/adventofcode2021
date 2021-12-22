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
		left  string
		right string
	}

	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "adds two nodes smallest example",
			args: args{
				left:  "[[[[4,3],4],4],[7,[[8,4],9]]]",
				right: "[1,1]",
			},
			want: "[[[[0,7],4],[[7,8],[6,0]]],[8,1]]",
		},
		{
			name: "adds two nodes together",
			args: args{
				left:  "[[[0,[4,5]],[0,0]],[[[4,5],[2,6]],[9,5]]]",
				right: "[7,[[[3,7],[4,3]],[[6,3],[8,8]]]]",
			},
			want: "[[[[4,0],[5,4]],[[7,7],[6,0]]],[[8,[7,7]],[[7,9],[5,0]]]]",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			lNode := parse(tt.args.left)
			rNode := parse(tt.args.right)
			assert.Equalf(t, tt.want, addNodes(lNode, rNode).String(), "addNodes(%v, %v)", tt.args.left, tt.args.right)
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

func Test_node_String(t *testing.T) {
	type fields struct {
		value int
		left  *node
		right *node
	}

	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "returns the string representation of a tree",
			fields: fields{
				value: 0,
				left: &node{
					left: &node{value: 3},
					right: &node{
						left: &node{value: 2},
						right: &node{
							left:  &node{value: 7},
							right: &node{value: 4},
						},
					},
				},
				right: &node{
					left:  &node{value: 5},
					right: &node{value: 8},
				},
			},
			want: "[[3,[2,[7,4]]],[5,8]]",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			n := &node{
				value: tt.fields.value,
				left:  tt.fields.left,
				right: tt.fields.right,
			}
			assert.Equalf(t, tt.want, n.String(), "String()")
			assert.Equalf(t, n, parse(n.String()), "re parsing the node from the string")
		})
	}
}

func Test_runExplosion(t *testing.T) {
	tests := []struct {
		name string
		in   string
		out  string
	}{
		{
			name: "single explosion example - 1",
			in:   "[[[[[9,8],1],2],3],4]",
			out:  "[[[[0,9],2],3],4]",
		},
		{
			name: "single explosion example - 2",
			in:   "[7,[6,[5,[4,[3,2]]]]]",
			out:  "[7,[6,[5,[7,0]]]]",
		},
		{
			name: "single explosion example - 3",
			in:   "[[6,[5,[4,[3,2]]]],1]",
			out:  "[[6,[5,[7,0]]],3]",
		},
		{
			name: "single explosion example - 4",
			in:   "[[3,[2,[1,[7,3]]]],[6,[5,[4,[3,2]]]]]",
			out:  "[[3,[2,[8,0]]],[9,[5,[4,[3,2]]]]]",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := parse(tt.in)
			runExplosion(r)

			assert.Equalf(t, tt.out, r.String(), "run explosion on tree %s", tt.in)
		})
	}
}

func Test_runSplit(t *testing.T) {
	tests := []struct {
		name string
		in   string
		out  string
	}{
		{
			name: "splits example - 1",
			in:   "[[[[0,7],4],[15,[0,13]]],[1,1]]",
			out:  "[[[[0,7],4],[[7,8],[0,13]]],[1,1]]",
		},
		{
			name: "splits example - 2",
			in:   "[[[[0,7],4],[[7,8],[0,13]]],[1,1]]",
			out:  "[[[[0,7],4],[[7,8],[0,[6,7]]]],[1,1]]",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			root := parse(tt.in)
			runSplit(root)
			assert.Equalf(t, tt.out, root.String(), "running split on %s", tt.in)
		})
	}
}

func Test_node_Magnitude(t *testing.T) {
	tests := []struct {
		name string
		in   string
		want int
	}{
		{
			name: "[9,1]",
			in:   "[9,1]",
			want: 29,
		},
		{
			name: "[[[[0,7],4],[[7,8],[6,0]]],[8,1]]",
			in:   "[[[[0,7],4],[[7,8],[6,0]]],[8,1]]",
			want: 1384,
		},
		{
			name: "[[[[8,7],[7,7]],[[8,6],[7,7]]],[[[0,7],[6,6]],[8,7]]]",
			in:   "[[[[8,7],[7,7]],[[8,6],[7,7]]],[[[0,7],[6,6]],[8,7]]]",
			want: 3488,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			n := parse(tt.in)
			assert.Equalf(t, tt.want, n.Magnitude(), "Magnitude()")
		})
	}
}
