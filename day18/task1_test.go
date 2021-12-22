package day18

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_reduce(t *testing.T) {
	tests := []struct {
		name string
		in   string
		out  string
	}{
		{
			name: "reduces example btree",
			in:   "[[[[[4,3],4],4],[7,[[8,4],9]]],[1,1]]",
			out:  "[[[[0,7],4],[[7,8],[6,0]]],[8,1]]",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tree := parse(tt.in)
			reduce(tree)

			assert.Equalf(t, tt.out, tree.String(), "reduce(%s)", tt.in)
		})
	}
}

func Test_task1(t *testing.T) {
	type args struct {
		input []string
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "task 1 example",
			args: args{input: []string{
				"[[[0,[5,8]],[[1,7],[9,6]]],[[4,[1,2]],[[1,4],2]]]",
				"[[[5,[2,8]],4],[5,[[9,9],0]]]",
				"[6,[[[6,2],[5,6]],[[7,6],[4,7]]]]",
				"[[[6,[0,7]],[0,9]],[4,[9,[9,0]]]]",
				"[[[7,[6,4]],[3,[1,3]]],[[[5,5],1],9]]",
				"[[6,[[7,3],[3,2]]],[[[3,8],[5,7]],4]]",
				"[[[[5,4],[7,7]],8],[[8,3],8]]",
				"[[9,3],[[9,9],[6,[4,9]]]]",
				"[[2,[[7,7],7]],[[5,8],[[9,3],[0,2]]]]",
				"[[[[5,2],5],[8,[3,7]]],[[5,[7,5]],[4,4]]]",
			}},
			want: 4140,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, task1(tt.args.input), "task1(%v)", tt.args.input)
		})
	}
}

func Test_addNodesInOrder(t *testing.T) {
	type args struct {
		forest []string
	}

	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "adds together nodes",
			args: args{forest: []string{
				"[1,1]",
				"[2,2]",
				"[3,3]",
				"[4,4]",
			}},
			want: "[[[[1,1],[2,2]],[3,3]],[4,4]]",
		},
		{
			name: "adds together nodes",
			args: args{forest: []string{
				"[1,1]",
				"[2,2]",
				"[3,3]",
				"[4,4]",
				"[5,5]",
			}},
			want: "[[[[3,0],[5,3]],[4,4]],[5,5]]",
		},
		{
			name: "adds together nodes",
			args: args{forest: []string{
				"[1,1]",
				"[2,2]",
				"[3,3]",
				"[4,4]",
				"[5,5]",
				"[6,6]",
			}},
			want: "[[[[5,0],[7,4]],[5,5]],[6,6]]",
		},
		{
			name: "adds together nodes",
			args: args{forest: []string{
				"[[[0,[4,5]],[0,0]],[[[4,5],[2,6]],[9,5]]]",
				"[7,[[[3,7],[4,3]],[[6,3],[8,8]]]]",
				"[[2,[[0,8],[3,4]]],[[[6,7],1],[7,[1,6]]]]",
				"[[[[2,4],7],[6,[0,5]]],[[[6,8],[2,8]],[[2,1],[4,5]]]]",
				"[7,[5,[[3,8],[1,4]]]]",
				"[[2,[2,2]],[8,[8,1]]]",
				"[2,9]",
				"[1,[[[9,3],9],[[9,0],[0,7]]]]",
				"[[[5,[7,4]],7],1]",
				"[[[[4,2],2],6],[8,7]]",
			}},
			want: "[[[[8,7],[7,7]],[[8,6],[7,7]]],[[[0,7],[6,6]],[8,7]]]",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, addNodesInOrder(tt.args.forest).String(), "addNodesInOrder(%v)", tt.args.forest)
		})
	}
}
