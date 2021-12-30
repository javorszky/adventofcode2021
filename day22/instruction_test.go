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
				xFrom: 1,
				xTo:   1,
				yFrom: 1,
				yTo:   1,
				zFrom: 1,
				zTo:   1,
				flip:  on,
			},
			want: 1,
		},
		{
			name: "calculates volume for cuboid",
			i: instruction{
				xFrom: 1,
				xTo:   3,
				yFrom: 1,
				yTo:   3,
				zFrom: 1,
				zTo:   3,
				flip:  on,
			},
			want: 27,
		},
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
		want    instruction
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
			want:    instruction{},
			wantErr: assert.Error,
		},
		{
			name: "finds overlapping box between the two boxes where box is fully contained",
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
					xFrom: -20,
					xTo:   20,
					yFrom: -20,
					yTo:   20,
					zFrom: -20,
					zTo:   20,
					flip:  off,
				},
			},
			want: instruction{
				xFrom: -10,
				xTo:   10,
				yFrom: -10,
				yTo:   10,
				zFrom: -10,
				zTo:   10,
				flip:  off,
			},
			wantErr: assert.NoError,
		},
		{
			name: "finds overlapping box between the two boxes where only a corner matches",
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
					xFrom: 5,
					xTo:   15,
					yFrom: 5,
					yTo:   15,
					zFrom: 5,
					zTo:   15,
					flip:  off,
				},
			},
			want: instruction{
				xFrom: 5,
				xTo:   10,
				yFrom: 5,
				yTo:   10,
				zFrom: 5,
				zTo:   10,
				flip:  off,
			},
			wantErr: assert.NoError,
		},
		{
			name: "finds overlapping box between the two boxes where only a single cube",
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
					xFrom: 10,
					xTo:   12,
					yFrom: 10,
					yTo:   12,
					zFrom: 10,
					zTo:   12,
					flip:  off,
				},
			},
			want: instruction{
				xFrom: 10,
				xTo:   10,
				yFrom: 10,
				yTo:   10,
				zFrom: 10,
				zTo:   10,
				flip:  off,
			},
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
					zFrom: 11,
					zTo:   20,
					flip:  off,
				},
			},
		},
		{
			name: "returns a single layer of top face",
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
					zTo:   19,
					flip:  on,
				},
			},
			want: []instruction{
				{
					xFrom: -10,
					xTo:   10,
					yFrom: -10,
					yTo:   10,
					zFrom: 20,
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
					zTo:   -11,
					flip:  off,
				},
			},
		},
		{
			name: "returns a single layer bottom face",
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
					zFrom: -19,
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
					zTo:   -20,
					flip:  off,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, findBottomFace(tt.args.box, tt.args.overlapBox),
				"findBottomFace(%v, %v)", tt.args.box, tt.args.overlapBox)
		})
	}
}

func Test_findFrontFace(t *testing.T) {
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
			name: "returns nil if overlapbox is at the front edge",
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
					xFrom: 10,
					xTo:   20,
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
			name: "returns front face box if overlapbox is not at the front edge",
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
					xFrom: 11,
					xTo:   20,
					yFrom: -10,
					yTo:   10,
					zFrom: -10,
					zTo:   10,
					flip:  off,
				},
			},
		},
		{
			name: "returns single layer front face box",
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
					xTo:   19,
					yFrom: -10,
					yTo:   10,
					zFrom: -10,
					zTo:   10,
					flip:  on,
				},
			},
			want: []instruction{
				{
					xFrom: 20,
					xTo:   20,
					yFrom: -10,
					yTo:   10,
					zFrom: -10,
					zTo:   10,
					flip:  off,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, findFrontFace(tt.args.box, tt.args.overlapBox),
				"findFrontFace(%v, %v)", tt.args.box, tt.args.overlapBox)
		})
	}
}

func Test_findBackFace(t *testing.T) {
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
			name: "returns nil if overlapbox is at the back edge",
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
					xFrom: -20,
					xTo:   -10,
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
			name: "returns back face box if overlapbox is not at the back edge",
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
					xFrom: -20,
					xTo:   -11,
					yFrom: -10,
					yTo:   10,
					zFrom: -10,
					zTo:   10,
					flip:  off,
				},
			},
		},
		{
			name: "returns single layer back face",
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
					xFrom: -19,
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
					xFrom: -20,
					xTo:   -20,
					yFrom: -10,
					yTo:   10,
					zFrom: -10,
					zTo:   10,
					flip:  off,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, findBackFace(tt.args.box, tt.args.overlapBox),
				"findBackFace(%v, %v)", tt.args.box, tt.args.overlapBox)
		})
	}
}

func Test_findLeftFace(t *testing.T) {
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
			name: "returns nil if overlapbox is at the left edge",
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
					yFrom: -20,
					yTo:   -10,
					zFrom: -10,
					zTo:   10,
					flip:  on,
				},
			},
			want: nil,
		},
		{
			name: "returns left face box if overlapbox is not at the left edge",
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
					yFrom: -20,
					yTo:   -11,
					zFrom: -10,
					zTo:   10,
					flip:  off,
				},
			},
		},
		{
			name: "returns single layer left face",
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
					yFrom: -19,
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
					yFrom: -20,
					yTo:   -20,
					zFrom: -10,
					zTo:   10,
					flip:  off,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, findLeftFace(tt.args.box, tt.args.overlapBox),
				"findLeftFace(%v, %v)", tt.args.box, tt.args.overlapBox)
		})
	}
}

func Test_findRightFace(t *testing.T) {
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
			name: "returns nil if overlapbox is at the right edge",
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
					yFrom: 10,
					yTo:   20,
					zFrom: -10,
					zTo:   10,
					flip:  on,
				},
			},
			want: nil,
		},
		{
			name: "returns right face box if overlapbox is not at the right edge",
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
					yFrom: 11,
					yTo:   20,
					zFrom: -10,
					zTo:   10,
					flip:  off,
				},
			},
		},
		{
			name: "returns single layer right face",
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
					yTo:   19,
					zFrom: -10,
					zTo:   10,
					flip:  on,
				},
			},
			want: []instruction{
				{
					xFrom: -10,
					xTo:   10,
					yFrom: 20,
					yTo:   20,
					zFrom: -10,
					zTo:   10,
					flip:  off,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, findRightFace(tt.args.box, tt.args.overlapBox),
				"findRightFace(%v, %v)", tt.args.box, tt.args.overlapBox)
		})
	}
}

func Test_findTopLeftEdge(t *testing.T) {
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
			name: "returns nil if overlapbox is at the left edge",
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
					yFrom: -20,
					yTo:   10,
					zFrom: -10,
					zTo:   10,
					flip:  on,
				},
			},
			want: nil,
		},
		{
			name: "returns nil if overlapbox is at both the top and the left edge",
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
					yFrom: -20,
					yTo:   10,
					zFrom: -10,
					zTo:   20,
					flip:  on,
				},
			},
			want: nil,
		},
		{
			name: "returns top left edge box if overlapbox is not at the top or left edges",
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
					yFrom: -20,
					yTo:   -11,
					zFrom: 11,
					zTo:   20,
					flip:  off,
				},
			},
		},
		{
			name: "returns single cube top left edge box",
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
					yFrom: -19,
					yTo:   10,
					zFrom: -10,
					zTo:   19,
					flip:  on,
				},
			},
			want: []instruction{
				{
					xFrom: -10,
					xTo:   10,
					yFrom: -20,
					yTo:   -20,
					zFrom: 20,
					zTo:   20,
					flip:  off,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, findTopLeftEdge(tt.args.box, tt.args.overlapBox),
				"findTopLeftEdge(%v, %v)", tt.args.box, tt.args.overlapBox)
		})
	}
}

func Test_findTopBackEdge(t *testing.T) {
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
			name: "returns nil if overlapbox is at the back edge",
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
					xFrom: -20,
					xTo:   10,
					yFrom: -10,
					yTo:   10,
					zFrom: -10,
					zTo:   10,
					flip:  on,
				},
			},
			want: nil,
		},
		{
			name: "returns nil if overlapbox is at both the top and the back edge",
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
					xFrom: -20,
					xTo:   10,
					yFrom: -10,
					yTo:   10,
					zFrom: -10,
					zTo:   20,
					flip:  on,
				},
			},
			want: nil,
		},
		{
			name: "returns top right edge box if overlapbox is not at the top or back edge",
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
					xFrom: -20,
					xTo:   -11,
					yFrom: -10,
					yTo:   10,
					zFrom: 11,
					zTo:   20,
					flip:  off,
				},
			},
		},
		{
			name: "returns single line top right edge box",
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
					xFrom: -19,
					xTo:   10,
					yFrom: -10,
					yTo:   10,
					zFrom: -10,
					zTo:   19,
					flip:  on,
				},
			},
			want: []instruction{
				{
					xFrom: -20,
					xTo:   -20,
					yFrom: -10,
					yTo:   10,
					zFrom: 20,
					zTo:   20,
					flip:  off,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, findTopBackEdge(tt.args.box, tt.args.overlapBox),
				"findTopBackEdge(%v, %v)", tt.args.box, tt.args.overlapBox)
		})
	}
}

func Test_findTopRightEdge(t *testing.T) {
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
			name: "returns nil if overlapbox is at the right edge",
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
					yTo:   20,
					zFrom: -10,
					zTo:   10,
					flip:  on,
				},
			},
			want: nil,
		},
		{
			name: "returns nil if overlapbox is at both the top and the right edge",
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
					yTo:   20,
					zFrom: -10,
					zTo:   20,
					flip:  on,
				},
			},
			want: nil,
		},
		{
			name: "returns top left edge box if overlapbox is not at the top or left edges",
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
					yFrom: 11,
					yTo:   20,
					zFrom: 11,
					zTo:   20,
					flip:  off,
				},
			},
		},
		{
			name: "returns single line top left edge box",
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
					yTo:   19,
					zFrom: -10,
					zTo:   19,
					flip:  on,
				},
			},
			want: []instruction{
				{
					xFrom: -10,
					xTo:   10,
					yFrom: 20,
					yTo:   20,
					zFrom: 20,
					zTo:   20,
					flip:  off,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, findTopRightEdge(tt.args.box, tt.args.overlapBox),
				"findTopRightEdge(%v, %v)", tt.args.box, tt.args.overlapBox)
		})
	}
}

func Test_findTopFrontEdge(t *testing.T) {
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
			name: "returns nil if overlapbox is at the front edge",
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
					xTo:   20,
					yFrom: -10,
					yTo:   10,
					zFrom: -10,
					zTo:   10,
					flip:  on,
				},
			},
			want: nil,
		},
		{
			name: "returns nil if overlapbox is at both the top and the front edge",
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
					xTo:   20,
					yFrom: -10,
					yTo:   10,
					zFrom: -10,
					zTo:   20,
					flip:  on,
				},
			},
			want: nil,
		},
		{
			name: "returns top right edge box if overlapbox is not at the top or front edge",
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
					xFrom: 11,
					xTo:   20,
					yFrom: -10,
					yTo:   10,
					zFrom: 11,
					zTo:   20,
					flip:  off,
				},
			},
		},
		{
			name: "returns single line top right edge box",
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
					xTo:   19,
					yFrom: -10,
					yTo:   10,
					zFrom: -10,
					zTo:   19,
					flip:  on,
				},
			},
			want: []instruction{
				{
					xFrom: 20,
					xTo:   20,
					yFrom: -10,
					yTo:   10,
					zFrom: 20,
					zTo:   20,
					flip:  off,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, findTopFrontEdge(tt.args.box, tt.args.overlapBox),
				"findTopFrontEdge(%v, %v)", tt.args.box, tt.args.overlapBox)
		})
	}
}

func Test_findBottomLeftEdge(t *testing.T) {
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
					zTo:   10,
					flip:  on,
				},
			},
			want: nil,
		},
		{
			name: "returns nil if overlapbox is at the left edge",
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
					yFrom: -20,
					yTo:   10,
					zFrom: -10,
					zTo:   10,
					flip:  on,
				},
			},
			want: nil,
		},
		{
			name: "returns nil if overlapbox is at both the bottom and the left edge",
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
					yFrom: -20,
					yTo:   10,
					zFrom: -20,
					zTo:   10,
					flip:  on,
				},
			},
			want: nil,
		},
		{
			name: "returns bottom left edge box if overlapbox is not at the bottom or left edges",
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
					yFrom: -20,
					yTo:   -11,
					zFrom: -20,
					zTo:   -11,
					flip:  off,
				},
			},
		},
		{
			name: "returns single line bottom left edge",
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
					yFrom: -19,
					yTo:   10,
					zFrom: -19,
					zTo:   10,
					flip:  on,
				},
			},
			want: []instruction{
				{
					xFrom: -10,
					xTo:   10,
					yFrom: -20,
					yTo:   -20,
					zFrom: -20,
					zTo:   -20,
					flip:  off,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, findBottomLeftEdge(tt.args.box, tt.args.overlapBox),
				"findBottomLeftEdge(%v, %v)", tt.args.box, tt.args.overlapBox)
		})
	}
}

func Test_findBottomBackEdge(t *testing.T) {
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
					zTo:   10,
					flip:  on,
				},
			},
			want: nil,
		},
		{
			name: "returns nil if overlapbox is at the back edge",
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
					xFrom: -20,
					xTo:   10,
					yFrom: -10,
					yTo:   10,
					zFrom: -10,
					zTo:   10,
					flip:  on,
				},
			},
			want: nil,
		},
		{
			name: "returns nil if overlapbox is at both the bottom and the back edge",
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
					xFrom: -20,
					xTo:   10,
					yFrom: -10,
					yTo:   10,
					zFrom: -20,
					zTo:   10,
					flip:  on,
				},
			},
			want: nil,
		},
		{
			name: "returns bottom right edge box if overlapbox is not at the bottom or back edge",
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
					xFrom: -20,
					xTo:   -11,
					yFrom: -10,
					yTo:   10,
					zFrom: -20,
					zTo:   -11,
					flip:  off,
				},
			},
		},
		{
			name: "returns single line bottom right edge",
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
					xFrom: -19,
					xTo:   10,
					yFrom: -10,
					yTo:   10,
					zFrom: -19,
					zTo:   10,
					flip:  on,
				},
			},
			want: []instruction{
				{
					xFrom: -20,
					xTo:   -20,
					yFrom: -10,
					yTo:   10,
					zFrom: -20,
					zTo:   -20,
					flip:  off,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, findBottomBackEdge(tt.args.box, tt.args.overlapBox),
				"findBottomBackEdge(%v, %v)", tt.args.box, tt.args.overlapBox)
		})
	}
}

func Test_findBottomRightEdge(t *testing.T) {
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
					zTo:   10,
					flip:  on,
				},
			},
			want: nil,
		},
		{
			name: "returns nil if overlapbox is at the right edge",
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
					yTo:   20,
					zFrom: -10,
					zTo:   10,
					flip:  on,
				},
			},
			want: nil,
		},
		{
			name: "returns nil if overlapbox is at both the bottom and the right edge",
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
					yTo:   20,
					zFrom: -20,
					zTo:   10,
					flip:  on,
				},
			},
			want: nil,
		},
		{
			name: "returns bottom left edge box if overlapbox is not at the bottom or left edges",
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
					yFrom: 10,
					yTo:   20,
					zFrom: -20,
					zTo:   -10,
					flip:  off,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, findBottomRightEdge(tt.args.box, tt.args.overlapBox),
				"findBottomRightEdge(%v, %v)", tt.args.box, tt.args.overlapBox)
		})
	}
}

func Test_findBottomFrontEdge(t *testing.T) {
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
					zTo:   10,
					flip:  on,
				},
			},
			want: nil,
		},
		{
			name: "returns nil if overlapbox is at the front edge",
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
					xTo:   20,
					yFrom: -10,
					yTo:   10,
					zFrom: -10,
					zTo:   10,
					flip:  on,
				},
			},
			want: nil,
		},
		{
			name: "returns nil if overlapbox is at both the bottom and the front edge",
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
					xTo:   20,
					yFrom: -10,
					yTo:   10,
					zFrom: -20,
					zTo:   10,
					flip:  on,
				},
			},
			want: nil,
		},
		{
			name: "returns bottom right edge box if overlapbox is not at the bottom or front edge",
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
					xFrom: 10,
					xTo:   20,
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
			assert.Equalf(t, tt.want, findBottomFrontEdge(tt.args.box, tt.args.overlapBox),
				"findBottomFrontEdge(%v, %v)", tt.args.box, tt.args.overlapBox)
		})
	}
}

func Test_findFrontLeftEdge(t *testing.T) {
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
			name: "returns nil if overlapbox is at the front edge",
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
					xTo:   20,
					yFrom: -10,
					yTo:   10,
					zFrom: -10,
					zTo:   10,
					flip:  on,
				},
			},
			want: nil,
		},
		{
			name: "returns nil if overlapbox is at the left edge",
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
					yFrom: -20,
					yTo:   10,
					zFrom: -10,
					zTo:   10,
					flip:  on,
				},
			},
			want: nil,
		},
		{
			name: "returns nil if overlapbox is at both the front and the left edge",
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
					xTo:   20,
					yFrom: -20,
					yTo:   10,
					zFrom: -10,
					zTo:   10,
					flip:  on,
				},
			},
			want: nil,
		},
		{
			name: "returns front left edge box if overlapbox is not at the front or left edges",
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
					xFrom: 10,
					xTo:   20,
					yFrom: -20,
					yTo:   -10,
					zFrom: -10,
					zTo:   10,
					flip:  off,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, findFrontLeftEdge(tt.args.box, tt.args.overlapBox),
				"findFrontLeftEdge(%v, %v)", tt.args.box, tt.args.overlapBox)
		})
	}
}

func Test_findFrontRightEdge(t *testing.T) {
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
			name: "returns nil if overlapbox is at the front edge",
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
					xTo:   20,
					yFrom: -10,
					yTo:   10,
					zFrom: -10,
					zTo:   10,
					flip:  on,
				},
			},
			want: nil,
		},
		{
			name: "returns nil if overlapbox is at the right edge",
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
					yTo:   20,
					zFrom: -10,
					zTo:   10,
					flip:  on,
				},
			},
			want: nil,
		},
		{
			name: "returns nil if overlapbox is at both the front and the right edge",
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
					xTo:   20,
					yFrom: -10,
					yTo:   20,
					zFrom: -10,
					zTo:   10,
					flip:  on,
				},
			},
			want: nil,
		},
		{
			name: "returns front right edge box if overlapbox is not at the front or right edges",
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
					xFrom: 10,
					xTo:   20,
					yFrom: 10,
					yTo:   20,
					zFrom: -10,
					zTo:   10,
					flip:  off,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, findFrontRightEdge(tt.args.box, tt.args.overlapBox),
				"findFrontRightEdge(%v, %v)", tt.args.box, tt.args.overlapBox)
		})
	}
}

func Test_findBackLeftEdge(t *testing.T) {
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
			name: "returns nil if overlapbox is at the back edge",
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
					xFrom: -20,
					xTo:   10,
					yFrom: -10,
					yTo:   10,
					zFrom: -10,
					zTo:   10,
					flip:  on,
				},
			},
			want: nil,
		},
		{
			name: "returns nil if overlapbox is at the left edge",
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
					yFrom: -20,
					yTo:   10,
					zFrom: -10,
					zTo:   10,
					flip:  on,
				},
			},
			want: nil,
		},
		{
			name: "returns nil if overlapbox is at both the back and the left edge",
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
					xFrom: -20,
					xTo:   10,
					yFrom: -20,
					yTo:   10,
					zFrom: -10,
					zTo:   10,
					flip:  on,
				},
			},
			want: nil,
		},
		{
			name: "returns back left edge box if overlapbox is not at the back or left edges",
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
					xFrom: -20,
					xTo:   -10,
					yFrom: -20,
					yTo:   -10,
					zFrom: -10,
					zTo:   10,
					flip:  off,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, findBackLeftEdge(tt.args.box, tt.args.overlapBox),
				"findBackLeftEdge(%v, %v)", tt.args.box, tt.args.overlapBox)
		})
	}
}

func Test_findBackRightEdge(t *testing.T) {
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
			name: "returns nil if overlapbox is at the back edge",
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
					xFrom: -20,
					xTo:   10,
					yFrom: -10,
					yTo:   10,
					zFrom: -10,
					zTo:   10,
					flip:  on,
				},
			},
			want: nil,
		},
		{
			name: "returns nil if overlapbox is at the right edge",
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
					yTo:   20,
					zFrom: -10,
					zTo:   10,
					flip:  on,
				},
			},
			want: nil,
		},
		{
			name: "returns nil if overlapbox is at both the back and the right edge",
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
					xFrom: -20,
					xTo:   10,
					yFrom: -10,
					yTo:   20,
					zFrom: -10,
					zTo:   10,
					flip:  on,
				},
			},
			want: nil,
		},
		{
			name: "returns back right edge box if overlapbox is not at the back or right edges",
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
					xFrom: -20,
					xTo:   -10,
					yFrom: 10,
					yTo:   20,
					zFrom: -10,
					zTo:   10,
					flip:  off,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, findBackRightEdge(tt.args.box, tt.args.overlapBox),
				"findBackRightEdge(%v, %v)", tt.args.box, tt.args.overlapBox)
		})
	}
}

func Test_findTopBackLeftCorner(t *testing.T) {
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
			name: "returns nil if overlapbox is at the back edge",
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
					xFrom: -20,
					xTo:   10,
					yFrom: -10,
					yTo:   10,
					zFrom: -10,
					zTo:   10,
					flip:  on,
				},
			},
			want: nil,
		},
		{
			name: "returns nil if overlapbox is at the left edge",
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
					yFrom: -20,
					yTo:   10,
					zFrom: -10,
					zTo:   10,
					flip:  on,
				},
			},
			want: nil,
		},
		{
			name: "returns nil if overlapbox is at both the back and the left edge",
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
					xFrom: -20,
					xTo:   10,
					yFrom: -20,
					yTo:   10,
					zFrom: -10,
					zTo:   10,
					flip:  on,
				},
			},
			want: nil,
		},
		{
			name: "returns nil if overlapbox is at both the top and the left edge",
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
					yFrom: -20,
					yTo:   10,
					zFrom: -10,
					zTo:   20,
					flip:  on,
				},
			},
			want: nil,
		},
		{
			name: "returns nil if overlapbox is at both the top and the back edge",
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
					xFrom: -20,
					xTo:   10,
					yFrom: -10,
					yTo:   10,
					zFrom: -10,
					zTo:   20,
					flip:  on,
				},
			},
			want: nil,
		},
		{
			name: "returns top back right corner box if overlapbox is not at the top, back, or left edges",
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
					xFrom: -20,
					xTo:   -10,
					yFrom: -20,
					yTo:   -10,
					zFrom: 10,
					zTo:   20,
					flip:  off,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, findTopBackLeftCorner(tt.args.box, tt.args.overlapBox),
				"findTopBackLeftCorner(%v, %v)", tt.args.box, tt.args.overlapBox)
		})
	}
}

func Test_findTopBackRightCorner(t *testing.T) {
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
			name: "returns nil if overlapbox is at the back edge",
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
					xFrom: -20,
					xTo:   10,
					yFrom: -10,
					yTo:   10,
					zFrom: -10,
					zTo:   10,
					flip:  on,
				},
			},
			want: nil,
		},
		{
			name: "returns nil if overlapbox is at the right edge",
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
					yTo:   20,
					zFrom: -10,
					zTo:   10,
					flip:  on,
				},
			},
			want: nil,
		},
		{
			name: "returns nil if overlapbox is at both the back and the right edge",
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
					xFrom: -20,
					xTo:   10,
					yFrom: -10,
					yTo:   20,
					zFrom: -10,
					zTo:   10,
					flip:  on,
				},
			},
			want: nil,
		},
		{
			name: "returns nil if overlapbox is at both the top and the right edge",
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
					yTo:   20,
					zFrom: -10,
					zTo:   20,
					flip:  on,
				},
			},
			want: nil,
		},
		{
			name: "returns nil if overlapbox is at both the top and the back edge",
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
					xFrom: -20,
					xTo:   10,
					yFrom: -10,
					yTo:   10,
					zFrom: -10,
					zTo:   20,
					flip:  on,
				},
			},
			want: nil,
		},
		{
			name: "returns top back right corner box if overlapbox is not at the top, back, or right edges",
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
					xFrom: -20,
					xTo:   -10,
					yFrom: 10,
					yTo:   20,
					zFrom: 10,
					zTo:   20,
					flip:  off,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, findTopBackRightCorner(tt.args.box, tt.args.overlapBox),
				"findTopBackRightCorner(%v, %v)", tt.args.box, tt.args.overlapBox)
		})
	}
}

func Test_findTopFrontLeftCorner(t *testing.T) {
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
			name: "returns nil if overlapbox is at the front edge",
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
					xTo:   20,
					yFrom: -10,
					yTo:   10,
					zFrom: -10,
					zTo:   10,
					flip:  on,
				},
			},
			want: nil,
		},
		{
			name: "returns nil if overlapbox is at the left edge",
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
					yFrom: -20,
					yTo:   10,
					zFrom: -10,
					zTo:   10,
					flip:  on,
				},
			},
			want: nil,
		},
		{
			name: "returns nil if overlapbox is at both the front and the left edge",
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
					xTo:   20,
					yFrom: -20,
					yTo:   10,
					zFrom: -10,
					zTo:   10,
					flip:  on,
				},
			},
			want: nil,
		},
		{
			name: "returns nil if overlapbox is at both the top and the left edge",
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
					yFrom: -20,
					yTo:   10,
					zFrom: -10,
					zTo:   20,
					flip:  on,
				},
			},
			want: nil,
		},
		{
			name: "returns nil if overlapbox is at both the top and the front edge",
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
					xTo:   20,
					yFrom: -10,
					yTo:   10,
					zFrom: -10,
					zTo:   20,
					flip:  on,
				},
			},
			want: nil,
		},
		{
			name: "returns top front right corner box if overlapbox is not at the top, front, or left edges",
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
					xFrom: 10,
					xTo:   20,
					yFrom: -20,
					yTo:   -10,
					zFrom: 10,
					zTo:   20,
					flip:  off,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, findTopFrontLeftCorner(tt.args.box, tt.args.overlapBox),
				"findTopFrontLeftCorner(%v, %v)", tt.args.box, tt.args.overlapBox)
		})
	}
}

func Test_findTopFrontRightCorner(t *testing.T) {
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
			name: "returns nil if overlapbox is at the front edge",
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
					xTo:   20,
					yFrom: -10,
					yTo:   10,
					zFrom: -10,
					zTo:   10,
					flip:  on,
				},
			},
			want: nil,
		},
		{
			name: "returns nil if overlapbox is at the right edge",
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
					yTo:   20,
					zFrom: -10,
					zTo:   10,
					flip:  on,
				},
			},
			want: nil,
		},
		{
			name: "returns nil if overlapbox is at both the front and the right edge",
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
					xTo:   20,
					yFrom: -10,
					yTo:   20,
					zFrom: -10,
					zTo:   10,
					flip:  on,
				},
			},
			want: nil,
		},
		{
			name: "returns nil if overlapbox is at both the top and the right edge",
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
					yTo:   20,
					zFrom: -10,
					zTo:   20,
					flip:  on,
				},
			},
			want: nil,
		},
		{
			name: "returns nil if overlapbox is at both the top and the front edge",
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
					xTo:   20,
					yFrom: -10,
					yTo:   10,
					zFrom: -10,
					zTo:   20,
					flip:  on,
				},
			},
			want: nil,
		},
		{
			name: "returns top front right corner box if overlapbox is not at the top, front, or right edges",
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
					xFrom: 10,
					xTo:   20,
					yFrom: 10,
					yTo:   20,
					zFrom: 10,
					zTo:   20,
					flip:  off,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, findTopFrontRightCorner(tt.args.box, tt.args.overlapBox),
				"findTopFrontRightCorner(%v, %v)", tt.args.box, tt.args.overlapBox)
		})
	}
}

func Test_findBottomBackLeftCorner(t *testing.T) {
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
			name: "returns nil if overlapbox is at the back edge",
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
					xFrom: -20,
					xTo:   10,
					yFrom: -10,
					yTo:   10,
					zFrom: -10,
					zTo:   10,
					flip:  on,
				},
			},
			want: nil,
		},
		{
			name: "returns nil if overlapbox is at the left edge",
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
					yFrom: -20,
					yTo:   10,
					zFrom: -10,
					zTo:   10,
					flip:  on,
				},
			},
			want: nil,
		},
		{
			name: "returns nil if overlapbox is at both the back and the left edge",
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
					xFrom: -20,
					xTo:   10,
					yFrom: -20,
					yTo:   10,
					zFrom: -10,
					zTo:   10,
					flip:  on,
				},
			},
			want: nil,
		},
		{
			name: "returns nil if overlapbox is at both the bottom and the left edge",
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
					yFrom: -20,
					yTo:   10,
					zFrom: -20,
					zTo:   10,
					flip:  on,
				},
			},
			want: nil,
		},
		{
			name: "returns nil if overlapbox is at both the bottom and the back edge",
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
					xFrom: -20,
					xTo:   10,
					yFrom: -10,
					yTo:   10,
					zFrom: -20,
					zTo:   10,
					flip:  on,
				},
			},
			want: nil,
		},
		{
			name: "returns bottom back right corner box if overlapbox is not at the bottom, back, or left edges",
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
					xFrom: -20,
					xTo:   -10,
					yFrom: -20,
					yTo:   -10,
					zFrom: -20,
					zTo:   -10,
					flip:  off,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, findBottomBackLeftCorner(tt.args.box, tt.args.overlapBox),
				"findBottomBackLeftCorner(%v, %v)", tt.args.box, tt.args.overlapBox)
		})
	}
}

func Test_findBottomBackRightCorner(t *testing.T) {
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
			name: "returns nil if overlapbox is at the back edge",
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
					xFrom: -20,
					xTo:   10,
					yFrom: -10,
					yTo:   10,
					zFrom: -10,
					zTo:   10,
					flip:  on,
				},
			},
			want: nil,
		},
		{
			name: "returns nil if overlapbox is at the right edge",
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
					yTo:   20,
					zFrom: -10,
					zTo:   10,
					flip:  on,
				},
			},
			want: nil,
		},
		{
			name: "returns nil if overlapbox is at both the back and the right edge",
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
					xFrom: -20,
					xTo:   10,
					yFrom: -10,
					yTo:   20,
					zFrom: -10,
					zTo:   10,
					flip:  on,
				},
			},
			want: nil,
		},
		{
			name: "returns nil if overlapbox is at both the bottom and the right edge",
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
					yTo:   20,
					zFrom: -20,
					zTo:   10,
					flip:  on,
				},
			},
			want: nil,
		},
		{
			name: "returns nil if overlapbox is at both the bottom and the back edge",
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
					xFrom: -20,
					xTo:   10,
					yFrom: -10,
					yTo:   10,
					zFrom: -20,
					zTo:   10,
					flip:  on,
				},
			},
			want: nil,
		},
		{
			name: "returns bottom back right corner box if overlapbox is not at the bottom, back, or right edges",
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
					xFrom: -20,
					xTo:   -10,
					yFrom: 10,
					yTo:   20,
					zFrom: -20,
					zTo:   -10,
					flip:  off,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, findBottomBackRightCorner(tt.args.box, tt.args.overlapBox),
				"findBottomBackRightCorner(%v, %v)", tt.args.box, tt.args.overlapBox)
		})
	}
}

func Test_findBottomFrontLeftCorner(t *testing.T) {
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
			name: "returns nil if overlapbox is at the front edge",
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
					xTo:   20,
					yFrom: -10,
					yTo:   10,
					zFrom: -10,
					zTo:   10,
					flip:  on,
				},
			},
			want: nil,
		},
		{
			name: "returns nil if overlapbox is at the left edge",
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
					yFrom: -20,
					yTo:   10,
					zFrom: -10,
					zTo:   10,
					flip:  on,
				},
			},
			want: nil,
		},
		{
			name: "returns nil if overlapbox is at both the front and the left edge",
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
					xTo:   20,
					yFrom: -20,
					yTo:   10,
					zFrom: -10,
					zTo:   10,
					flip:  on,
				},
			},
			want: nil,
		},
		{
			name: "returns nil if overlapbox is at both the bottom and the left edge",
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
					yFrom: -20,
					yTo:   10,
					zFrom: -20,
					zTo:   10,
					flip:  on,
				},
			},
			want: nil,
		},
		{
			name: "returns nil if overlapbox is at both the bottom and the front edge",
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
					xTo:   20,
					yFrom: -10,
					yTo:   10,
					zFrom: -20,
					zTo:   10,
					flip:  on,
				},
			},
			want: nil,
		},
		{
			name: "returns bottom front right corner box if overlapbox is not at the bottom, front, or left edges",
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
					xFrom: 10,
					xTo:   20,
					yFrom: -20,
					yTo:   -10,
					zFrom: -20,
					zTo:   -10,
					flip:  off,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, findBottomFrontLeftCorner(tt.args.box, tt.args.overlapBox),
				"findBottomFrontLeftCorner(%v, %v)", tt.args.box, tt.args.overlapBox)
		})
	}
}

func Test_findBottomFrontRightCorner(t *testing.T) {
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
			name: "returns nil if overlapbox is at the front edge",
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
					xTo:   20,
					yFrom: -10,
					yTo:   10,
					zFrom: -10,
					zTo:   10,
					flip:  on,
				},
			},
			want: nil,
		},
		{
			name: "returns nil if overlapbox is at the right edge",
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
					yTo:   20,
					zFrom: -10,
					zTo:   10,
					flip:  on,
				},
			},
			want: nil,
		},
		{
			name: "returns nil if overlapbox is at both the front and the right edge",
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
					xTo:   20,
					yFrom: -10,
					yTo:   20,
					zFrom: -10,
					zTo:   10,
					flip:  on,
				},
			},
			want: nil,
		},
		{
			name: "returns nil if overlapbox is at both the bottom and the right edge",
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
					yTo:   20,
					zFrom: -20,
					zTo:   10,
					flip:  on,
				},
			},
			want: nil,
		},
		{
			name: "returns nil if overlapbox is at both the bottom and the front edge",
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
					xTo:   20,
					yFrom: -10,
					yTo:   10,
					zFrom: -20,
					zTo:   10,
					flip:  on,
				},
			},
			want: nil,
		},
		{
			name: "returns bottom front right corner box if overlapbox is not at the bottom, front, or right edges",
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
					xFrom: 10,
					xTo:   20,
					yFrom: 10,
					yTo:   20,
					zFrom: -20,
					zTo:   -10,
					flip:  off,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, findBottomFrontRightCorner(tt.args.box, tt.args.overlapBox),
				"findBottomFrontRightCorner(%v, %v)", tt.args.box, tt.args.overlapBox)
		})
	}
}

func Test_overlap(t *testing.T) {
	type args struct {
		box      instruction
		otherBox instruction
	}

	tests := []struct {
		name string
		args args
		want []instruction
	}{
		{
			name: "returns only the two boxes if they do not overlap",
			args: args{
				box: instruction{
					xFrom: -20,
					xTo:   -10,
					yFrom: -20,
					yTo:   -10,
					zFrom: -20,
					zTo:   -10,
					flip:  on,
				},
				otherBox: instruction{
					xFrom: 10,
					xTo:   20,
					yFrom: 10,
					yTo:   20,
					zFrom: 10,
					zTo:   20,
					flip:  off,
				},
			},
			want: []instruction{
				{
					xFrom: -20,
					xTo:   -10,
					yFrom: -20,
					yTo:   -10,
					zFrom: -20,
					zTo:   -10,
					flip:  on,
				},
				{
					xFrom: 10,
					xTo:   20,
					yFrom: 10,
					yTo:   20,
					zFrom: 10,
					zTo:   20,
					flip:  off,
				},
			},
		},
		{
			name: "returns overlapbox, and all other pieces from both of the boxes",
			args: args{
				box: instruction{
					xFrom: -20,
					xTo:   5,
					yFrom: -20,
					yTo:   5,
					zFrom: -20,
					zTo:   5,
					flip:  on,
				},
				otherBox: instruction{
					xFrom: -5,
					xTo:   20,
					yFrom: -5,
					yTo:   20,
					zFrom: -5,
					zTo:   20,
					flip:  off,
				},
			},
			want: []instruction{
				// the overlapbox
				{
					xFrom: -5,
					xTo:   5,
					yFrom: -5,
					yTo:   5,
					zFrom: -5,
					zTo:   5,
					flip:  off,
				},
				// box back face
				{
					xFrom: -20,
					xTo:   -5,
					yFrom: -5,
					yTo:   5,
					zFrom: -5,
					zTo:   5,
					flip:  on,
				},
				// box bottom face
				{
					xFrom: -5,
					xTo:   5,
					yFrom: -5,
					yTo:   5,
					zFrom: -20,
					zTo:   -5,
					flip:  on,
				},
				// box left face
				{
					xFrom: -5,
					xTo:   5,
					yFrom: -20,
					yTo:   -5,
					zFrom: -5,
					zTo:   5,
					flip:  on,
				},
				// box back left edge
				{
					xFrom: -20,
					xTo:   -5,
					yFrom: -20,
					yTo:   -5,
					zFrom: -5,
					zTo:   5,
					flip:  on,
				},
				// box back bottom edge
				{
					xFrom: -20,
					xTo:   -5,
					yFrom: -5,
					yTo:   5,
					zFrom: -20,
					zTo:   -5,
					flip:  on,
				},
				// box left bottom edge
				{
					xFrom: -5,
					xTo:   5,
					yFrom: -20,
					yTo:   -5,
					zFrom: -20,
					zTo:   -5,
					flip:  on,
				},
				// box bottom back left corner
				{
					xFrom: -20,
					xTo:   -5,
					yFrom: -20,
					yTo:   -5,
					zFrom: -20,
					zTo:   -5,
					flip:  on,
				},

				// otherBox front face
				{
					xFrom: 5,
					xTo:   20,
					yFrom: -5,
					yTo:   5,
					zFrom: -5,
					zTo:   5,
					flip:  off,
				},
				// otherBox right face
				{
					xFrom: -5,
					xTo:   5,
					yFrom: 5,
					yTo:   20,
					zFrom: -5,
					zTo:   5,
					flip:  off,
				},
				// otherBox top face
				{
					xFrom: -5,
					xTo:   5,
					yFrom: -5,
					yTo:   5,
					zFrom: 5,
					zTo:   20,
					flip:  off,
				},
				// otherBox front right edge
				{
					xFrom: 5,
					xTo:   20,
					yFrom: 5,
					yTo:   20,
					zFrom: -5,
					zTo:   5,
					flip:  off,
				},
				// otherBox front top edge
				{
					xFrom: 5,
					xTo:   20,
					yFrom: -5,
					yTo:   5,
					zFrom: 5,
					zTo:   20,
					flip:  off,
				},
				// otherBox right top edge
				{
					xFrom: -5,
					xTo:   5,
					yFrom: 5,
					yTo:   20,
					zFrom: 5,
					zTo:   20,
					flip:  off,
				},
				// otherBox top right front corner
				{
					xFrom: 5,
					xTo:   20,
					yFrom: 5,
					yTo:   20,
					zFrom: 5,
					zTo:   20,
					flip:  off,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.ElementsMatchf(t, tt.want, overlap(tt.args.box, tt.args.otherBox),
				"overlap(%v, %v)", tt.args.box, tt.args.otherBox)
		})
	}
}

func Test_mergeBoxes(t *testing.T) {
	standardBox := instruction{
		xFrom: 0,
		xTo:   10,
		yFrom: 0,
		yTo:   10,
		zFrom: 0,
		zTo:   10,
		flip:  on,
	}

	type args struct {
		box      instruction
		otherBox instruction
	}

	tests := []struct {
		name string
		args args
		want []instruction
	}{
		{
			name: "merges two boxes on x, touching",
			args: args{
				box: standardBox,
				otherBox: instruction{
					xFrom: 10,
					xTo:   20,
					yFrom: 0,
					yTo:   10,
					zFrom: 0,
					zTo:   10,
					flip:  on,
				},
			},
			want: []instruction{
				{
					xFrom: 0,
					xTo:   20,
					yFrom: 0,
					yTo:   10,
					zFrom: 0,
					zTo:   10,
					flip:  on,
				},
			},
		},
		{
			name: "merges two boxes on x, overlapping",
			args: args{
				box: standardBox,
				otherBox: instruction{
					xFrom: 5,
					xTo:   15,
					yFrom: 0,
					yTo:   10,
					zFrom: 0,
					zTo:   10,
					flip:  on,
				},
			},
			want: []instruction{
				{
					xFrom: 0,
					xTo:   15,
					yFrom: 0,
					yTo:   10,
					zFrom: 0,
					zTo:   10,
					flip:  on,
				},
			},
		},
		{
			name: "merges two boxes on x, fully enclosed",
			args: args{
				box: standardBox,
				otherBox: instruction{
					xFrom: -6,
					xTo:   17,
					yFrom: 0,
					yTo:   10,
					zFrom: 0,
					zTo:   10,
					flip:  on,
				},
			},
			want: []instruction{
				{
					xFrom: -6,
					xTo:   17,
					yFrom: 0,
					yTo:   10,
					zFrom: 0,
					zTo:   10,
					flip:  on,
				},
			},
		},
		{
			name: "fails merging two boxes on x, do not touch",
			args: args{
				box: standardBox,
				otherBox: instruction{
					xFrom: 12,
					xTo:   17,
					yFrom: 0,
					yTo:   10,
					zFrom: 0,
					zTo:   10,
					flip:  on,
				},
			},
			want: []instruction{
				standardBox,
				{
					xFrom: 12,
					xTo:   17,
					yFrom: 0,
					yTo:   10,
					zFrom: 0,
					zTo:   10,
					flip:  on,
				},
			},
		},
		{
			name: "merges two boxes on y, touching",
			args: args{
				box: standardBox,
				otherBox: instruction{
					xFrom: 0,
					xTo:   10,
					yFrom: 10,
					yTo:   20,
					zFrom: 0,
					zTo:   10,
					flip:  on,
				},
			},
			want: []instruction{
				{
					xFrom: 0,
					xTo:   10,
					yFrom: 0,
					yTo:   20,
					zFrom: 0,
					zTo:   10,
					flip:  on,
				},
			},
		},
		{
			name: "merges two boxes on y, overlapping",
			args: args{
				box: standardBox,
				otherBox: instruction{
					xFrom: 0,
					xTo:   10,
					yFrom: 5,
					yTo:   15,
					zFrom: 0,
					zTo:   10,
					flip:  on,
				},
			},
			want: []instruction{
				{
					xFrom: 0,
					xTo:   10,
					yFrom: 0,
					yTo:   15,
					zFrom: 0,
					zTo:   10,
					flip:  on,
				},
			},
		},
		{
			name: "merges two boxes on y, fully enclosed",
			args: args{
				box: standardBox,
				otherBox: instruction{
					xFrom: 0,
					xTo:   10,
					yFrom: -6,
					yTo:   17,
					zFrom: 0,
					zTo:   10,
					flip:  on,
				},
			},
			want: []instruction{
				{
					xFrom: 0,
					xTo:   10,
					yFrom: -6,
					yTo:   17,
					zFrom: 0,
					zTo:   10,
					flip:  on,
				},
			},
		},
		{
			name: "fails merging two boxes on y, do not touch",
			args: args{
				box: standardBox,
				otherBox: instruction{
					xFrom: 0,
					xTo:   10,
					yFrom: 12,
					yTo:   17,
					zFrom: 0,
					zTo:   10,
					flip:  on,
				},
			},
			want: []instruction{
				standardBox,
				{
					xFrom: 0,
					xTo:   10,
					yFrom: 12,
					yTo:   17,
					zFrom: 0,
					zTo:   10,
					flip:  on,
				},
			},
		},
		{
			name: "merges two boxes on z, touching",
			args: args{
				box: standardBox,
				otherBox: instruction{
					xFrom: 0,
					xTo:   10,
					yFrom: 0,
					yTo:   10,
					zFrom: 10,
					zTo:   20,
					flip:  on,
				},
			},
			want: []instruction{
				{
					xFrom: 0,
					xTo:   10,
					yFrom: 0,
					yTo:   10,
					zFrom: 0,
					zTo:   20,
					flip:  on,
				},
			},
		},
		{
			name: "merges two boxes on z, overlapping",
			args: args{
				box: standardBox,
				otherBox: instruction{
					xFrom: 0,
					xTo:   10,
					yFrom: 0,
					yTo:   10,
					zFrom: 5,
					zTo:   15,
					flip:  on,
				},
			},
			want: []instruction{
				{
					xFrom: 0,
					xTo:   10,
					yFrom: 0,
					yTo:   10,
					zFrom: 0,
					zTo:   15,
					flip:  on,
				},
			},
		},
		{
			name: "merges two boxes on z, fully enclosed",
			args: args{
				box: standardBox,
				otherBox: instruction{
					xFrom: 0,
					xTo:   10,
					yFrom: 0,
					yTo:   10,
					zFrom: -6,
					zTo:   17,
					flip:  on,
				},
			},
			want: []instruction{
				{
					xFrom: 0,
					xTo:   10,
					yFrom: 0,
					yTo:   10,
					zFrom: -6,
					zTo:   17,
					flip:  on,
				},
			},
		},
		{
			name: "fails merging two boxes on z, do not touch",
			args: args{
				box: standardBox,
				otherBox: instruction{
					xFrom: 0,
					xTo:   10,
					yFrom: 0,
					yTo:   10,
					zFrom: 12,
					zTo:   17,
					flip:  on,
				},
			},
			want: []instruction{
				standardBox,
				{
					xFrom: 0,
					xTo:   10,
					yFrom: 0,
					yTo:   10,
					zFrom: 12,
					zTo:   17,
					flip:  on,
				},
			},
		},
		{
			name: "fails merging two boxes, two axis values do not match",
			args: args{
				box: standardBox,
				otherBox: instruction{
					xFrom: 0,
					xTo:   10,
					yFrom: 3,
					yTo:   10,
					zFrom: 5,
					zTo:   15,
					flip:  on,
				},
			},
			want: []instruction{
				standardBox,
				{
					xFrom: 0,
					xTo:   10,
					yFrom: 3,
					yTo:   10,
					zFrom: 5,
					zTo:   15,
					flip:  on,
				},
			},
		},
		{
			name: "fails merging two boxes, on off switches differ",
			args: args{
				box: standardBox,
				otherBox: instruction{
					xFrom: 0,
					xTo:   10,
					yFrom: 0,
					yTo:   10,
					zFrom: 5,
					zTo:   15,
					flip:  off,
				},
			},
			want: []instruction{
				standardBox,
				{
					xFrom: 0,
					xTo:   10,
					yFrom: 0,
					yTo:   10,
					zFrom: 5,
					zTo:   15,
					flip:  off,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.ElementsMatchf(t, tt.want, mergeBoxes(tt.args.box, tt.args.otherBox),
				"mergeBoxes(%v, %v)", tt.args.box, tt.args.otherBox)
		})
	}
}

func Test_overlapAndMerge(t *testing.T) {
	type args struct {
		box      instruction
		otherBox instruction
	}

	tests := []struct {
		name string
		args args
		want map[string]instruction
	}{
		{
			name: "overlaps and merges two boxes",
			args: args{
				box: instruction{
					xFrom: -20,
					xTo:   5,
					yFrom: -20,
					yTo:   5,
					zFrom: -20,
					zTo:   5,
					flip:  on,
				},
				otherBox: instruction{
					xFrom: -5,
					xTo:   20,
					yFrom: -5,
					yTo:   20,
					zFrom: -5,
					zTo:   20,
					flip:  off,
				},
			},
			want: map[string]instruction{
				"-20/-5/-20/-5/-20/-5/on": {xFrom: -20, xTo: -5, yFrom: -20, yTo: -5, zFrom: -20, zTo: -5, flip: 1},
				"-5/5/-20/5/-20/-5/on":    {xFrom: -5, xTo: 5, yFrom: -20, yTo: 5, zFrom: -20, zTo: -5, flip: 1},
				"-20/5/-20/-5/-5/5/on":    {xFrom: -20, xTo: 5, yFrom: -20, yTo: -5, zFrom: -5, zTo: 5, flip: 1},
				"-20/-5/-5/5/-20/5/on":    {xFrom: -20, xTo: -5, yFrom: -5, yTo: 5, zFrom: -20, zTo: 5, flip: 1},
				"-5/20/-5/20/-5/20/off":   {xFrom: -5, xTo: 20, yFrom: -5, yTo: 20, zFrom: -5, zTo: 20, flip: 2},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, overlapAndMerge(tt.args.box, tt.args.otherBox),
				"overlapAndMerge(%v, %v)", tt.args.box, tt.args.otherBox)
		})
	}
}
