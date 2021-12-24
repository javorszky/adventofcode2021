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
			name: "tests baby 3d example",
			args: args{
				fn: `--- scanner 0 ---
9,2,-3
8,5,1
-2,6,3
0,0,11`,
			},
			want: []probe{
				{
					name: "scanner 0",
					beacons: beacons{
						{x: 9, y: 2, z: -3},
						{x: 8, y: 5, z: 1},
						{x: -2, y: 6, z: 3},
						{x: 0, y: 0, z: 11},
					},
					distances: map[int][][2]position{
						26: {
							{
								{x: 9, y: 2, z: -3},
								{x: 8, y: 5, z: 1},
							},
						},
						173: {
							{
								{x: 9, y: 2, z: -3},
								{x: -2, y: 6, z: 3},
							},
						},
						281: {
							{
								{x: 9, y: 2, z: -3},
								{x: 0, y: 0, z: 11},
							},
						},
						189: {
							{
								{x: 8, y: 5, z: 1},
								{x: 0, y: 0, z: 11},
							},
						},
						105: {
							{
								{x: 8, y: 5, z: 1},
								{x: -2, y: 6, z: 3},
							},
						},
						104: {
							{
								{x: -2, y: 6, z: 3},
								{x: 0, y: 0, z: 11},
							},
						},
					},
				},
			},
		},
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
					distances: map[int][][2]position{
						5: {
							{
								// double checked
								{x: -1, y: -1, z: 1},
								{x: -2, y: -3, z: 1},
							},
							{
								// double checked
								{x: -3, y: -3, z: 3},
								{x: -2, y: -3, z: 1},
							},
						},
						3: {
							{
								// double checked
								{x: -1, y: -1, z: 1},
								{x: -2, y: -2, z: 2},
							},
							{
								// double checked
								{x: -2, y: -2, z: 2},
								{x: -3, y: -3, z: 3},
							},
						},
						12: {
							{
								// double checked
								{x: -1, y: -1, z: 1},
								{x: -3, y: -3, z: 3},
							},
						},
						194: {
							{
								// double checked
								{x: -3, y: -3, z: 3},
								{x: 5, y: 6, z: -4},
							},
						},
						146: {
							{
								// double checked
								{x: -3, y: -3, z: 3},
								{x: 8, y: 0, z: 7},
							},
						},
						145: {
							{
								// double checked
								{x: -2, y: -3, z: 1},
								{x: 8, y: 0, z: 7},
							},
						},
						2: {
							{
								// double checked
								{x: -2, y: -2, z: 2},
								{x: -2, y: -3, z: 1},
							},
						},
						155: {
							{
								// double checked
								{x: -2, y: -3, z: 1},
								{x: 5, y: 6, z: -4},
							},
						},
						149: {
							{
								// double checked
								{x: -2, y: -2, z: 2},
								{x: 5, y: 6, z: -4},
							},
						},
						129: {
							{
								// double checked
								{x: -2, y: -2, z: 2},
								{x: 8, y: 0, z: 7},
							},
						},
						118: {
							{
								// double checked
								{x: -1, y: -1, z: 1},
								{x: 8, y: 0, z: 7},
							},
						},
						110: {
							{
								// double checked
								{x: -1, y: -1, z: 1},
								{x: 5, y: 6, z: -4},
							},
						},
						166: {
							{
								// double checked
								{x: 5, y: 6, z: -4},
								{x: 8, y: 0, z: 7},
							},
						},
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
					distances: map[int][][2]position{
						3: {
							{
								// double checked
								{x: 1, y: -1, z: 1},
								{x: 2, y: -2, z: 2},
							},
							{
								// double checked
								{x: 2, y: -2, z: 2},
								{x: 3, y: -3, z: 3},
							},
						},
						12: {
							{
								// double checked
								{x: 1, y: -1, z: 1},
								{x: 3, y: -3, z: 3},
							},
						},
						5: {
							{
								// double checked
								{x: 1, y: -1, z: 1},
								{x: 2, y: -1, z: 3},
							},
							{
								// double checked
								{x: 3, y: -3, z: 3},
								{x: 2, y: -1, z: 3},
							},
						},
						110: {
							{
								// double checked
								{x: 1, y: -1, z: 1},
								{x: -5, y: 4, z: -6},
							},
						},
						118: {
							{
								// double checked
								{x: 1, y: -1, z: 1},
								{x: -8, y: -7, z: 0},
							},
						},
						2: {
							{
								// double checked
								{x: 2, y: -2, z: 2},
								{x: 2, y: -1, z: 3},
							},
						},
						149: {
							{
								// double checked
								{x: 2, y: -2, z: 2},
								{x: -5, y: 4, z: -6},
							},
						},
						129: {
							{
								// double checked
								{x: 2, y: -2, z: 2},
								{x: -8, y: -7, z: 0},
							},
						},
						146: {
							{
								// double checked
								{x: 3, y: -3, z: 3},
								{x: -8, y: -7, z: 0},
							},
						},
						194: {
							{
								// double checked
								{x: 3, y: -3, z: 3},
								{x: -5, y: 4, z: -6},
							},
						},
						155: {
							{
								// double checked
								{x: 2, y: -1, z: 3},
								{x: -5, y: 4, z: -6},
							},
						},
						145: {
							{
								// double checked
								{x: 2, y: -1, z: 3},
								{x: -8, y: -7, z: 0},
							},
						},
						166: {
							{
								// double checked
								{x: -5, y: 4, z: -6},
								{x: -8, y: -7, z: 0},
							},
						},
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
		want map[int][][2]position
	}{
		{
			name: "parses distances for three beacons",
			args: args{beaconSlice: []position{
				{x: 3, y: 9, z: 0},
				{x: 3, y: 1, z: -5},
				{x: -3, y: -2, z: 7},
			}},
			want: map[int][][2]position{
				89: {
					{
						{x: 3, y: 9, z: 0},
						{x: 3, y: 1, z: -5},
					},
				},
				189: {
					{
						{x: 3, y: 1, z: -5},
						{x: -3, y: -2, z: 7},
					},
				},
				206: {
					{
						{x: 3, y: 9, z: 0},
						{x: -3, y: -2, z: 7},
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, parseDistances(tt.args.beaconSlice),
				"parseDistances(%v)", tt.args.beaconSlice)
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
