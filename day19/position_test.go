package day19

import (
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
				x: 1,
				y: 2,
				z: 3,
			},
			want: []string{
				"1, 2, 3", // xpzp
				"1, 3, -2",
				"1, -2, -3",
				"1, -3, 2",

				"-1, -2, 3", // xmzp
				"-1, -3, -2",
				"-1, 2, -3",
				"-1, 3, 2",

				"-2, 1, 3", // ypzp
				"-3, 1, -2",
				"2, 1, -3",
				"3, 1, 2",

				"2, -1, 3", // ymzp
				"-3, -1, 2",
				"-2, -1, -3",
				"3, -1, -1",

				"-3, 2, 1", // zpxm
				"2, 3, 1",
				"3, -2, 1",
				"-2, -3, -1",

				"3, 2, 1", // zmxp
				"2, -3, -1",
				"-3, -2, -1",
				"-2, 3, -1",
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
