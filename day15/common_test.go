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
