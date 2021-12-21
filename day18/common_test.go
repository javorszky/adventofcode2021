package day18

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_leaf(t *testing.T) {
	type args struct {
		value  int
		parent *node
	}

	tests := []struct {
		name string
		args args
		want *node
	}{
		{
			name: "creates top parent node",
			args: args{
				value:  5,
				parent: nil,
			},
			want: &node{
				parent: nil,
				value:  5,
				left:   nil,
				right:  nil,
				depth:  0,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, leaf(tt.args.value, tt.args.parent), tt.want)
		})
	}
}

func Test_branch(t *testing.T) {
	type args struct {
		left   *node
		right  *node
		parent *node
	}

	tests := []struct {
		name    string
		args    args
		want    func() *node
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name: "creates a branch with no parent and both sides",
			args: args{
				left: &node{
					parent: nil,
					value:  5,
					left:   nil,
					right:  nil,
					depth:  0,
				},
				right: &node{
					parent: nil,
					value:  9,
					left:   nil,
					right:  nil,
					depth:  0,
				},
				parent: nil,
			},
			want: func() *node {
				leftNode := &node{
					parent: nil,
					value:  5,
					left:   nil,
					right:  nil,
					depth:  0,
				}

				rightNode := &node{
					parent: nil,
					value:  9,
					left:   nil,
					right:  nil,
					depth:  0,
				}
				bNode := &node{
					parent: nil,
					value:  0,
					left:   leftNode,
					right:  rightNode,
					depth:  0,
				}

				leftNode.parent = bNode
				rightNode.parent = bNode
				leftNode.depth = 1
				rightNode.depth = 1

				return bNode
			},
			wantErr: assert.NoError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := branch(tt.args.left, tt.args.right, tt.args.parent)
			if !tt.wantErr(t, err, fmt.Sprintf("branch(%v, %v, %v)", tt.args.left, tt.args.right, tt.args.parent)) {
				return
			}
			assert.Equalf(t, tt.want(), got, "branch(%v, %v, %v)", tt.args.left, tt.args.right, tt.args.parent)
		})
	}
}
