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
