package day15

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_makeMap(t *testing.T) {
	type args struct {
		input []string
	}

	tests := []struct {
		name string
		args args
		want map[int]int
	}{
		{
			name: "makes a map out of string slice",
			args: args{input: []string{
				"1234",
				"5678",
				"9012",
				"3456",
			}},
			want: map[int]int{
				0:   1,
				1:   2,
				2:   3,
				3:   4,
				100: 5,
				101: 6,
				102: 7,
				103: 8,
				200: 9,
				201: 0,
				202: 1,
				203: 2,
				300: 3,
				301: 4,
				302: 5,
				303: 6,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, makeMap(tt.args.input))
		})
	}
}

func Test_makeWalkOrder(t *testing.T) {
	type args struct {
		in   map[int]int
		edge int
	}

	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "creates a walk order for a 4x4 square",
			args: args{
				in: map[int]int{
					0:   1,
					1:   2,
					2:   3,
					3:   4,
					100: 5,
					101: 6,
					102: 7,
					103: 8,
					200: 9,
					201: 0,
					202: 1,
					203: 2,
					300: 3,
					301: 4,
					302: 5,
					303: 6,
				},
				edge: 3,
			},
			want: []int{0, 1, 100, 2, 101, 200, 3, 102, 201, 300, 103, 202, 301, 203, 302, 303},
		},
		{
			name: "creates a walk order for a 6x6 square",
			args: args{
				in: map[int]int{
					0:   1,
					1:   2,
					2:   3,
					3:   4,
					4:   3,
					5:   9,
					100: 5,
					101: 6,
					102: 7,
					103: 8,
					104: 2,
					105: 3,
					200: 9,
					201: 0,
					202: 1,
					203: 2,
					204: 0,
					205: 3,
					300: 3,
					301: 4,
					302: 5,
					303: 6,
					304: 5,
					305: 3,
					400: 3,
					401: 2,
					402: 8,
					403: 1,
					404: 0,
					405: 3,
					500: 1,
					501: 2,
					502: 3,
					503: 2,
					504: 6,
					505: 2,
				},
				edge: 5,
			},
			want: []int{
				0,
				1, 100,
				2, 101, 200,
				3, 102, 201, 300,
				4, 103, 202, 301, 400,
				5, 104, 203, 302, 401, 500,
				105, 204, 303, 402, 501,
				205, 304, 403, 502,
				305, 404, 503,
				405, 504,
				505,
			},
		},
		{
			name: "creates a walk order for a 6x6 from string slice",
			args: args{
				in: makeMap([]string{
					"983345",
					"920343",
					"993421",
					"224432",
					"002933",
					"142321",
				}),
				edge: 5,
			},
			want: []int{
				0,
				1, 100,
				2, 101, 200,
				3, 102, 201, 300,
				4, 103, 202, 301, 400,
				5, 104, 203, 302, 401, 500,
				105, 204, 303, 402, 501,
				205, 304, 403, 502,
				305, 404, 503,
				405, 504,
				505,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, makeWalkOrder(tt.args.in, tt.args.edge),
				"makeWalkOrder(%v, %v)", tt.args.in, tt.args.edge)
		})
	}
}

func Test_split(t *testing.T) {
	type args struct {
		in int
	}

	tests := []struct {
		name  string
		args  args
		want  int
		want1 int
	}{
		{
			name: "deals with 505",
			args: args{
				in: 505,
			},
			want:  5,
			want1: 5,
		},
		{
			name:  "deals with 99",
			args:  args{in: 99},
			want:  0,
			want1: 99,
		},
		{
			name:  "deals with 0",
			args:  args{in: 0},
			want:  0,
			want1: 0,
		},
		{
			name:  "deals with 100",
			args:  args{in: 100},
			want:  1,
			want1: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := split(tt.args.in)
			assert.Equalf(t, tt.want, got, "split(%v)", tt.args.in)
			assert.Equalf(t, tt.want1, got1, "split(%v)", tt.args.in)
		})
	}
}
