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
