package day22

import (
	"fmt"
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

func Test_instruction_Lights(t *testing.T) {
	tests := []struct {
		name string
		i    instruction
		want int
	}{
		{
			name: "calculates volume for cuboid",
			i: instruction{
				xFrom: -2,
				xTo:   2,
				yFrom: -8,
				yTo:   -4,
				zFrom: 5,
				zTo:   10,
				flip:  on,
			},
			want: 150,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, tt.i.Lights(), "Lights()")
		})
	}
}

func Test_findOverlapBox(t *testing.T) {
	type args struct {
		box      instruction
		otherBox instruction
	}

	tests := []struct {
		name    string
		args    args
		want    []instruction
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name: "finds no overlapping box between the two boxes",
			args: args{
				box: instruction{
					xFrom: -10,
					xTo:   10,
					yFrom: -10,
					yTo:   10,
					zFrom: -10,
					zTo:   10,
					flip:  on,
				},
				otherBox: instruction{
					xFrom: 15,
					xTo:   20,
					yFrom: -10,
					yTo:   10,
					zFrom: -10,
					zTo:   10,
					flip:  off,
				},
			},
			want:    nil,
			wantErr: assert.NoError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			overlapBox, err := findOverlapBox(tt.args.box, tt.args.otherBox)
			if !tt.wantErr(t, err, fmt.Sprintf("findOverlapBox(%v, %v)", tt.args.box, tt.args.otherBox)) {
				return
			}
			assert.Equalf(t, tt.want, overlapBox,
				"findOverlapBox(%v, %v)", tt.args.box, tt.args.otherBox)
		})
	}
}

func Test_findTopFace(t *testing.T) {
	type args struct {
		box        instruction
		overlapBox instruction
	}

	tests := []struct {
		name string
		args args
		want []instruction
	}{
		{
			name: "returns nil if overlapbox is at the top edge",
			args: args{
				box: instruction{
					xFrom: -20,
					xTo:   20,
					yFrom: -20,
					yTo:   20,
					zFrom: -20,
					zTo:   20,
					flip:  off,
				},
				overlapBox: instruction{
					xFrom: -10,
					xTo:   10,
					yFrom: -10,
					yTo:   10,
					zFrom: 10,
					zTo:   20,
					flip:  on,
				},
			},
			want: nil,
		},
		{
			name: "returns top face box if overlapbox is not at the top edge",
			args: args{
				box: instruction{
					xFrom: -20,
					xTo:   20,
					yFrom: -20,
					yTo:   20,
					zFrom: -20,
					zTo:   20,
					flip:  off,
				},
				overlapBox: instruction{
					xFrom: -10,
					xTo:   10,
					yFrom: -10,
					yTo:   10,
					zFrom: -10,
					zTo:   10,
					flip:  on,
				},
			},
			want: []instruction{
				{
					xFrom: -10,
					xTo:   10,
					yFrom: -10,
					yTo:   10,
					zFrom: 10,
					zTo:   20,
					flip:  off,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, findTopFace(tt.args.box, tt.args.overlapBox),
				"findTopFace(%v, %v)", tt.args.box, tt.args.overlapBox)
		})
	}
}

func Test_findBottomFace(t *testing.T) {
	type args struct {
		box        instruction
		overlapBox instruction
	}

	tests := []struct {
		name string
		args args
		want []instruction
	}{
		{
			name: "returns nil if overlapbox is at the bottom edge",
			args: args{
				box: instruction{
					xFrom: -20,
					xTo:   20,
					yFrom: -20,
					yTo:   20,
					zFrom: -20,
					zTo:   20,
					flip:  off,
				},
				overlapBox: instruction{
					xFrom: -10,
					xTo:   10,
					yFrom: -10,
					yTo:   10,
					zFrom: -20,
					zTo:   -10,
					flip:  on,
				},
			},
			want: nil,
		},
		{
			name: "returns bottom face box if overlapbox is not at the bottom edge",
			args: args{
				box: instruction{
					xFrom: -20,
					xTo:   20,
					yFrom: -20,
					yTo:   20,
					zFrom: -20,
					zTo:   20,
					flip:  off,
				},
				overlapBox: instruction{
					xFrom: -10,
					xTo:   10,
					yFrom: -10,
					yTo:   10,
					zFrom: -10,
					zTo:   10,
					flip:  on,
				},
			},
			want: []instruction{
				{
					xFrom: -10,
					xTo:   10,
					yFrom: -10,
					yTo:   10,
					zFrom: -20,
					zTo:   -10,
					flip:  off,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, findBottomFace(tt.args.box, tt.args.overlapBox), "findBottomFace(%v, %v)", tt.args.box, tt.args.overlapBox)
		})
	}
}
