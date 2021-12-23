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

func Test_parseProbes(t *testing.T) {
	type args struct {
		fn string
	}

	tests := []struct {
		name string
		args args
		want []probe
	}{
		{
			name: "tests small 3d example",
			args: args{
				fn: `--- scanner 0 ---
-1,-1,1
-2,-2,2
-3,-3,3
-2,-3,1
5,6,-4
8,0,7

--- scanner 1 ---
1,-1,1
2,-2,2
3,-3,3
2,-1,3
-5,4,-6
-8,-7,0`,
			},
			want: []probe{
				{
					name: "scanner 0",
					beacons: []position{
						{x: -1, y: -1, z: 1},
						{x: -2, y: -2, z: 2},
						{x: -3, y: -3, z: 3},
						{x: -2, y: -3, z: 1},
						{x: 5, y: 6, z: -4},
						{x: 8, y: 0, z: 7},
					},
				},
				{
					name: "scanner 1",
					beacons: []position{
						{x: 1, y: -1, z: 1},
						{x: 2, y: -2, z: 2},
						{x: 3, y: -3, z: 3},
						{x: 2, y: -1, z: 3},
						{x: -5, y: 4, z: -6},
						{x: -8, y: -7, z: 0},
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, parseProbes(tt.args.fn), "getInputs(%v)", tt.args.fn)
		})
	}
}
