package day19

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

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

func Test_shiftPositionBy(t *testing.T) {
	type args struct {
		shiftThis position
		by        position
	}

	tests := []struct {
		name string
		args args
		want position
	}{
		{
			name: "normalizes position by itself to yield 0,0,0",
			args: args{
				shiftThis: position{
					x: 43,
					y: 22,
					z: -98,
				},
				by: position{
					x: 43,
					y: 22,
					z: -98,
				},
			},
			want: position{
				x: 0,
				y: 0,
				z: 0,
			},
		},
		{
			name: "normalizes position by another position",
			args: args{
				shiftThis: position{
					x: 43,
					y: 22,
					z: -98,
				},
				by: position{
					x: -3,
					y: 254,
					z: 98,
				},
			},
			want: position{
				x: 46,
				y: -232,
				z: -196,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, shiftPositionBy(tt.args.shiftThis, tt.args.by),
				"shiftPositionBy(%v, %v)", tt.args.shiftThis, tt.args.by)
		})
	}
}

func Test_beaconsOrder(t *testing.T) {
	tests := []struct {
		name    string
		beacons beacons
		want    beacons
	}{
		{
			name: "orders beacons by their x coordinate",
			beacons: beacons{
				{x: 28, y: 44, z: 55},
				{x: 45, y: 44, z: 55},
				{x: 32, y: 44, z: 55},
				{x: 98, y: 44, z: 55},
				{x: -8, y: 44, z: 55},
				{x: 31, y: 44, z: 55},
			},
			want: beacons{
				{x: -8, y: 44, z: 55},
				{x: 28, y: 44, z: 55},
				{x: 31, y: 44, z: 55},
				{x: 32, y: 44, z: 55},
				{x: 45, y: 44, z: 55},
				{x: 98, y: 44, z: 55},
			},
		},
		{
			name: "orders beacons by their y coordinate",
			beacons: beacons{
				{x: 0, y: 93, z: 55},
				{x: 0, y: 42, z: 55},
				{x: 0, y: 4, z: 55},
				{x: 0, y: 424, z: 55},
				{x: 0, y: 92, z: 55},
				{x: 0, y: -344, z: 55},
			},
			want: beacons{
				{x: 0, y: -344, z: 55},
				{x: 0, y: 4, z: 55},
				{x: 0, y: 42, z: 55},
				{x: 0, y: 92, z: 55},
				{x: 0, y: 93, z: 55},
				{x: 0, y: 424, z: 55},
			},
		},
		{
			name: "orders beacons by their z coordinate",
			beacons: beacons{
				{x: 0, y: 0, z: 12},
				{x: 0, y: 0, z: 31},
				{x: 0, y: 0, z: -2},
				{x: 0, y: 0, z: -0},
				{x: 0, y: 0, z: 444},
				{x: 0, y: 0, z: 3},
			},
			want: beacons{
				{x: 0, y: 0, z: -2},
				{x: 0, y: 0, z: -0},
				{x: 0, y: 0, z: 3},
				{x: 0, y: 0, z: 12},
				{x: 0, y: 0, z: 31},
				{x: 0, y: 0, z: 444},
			},
		},
		{
			name: "orders beacons by their all",
			beacons: beacons{
				{x: 21, y: -4, z: 12},
				{x: 21, y: 3, z: 31},
				{x: 21, y: 5, z: -2},
				{x: 21, y: 5, z: -0},
				{x: 49, y: 833, z: 444},
				{x: -201, y: 2, z: 3},
			},
			want: beacons{
				{x: -201, y: 2, z: 3},
				{x: 21, y: -4, z: 12},
				{x: 21, y: 3, z: 31},
				{x: 21, y: 5, z: -2},
				{x: 21, y: 5, z: -0},
				{x: 49, y: 833, z: 444},
			},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			sort.Sort(test.beacons)
			assert.Exactlyf(t, test.want, test.beacons, "sort.Sort(%v)", test.beacons)
		})
	}
}

func Test_parseDistances(t *testing.T) {
	type args struct {
		beaconSlice beacons
	}

	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "parses distances for three beacons",
			args: args{beaconSlice: []position{
				{x: 3, y: 9, z: 0},
				{x: 3, y: 1, z: -5},
				{x: -3, y: -2, z: 7},
			}},
			want: []int{
				89,
				189,
				206,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, parseDistances(tt.args.beaconSlice), "parseDistances(%v)", tt.args.beaconSlice)
		})
	}
}

func Test_findCenterPoint(t *testing.T) {
	type args struct {
		beacons beacons
	}

	tests := []struct {
		name string
		args args
		want position
	}{
		{
			name: "finds centerpoint of 6 positions",
			args: args{beacons: beacons{
				{x: 90, y: 0, z: 0},
				{x: -90, y: 0, z: 0},
				{x: 0, y: -90, z: 0},
				{x: 0, y: 90, z: 0},
				{x: 0, y: 0, z: -90},
				{x: 0, y: 0, z: 90},
			}},
			want: position{
				x: 0,
				y: 0,
				z: 0,
			},
		},
		{
			name: "finds centerpoint of 3 positions",
			args: args{beacons: beacons{
				{x: 12, y: 34, z: -9},
				{x: -90, y: 17, z: 21},
				{x: 7, y: -56, z: 53},
			}},
			want: position{
				x: -23,
				y: -1,
				z: 21,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, findCenterPoint(tt.args.beacons), "findCenterPoint(%v)", tt.args.beacons)
		})
	}
}

func Test_parseDistancesFromCenterpoint(t *testing.T) {
	type args struct {
		beaconSlice beacons
		transform   func(beacons) beacons
	}

	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "parses distances from center point",
			args: args{
				beaconSlice: beacons{
					{x: 5, y: 5, z: 5},
					{x: 4, y: 2, z: -4},
					{x: 1, y: -5, z: 0},
					{x: -10, y: 3, z: -1},
				},
				transform: func(b beacons) beacons {
					return b
				},
			},
			want: []int{
				59,
				32,
				50,
				102,
			},
		},
		{
			name: "parses distances from center point shifted by some value",
			args: args{
				beaconSlice: beacons{
					{x: 5, y: 5, z: 5},
					{x: 4, y: 2, z: -4},
					{x: 1, y: -5, z: 0},
					{x: -10, y: 3, z: -1},
				},
				transform: func(b beacons) beacons {
					p := position{
						x: -324,
						y: 992,
						z: 45,
					}

					newBeacons := make(beacons, len(b))
					for i, oldB := range b {
						newBeacons[i] = shiftPositionBy(oldB, p)
					}

					return newBeacons
				},
			},
			want: []int{
				59,
				32,
				50,
				102,
			},
		},
		{
			name: "parses distances from center point shifted and rotated by some value",
			args: args{
				beaconSlice: beacons{
					{x: 5, y: 5, z: 5},
					{x: 4, y: 2, z: -4},
					{x: 1, y: -5, z: 0},
					{x: -10, y: 3, z: -1},
				},
				transform: func(b beacons) beacons {
					p := position{
						x: -324,
						y: 992,
						z: 45,
					}

					newBeacons := make(beacons, len(b))
					for i, oldB := range b {
						newBeacons[i] = shiftPositionBy(oldB, p)
					}

					newBeaconsRotated := make(beacons, len(b))
					for i, oldB := range newBeacons {
						rs := oldB.rotations()
						newBeaconsRotated[i] = rs[9]
					}

					return newBeaconsRotated
				},
			},
			want: []int{
				59,
				32,
				50,
				102,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.ElementsMatchf(t, tt.want, parseDistancesFromCenterpoint(tt.args.beaconSlice),
				"parseDistancesFromCenterpoint(%v)", tt.args.beaconSlice)
		})
	}
}
