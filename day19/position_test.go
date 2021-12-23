package day19

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_position_String(t *testing.T) {
	type fields struct {
		x int
		y int
		z int
	}

	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "gets coordinates for position",
			fields: fields{
				x: -1,
				y: 2,
				z: -3,
			},
			want: "-1, 2, -3",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := position{
				x: tt.fields.x,
				y: tt.fields.y,
				z: tt.fields.z,
			}

			assert.Equalf(t, tt.want, p.String(), "position.string()")
		})
	}
}

func Test_position_rotations(t *testing.T) {
	type fields struct {
		x int
		y int
		z int
	}

	tests := []struct {
		name   string
		fields fields
		want   []string
	}{
		{
			name: "rotates point 24 times around 0,0,0",
			fields: fields{
				x: 3,
				y: 2,
				z: 4,
			},
			want: []string{
				"3, 2, 4", // xpzp
				"3, -4, 2",
				"3, -2, -4",
				"3, 4, -2",

				"-3, -2, 4", //xmzp
				"-3, -4, -2",
				"-3, 2, -4",
				"-3, 4, 2",

				"-2, 3, 4", // ypzp
				"-4, 3, -2",
				"2, 3, -4",
				"4, 3, 2",

				"2, -3, 4", // ymzp
				"-4, -3, 2",
				"-2, -3, -4",
				"4, -3, -2",

				"-4, 2, 3", // zpxm
				"-2, -4, 3",
				"4, -2, 3",
				"2, 4, 3",

				"4, 2, -3", // zmxp
				"-2, 4, -3",
				"-4, -2, -3",
				"2, -4, -3",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := position{
				x: tt.fields.x,
				y: tt.fields.y,
				z: tt.fields.z,
			}

			got := p.rotations()
			gotStrings := make([]string, len(got))
			for i, pos := range got {
				gotStrings[i] = pos.String()
			}

			assert.ElementsMatchf(t, tt.want, gotStrings, "position.rotations()")
		})
	}
}

func Test_distance(t *testing.T) {
	type args struct {
		p1 position
		p2 position
	}

	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "negatives",
			args: args{
				p1: position{
					x: 1,
					y: 2,
					z: 3,
				},
				p2: position{
					x: 3,
					y: 4,
					z: 5,
				},
			},
			want: math.Sqrt(12),
		},
		{
			name: "positives",
			args: args{
				p1: position{
					x: 12,
					y: 22,
					z: 53,
				},
				p2: position{
					x: 3,
					y: 4,
					z: 5,
				},
			},
			want: math.Sqrt(2709),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, distance(tt.args.p1, tt.args.p2), "distance(%v, %v)", tt.args.p1, tt.args.p2)
		})
	}
}

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
