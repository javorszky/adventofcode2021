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
						XFrom: -10,
						XTo:   10,
						YFrom: -10,
						YTo:   10,
						ZFrom: -10,
						ZTo:   10,
						Flip:  on,
					},
					{
						XFrom: -30,
						XTo:   -20,
						YFrom: -30,
						YTo:   25,
						ZFrom: 12,
						ZTo:   43,
						Flip:  off,
					},
				},
			},
			want: instruction{
				XFrom: -30,
				XTo:   10,
				YFrom: -30,
				YTo:   25,
				ZFrom: -10,
				ZTo:   43,
				Flip:  off,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, findBoundaries(tt.args.in), "findBoundaries(%v)", tt.args.in)
		})
	}
}
