package day22

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_findBoundaries(t *testing.T) {
	type args struct {
		in []instruction
	}

	tests := []struct {
		name string
		args args
		want instruction
	}{
		{
			name: "calculates total cubespace",
			args: args{
				in: []instruction{
					{
						xFrom: -10,
						xTo:   10,
						yFrom: -10,
						yTo:   10,
						zFrom: -10,
						zTo:   10,
						flip:  on,
					},
					{
						xFrom: -30,
						xTo:   -20,
						yFrom: -30,
						yTo:   25,
						zFrom: 12,
						zTo:   43,
						flip:  off,
					},
				},
			},
			want: instruction{
				xFrom: -30,
				xTo:   10,
				yFrom: -30,
				yTo:   25,
				zFrom: -10,
				zTo:   43,
				flip:  off,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, findBoundaries(tt.args.in), "findBoundaries(%v)", tt.args.in)
		})
	}
}
