package day19

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_parseBeacon(t *testing.T) {
	type args struct {
		s string
	}

	tests := []struct {
		name string
		args args
		want position
	}{
		{
			name: "Parses a beacon",
			args: args{s: "0,0,0"},
			want: position{
				x: 0,
				y: 0,
				z: 0,
			},
		},
		{
			name: "parses negatives",
			args: args{s: "-143,23,-98"},
			want: position{
				x: -143,
				y: 23,
				z: -98,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, parseBeacon(tt.args.s), "parseBeacon(%v)", tt.args.s)
		})
	}
}
