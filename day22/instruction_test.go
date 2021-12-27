package day22

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_parseInstruction(t *testing.T) {
	type args struct {
		s string
	}

	tests := []struct {
		name string
		args args
		want instruction
	}{
		{
			name: "parses string correctly",
			args: args{
				s: "on x=967..23432,y=45373..81175,z=27513..53682",
			},
			want: instruction{
				xFrom: 967,
				xTo:   23432,
				yFrom: 45373,
				yTo:   81175,
				zFrom: 27513,
				zTo:   53682,
				flip:  on,
			},
		},
		{
			name: "parses string with negative values correctly",
			args: args{
				s: "on x=-54112..-39298,y=-85059..-49293,z=-27449..7877",
			},
			want: instruction{
				xFrom: -54112,
				xTo:   -39298,
				yFrom: -85059,
				yTo:   -49293,
				zFrom: -27449,
				zTo:   7877,
				flip:  on,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, parseInstruction(tt.args.s))
		})
	}
}
