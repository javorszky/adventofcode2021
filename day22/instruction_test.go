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
				XFrom: 967,
				XTo:   23432,
				YFrom: 45373,
				YTo:   81175,
				ZFrom: 27513,
				ZTo:   53682,
				Flip:  on,
			},
		},
		{
			name: "parses string with negative values correctly",
			args: args{
				s: "on x=-54112..-39298,y=-85059..-49293,z=-27449..7877",
			},
			want: instruction{
				XFrom: -54112,
				XTo:   -39298,
				YFrom: -85059,
				YTo:   -49293,
				ZFrom: -27449,
				ZTo:   7877,
				Flip:  on,
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
				XFrom: 1,
				XTo:   1,
				YFrom: 1,
				YTo:   1,
				ZFrom: 1,
				ZTo:   1,
				Flip:  on,
			},
			want: 1,
		},
		{
			name: "calculates volume for cuboid",
			i: instruction{
				XFrom: 1,
				XTo:   3,
				YFrom: 1,
				YTo:   3,
				ZFrom: 1,
				ZTo:   3,
				Flip:  on,
			},
			want: 27,
		},
		{
			name: "calculates volume for cuboid",
			i: instruction{
				XFrom: -2,
				XTo:   2,
				YFrom: -8,
				YTo:   -4,
				ZFrom: 5,
				ZTo:   10,
				Flip:  on,
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
					XFrom: -10,
					XTo:   10,
					YFrom: -10,
					YTo:   10,
					ZFrom: -10,
					ZTo:   10,
					Flip:  on,
				},
				otherBox: instruction{
					XFrom: 15,
					XTo:   20,
					YFrom: -10,
					YTo:   10,
					ZFrom: -10,
					ZTo:   10,
					Flip:  off,
				},
			},
			want:    instruction{},
			wantErr: assert.Error,
		},
		{
			name: "finds overlapping box between the two boxes where box is fully contained",
			args: args{
				box: instruction{
					XFrom: -10,
					XTo:   10,
					YFrom: -10,
					YTo:   10,
					ZFrom: -10,
					ZTo:   10,
					Flip:  on,
				},
				otherBox: instruction{
					XFrom: -20,
					XTo:   20,
					YFrom: -20,
					YTo:   20,
					ZFrom: -20,
					ZTo:   20,
					Flip:  off,
				},
			},
			want: instruction{
				XFrom: -10,
				XTo:   10,
				YFrom: -10,
				YTo:   10,
				ZFrom: -10,
				ZTo:   10,
				Flip:  off,
			},
			wantErr: assert.NoError,
		},
		{
			name: "finds overlapping box between the two boxes where only a corner matches",
			args: args{
				box: instruction{
					XFrom: -10,
					XTo:   10,
					YFrom: -10,
					YTo:   10,
					ZFrom: -10,
					ZTo:   10,
					Flip:  on,
				},
				otherBox: instruction{
					XFrom: 5,
					XTo:   15,
					YFrom: 5,
					YTo:   15,
					ZFrom: 5,
					ZTo:   15,
					Flip:  off,
				},
			},
			want: instruction{
				XFrom: 5,
				XTo:   10,
				YFrom: 5,
				YTo:   10,
				ZFrom: 5,
				ZTo:   10,
				Flip:  off,
			},
			wantErr: assert.NoError,
		},
		{
			name: "finds overlapping box between the two boxes where only a single cube",
			args: args{
				box: instruction{
					XFrom: -10,
					XTo:   10,
					YFrom: -10,
					YTo:   10,
					ZFrom: -10,
					ZTo:   10,
					Flip:  on,
				},
				otherBox: instruction{
					XFrom: 10,
					XTo:   12,
					YFrom: 10,
					YTo:   12,
					ZFrom: 10,
					ZTo:   12,
					Flip:  off,
				},
			},
			want: instruction{
				XFrom: 10,
				XTo:   10,
				YFrom: 10,
				YTo:   10,
				ZFrom: 10,
				ZTo:   10,
				Flip:  off,
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
					XFrom: -20,
					XTo:   20,
					YFrom: -20,
					YTo:   20,
					ZFrom: -20,
					ZTo:   20,
					Flip:  off,
				},
				overlapBox: instruction{
					XFrom: -10,
					XTo:   10,
					YFrom: -10,
					YTo:   10,
					ZFrom: 10,
					ZTo:   20,
					Flip:  on,
				},
			},
			want: nil,
		},
		{
			name: "returns top face box if overlapbox is not at the top edge",
			args: args{
				box: instruction{
					XFrom: -20,
					XTo:   20,
					YFrom: -20,
					YTo:   20,
					ZFrom: -20,
					ZTo:   20,
					Flip:  off,
				},
				overlapBox: instruction{
					XFrom: -10,
					XTo:   10,
					YFrom: -10,
					YTo:   10,
					ZFrom: -10,
					ZTo:   10,
					Flip:  on,
				},
			},
			want: []instruction{
				{
					XFrom: -10,
					XTo:   10,
					YFrom: -10,
					YTo:   10,
					ZFrom: 11,
					ZTo:   20,
					Flip:  off,
				},
			},
		},
		{
			name: "returns a single layer of top face",
			args: args{
				box: instruction{
					XFrom: -20,
					XTo:   20,
					YFrom: -20,
					YTo:   20,
					ZFrom: -20,
					ZTo:   20,
					Flip:  off,
				},
				overlapBox: instruction{
					XFrom: -10,
					XTo:   10,
					YFrom: -10,
					YTo:   10,
					ZFrom: -10,
					ZTo:   19,
					Flip:  on,
				},
			},
			want: []instruction{
				{
					XFrom: -10,
					XTo:   10,
					YFrom: -10,
					YTo:   10,
					ZFrom: 20,
					ZTo:   20,
					Flip:  off,
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
					XFrom: -20,
					XTo:   20,
					YFrom: -20,
					YTo:   20,
					ZFrom: -20,
					ZTo:   20,
					Flip:  off,
				},
				overlapBox: instruction{
					XFrom: -10,
					XTo:   10,
					YFrom: -10,
					YTo:   10,
					ZFrom: -20,
					ZTo:   -10,
					Flip:  on,
				},
			},
			want: nil,
		},
		{
			name: "returns bottom face box if overlapbox is not at the bottom edge",
			args: args{
				box: instruction{
					XFrom: -20,
					XTo:   20,
					YFrom: -20,
					YTo:   20,
					ZFrom: -20,
					ZTo:   20,
					Flip:  off,
				},
				overlapBox: instruction{
					XFrom: -10,
					XTo:   10,
					YFrom: -10,
					YTo:   10,
					ZFrom: -10,
					ZTo:   10,
					Flip:  on,
				},
			},
			want: []instruction{
				{
					XFrom: -10,
					XTo:   10,
					YFrom: -10,
					YTo:   10,
					ZFrom: -20,
					ZTo:   -11,
					Flip:  off,
				},
			},
		},
		{
			name: "returns a single layer bottom face",
			args: args{
				box: instruction{
					XFrom: -20,
					XTo:   20,
					YFrom: -20,
					YTo:   20,
					ZFrom: -20,
					ZTo:   20,
					Flip:  off,
				},
				overlapBox: instruction{
					XFrom: -10,
					XTo:   10,
					YFrom: -10,
					YTo:   10,
					ZFrom: -19,
					ZTo:   10,
					Flip:  on,
				},
			},
			want: []instruction{
				{
					XFrom: -10,
					XTo:   10,
					YFrom: -10,
					YTo:   10,
					ZFrom: -20,
					ZTo:   -20,
					Flip:  off,
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
					XFrom: -20,
					XTo:   20,
					YFrom: -20,
					YTo:   20,
					ZFrom: -20,
					ZTo:   20,
					Flip:  off,
				},
				overlapBox: instruction{
					XFrom: 10,
					XTo:   20,
					YFrom: -10,
					YTo:   10,
					ZFrom: -20,
					ZTo:   -10,
					Flip:  on,
				},
			},
			want: nil,
		},
		{
			name: "returns front face box if overlapbox is not at the front edge",
			args: args{
				box: instruction{
					XFrom: -20,
					XTo:   20,
					YFrom: -20,
					YTo:   20,
					ZFrom: -20,
					ZTo:   20,
					Flip:  off,
				},
				overlapBox: instruction{
					XFrom: -10,
					XTo:   10,
					YFrom: -10,
					YTo:   10,
					ZFrom: -10,
					ZTo:   10,
					Flip:  on,
				},
			},
			want: []instruction{
				{
					XFrom: 11,
					XTo:   20,
					YFrom: -10,
					YTo:   10,
					ZFrom: -10,
					ZTo:   10,
					Flip:  off,
				},
			},
		},
		{
			name: "returns single layer front face box",
			args: args{
				box: instruction{
					XFrom: -20,
					XTo:   20,
					YFrom: -20,
					YTo:   20,
					ZFrom: -20,
					ZTo:   20,
					Flip:  off,
				},
				overlapBox: instruction{
					XFrom: -10,
					XTo:   19,
					YFrom: -10,
					YTo:   10,
					ZFrom: -10,
					ZTo:   10,
					Flip:  on,
				},
			},
			want: []instruction{
				{
					XFrom: 20,
					XTo:   20,
					YFrom: -10,
					YTo:   10,
					ZFrom: -10,
					ZTo:   10,
					Flip:  off,
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
					XFrom: -20,
					XTo:   20,
					YFrom: -20,
					YTo:   20,
					ZFrom: -20,
					ZTo:   20,
					Flip:  off,
				},
				overlapBox: instruction{
					XFrom: -20,
					XTo:   -10,
					YFrom: -10,
					YTo:   10,
					ZFrom: -20,
					ZTo:   -10,
					Flip:  on,
				},
			},
			want: nil,
		},
		{
			name: "returns back face box if overlapbox is not at the back edge",
			args: args{
				box: instruction{
					XFrom: -20,
					XTo:   20,
					YFrom: -20,
					YTo:   20,
					ZFrom: -20,
					ZTo:   20,
					Flip:  off,
				},
				overlapBox: instruction{
					XFrom: -10,
					XTo:   10,
					YFrom: -10,
					YTo:   10,
					ZFrom: -10,
					ZTo:   10,
					Flip:  on,
				},
			},
			want: []instruction{
				{
					XFrom: -20,
					XTo:   -11,
					YFrom: -10,
					YTo:   10,
					ZFrom: -10,
					ZTo:   10,
					Flip:  off,
				},
			},
		},
		{
			name: "returns single layer back face",
			args: args{
				box: instruction{
					XFrom: -20,
					XTo:   20,
					YFrom: -20,
					YTo:   20,
					ZFrom: -20,
					ZTo:   20,
					Flip:  off,
				},
				overlapBox: instruction{
					XFrom: -19,
					XTo:   10,
					YFrom: -10,
					YTo:   10,
					ZFrom: -10,
					ZTo:   10,
					Flip:  on,
				},
			},
			want: []instruction{
				{
					XFrom: -20,
					XTo:   -20,
					YFrom: -10,
					YTo:   10,
					ZFrom: -10,
					ZTo:   10,
					Flip:  off,
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
					XFrom: -20,
					XTo:   20,
					YFrom: -20,
					YTo:   20,
					ZFrom: -20,
					ZTo:   20,
					Flip:  off,
				},
				overlapBox: instruction{
					XFrom: -10,
					XTo:   10,
					YFrom: -20,
					YTo:   -10,
					ZFrom: -10,
					ZTo:   10,
					Flip:  on,
				},
			},
			want: nil,
		},
		{
			name: "returns left face box if overlapbox is not at the left edge",
			args: args{
				box: instruction{
					XFrom: -20,
					XTo:   20,
					YFrom: -20,
					YTo:   20,
					ZFrom: -20,
					ZTo:   20,
					Flip:  off,
				},
				overlapBox: instruction{
					XFrom: -10,
					XTo:   10,
					YFrom: -10,
					YTo:   10,
					ZFrom: -10,
					ZTo:   10,
					Flip:  on,
				},
			},
			want: []instruction{
				{
					XFrom: -10,
					XTo:   10,
					YFrom: -20,
					YTo:   -11,
					ZFrom: -10,
					ZTo:   10,
					Flip:  off,
				},
			},
		},
		{
			name: "returns single layer left face",
			args: args{
				box: instruction{
					XFrom: -20,
					XTo:   20,
					YFrom: -20,
					YTo:   20,
					ZFrom: -20,
					ZTo:   20,
					Flip:  off,
				},
				overlapBox: instruction{
					XFrom: -10,
					XTo:   10,
					YFrom: -19,
					YTo:   10,
					ZFrom: -10,
					ZTo:   10,
					Flip:  on,
				},
			},
			want: []instruction{
				{
					XFrom: -10,
					XTo:   10,
					YFrom: -20,
					YTo:   -20,
					ZFrom: -10,
					ZTo:   10,
					Flip:  off,
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
					XFrom: -20,
					XTo:   20,
					YFrom: -20,
					YTo:   20,
					ZFrom: -20,
					ZTo:   20,
					Flip:  off,
				},
				overlapBox: instruction{
					XFrom: -10,
					XTo:   10,
					YFrom: 10,
					YTo:   20,
					ZFrom: -10,
					ZTo:   10,
					Flip:  on,
				},
			},
			want: nil,
		},
		{
			name: "returns right face box if overlapbox is not at the right edge",
			args: args{
				box: instruction{
					XFrom: -20,
					XTo:   20,
					YFrom: -20,
					YTo:   20,
					ZFrom: -20,
					ZTo:   20,
					Flip:  off,
				},
				overlapBox: instruction{
					XFrom: -10,
					XTo:   10,
					YFrom: -10,
					YTo:   10,
					ZFrom: -10,
					ZTo:   10,
					Flip:  on,
				},
			},
			want: []instruction{
				{
					XFrom: -10,
					XTo:   10,
					YFrom: 11,
					YTo:   20,
					ZFrom: -10,
					ZTo:   10,
					Flip:  off,
				},
			},
		},
		{
			name: "returns single layer right face",
			args: args{
				box: instruction{
					XFrom: -20,
					XTo:   20,
					YFrom: -20,
					YTo:   20,
					ZFrom: -20,
					ZTo:   20,
					Flip:  off,
				},
				overlapBox: instruction{
					XFrom: -10,
					XTo:   10,
					YFrom: -10,
					YTo:   19,
					ZFrom: -10,
					ZTo:   10,
					Flip:  on,
				},
			},
			want: []instruction{
				{
					XFrom: -10,
					XTo:   10,
					YFrom: 20,
					YTo:   20,
					ZFrom: -10,
					ZTo:   10,
					Flip:  off,
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
					XFrom: -20,
					XTo:   20,
					YFrom: -20,
					YTo:   20,
					ZFrom: -20,
					ZTo:   20,
					Flip:  off,
				},
				overlapBox: instruction{
					XFrom: -10,
					XTo:   10,
					YFrom: -10,
					YTo:   10,
					ZFrom: 10,
					ZTo:   20,
					Flip:  on,
				},
			},
			want: nil,
		},
		{
			name: "returns nil if overlapbox is at the left edge",
			args: args{
				box: instruction{
					XFrom: -20,
					XTo:   20,
					YFrom: -20,
					YTo:   20,
					ZFrom: -20,
					ZTo:   20,
					Flip:  off,
				},
				overlapBox: instruction{
					XFrom: -10,
					XTo:   10,
					YFrom: -20,
					YTo:   10,
					ZFrom: -10,
					ZTo:   10,
					Flip:  on,
				},
			},
			want: nil,
		},
		{
			name: "returns nil if overlapbox is at both the top and the left edge",
			args: args{
				box: instruction{
					XFrom: -20,
					XTo:   20,
					YFrom: -20,
					YTo:   20,
					ZFrom: -20,
					ZTo:   20,
					Flip:  off,
				},
				overlapBox: instruction{
					XFrom: -10,
					XTo:   10,
					YFrom: -20,
					YTo:   10,
					ZFrom: -10,
					ZTo:   20,
					Flip:  on,
				},
			},
			want: nil,
		},
		{
			name: "returns top left edge box if overlapbox is not at the top or left edges",
			args: args{
				box: instruction{
					XFrom: -20,
					XTo:   20,
					YFrom: -20,
					YTo:   20,
					ZFrom: -20,
					ZTo:   20,
					Flip:  off,
				},
				overlapBox: instruction{
					XFrom: -10,
					XTo:   10,
					YFrom: -10,
					YTo:   10,
					ZFrom: -10,
					ZTo:   10,
					Flip:  on,
				},
			},
			want: []instruction{
				{
					XFrom: -10,
					XTo:   10,
					YFrom: -20,
					YTo:   -11,
					ZFrom: 11,
					ZTo:   20,
					Flip:  off,
				},
			},
		},
		{
			name: "returns single cube top left edge box",
			args: args{
				box: instruction{
					XFrom: -20,
					XTo:   20,
					YFrom: -20,
					YTo:   20,
					ZFrom: -20,
					ZTo:   20,
					Flip:  off,
				},
				overlapBox: instruction{
					XFrom: -10,
					XTo:   10,
					YFrom: -19,
					YTo:   10,
					ZFrom: -10,
					ZTo:   19,
					Flip:  on,
				},
			},
			want: []instruction{
				{
					XFrom: -10,
					XTo:   10,
					YFrom: -20,
					YTo:   -20,
					ZFrom: 20,
					ZTo:   20,
					Flip:  off,
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
					XFrom: -20,
					XTo:   20,
					YFrom: -20,
					YTo:   20,
					ZFrom: -20,
					ZTo:   20,
					Flip:  off,
				},
				overlapBox: instruction{
					XFrom: -10,
					XTo:   10,
					YFrom: -10,
					YTo:   10,
					ZFrom: 10,
					ZTo:   20,
					Flip:  on,
				},
			},
			want: nil,
		},
		{
			name: "returns nil if overlapbox is at the back edge",
			args: args{
				box: instruction{
					XFrom: -20,
					XTo:   20,
					YFrom: -20,
					YTo:   20,
					ZFrom: -20,
					ZTo:   20,
					Flip:  off,
				},
				overlapBox: instruction{
					XFrom: -20,
					XTo:   10,
					YFrom: -10,
					YTo:   10,
					ZFrom: -10,
					ZTo:   10,
					Flip:  on,
				},
			},
			want: nil,
		},
		{
			name: "returns nil if overlapbox is at both the top and the back edge",
			args: args{
				box: instruction{
					XFrom: -20,
					XTo:   20,
					YFrom: -20,
					YTo:   20,
					ZFrom: -20,
					ZTo:   20,
					Flip:  off,
				},
				overlapBox: instruction{
					XFrom: -20,
					XTo:   10,
					YFrom: -10,
					YTo:   10,
					ZFrom: -10,
					ZTo:   20,
					Flip:  on,
				},
			},
			want: nil,
		},
		{
			name: "returns top right edge box if overlapbox is not at the top or back edge",
			args: args{
				box: instruction{
					XFrom: -20,
					XTo:   20,
					YFrom: -20,
					YTo:   20,
					ZFrom: -20,
					ZTo:   20,
					Flip:  off,
				},
				overlapBox: instruction{
					XFrom: -10,
					XTo:   10,
					YFrom: -10,
					YTo:   10,
					ZFrom: -10,
					ZTo:   10,
					Flip:  on,
				},
			},
			want: []instruction{
				{
					XFrom: -20,
					XTo:   -11,
					YFrom: -10,
					YTo:   10,
					ZFrom: 11,
					ZTo:   20,
					Flip:  off,
				},
			},
		},
		{
			name: "returns single line top right edge box",
			args: args{
				box: instruction{
					XFrom: -20,
					XTo:   20,
					YFrom: -20,
					YTo:   20,
					ZFrom: -20,
					ZTo:   20,
					Flip:  off,
				},
				overlapBox: instruction{
					XFrom: -19,
					XTo:   10,
					YFrom: -10,
					YTo:   10,
					ZFrom: -10,
					ZTo:   19,
					Flip:  on,
				},
			},
			want: []instruction{
				{
					XFrom: -20,
					XTo:   -20,
					YFrom: -10,
					YTo:   10,
					ZFrom: 20,
					ZTo:   20,
					Flip:  off,
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
					XFrom: -20,
					XTo:   20,
					YFrom: -20,
					YTo:   20,
					ZFrom: -20,
					ZTo:   20,
					Flip:  off,
				},
				overlapBox: instruction{
					XFrom: -10,
					XTo:   10,
					YFrom: -10,
					YTo:   10,
					ZFrom: 10,
					ZTo:   20,
					Flip:  on,
				},
			},
			want: nil,
		},
		{
			name: "returns nil if overlapbox is at the right edge",
			args: args{
				box: instruction{
					XFrom: -20,
					XTo:   20,
					YFrom: -20,
					YTo:   20,
					ZFrom: -20,
					ZTo:   20,
					Flip:  off,
				},
				overlapBox: instruction{
					XFrom: -10,
					XTo:   10,
					YFrom: -10,
					YTo:   20,
					ZFrom: -10,
					ZTo:   10,
					Flip:  on,
				},
			},
			want: nil,
		},
		{
			name: "returns nil if overlapbox is at both the top and the right edge",
			args: args{
				box: instruction{
					XFrom: -20,
					XTo:   20,
					YFrom: -20,
					YTo:   20,
					ZFrom: -20,
					ZTo:   20,
					Flip:  off,
				},
				overlapBox: instruction{
					XFrom: -10,
					XTo:   10,
					YFrom: -10,
					YTo:   20,
					ZFrom: -10,
					ZTo:   20,
					Flip:  on,
				},
			},
			want: nil,
		},
		{
			name: "returns top left edge box if overlapbox is not at the top or left edges",
			args: args{
				box: instruction{
					XFrom: -20,
					XTo:   20,
					YFrom: -20,
					YTo:   20,
					ZFrom: -20,
					ZTo:   20,
					Flip:  off,
				},
				overlapBox: instruction{
					XFrom: -10,
					XTo:   10,
					YFrom: -10,
					YTo:   10,
					ZFrom: -10,
					ZTo:   10,
					Flip:  on,
				},
			},
			want: []instruction{
				{
					XFrom: -10,
					XTo:   10,
					YFrom: 11,
					YTo:   20,
					ZFrom: 11,
					ZTo:   20,
					Flip:  off,
				},
			},
		},
		{
			name: "returns single line top left edge box",
			args: args{
				box: instruction{
					XFrom: -20,
					XTo:   20,
					YFrom: -20,
					YTo:   20,
					ZFrom: -20,
					ZTo:   20,
					Flip:  off,
				},
				overlapBox: instruction{
					XFrom: -10,
					XTo:   10,
					YFrom: -10,
					YTo:   19,
					ZFrom: -10,
					ZTo:   19,
					Flip:  on,
				},
			},
			want: []instruction{
				{
					XFrom: -10,
					XTo:   10,
					YFrom: 20,
					YTo:   20,
					ZFrom: 20,
					ZTo:   20,
					Flip:  off,
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
					XFrom: -20,
					XTo:   20,
					YFrom: -20,
					YTo:   20,
					ZFrom: -20,
					ZTo:   20,
					Flip:  off,
				},
				overlapBox: instruction{
					XFrom: -10,
					XTo:   10,
					YFrom: -10,
					YTo:   10,
					ZFrom: 10,
					ZTo:   20,
					Flip:  on,
				},
			},
			want: nil,
		},
		{
			name: "returns nil if overlapbox is at the front edge",
			args: args{
				box: instruction{
					XFrom: -20,
					XTo:   20,
					YFrom: -20,
					YTo:   20,
					ZFrom: -20,
					ZTo:   20,
					Flip:  off,
				},
				overlapBox: instruction{
					XFrom: -10,
					XTo:   20,
					YFrom: -10,
					YTo:   10,
					ZFrom: -10,
					ZTo:   10,
					Flip:  on,
				},
			},
			want: nil,
		},
		{
			name: "returns nil if overlapbox is at both the top and the front edge",
			args: args{
				box: instruction{
					XFrom: -20,
					XTo:   20,
					YFrom: -20,
					YTo:   20,
					ZFrom: -20,
					ZTo:   20,
					Flip:  off,
				},
				overlapBox: instruction{
					XFrom: -10,
					XTo:   20,
					YFrom: -10,
					YTo:   10,
					ZFrom: -10,
					ZTo:   20,
					Flip:  on,
				},
			},
			want: nil,
		},
		{
			name: "returns top right edge box if overlapbox is not at the top or front edge",
			args: args{
				box: instruction{
					XFrom: -20,
					XTo:   20,
					YFrom: -20,
					YTo:   20,
					ZFrom: -20,
					ZTo:   20,
					Flip:  off,
				},
				overlapBox: instruction{
					XFrom: -10,
					XTo:   10,
					YFrom: -10,
					YTo:   10,
					ZFrom: -10,
					ZTo:   10,
					Flip:  on,
				},
			},
			want: []instruction{
				{
					XFrom: 11,
					XTo:   20,
					YFrom: -10,
					YTo:   10,
					ZFrom: 11,
					ZTo:   20,
					Flip:  off,
				},
			},
		},
		{
			name: "returns single line top right edge box",
			args: args{
				box: instruction{
					XFrom: -20,
					XTo:   20,
					YFrom: -20,
					YTo:   20,
					ZFrom: -20,
					ZTo:   20,
					Flip:  off,
				},
				overlapBox: instruction{
					XFrom: -10,
					XTo:   19,
					YFrom: -10,
					YTo:   10,
					ZFrom: -10,
					ZTo:   19,
					Flip:  on,
				},
			},
			want: []instruction{
				{
					XFrom: 20,
					XTo:   20,
					YFrom: -10,
					YTo:   10,
					ZFrom: 20,
					ZTo:   20,
					Flip:  off,
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
					XFrom: -20,
					XTo:   20,
					YFrom: -20,
					YTo:   20,
					ZFrom: -20,
					ZTo:   20,
					Flip:  off,
				},
				overlapBox: instruction{
					XFrom: -10,
					XTo:   10,
					YFrom: -10,
					YTo:   10,
					ZFrom: -20,
					ZTo:   10,
					Flip:  on,
				},
			},
			want: nil,
		},
		{
			name: "returns nil if overlapbox is at the left edge",
			args: args{
				box: instruction{
					XFrom: -20,
					XTo:   20,
					YFrom: -20,
					YTo:   20,
					ZFrom: -20,
					ZTo:   20,
					Flip:  off,
				},
				overlapBox: instruction{
					XFrom: -10,
					XTo:   10,
					YFrom: -20,
					YTo:   10,
					ZFrom: -10,
					ZTo:   10,
					Flip:  on,
				},
			},
			want: nil,
		},
		{
			name: "returns nil if overlapbox is at both the bottom and the left edge",
			args: args{
				box: instruction{
					XFrom: -20,
					XTo:   20,
					YFrom: -20,
					YTo:   20,
					ZFrom: -20,
					ZTo:   20,
					Flip:  off,
				},
				overlapBox: instruction{
					XFrom: -10,
					XTo:   10,
					YFrom: -20,
					YTo:   10,
					ZFrom: -20,
					ZTo:   10,
					Flip:  on,
				},
			},
			want: nil,
		},
		{
			name: "returns bottom left edge box if overlapbox is not at the bottom or left edges",
			args: args{
				box: instruction{
					XFrom: -20,
					XTo:   20,
					YFrom: -20,
					YTo:   20,
					ZFrom: -20,
					ZTo:   20,
					Flip:  off,
				},
				overlapBox: instruction{
					XFrom: -10,
					XTo:   10,
					YFrom: -10,
					YTo:   10,
					ZFrom: -10,
					ZTo:   10,
					Flip:  on,
				},
			},
			want: []instruction{
				{
					XFrom: -10,
					XTo:   10,
					YFrom: -20,
					YTo:   -11,
					ZFrom: -20,
					ZTo:   -11,
					Flip:  off,
				},
			},
		},
		{
			name: "returns single line bottom left edge",
			args: args{
				box: instruction{
					XFrom: -20,
					XTo:   20,
					YFrom: -20,
					YTo:   20,
					ZFrom: -20,
					ZTo:   20,
					Flip:  off,
				},
				overlapBox: instruction{
					XFrom: -10,
					XTo:   10,
					YFrom: -19,
					YTo:   10,
					ZFrom: -19,
					ZTo:   10,
					Flip:  on,
				},
			},
			want: []instruction{
				{
					XFrom: -10,
					XTo:   10,
					YFrom: -20,
					YTo:   -20,
					ZFrom: -20,
					ZTo:   -20,
					Flip:  off,
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
					XFrom: -20,
					XTo:   20,
					YFrom: -20,
					YTo:   20,
					ZFrom: -20,
					ZTo:   20,
					Flip:  off,
				},
				overlapBox: instruction{
					XFrom: -10,
					XTo:   10,
					YFrom: -10,
					YTo:   10,
					ZFrom: -20,
					ZTo:   10,
					Flip:  on,
				},
			},
			want: nil,
		},
		{
			name: "returns nil if overlapbox is at the back edge",
			args: args{
				box: instruction{
					XFrom: -20,
					XTo:   20,
					YFrom: -20,
					YTo:   20,
					ZFrom: -20,
					ZTo:   20,
					Flip:  off,
				},
				overlapBox: instruction{
					XFrom: -20,
					XTo:   10,
					YFrom: -10,
					YTo:   10,
					ZFrom: -10,
					ZTo:   10,
					Flip:  on,
				},
			},
			want: nil,
		},
		{
			name: "returns nil if overlapbox is at both the bottom and the back edge",
			args: args{
				box: instruction{
					XFrom: -20,
					XTo:   20,
					YFrom: -20,
					YTo:   20,
					ZFrom: -20,
					ZTo:   20,
					Flip:  off,
				},
				overlapBox: instruction{
					XFrom: -20,
					XTo:   10,
					YFrom: -10,
					YTo:   10,
					ZFrom: -20,
					ZTo:   10,
					Flip:  on,
				},
			},
			want: nil,
		},
		{
			name: "returns bottom right edge box if overlapbox is not at the bottom or back edge",
			args: args{
				box: instruction{
					XFrom: -20,
					XTo:   20,
					YFrom: -20,
					YTo:   20,
					ZFrom: -20,
					ZTo:   20,
					Flip:  off,
				},
				overlapBox: instruction{
					XFrom: -10,
					XTo:   10,
					YFrom: -10,
					YTo:   10,
					ZFrom: -10,
					ZTo:   10,
					Flip:  on,
				},
			},
			want: []instruction{
				{
					XFrom: -20,
					XTo:   -11,
					YFrom: -10,
					YTo:   10,
					ZFrom: -20,
					ZTo:   -11,
					Flip:  off,
				},
			},
		},
		{
			name: "returns single line bottom right edge",
			args: args{
				box: instruction{
					XFrom: -20,
					XTo:   20,
					YFrom: -20,
					YTo:   20,
					ZFrom: -20,
					ZTo:   20,
					Flip:  off,
				},
				overlapBox: instruction{
					XFrom: -19,
					XTo:   10,
					YFrom: -10,
					YTo:   10,
					ZFrom: -19,
					ZTo:   10,
					Flip:  on,
				},
			},
			want: []instruction{
				{
					XFrom: -20,
					XTo:   -20,
					YFrom: -10,
					YTo:   10,
					ZFrom: -20,
					ZTo:   -20,
					Flip:  off,
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
					XFrom: -20,
					XTo:   20,
					YFrom: -20,
					YTo:   20,
					ZFrom: -20,
					ZTo:   20,
					Flip:  off,
				},
				overlapBox: instruction{
					XFrom: -10,
					XTo:   10,
					YFrom: -10,
					YTo:   10,
					ZFrom: -20,
					ZTo:   10,
					Flip:  on,
				},
			},
			want: nil,
		},
		{
			name: "returns nil if overlapbox is at the right edge",
			args: args{
				box: instruction{
					XFrom: -20,
					XTo:   20,
					YFrom: -20,
					YTo:   20,
					ZFrom: -20,
					ZTo:   20,
					Flip:  off,
				},
				overlapBox: instruction{
					XFrom: -10,
					XTo:   10,
					YFrom: -10,
					YTo:   20,
					ZFrom: -10,
					ZTo:   10,
					Flip:  on,
				},
			},
			want: nil,
		},
		{
			name: "returns nil if overlapbox is at both the bottom and the right edge",
			args: args{
				box: instruction{
					XFrom: -20,
					XTo:   20,
					YFrom: -20,
					YTo:   20,
					ZFrom: -20,
					ZTo:   20,
					Flip:  off,
				},
				overlapBox: instruction{
					XFrom: -10,
					XTo:   10,
					YFrom: -10,
					YTo:   20,
					ZFrom: -20,
					ZTo:   10,
					Flip:  on,
				},
			},
			want: nil,
		},
		{
			name: "returns bottom left edge box if overlapbox is not at the bottom or left edges",
			args: args{
				box: instruction{
					XFrom: -20,
					XTo:   20,
					YFrom: -20,
					YTo:   20,
					ZFrom: -20,
					ZTo:   20,
					Flip:  off,
				},
				overlapBox: instruction{
					XFrom: -10,
					XTo:   10,
					YFrom: -10,
					YTo:   10,
					ZFrom: -10,
					ZTo:   10,
					Flip:  on,
				},
			},
			want: []instruction{
				{
					XFrom: -10,
					XTo:   10,
					YFrom: 11,
					YTo:   20,
					ZFrom: -20,
					ZTo:   -11,
					Flip:  off,
				},
			},
		},
		{
			name: "returns single line bottom left edge",
			args: args{
				box: instruction{
					XFrom: -20,
					XTo:   20,
					YFrom: -20,
					YTo:   20,
					ZFrom: -20,
					ZTo:   20,
					Flip:  off,
				},
				overlapBox: instruction{
					XFrom: -10,
					XTo:   10,
					YFrom: -10,
					YTo:   19,
					ZFrom: -19,
					ZTo:   10,
					Flip:  on,
				},
			},
			want: []instruction{
				{
					XFrom: -10,
					XTo:   10,
					YFrom: 20,
					YTo:   20,
					ZFrom: -20,
					ZTo:   -20,
					Flip:  off,
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
					XFrom: -20,
					XTo:   20,
					YFrom: -20,
					YTo:   20,
					ZFrom: -20,
					ZTo:   20,
					Flip:  off,
				},
				overlapBox: instruction{
					XFrom: -10,
					XTo:   10,
					YFrom: -10,
					YTo:   10,
					ZFrom: -20,
					ZTo:   10,
					Flip:  on,
				},
			},
			want: nil,
		},
		{
			name: "returns nil if overlapbox is at the front edge",
			args: args{
				box: instruction{
					XFrom: -20,
					XTo:   20,
					YFrom: -20,
					YTo:   20,
					ZFrom: -20,
					ZTo:   20,
					Flip:  off,
				},
				overlapBox: instruction{
					XFrom: -10,
					XTo:   20,
					YFrom: -10,
					YTo:   10,
					ZFrom: -10,
					ZTo:   10,
					Flip:  on,
				},
			},
			want: nil,
		},
		{
			name: "returns nil if overlapbox is at both the bottom and the front edge",
			args: args{
				box: instruction{
					XFrom: -20,
					XTo:   20,
					YFrom: -20,
					YTo:   20,
					ZFrom: -20,
					ZTo:   20,
					Flip:  off,
				},
				overlapBox: instruction{
					XFrom: -10,
					XTo:   20,
					YFrom: -10,
					YTo:   10,
					ZFrom: -20,
					ZTo:   10,
					Flip:  on,
				},
			},
			want: nil,
		},
		{
			name: "returns bottom right edge box if overlapbox is not at the bottom or front edge",
			args: args{
				box: instruction{
					XFrom: -20,
					XTo:   20,
					YFrom: -20,
					YTo:   20,
					ZFrom: -20,
					ZTo:   20,
					Flip:  off,
				},
				overlapBox: instruction{
					XFrom: -10,
					XTo:   10,
					YFrom: -10,
					YTo:   10,
					ZFrom: -10,
					ZTo:   10,
					Flip:  on,
				},
			},
			want: []instruction{
				{
					XFrom: 11,
					XTo:   20,
					YFrom: -10,
					YTo:   10,
					ZFrom: -20,
					ZTo:   -11,
					Flip:  off,
				},
			},
		},
		{
			name: "returns single line bottom right edge",
			args: args{
				box: instruction{
					XFrom: -20,
					XTo:   20,
					YFrom: -20,
					YTo:   20,
					ZFrom: -20,
					ZTo:   20,
					Flip:  off,
				},
				overlapBox: instruction{
					XFrom: -10,
					XTo:   10,
					YFrom: -10,
					YTo:   10,
					ZFrom: -10,
					ZTo:   10,
					Flip:  on,
				},
			},
			want: []instruction{
				{
					XFrom: 11,
					XTo:   20,
					YFrom: -10,
					YTo:   10,
					ZFrom: -20,
					ZTo:   -11,
					Flip:  off,
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
					XFrom: -20,
					XTo:   20,
					YFrom: -20,
					YTo:   20,
					ZFrom: -20,
					ZTo:   20,
					Flip:  off,
				},
				overlapBox: instruction{
					XFrom: -10,
					XTo:   20,
					YFrom: -10,
					YTo:   10,
					ZFrom: -10,
					ZTo:   10,
					Flip:  on,
				},
			},
			want: nil,
		},
		{
			name: "returns nil if overlapbox is at the left edge",
			args: args{
				box: instruction{
					XFrom: -20,
					XTo:   20,
					YFrom: -20,
					YTo:   20,
					ZFrom: -20,
					ZTo:   20,
					Flip:  off,
				},
				overlapBox: instruction{
					XFrom: -10,
					XTo:   10,
					YFrom: -20,
					YTo:   10,
					ZFrom: -10,
					ZTo:   10,
					Flip:  on,
				},
			},
			want: nil,
		},
		{
			name: "returns nil if overlapbox is at both the front and the left edge",
			args: args{
				box: instruction{
					XFrom: -20,
					XTo:   20,
					YFrom: -20,
					YTo:   20,
					ZFrom: -20,
					ZTo:   20,
					Flip:  off,
				},
				overlapBox: instruction{
					XFrom: -10,
					XTo:   20,
					YFrom: -20,
					YTo:   10,
					ZFrom: -10,
					ZTo:   10,
					Flip:  on,
				},
			},
			want: nil,
		},
		{
			name: "returns front left edge box if overlapbox is not at the front or left edges",
			args: args{
				box: instruction{
					XFrom: -20,
					XTo:   20,
					YFrom: -20,
					YTo:   20,
					ZFrom: -20,
					ZTo:   20,
					Flip:  off,
				},
				overlapBox: instruction{
					XFrom: -10,
					XTo:   10,
					YFrom: -10,
					YTo:   10,
					ZFrom: -10,
					ZTo:   10,
					Flip:  on,
				},
			},
			want: []instruction{
				{
					XFrom: 11,
					XTo:   20,
					YFrom: -20,
					YTo:   -11,
					ZFrom: -10,
					ZTo:   10,
					Flip:  off,
				},
			},
		},
		{
			name: "returns single line front left edge",
			args: args{
				box: instruction{
					XFrom: -20,
					XTo:   20,
					YFrom: -20,
					YTo:   20,
					ZFrom: -20,
					ZTo:   20,
					Flip:  off,
				},
				overlapBox: instruction{
					XFrom: -10,
					XTo:   19,
					YFrom: -19,
					YTo:   10,
					ZFrom: -10,
					ZTo:   10,
					Flip:  on,
				},
			},
			want: []instruction{
				{
					XFrom: 20,
					XTo:   20,
					YFrom: -20,
					YTo:   -20,
					ZFrom: -10,
					ZTo:   10,
					Flip:  off,
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
					XFrom: -20,
					XTo:   20,
					YFrom: -20,
					YTo:   20,
					ZFrom: -20,
					ZTo:   20,
					Flip:  off,
				},
				overlapBox: instruction{
					XFrom: -10,
					XTo:   20,
					YFrom: -10,
					YTo:   10,
					ZFrom: -10,
					ZTo:   10,
					Flip:  on,
				},
			},
			want: nil,
		},
		{
			name: "returns nil if overlapbox is at the right edge",
			args: args{
				box: instruction{
					XFrom: -20,
					XTo:   20,
					YFrom: -20,
					YTo:   20,
					ZFrom: -20,
					ZTo:   20,
					Flip:  off,
				},
				overlapBox: instruction{
					XFrom: -10,
					XTo:   10,
					YFrom: -10,
					YTo:   20,
					ZFrom: -10,
					ZTo:   10,
					Flip:  on,
				},
			},
			want: nil,
		},
		{
			name: "returns nil if overlapbox is at both the front and the right edge",
			args: args{
				box: instruction{
					XFrom: -20,
					XTo:   20,
					YFrom: -20,
					YTo:   20,
					ZFrom: -20,
					ZTo:   20,
					Flip:  off,
				},
				overlapBox: instruction{
					XFrom: -10,
					XTo:   20,
					YFrom: -10,
					YTo:   20,
					ZFrom: -10,
					ZTo:   10,
					Flip:  on,
				},
			},
			want: nil,
		},
		{
			name: "returns front right edge box if overlapbox is not at the front or right edges",
			args: args{
				box: instruction{
					XFrom: -20,
					XTo:   20,
					YFrom: -20,
					YTo:   20,
					ZFrom: -20,
					ZTo:   20,
					Flip:  off,
				},
				overlapBox: instruction{
					XFrom: -10,
					XTo:   10,
					YFrom: -10,
					YTo:   10,
					ZFrom: -10,
					ZTo:   10,
					Flip:  on,
				},
			},
			want: []instruction{
				{
					XFrom: 11,
					XTo:   20,
					YFrom: 11,
					YTo:   20,
					ZFrom: -10,
					ZTo:   10,
					Flip:  off,
				},
			},
		},
		{
			name: "returns single front right edge",
			args: args{
				box: instruction{
					XFrom: -20,
					XTo:   20,
					YFrom: -20,
					YTo:   20,
					ZFrom: -20,
					ZTo:   20,
					Flip:  off,
				},
				overlapBox: instruction{
					XFrom: -10,
					XTo:   19,
					YFrom: -10,
					YTo:   19,
					ZFrom: -10,
					ZTo:   10,
					Flip:  on,
				},
			},
			want: []instruction{
				{
					XFrom: 20,
					XTo:   20,
					YFrom: 20,
					YTo:   20,
					ZFrom: -10,
					ZTo:   10,
					Flip:  off,
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
					XFrom: -20,
					XTo:   20,
					YFrom: -20,
					YTo:   20,
					ZFrom: -20,
					ZTo:   20,
					Flip:  off,
				},
				overlapBox: instruction{
					XFrom: -20,
					XTo:   10,
					YFrom: -10,
					YTo:   10,
					ZFrom: -10,
					ZTo:   10,
					Flip:  on,
				},
			},
			want: nil,
		},
		{
			name: "returns nil if overlapbox is at the left edge",
			args: args{
				box: instruction{
					XFrom: -20,
					XTo:   20,
					YFrom: -20,
					YTo:   20,
					ZFrom: -20,
					ZTo:   20,
					Flip:  off,
				},
				overlapBox: instruction{
					XFrom: -10,
					XTo:   10,
					YFrom: -20,
					YTo:   10,
					ZFrom: -10,
					ZTo:   10,
					Flip:  on,
				},
			},
			want: nil,
		},
		{
			name: "returns nil if overlapbox is at both the back and the left edge",
			args: args{
				box: instruction{
					XFrom: -20,
					XTo:   20,
					YFrom: -20,
					YTo:   20,
					ZFrom: -20,
					ZTo:   20,
					Flip:  off,
				},
				overlapBox: instruction{
					XFrom: -20,
					XTo:   10,
					YFrom: -20,
					YTo:   10,
					ZFrom: -10,
					ZTo:   10,
					Flip:  on,
				},
			},
			want: nil,
		},
		{
			name: "returns back left edge box if overlapbox is not at the back or left edges",
			args: args{
				box: instruction{
					XFrom: -20,
					XTo:   20,
					YFrom: -20,
					YTo:   20,
					ZFrom: -20,
					ZTo:   20,
					Flip:  off,
				},
				overlapBox: instruction{
					XFrom: -10,
					XTo:   10,
					YFrom: -10,
					YTo:   10,
					ZFrom: -10,
					ZTo:   10,
					Flip:  on,
				},
			},
			want: []instruction{
				{
					XFrom: -20,
					XTo:   -11,
					YFrom: -20,
					YTo:   -11,
					ZFrom: -10,
					ZTo:   10,
					Flip:  off,
				},
			},
		},
		{
			name: "returns single line back left edge",
			args: args{
				box: instruction{
					XFrom: -20,
					XTo:   20,
					YFrom: -20,
					YTo:   20,
					ZFrom: -20,
					ZTo:   20,
					Flip:  off,
				},
				overlapBox: instruction{
					XFrom: -19,
					XTo:   10,
					YFrom: -19,
					YTo:   10,
					ZFrom: -10,
					ZTo:   10,
					Flip:  on,
				},
			},
			want: []instruction{
				{
					XFrom: -20,
					XTo:   -20,
					YFrom: -20,
					YTo:   -20,
					ZFrom: -10,
					ZTo:   10,
					Flip:  off,
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
					XFrom: -20,
					XTo:   20,
					YFrom: -20,
					YTo:   20,
					ZFrom: -20,
					ZTo:   20,
					Flip:  off,
				},
				overlapBox: instruction{
					XFrom: -20,
					XTo:   10,
					YFrom: -10,
					YTo:   10,
					ZFrom: -10,
					ZTo:   10,
					Flip:  on,
				},
			},
			want: nil,
		},
		{
			name: "returns nil if overlapbox is at the right edge",
			args: args{
				box: instruction{
					XFrom: -20,
					XTo:   20,
					YFrom: -20,
					YTo:   20,
					ZFrom: -20,
					ZTo:   20,
					Flip:  off,
				},
				overlapBox: instruction{
					XFrom: -10,
					XTo:   10,
					YFrom: -10,
					YTo:   20,
					ZFrom: -10,
					ZTo:   10,
					Flip:  on,
				},
			},
			want: nil,
		},
		{
			name: "returns nil if overlapbox is at both the back and the right edge",
			args: args{
				box: instruction{
					XFrom: -20,
					XTo:   20,
					YFrom: -20,
					YTo:   20,
					ZFrom: -20,
					ZTo:   20,
					Flip:  off,
				},
				overlapBox: instruction{
					XFrom: -20,
					XTo:   10,
					YFrom: -10,
					YTo:   20,
					ZFrom: -10,
					ZTo:   10,
					Flip:  on,
				},
			},
			want: nil,
		},
		{
			name: "returns back right edge box if overlapbox is not at the back or right edges",
			args: args{
				box: instruction{
					XFrom: -20,
					XTo:   20,
					YFrom: -20,
					YTo:   20,
					ZFrom: -20,
					ZTo:   20,
					Flip:  off,
				},
				overlapBox: instruction{
					XFrom: -10,
					XTo:   10,
					YFrom: -10,
					YTo:   10,
					ZFrom: -10,
					ZTo:   10,
					Flip:  on,
				},
			},
			want: []instruction{
				{
					XFrom: -20,
					XTo:   -11,
					YFrom: 11,
					YTo:   20,
					ZFrom: -10,
					ZTo:   10,
					Flip:  off,
				},
			},
		},
		{
			name: "returns single line back right edge",
			args: args{
				box: instruction{
					XFrom: -20,
					XTo:   20,
					YFrom: -20,
					YTo:   20,
					ZFrom: -20,
					ZTo:   20,
					Flip:  off,
				},
				overlapBox: instruction{
					XFrom: -19,
					XTo:   10,
					YFrom: -10,
					YTo:   19,
					ZFrom: -10,
					ZTo:   10,
					Flip:  on,
				},
			},
			want: []instruction{
				{
					XFrom: -20,
					XTo:   -20,
					YFrom: 20,
					YTo:   20,
					ZFrom: -10,
					ZTo:   10,
					Flip:  off,
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
					XFrom: -20,
					XTo:   20,
					YFrom: -20,
					YTo:   20,
					ZFrom: -20,
					ZTo:   20,
					Flip:  off,
				},
				overlapBox: instruction{
					XFrom: -20,
					XTo:   10,
					YFrom: -10,
					YTo:   10,
					ZFrom: -10,
					ZTo:   10,
					Flip:  on,
				},
			},
			want: nil,
		},
		{
			name: "returns nil if overlapbox is at the top edge",
			args: args{
				box: instruction{
					XFrom: -20,
					XTo:   20,
					YFrom: -20,
					YTo:   20,
					ZFrom: -20,
					ZTo:   20,
					Flip:  off,
				},
				overlapBox: instruction{
					XFrom: -10,
					XTo:   10,
					YFrom: -10,
					YTo:   10,
					ZFrom: -10,
					ZTo:   20,
					Flip:  on,
				},
			},
			want: nil,
		},
		{
			name: "returns nil if overlapbox is at the left edge",
			args: args{
				box: instruction{
					XFrom: -20,
					XTo:   20,
					YFrom: -20,
					YTo:   20,
					ZFrom: -20,
					ZTo:   20,
					Flip:  off,
				},
				overlapBox: instruction{
					XFrom: -10,
					XTo:   10,
					YFrom: -20,
					YTo:   10,
					ZFrom: -10,
					ZTo:   10,
					Flip:  on,
				},
			},
			want: nil,
		},
		{
			name: "returns nil if overlapbox is at both the back and the left edge",
			args: args{
				box: instruction{
					XFrom: -20,
					XTo:   20,
					YFrom: -20,
					YTo:   20,
					ZFrom: -20,
					ZTo:   20,
					Flip:  off,
				},
				overlapBox: instruction{
					XFrom: -20,
					XTo:   10,
					YFrom: -20,
					YTo:   10,
					ZFrom: -10,
					ZTo:   10,
					Flip:  on,
				},
			},
			want: nil,
		},
		{
			name: "returns nil if overlapbox is at both the top and the left edge",
			args: args{
				box: instruction{
					XFrom: -20,
					XTo:   20,
					YFrom: -20,
					YTo:   20,
					ZFrom: -20,
					ZTo:   20,
					Flip:  off,
				},
				overlapBox: instruction{
					XFrom: -10,
					XTo:   10,
					YFrom: -20,
					YTo:   10,
					ZFrom: -10,
					ZTo:   20,
					Flip:  on,
				},
			},
			want: nil,
		},
		{
			name: "returns nil if overlapbox is at both the top and the back edge",
			args: args{
				box: instruction{
					XFrom: -20,
					XTo:   20,
					YFrom: -20,
					YTo:   20,
					ZFrom: -20,
					ZTo:   20,
					Flip:  off,
				},
				overlapBox: instruction{
					XFrom: -20,
					XTo:   10,
					YFrom: -10,
					YTo:   10,
					ZFrom: -10,
					ZTo:   20,
					Flip:  on,
				},
			},
			want: nil,
		},
		{
			name: "returns top back right corner box if overlapbox is not at the top, back, or left edges",
			args: args{
				box: instruction{
					XFrom: -20,
					XTo:   20,
					YFrom: -20,
					YTo:   20,
					ZFrom: -20,
					ZTo:   20,
					Flip:  off,
				},
				overlapBox: instruction{
					XFrom: -10,
					XTo:   10,
					YFrom: -10,
					YTo:   10,
					ZFrom: -10,
					ZTo:   10,
					Flip:  on,
				},
			},
			want: []instruction{
				{
					XFrom: -20,
					XTo:   -11,
					YFrom: -20,
					YTo:   -11,
					ZFrom: 11,
					ZTo:   20,
					Flip:  off,
				},
			},
		},
		{
			name: "returns single top back right corner cube",
			args: args{
				box: instruction{
					XFrom: -20,
					XTo:   20,
					YFrom: -20,
					YTo:   20,
					ZFrom: -20,
					ZTo:   20,
					Flip:  off,
				},
				overlapBox: instruction{
					XFrom: -19,
					XTo:   10,
					YFrom: -19,
					YTo:   10,
					ZFrom: -10,
					ZTo:   19,
					Flip:  on,
				},
			},
			want: []instruction{
				{
					XFrom: -20,
					XTo:   -20,
					YFrom: -20,
					YTo:   -20,
					ZFrom: 20,
					ZTo:   20,
					Flip:  off,
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
					XFrom: -20,
					XTo:   20,
					YFrom: -20,
					YTo:   20,
					ZFrom: -20,
					ZTo:   20,
					Flip:  off,
				},
				overlapBox: instruction{
					XFrom: -20,
					XTo:   10,
					YFrom: -10,
					YTo:   10,
					ZFrom: -10,
					ZTo:   10,
					Flip:  on,
				},
			},
			want: nil,
		},
		{
			name: "returns nil if overlapbox is at the top edge",
			args: args{
				box: instruction{
					XFrom: -20,
					XTo:   20,
					YFrom: -20,
					YTo:   20,
					ZFrom: -20,
					ZTo:   20,
					Flip:  off,
				},
				overlapBox: instruction{
					XFrom: -10,
					XTo:   10,
					YFrom: -10,
					YTo:   10,
					ZFrom: -10,
					ZTo:   20,
					Flip:  on,
				},
			},
			want: nil,
		},
		{
			name: "returns nil if overlapbox is at the right edge",
			args: args{
				box: instruction{
					XFrom: -20,
					XTo:   20,
					YFrom: -20,
					YTo:   20,
					ZFrom: -20,
					ZTo:   20,
					Flip:  off,
				},
				overlapBox: instruction{
					XFrom: -10,
					XTo:   10,
					YFrom: -10,
					YTo:   20,
					ZFrom: -10,
					ZTo:   10,
					Flip:  on,
				},
			},
			want: nil,
		},
		{
			name: "returns nil if overlapbox is at both the back and the right edge",
			args: args{
				box: instruction{
					XFrom: -20,
					XTo:   20,
					YFrom: -20,
					YTo:   20,
					ZFrom: -20,
					ZTo:   20,
					Flip:  off,
				},
				overlapBox: instruction{
					XFrom: -20,
					XTo:   10,
					YFrom: -10,
					YTo:   20,
					ZFrom: -10,
					ZTo:   10,
					Flip:  on,
				},
			},
			want: nil,
		},
		{
			name: "returns nil if overlapbox is at both the top and the right edge",
			args: args{
				box: instruction{
					XFrom: -20,
					XTo:   20,
					YFrom: -20,
					YTo:   20,
					ZFrom: -20,
					ZTo:   20,
					Flip:  off,
				},
				overlapBox: instruction{
					XFrom: -10,
					XTo:   10,
					YFrom: -10,
					YTo:   20,
					ZFrom: -10,
					ZTo:   20,
					Flip:  on,
				},
			},
			want: nil,
		},
		{
			name: "returns nil if overlapbox is at both the top and the back edge",
			args: args{
				box: instruction{
					XFrom: -20,
					XTo:   20,
					YFrom: -20,
					YTo:   20,
					ZFrom: -20,
					ZTo:   20,
					Flip:  off,
				},
				overlapBox: instruction{
					XFrom: -20,
					XTo:   10,
					YFrom: -10,
					YTo:   10,
					ZFrom: -10,
					ZTo:   20,
					Flip:  on,
				},
			},
			want: nil,
		},
		{
			name: "returns top back right corner box if overlapbox is not at the top, back, or right edges",
			args: args{
				box: instruction{
					XFrom: -20,
					XTo:   20,
					YFrom: -20,
					YTo:   20,
					ZFrom: -20,
					ZTo:   20,
					Flip:  off,
				},
				overlapBox: instruction{
					XFrom: -10,
					XTo:   10,
					YFrom: -10,
					YTo:   10,
					ZFrom: -10,
					ZTo:   10,
					Flip:  on,
				},
			},
			want: []instruction{
				{
					XFrom: -20,
					XTo:   -11,
					YFrom: 11,
					YTo:   20,
					ZFrom: 11,
					ZTo:   20,
					Flip:  off,
				},
			},
		},
		{
			name: "returns single top back right corner",
			args: args{
				box: instruction{
					XFrom: -20,
					XTo:   20,
					YFrom: -20,
					YTo:   20,
					ZFrom: -20,
					ZTo:   20,
					Flip:  off,
				},
				overlapBox: instruction{
					XFrom: -19,
					XTo:   10,
					YFrom: -10,
					YTo:   19,
					ZFrom: -10,
					ZTo:   19,
					Flip:  on,
				},
			},
			want: []instruction{
				{
					XFrom: -20,
					XTo:   -20,
					YFrom: 20,
					YTo:   20,
					ZFrom: 20,
					ZTo:   20,
					Flip:  off,
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
					XFrom: -20,
					XTo:   20,
					YFrom: -20,
					YTo:   20,
					ZFrom: -20,
					ZTo:   20,
					Flip:  off,
				},
				overlapBox: instruction{
					XFrom: -10,
					XTo:   20,
					YFrom: -10,
					YTo:   10,
					ZFrom: -10,
					ZTo:   10,
					Flip:  on,
				},
			},
			want: nil,
		},
		{
			name: "returns nil if overlapbox is at the top edge",
			args: args{
				box: instruction{
					XFrom: -20,
					XTo:   20,
					YFrom: -20,
					YTo:   20,
					ZFrom: -20,
					ZTo:   20,
					Flip:  off,
				},
				overlapBox: instruction{
					XFrom: -10,
					XTo:   10,
					YFrom: -10,
					YTo:   10,
					ZFrom: -10,
					ZTo:   20,
					Flip:  on,
				},
			},
			want: nil,
		},
		{
			name: "returns nil if overlapbox is at the left edge",
			args: args{
				box: instruction{
					XFrom: -20,
					XTo:   20,
					YFrom: -20,
					YTo:   20,
					ZFrom: -20,
					ZTo:   20,
					Flip:  off,
				},
				overlapBox: instruction{
					XFrom: -10,
					XTo:   10,
					YFrom: -20,
					YTo:   10,
					ZFrom: -10,
					ZTo:   10,
					Flip:  on,
				},
			},
			want: nil,
		},
		{
			name: "returns nil if overlapbox is at both the front and the left edge",
			args: args{
				box: instruction{
					XFrom: -20,
					XTo:   20,
					YFrom: -20,
					YTo:   20,
					ZFrom: -20,
					ZTo:   20,
					Flip:  off,
				},
				overlapBox: instruction{
					XFrom: -10,
					XTo:   20,
					YFrom: -20,
					YTo:   10,
					ZFrom: -10,
					ZTo:   10,
					Flip:  on,
				},
			},
			want: nil,
		},
		{
			name: "returns nil if overlapbox is at both the top and the left edge",
			args: args{
				box: instruction{
					XFrom: -20,
					XTo:   20,
					YFrom: -20,
					YTo:   20,
					ZFrom: -20,
					ZTo:   20,
					Flip:  off,
				},
				overlapBox: instruction{
					XFrom: -10,
					XTo:   10,
					YFrom: -20,
					YTo:   10,
					ZFrom: -10,
					ZTo:   20,
					Flip:  on,
				},
			},
			want: nil,
		},
		{
			name: "returns nil if overlapbox is at both the top and the front edge",
			args: args{
				box: instruction{
					XFrom: -20,
					XTo:   20,
					YFrom: -20,
					YTo:   20,
					ZFrom: -20,
					ZTo:   20,
					Flip:  off,
				},
				overlapBox: instruction{
					XFrom: -10,
					XTo:   20,
					YFrom: -10,
					YTo:   10,
					ZFrom: -10,
					ZTo:   20,
					Flip:  on,
				},
			},
			want: nil,
		},
		{
			name: "returns top front right corner box if overlapbox is not at the top, front, or left edges",
			args: args{
				box: instruction{
					XFrom: -20,
					XTo:   20,
					YFrom: -20,
					YTo:   20,
					ZFrom: -20,
					ZTo:   20,
					Flip:  off,
				},
				overlapBox: instruction{
					XFrom: -10,
					XTo:   10,
					YFrom: -10,
					YTo:   10,
					ZFrom: -10,
					ZTo:   10,
					Flip:  on,
				},
			},
			want: []instruction{
				{
					XFrom: 11,
					XTo:   20,
					YFrom: -20,
					YTo:   -11,
					ZFrom: 11,
					ZTo:   20,
					Flip:  off,
				},
			},
		},
		{
			name: "returns single top front right corner",
			args: args{
				box: instruction{
					XFrom: -20,
					XTo:   20,
					YFrom: -20,
					YTo:   20,
					ZFrom: -20,
					ZTo:   20,
					Flip:  off,
				},
				overlapBox: instruction{
					XFrom: -10,
					XTo:   19,
					YFrom: -19,
					YTo:   10,
					ZFrom: -10,
					ZTo:   19,
					Flip:  on,
				},
			},
			want: []instruction{
				{
					XFrom: 20,
					XTo:   20,
					YFrom: -20,
					YTo:   -20,
					ZFrom: 20,
					ZTo:   20,
					Flip:  off,
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
					XFrom: -20,
					XTo:   20,
					YFrom: -20,
					YTo:   20,
					ZFrom: -20,
					ZTo:   20,
					Flip:  off,
				},
				overlapBox: instruction{
					XFrom: -10,
					XTo:   20,
					YFrom: -10,
					YTo:   10,
					ZFrom: -10,
					ZTo:   10,
					Flip:  on,
				},
			},
			want: nil,
		},
		{
			name: "returns nil if overlapbox is at the top edge",
			args: args{
				box: instruction{
					XFrom: -20,
					XTo:   20,
					YFrom: -20,
					YTo:   20,
					ZFrom: -20,
					ZTo:   20,
					Flip:  off,
				},
				overlapBox: instruction{
					XFrom: -10,
					XTo:   10,
					YFrom: -10,
					YTo:   10,
					ZFrom: -10,
					ZTo:   20,
					Flip:  on,
				},
			},
			want: nil,
		},
		{
			name: "returns nil if overlapbox is at the right edge",
			args: args{
				box: instruction{
					XFrom: -20,
					XTo:   20,
					YFrom: -20,
					YTo:   20,
					ZFrom: -20,
					ZTo:   20,
					Flip:  off,
				},
				overlapBox: instruction{
					XFrom: -10,
					XTo:   10,
					YFrom: -10,
					YTo:   20,
					ZFrom: -10,
					ZTo:   10,
					Flip:  on,
				},
			},
			want: nil,
		},
		{
			name: "returns nil if overlapbox is at both the front and the right edge",
			args: args{
				box: instruction{
					XFrom: -20,
					XTo:   20,
					YFrom: -20,
					YTo:   20,
					ZFrom: -20,
					ZTo:   20,
					Flip:  off,
				},
				overlapBox: instruction{
					XFrom: -10,
					XTo:   20,
					YFrom: -10,
					YTo:   20,
					ZFrom: -10,
					ZTo:   10,
					Flip:  on,
				},
			},
			want: nil,
		},
		{
			name: "returns nil if overlapbox is at both the top and the right edge",
			args: args{
				box: instruction{
					XFrom: -20,
					XTo:   20,
					YFrom: -20,
					YTo:   20,
					ZFrom: -20,
					ZTo:   20,
					Flip:  off,
				},
				overlapBox: instruction{
					XFrom: -10,
					XTo:   10,
					YFrom: -10,
					YTo:   20,
					ZFrom: -10,
					ZTo:   20,
					Flip:  on,
				},
			},
			want: nil,
		},
		{
			name: "returns nil if overlapbox is at both the top and the front edge",
			args: args{
				box: instruction{
					XFrom: -20,
					XTo:   20,
					YFrom: -20,
					YTo:   20,
					ZFrom: -20,
					ZTo:   20,
					Flip:  off,
				},
				overlapBox: instruction{
					XFrom: -10,
					XTo:   20,
					YFrom: -10,
					YTo:   10,
					ZFrom: -10,
					ZTo:   20,
					Flip:  on,
				},
			},
			want: nil,
		},
		{
			name: "returns top front right corner box if overlapbox is not at the top, front, or right edges",
			args: args{
				box: instruction{
					XFrom: -20,
					XTo:   20,
					YFrom: -20,
					YTo:   20,
					ZFrom: -20,
					ZTo:   20,
					Flip:  off,
				},
				overlapBox: instruction{
					XFrom: -10,
					XTo:   10,
					YFrom: -10,
					YTo:   10,
					ZFrom: -10,
					ZTo:   10,
					Flip:  on,
				},
			},
			want: []instruction{
				{
					XFrom: 11,
					XTo:   20,
					YFrom: 11,
					YTo:   20,
					ZFrom: 11,
					ZTo:   20,
					Flip:  off,
				},
			},
		},
		{
			name: "returns single top front right corner",
			args: args{
				box: instruction{
					XFrom: -20,
					XTo:   20,
					YFrom: -20,
					YTo:   20,
					ZFrom: -20,
					ZTo:   20,
					Flip:  off,
				},
				overlapBox: instruction{
					XFrom: -10,
					XTo:   19,
					YFrom: -10,
					YTo:   19,
					ZFrom: -10,
					ZTo:   19,
					Flip:  on,
				},
			},
			want: []instruction{
				{
					XFrom: 20,
					XTo:   20,
					YFrom: 20,
					YTo:   20,
					ZFrom: 20,
					ZTo:   20,
					Flip:  off,
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
					XFrom: -20,
					XTo:   20,
					YFrom: -20,
					YTo:   20,
					ZFrom: -20,
					ZTo:   20,
					Flip:  off,
				},
				overlapBox: instruction{
					XFrom: -20,
					XTo:   10,
					YFrom: -10,
					YTo:   10,
					ZFrom: -10,
					ZTo:   10,
					Flip:  on,
				},
			},
			want: nil,
		},
		{
			name: "returns nil if overlapbox is at the bottom edge",
			args: args{
				box: instruction{
					XFrom: -20,
					XTo:   20,
					YFrom: -20,
					YTo:   20,
					ZFrom: -20,
					ZTo:   20,
					Flip:  off,
				},
				overlapBox: instruction{
					XFrom: -10,
					XTo:   10,
					YFrom: -10,
					YTo:   10,
					ZFrom: -20,
					ZTo:   10,
					Flip:  on,
				},
			},
			want: nil,
		},
		{
			name: "returns nil if overlapbox is at the left edge",
			args: args{
				box: instruction{
					XFrom: -20,
					XTo:   20,
					YFrom: -20,
					YTo:   20,
					ZFrom: -20,
					ZTo:   20,
					Flip:  off,
				},
				overlapBox: instruction{
					XFrom: -10,
					XTo:   10,
					YFrom: -20,
					YTo:   10,
					ZFrom: -10,
					ZTo:   10,
					Flip:  on,
				},
			},
			want: nil,
		},
		{
			name: "returns nil if overlapbox is at both the back and the left edge",
			args: args{
				box: instruction{
					XFrom: -20,
					XTo:   20,
					YFrom: -20,
					YTo:   20,
					ZFrom: -20,
					ZTo:   20,
					Flip:  off,
				},
				overlapBox: instruction{
					XFrom: -20,
					XTo:   10,
					YFrom: -20,
					YTo:   10,
					ZFrom: -10,
					ZTo:   10,
					Flip:  on,
				},
			},
			want: nil,
		},
		{
			name: "returns nil if overlapbox is at both the bottom and the left edge",
			args: args{
				box: instruction{
					XFrom: -20,
					XTo:   20,
					YFrom: -20,
					YTo:   20,
					ZFrom: -20,
					ZTo:   20,
					Flip:  off,
				},
				overlapBox: instruction{
					XFrom: -10,
					XTo:   10,
					YFrom: -20,
					YTo:   10,
					ZFrom: -20,
					ZTo:   10,
					Flip:  on,
				},
			},
			want: nil,
		},
		{
			name: "returns nil if overlapbox is at both the bottom and the back edge",
			args: args{
				box: instruction{
					XFrom: -20,
					XTo:   20,
					YFrom: -20,
					YTo:   20,
					ZFrom: -20,
					ZTo:   20,
					Flip:  off,
				},
				overlapBox: instruction{
					XFrom: -20,
					XTo:   10,
					YFrom: -10,
					YTo:   10,
					ZFrom: -20,
					ZTo:   10,
					Flip:  on,
				},
			},
			want: nil,
		},
		{
			name: "returns bottom back right corner box if overlapbox is not at the bottom, back, or left edges",
			args: args{
				box: instruction{
					XFrom: -20,
					XTo:   20,
					YFrom: -20,
					YTo:   20,
					ZFrom: -20,
					ZTo:   20,
					Flip:  off,
				},
				overlapBox: instruction{
					XFrom: -10,
					XTo:   10,
					YFrom: -10,
					YTo:   10,
					ZFrom: -10,
					ZTo:   10,
					Flip:  on,
				},
			},
			want: []instruction{
				{
					XFrom: -20,
					XTo:   -11,
					YFrom: -20,
					YTo:   -11,
					ZFrom: -20,
					ZTo:   -11,
					Flip:  off,
				},
			},
		},
		{
			name: "returns single bottom back right corner",
			args: args{
				box: instruction{
					XFrom: -20,
					XTo:   20,
					YFrom: -20,
					YTo:   20,
					ZFrom: -20,
					ZTo:   20,
					Flip:  off,
				},
				overlapBox: instruction{
					XFrom: -19,
					XTo:   10,
					YFrom: -19,
					YTo:   10,
					ZFrom: -19,
					ZTo:   10,
					Flip:  on,
				},
			},
			want: []instruction{
				{
					XFrom: -20,
					XTo:   -20,
					YFrom: -20,
					YTo:   -20,
					ZFrom: -20,
					ZTo:   -20,
					Flip:  off,
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
					XFrom: -20,
					XTo:   20,
					YFrom: -20,
					YTo:   20,
					ZFrom: -20,
					ZTo:   20,
					Flip:  off,
				},
				overlapBox: instruction{
					XFrom: -20,
					XTo:   10,
					YFrom: -10,
					YTo:   10,
					ZFrom: -10,
					ZTo:   10,
					Flip:  on,
				},
			},
			want: nil,
		},
		{
			name: "returns nil if overlapbox is at the bottom edge",
			args: args{
				box: instruction{
					XFrom: -20,
					XTo:   20,
					YFrom: -20,
					YTo:   20,
					ZFrom: -20,
					ZTo:   20,
					Flip:  off,
				},
				overlapBox: instruction{
					XFrom: -10,
					XTo:   10,
					YFrom: -10,
					YTo:   10,
					ZFrom: -20,
					ZTo:   10,
					Flip:  on,
				},
			},
			want: nil,
		},
		{
			name: "returns nil if overlapbox is at the right edge",
			args: args{
				box: instruction{
					XFrom: -20,
					XTo:   20,
					YFrom: -20,
					YTo:   20,
					ZFrom: -20,
					ZTo:   20,
					Flip:  off,
				},
				overlapBox: instruction{
					XFrom: -10,
					XTo:   10,
					YFrom: -10,
					YTo:   20,
					ZFrom: -10,
					ZTo:   10,
					Flip:  on,
				},
			},
			want: nil,
		},
		{
			name: "returns nil if overlapbox is at both the back and the right edge",
			args: args{
				box: instruction{
					XFrom: -20,
					XTo:   20,
					YFrom: -20,
					YTo:   20,
					ZFrom: -20,
					ZTo:   20,
					Flip:  off,
				},
				overlapBox: instruction{
					XFrom: -20,
					XTo:   10,
					YFrom: -10,
					YTo:   20,
					ZFrom: -10,
					ZTo:   10,
					Flip:  on,
				},
			},
			want: nil,
		},
		{
			name: "returns nil if overlapbox is at both the bottom and the right edge",
			args: args{
				box: instruction{
					XFrom: -20,
					XTo:   20,
					YFrom: -20,
					YTo:   20,
					ZFrom: -20,
					ZTo:   20,
					Flip:  off,
				},
				overlapBox: instruction{
					XFrom: -10,
					XTo:   10,
					YFrom: -10,
					YTo:   20,
					ZFrom: -20,
					ZTo:   10,
					Flip:  on,
				},
			},
			want: nil,
		},
		{
			name: "returns nil if overlapbox is at both the bottom and the back edge",
			args: args{
				box: instruction{
					XFrom: -20,
					XTo:   20,
					YFrom: -20,
					YTo:   20,
					ZFrom: -20,
					ZTo:   20,
					Flip:  off,
				},
				overlapBox: instruction{
					XFrom: -20,
					XTo:   10,
					YFrom: -10,
					YTo:   10,
					ZFrom: -20,
					ZTo:   10,
					Flip:  on,
				},
			},
			want: nil,
		},
		{
			name: "returns bottom back right corner box if overlapbox is not at the bottom, back, or right edges",
			args: args{
				box: instruction{
					XFrom: -20,
					XTo:   20,
					YFrom: -20,
					YTo:   20,
					ZFrom: -20,
					ZTo:   20,
					Flip:  off,
				},
				overlapBox: instruction{
					XFrom: -10,
					XTo:   10,
					YFrom: -10,
					YTo:   10,
					ZFrom: -10,
					ZTo:   10,
					Flip:  on,
				},
			},
			want: []instruction{
				{
					XFrom: -20,
					XTo:   -11,
					YFrom: 11,
					YTo:   20,
					ZFrom: -20,
					ZTo:   -11,
					Flip:  off,
				},
			},
		},
		{
			name: "returns single bottom back right corner box",
			args: args{
				box: instruction{
					XFrom: -20,
					XTo:   20,
					YFrom: -20,
					YTo:   20,
					ZFrom: -20,
					ZTo:   20,
					Flip:  off,
				},
				overlapBox: instruction{
					XFrom: -19,
					XTo:   10,
					YFrom: -10,
					YTo:   19,
					ZFrom: -19,
					ZTo:   10,
					Flip:  on,
				},
			},
			want: []instruction{
				{
					XFrom: -20,
					XTo:   -20,
					YFrom: 20,
					YTo:   20,
					ZFrom: -20,
					ZTo:   -20,
					Flip:  off,
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
					XFrom: -20,
					XTo:   20,
					YFrom: -20,
					YTo:   20,
					ZFrom: -20,
					ZTo:   20,
					Flip:  off,
				},
				overlapBox: instruction{
					XFrom: -10,
					XTo:   20,
					YFrom: -10,
					YTo:   10,
					ZFrom: -10,
					ZTo:   10,
					Flip:  on,
				},
			},
			want: nil,
		},
		{
			name: "returns nil if overlapbox is at the bottom edge",
			args: args{
				box: instruction{
					XFrom: -20,
					XTo:   20,
					YFrom: -20,
					YTo:   20,
					ZFrom: -20,
					ZTo:   20,
					Flip:  off,
				},
				overlapBox: instruction{
					XFrom: -10,
					XTo:   10,
					YFrom: -10,
					YTo:   10,
					ZFrom: -20,
					ZTo:   10,
					Flip:  on,
				},
			},
			want: nil,
		},
		{
			name: "returns nil if overlapbox is at the left edge",
			args: args{
				box: instruction{
					XFrom: -20,
					XTo:   20,
					YFrom: -20,
					YTo:   20,
					ZFrom: -20,
					ZTo:   20,
					Flip:  off,
				},
				overlapBox: instruction{
					XFrom: -10,
					XTo:   10,
					YFrom: -20,
					YTo:   10,
					ZFrom: -10,
					ZTo:   10,
					Flip:  on,
				},
			},
			want: nil,
		},
		{
			name: "returns nil if overlapbox is at both the front and the left edge",
			args: args{
				box: instruction{
					XFrom: -20,
					XTo:   20,
					YFrom: -20,
					YTo:   20,
					ZFrom: -20,
					ZTo:   20,
					Flip:  off,
				},
				overlapBox: instruction{
					XFrom: -10,
					XTo:   20,
					YFrom: -20,
					YTo:   10,
					ZFrom: -10,
					ZTo:   10,
					Flip:  on,
				},
			},
			want: nil,
		},
		{
			name: "returns nil if overlapbox is at both the bottom and the left edge",
			args: args{
				box: instruction{
					XFrom: -20,
					XTo:   20,
					YFrom: -20,
					YTo:   20,
					ZFrom: -20,
					ZTo:   20,
					Flip:  off,
				},
				overlapBox: instruction{
					XFrom: -10,
					XTo:   10,
					YFrom: -20,
					YTo:   10,
					ZFrom: -20,
					ZTo:   10,
					Flip:  on,
				},
			},
			want: nil,
		},
		{
			name: "returns nil if overlapbox is at both the bottom and the front edge",
			args: args{
				box: instruction{
					XFrom: -20,
					XTo:   20,
					YFrom: -20,
					YTo:   20,
					ZFrom: -20,
					ZTo:   20,
					Flip:  off,
				},
				overlapBox: instruction{
					XFrom: -10,
					XTo:   20,
					YFrom: -10,
					YTo:   10,
					ZFrom: -20,
					ZTo:   10,
					Flip:  on,
				},
			},
			want: nil,
		},
		{
			name: "returns bottom front right corner box if overlapbox is not at the bottom, front, or left edges",
			args: args{
				box: instruction{
					XFrom: -20,
					XTo:   20,
					YFrom: -20,
					YTo:   20,
					ZFrom: -20,
					ZTo:   20,
					Flip:  off,
				},
				overlapBox: instruction{
					XFrom: -10,
					XTo:   10,
					YFrom: -10,
					YTo:   10,
					ZFrom: -10,
					ZTo:   10,
					Flip:  on,
				},
			},
			want: []instruction{
				{
					XFrom: 11,
					XTo:   20,
					YFrom: -20,
					YTo:   -11,
					ZFrom: -20,
					ZTo:   -11,
					Flip:  off,
				},
			},
		},
		{
			name: "returns single bottom front right corner",
			args: args{
				box: instruction{
					XFrom: -20,
					XTo:   20,
					YFrom: -20,
					YTo:   20,
					ZFrom: -20,
					ZTo:   20,
					Flip:  off,
				},
				overlapBox: instruction{
					XFrom: -10,
					XTo:   19,
					YFrom: -19,
					YTo:   10,
					ZFrom: -19,
					ZTo:   10,
					Flip:  on,
				},
			},
			want: []instruction{
				{
					XFrom: 20,
					XTo:   20,
					YFrom: -20,
					YTo:   -20,
					ZFrom: -20,
					ZTo:   -20,
					Flip:  off,
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
					XFrom: -20,
					XTo:   20,
					YFrom: -20,
					YTo:   20,
					ZFrom: -20,
					ZTo:   20,
					Flip:  off,
				},
				overlapBox: instruction{
					XFrom: -10,
					XTo:   20,
					YFrom: -10,
					YTo:   10,
					ZFrom: -10,
					ZTo:   10,
					Flip:  on,
				},
			},
			want: nil,
		},
		{
			name: "returns nil if overlapbox is at the bottom edge",
			args: args{
				box: instruction{
					XFrom: -20,
					XTo:   20,
					YFrom: -20,
					YTo:   20,
					ZFrom: -20,
					ZTo:   20,
					Flip:  off,
				},
				overlapBox: instruction{
					XFrom: -10,
					XTo:   10,
					YFrom: -10,
					YTo:   10,
					ZFrom: -20,
					ZTo:   10,
					Flip:  on,
				},
			},
			want: nil,
		},
		{
			name: "returns nil if overlapbox is at the right edge",
			args: args{
				box: instruction{
					XFrom: -20,
					XTo:   20,
					YFrom: -20,
					YTo:   20,
					ZFrom: -20,
					ZTo:   20,
					Flip:  off,
				},
				overlapBox: instruction{
					XFrom: -10,
					XTo:   10,
					YFrom: -10,
					YTo:   20,
					ZFrom: -10,
					ZTo:   10,
					Flip:  on,
				},
			},
			want: nil,
		},
		{
			name: "returns nil if overlapbox is at both the front and the right edge",
			args: args{
				box: instruction{
					XFrom: -20,
					XTo:   20,
					YFrom: -20,
					YTo:   20,
					ZFrom: -20,
					ZTo:   20,
					Flip:  off,
				},
				overlapBox: instruction{
					XFrom: -10,
					XTo:   20,
					YFrom: -10,
					YTo:   20,
					ZFrom: -10,
					ZTo:   10,
					Flip:  on,
				},
			},
			want: nil,
		},
		{
			name: "returns nil if overlapbox is at both the bottom and the right edge",
			args: args{
				box: instruction{
					XFrom: -20,
					XTo:   20,
					YFrom: -20,
					YTo:   20,
					ZFrom: -20,
					ZTo:   20,
					Flip:  off,
				},
				overlapBox: instruction{
					XFrom: -10,
					XTo:   10,
					YFrom: -10,
					YTo:   20,
					ZFrom: -20,
					ZTo:   10,
					Flip:  on,
				},
			},
			want: nil,
		},
		{
			name: "returns nil if overlapbox is at both the bottom and the front edge",
			args: args{
				box: instruction{
					XFrom: -20,
					XTo:   20,
					YFrom: -20,
					YTo:   20,
					ZFrom: -20,
					ZTo:   20,
					Flip:  off,
				},
				overlapBox: instruction{
					XFrom: -10,
					XTo:   20,
					YFrom: -10,
					YTo:   10,
					ZFrom: -20,
					ZTo:   10,
					Flip:  on,
				},
			},
			want: nil,
		},
		{
			name: "returns bottom front right corner box if overlapbox is not at the bottom, front, or right edges",
			args: args{
				box: instruction{
					XFrom: -20,
					XTo:   20,
					YFrom: -20,
					YTo:   20,
					ZFrom: -20,
					ZTo:   20,
					Flip:  off,
				},
				overlapBox: instruction{
					XFrom: -10,
					XTo:   10,
					YFrom: -10,
					YTo:   10,
					ZFrom: -10,
					ZTo:   10,
					Flip:  on,
				},
			},
			want: []instruction{
				{
					XFrom: 11,
					XTo:   20,
					YFrom: 11,
					YTo:   20,
					ZFrom: -20,
					ZTo:   -11,
					Flip:  off,
				},
			},
		},
		{
			name: "returns single bottom front right corner",
			args: args{
				box: instruction{
					XFrom: -20,
					XTo:   20,
					YFrom: -20,
					YTo:   20,
					ZFrom: -20,
					ZTo:   20,
					Flip:  off,
				},
				overlapBox: instruction{
					XFrom: -10,
					XTo:   19,
					YFrom: -10,
					YTo:   19,
					ZFrom: -19,
					ZTo:   10,
					Flip:  on,
				},
			},
			want: []instruction{
				{
					XFrom: 20,
					XTo:   20,
					YFrom: 20,
					YTo:   20,
					ZFrom: -20,
					ZTo:   -20,
					Flip:  off,
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
					XFrom: -20,
					XTo:   -10,
					YFrom: -20,
					YTo:   -10,
					ZFrom: -20,
					ZTo:   -10,
					Flip:  on,
				},
				otherBox: instruction{
					XFrom: 10,
					XTo:   20,
					YFrom: 10,
					YTo:   20,
					ZFrom: 10,
					ZTo:   20,
					Flip:  off,
				},
			},
			want: []instruction{
				{
					XFrom: -20,
					XTo:   -10,
					YFrom: -20,
					YTo:   -10,
					ZFrom: -20,
					ZTo:   -10,
					Flip:  on,
				},
				{
					XFrom: 10,
					XTo:   20,
					YFrom: 10,
					YTo:   20,
					ZFrom: 10,
					ZTo:   20,
					Flip:  off,
				},
			},
		},
		{
			name: "returns overlapbox, and all other pieces from both of the boxes",
			args: args{
				box: instruction{
					XFrom: -20,
					XTo:   5,
					YFrom: -20,
					YTo:   5,
					ZFrom: -20,
					ZTo:   5,
					Flip:  on,
				},
				otherBox: instruction{
					XFrom: -5,
					XTo:   20,
					YFrom: -5,
					YTo:   20,
					ZFrom: -5,
					ZTo:   20,
					Flip:  off,
				},
			},
			want: []instruction{
				// box back face
				{
					XFrom: -20,
					XTo:   -6,
					YFrom: -5,
					YTo:   5,
					ZFrom: -5,
					ZTo:   5,
					Flip:  on,
				},
				// box bottom face
				{
					XFrom: -5,
					XTo:   5,
					YFrom: -5,
					YTo:   5,
					ZFrom: -20,
					ZTo:   -6,
					Flip:  on,
				},
				// box left face
				{
					XFrom: -5,
					XTo:   5,
					YFrom: -20,
					YTo:   -6,
					ZFrom: -5,
					ZTo:   5,
					Flip:  on,
				},
				// box back left edge
				{
					XFrom: -20,
					XTo:   -6,
					YFrom: -20,
					YTo:   -6,
					ZFrom: -5,
					ZTo:   5,
					Flip:  on,
				},
				// box back bottom edge
				{
					XFrom: -20,
					XTo:   -6,
					YFrom: -5,
					YTo:   5,
					ZFrom: -20,
					ZTo:   -6,
					Flip:  on,
				},
				// box left bottom edge
				{
					XFrom: -5,
					XTo:   5,
					YFrom: -20,
					YTo:   -6,
					ZFrom: -20,
					ZTo:   -6,
					Flip:  on,
				},
				// box bottom back left corner
				{
					XFrom: -20,
					XTo:   -6,
					YFrom: -20,
					YTo:   -6,
					ZFrom: -20,
					ZTo:   -6,
					Flip:  on,
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
		XFrom: 0,
		XTo:   10,
		YFrom: 0,
		YTo:   10,
		ZFrom: 0,
		ZTo:   10,
		Flip:  on,
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
			name: "merges two boxes on x, touching above",
			args: args{
				box: standardBox,
				otherBox: instruction{
					XFrom: 11,
					XTo:   20,
					YFrom: 0,
					YTo:   10,
					ZFrom: 0,
					ZTo:   10,
					Flip:  on,
				},
			},
			want: []instruction{
				{
					XFrom: 0,
					XTo:   20,
					YFrom: 0,
					YTo:   10,
					ZFrom: 0,
					ZTo:   10,
					Flip:  on,
				},
			},
		},
		{
			name: "merges two boxes on x, overlaps on one above",
			args: args{
				box: standardBox,
				otherBox: instruction{
					XFrom: 10,
					XTo:   20,
					YFrom: 0,
					YTo:   10,
					ZFrom: 0,
					ZTo:   10,
					Flip:  on,
				},
			},
			want: []instruction{
				{
					XFrom: 0,
					XTo:   20,
					YFrom: 0,
					YTo:   10,
					ZFrom: 0,
					ZTo:   10,
					Flip:  on,
				},
			},
		},
		{
			name: "merges two boxes on x, overlaps wholly above",
			args: args{
				box: standardBox,
				otherBox: instruction{
					XFrom: 0,
					XTo:   20,
					YFrom: 0,
					YTo:   10,
					ZFrom: 0,
					ZTo:   10,
					Flip:  on,
				},
			},
			want: []instruction{
				{
					XFrom: 0,
					XTo:   20,
					YFrom: 0,
					YTo:   10,
					ZFrom: 0,
					ZTo:   10,
					Flip:  on,
				},
			},
		},
		{
			name: "merges two boxes on x, touching below",
			args: args{
				box: standardBox,
				otherBox: instruction{
					XFrom: -10,
					XTo:   -1,
					YFrom: 0,
					YTo:   10,
					ZFrom: 0,
					ZTo:   10,
					Flip:  on,
				},
			},
			want: []instruction{
				{
					XFrom: -10,
					XTo:   10,
					YFrom: 0,
					YTo:   10,
					ZFrom: 0,
					ZTo:   10,
					Flip:  on,
				},
			},
		},
		{
			name: "merges two boxes on x, fully enclosed",
			args: args{
				box: standardBox,
				otherBox: instruction{
					XFrom: -6,
					XTo:   17,
					YFrom: 0,
					YTo:   10,
					ZFrom: 0,
					ZTo:   10,
					Flip:  on,
				},
			},
			want: []instruction{
				{
					XFrom: -6,
					XTo:   17,
					YFrom: 0,
					YTo:   10,
					ZFrom: 0,
					ZTo:   10,
					Flip:  on,
				},
			},
		},
		{
			name: "fails merging two boxes on x, do not touch",
			args: args{
				box: standardBox,
				otherBox: instruction{
					XFrom: 12,
					XTo:   17,
					YFrom: 0,
					YTo:   10,
					ZFrom: 0,
					ZTo:   10,
					Flip:  on,
				},
			},
			want: []instruction{
				standardBox,
				{
					XFrom: 12,
					XTo:   17,
					YFrom: 0,
					YTo:   10,
					ZFrom: 0,
					ZTo:   10,
					Flip:  on,
				},
			},
		},
		{
			name: "merges two boxes on y, touching above",
			args: args{
				box: standardBox,
				otherBox: instruction{
					XFrom: 0,
					XTo:   10,
					YFrom: 11,
					YTo:   20,
					ZFrom: 0,
					ZTo:   10,
					Flip:  on,
				},
			},
			want: []instruction{
				{
					XFrom: 0,
					XTo:   10,
					YFrom: 0,
					YTo:   20,
					ZFrom: 0,
					ZTo:   10,
					Flip:  on,
				},
			},
		},
		{
			name: "merges two boxes on y, touching below",
			args: args{
				box: standardBox,
				otherBox: instruction{
					XFrom: 0,
					XTo:   10,
					YFrom: -10,
					YTo:   -1,
					ZFrom: 0,
					ZTo:   10,
					Flip:  on,
				},
			},
			want: []instruction{
				{
					XFrom: 0,
					XTo:   10,
					YFrom: -10,
					YTo:   10,
					ZFrom: 0,
					ZTo:   10,
					Flip:  on,
				},
			},
		},
		{
			name: "merges two boxes on y, overlapping",
			args: args{
				box: standardBox,
				otherBox: instruction{
					XFrom: 0,
					XTo:   10,
					YFrom: 5,
					YTo:   15,
					ZFrom: 0,
					ZTo:   10,
					Flip:  on,
				},
			},
			want: []instruction{
				{
					XFrom: 0,
					XTo:   10,
					YFrom: 0,
					YTo:   15,
					ZFrom: 0,
					ZTo:   10,
					Flip:  on,
				},
			},
		},
		{
			name: "merges two boxes on y, overlapping below",
			args: args{
				box: standardBox,
				otherBox: instruction{
					XFrom: 0,
					XTo:   10,
					YFrom: -5,
					YTo:   5,
					ZFrom: 0,
					ZTo:   10,
					Flip:  on,
				},
			},
			want: []instruction{
				{
					XFrom: 0,
					XTo:   10,
					YFrom: -5,
					YTo:   10,
					ZFrom: 0,
					ZTo:   10,
					Flip:  on,
				},
			},
		},
		{
			name: "merges two boxes on y, fully enclosed",
			args: args{
				box: standardBox,
				otherBox: instruction{
					XFrom: 0,
					XTo:   10,
					YFrom: -6,
					YTo:   17,
					ZFrom: 0,
					ZTo:   10,
					Flip:  on,
				},
			},
			want: []instruction{
				{
					XFrom: 0,
					XTo:   10,
					YFrom: -6,
					YTo:   17,
					ZFrom: 0,
					ZTo:   10,
					Flip:  on,
				},
			},
		},
		{
			name: "fails merging two boxes on y, do not touch",
			args: args{
				box: standardBox,
				otherBox: instruction{
					XFrom: 0,
					XTo:   10,
					YFrom: 12,
					YTo:   17,
					ZFrom: 0,
					ZTo:   10,
					Flip:  on,
				},
			},
			want: []instruction{
				standardBox,
				{
					XFrom: 0,
					XTo:   10,
					YFrom: 12,
					YTo:   17,
					ZFrom: 0,
					ZTo:   10,
					Flip:  on,
				},
			},
		},
		{
			name: "merges two boxes on z, touching above",
			args: args{
				box: standardBox,
				otherBox: instruction{
					XFrom: 0,
					XTo:   10,
					YFrom: 0,
					YTo:   10,
					ZFrom: 11,
					ZTo:   20,
					Flip:  on,
				},
			},
			want: []instruction{
				{
					XFrom: 0,
					XTo:   10,
					YFrom: 0,
					YTo:   10,
					ZFrom: 0,
					ZTo:   20,
					Flip:  on,
				},
			},
		},

		{
			name: "merges two boxes on z, touching below",
			args: args{
				box: standardBox,
				otherBox: instruction{
					XFrom: 0,
					XTo:   10,
					YFrom: 0,
					YTo:   10,
					ZFrom: -10,
					ZTo:   -1,
					Flip:  on,
				},
			},
			want: []instruction{
				{
					XFrom: 0,
					XTo:   10,
					YFrom: 0,
					YTo:   10,
					ZFrom: -10,
					ZTo:   10,
					Flip:  on,
				},
			},
		},
		{
			name: "merges two boxes on z, overlapping",
			args: args{
				box: standardBox,
				otherBox: instruction{
					XFrom: 0,
					XTo:   10,
					YFrom: 0,
					YTo:   10,
					ZFrom: 5,
					ZTo:   15,
					Flip:  on,
				},
			},
			want: []instruction{
				{
					XFrom: 0,
					XTo:   10,
					YFrom: 0,
					YTo:   10,
					ZFrom: 0,
					ZTo:   15,
					Flip:  on,
				},
			},
		},
		{
			name: "merges two boxes on z, overlapping below",
			args: args{
				box: standardBox,
				otherBox: instruction{
					XFrom: 0,
					XTo:   10,
					YFrom: 0,
					YTo:   10,
					ZFrom: -5,
					ZTo:   5,
					Flip:  on,
				},
			},
			want: []instruction{
				{
					XFrom: 0,
					XTo:   10,
					YFrom: 0,
					YTo:   10,
					ZFrom: -5,
					ZTo:   10,
					Flip:  on,
				},
			},
		},
		{
			name: "merges two boxes on z, fully enclosed",
			args: args{
				box: standardBox,
				otherBox: instruction{
					XFrom: 0,
					XTo:   10,
					YFrom: 0,
					YTo:   10,
					ZFrom: -6,
					ZTo:   17,
					Flip:  on,
				},
			},
			want: []instruction{
				{
					XFrom: 0,
					XTo:   10,
					YFrom: 0,
					YTo:   10,
					ZFrom: -6,
					ZTo:   17,
					Flip:  on,
				},
			},
		},
		{
			name: "fails merging two boxes on z, do not touch",
			args: args{
				box: standardBox,
				otherBox: instruction{
					XFrom: 0,
					XTo:   10,
					YFrom: 0,
					YTo:   10,
					ZFrom: 12,
					ZTo:   17,
					Flip:  on,
				},
			},
			want: []instruction{
				standardBox,
				{
					XFrom: 0,
					XTo:   10,
					YFrom: 0,
					YTo:   10,
					ZFrom: 12,
					ZTo:   17,
					Flip:  on,
				},
			},
		},
		{
			name: "fails merging two boxes, two axis values do not match",
			args: args{
				box: standardBox,
				otherBox: instruction{
					XFrom: 0,
					XTo:   10,
					YFrom: 3,
					YTo:   10,
					ZFrom: 5,
					ZTo:   15,
					Flip:  on,
				},
			},
			want: []instruction{
				standardBox,
				{
					XFrom: 0,
					XTo:   10,
					YFrom: 3,
					YTo:   10,
					ZFrom: 5,
					ZTo:   15,
					Flip:  on,
				},
			},
		},
		{
			name: "fails merging two boxes, on off switches differ",
			args: args{
				box: standardBox,
				otherBox: instruction{
					XFrom: 0,
					XTo:   10,
					YFrom: 0,
					YTo:   10,
					ZFrom: 5,
					ZTo:   15,
					Flip:  off,
				},
			},
			want: []instruction{
				standardBox,
				{
					XFrom: 0,
					XTo:   10,
					YFrom: 0,
					YTo:   10,
					ZFrom: 5,
					ZTo:   15,
					Flip:  off,
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
					XFrom: -20,
					XTo:   5,
					YFrom: -20,
					YTo:   5,
					ZFrom: -20,
					ZTo:   5,
					Flip:  on,
				},
				otherBox: instruction{
					XFrom: -5,
					XTo:   20,
					YFrom: -5,
					YTo:   20,
					ZFrom: -5,
					ZTo:   20,
					Flip:  off,
				},
			},
			want: map[string]instruction{
				"-20/-6/-20/-6/-20/-6/on": {XFrom: -20, XTo: -6, YFrom: -20, YTo: -6, ZFrom: -20, ZTo: -6, Flip: 1},
				"-20/-6/-5/5/-20/5/on":    {XFrom: -20, XTo: -6, YFrom: -5, YTo: 5, ZFrom: -20, ZTo: 5, Flip: 1},
				"-20/5/-20/-6/-5/5/on":    {XFrom: -20, XTo: 5, YFrom: -20, YTo: -6, ZFrom: -5, ZTo: 5, Flip: 1},
				"-5/5/-20/5/-20/-6/on":    {XFrom: -5, XTo: 5, YFrom: -20, YTo: 5, ZFrom: -20, ZTo: -6, Flip: 1},
			},
		},
		{
			name: "slices out a piece from a plane",
			args: args{
				box: instruction{
					XFrom: 10,
					XTo:   12,
					YFrom: 10,
					YTo:   10,
					ZFrom: 11,
					ZTo:   12,
					Flip:  on,
				},
				otherBox: instruction{
					XFrom: 9,
					XTo:   11,
					YFrom: 9,
					YTo:   11,
					ZFrom: 9,
					ZTo:   11,
					Flip:  off,
				},
			},
			want: map[string]instruction{
				"10/12/10/10/12/12/on": {XFrom: 10, XTo: 12, YFrom: 10, YTo: 10, ZFrom: 12, ZTo: 12, Flip: on},
				"12/12/10/10/11/11/on": {XFrom: 12, XTo: 12, YFrom: 10, YTo: 10, ZFrom: 11, ZTo: 11, Flip: on},
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

func Test_overlapMap(t *testing.T) {
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
			name: "returns only the two boxes if they do not overlap",
			args: args{
				box: instruction{
					XFrom: -20,
					XTo:   -10,
					YFrom: -20,
					YTo:   -10,
					ZFrom: -20,
					ZTo:   -10,
					Flip:  on,
				},
				otherBox: instruction{
					XFrom: 10,
					XTo:   20,
					YFrom: 10,
					YTo:   20,
					ZFrom: 10,
					ZTo:   20,
					Flip:  off,
				},
			},
			want: map[string]instruction{
				"-20/-10/-20/-10/-20/-10/on": {
					XFrom: -20,
					XTo:   -10,
					YFrom: -20,
					YTo:   -10,
					ZFrom: -20,
					ZTo:   -10,
					Flip:  on,
				},
				"10/20/10/20/10/20/off": {
					XFrom: 10,
					XTo:   20,
					YFrom: 10,
					YTo:   20,
					ZFrom: 10,
					ZTo:   20,
					Flip:  off,
				},
			},
		},
		{
			name: "returns overlapbox, and all other pieces from both of the boxes",
			args: args{
				box: instruction{
					XFrom: -20,
					XTo:   5,
					YFrom: -20,
					YTo:   5,
					ZFrom: -20,
					ZTo:   5,
					Flip:  on,
				},
				otherBox: instruction{
					XFrom: -5,
					XTo:   20,
					YFrom: -5,
					YTo:   20,
					ZFrom: -5,
					ZTo:   20,
					Flip:  off,
				},
			},
			want: map[string]instruction{
				// box back face
				"-20/-6/-5/5/-5/5/on": {
					XFrom: -20,
					XTo:   -6,
					YFrom: -5,
					YTo:   5,
					ZFrom: -5,
					ZTo:   5,
					Flip:  on,
				},
				// box bottom face
				"-5/5/-5/5/-20/-6/on": {
					XFrom: -5,
					XTo:   5,
					YFrom: -5,
					YTo:   5,
					ZFrom: -20,
					ZTo:   -6,
					Flip:  on,
				},
				// box left face
				"-5/5/-20/-6/-5/5/on": {
					XFrom: -5,
					XTo:   5,
					YFrom: -20,
					YTo:   -6,
					ZFrom: -5,
					ZTo:   5,
					Flip:  on,
				},
				// box back left edge
				"-20/-6/-20/-6/-5/5/on": {
					XFrom: -20,
					XTo:   -6,
					YFrom: -20,
					YTo:   -6,
					ZFrom: -5,
					ZTo:   5,
					Flip:  on,
				},
				// box back bottom edge
				"-20/-6/-5/5/-20/-6/on": {
					XFrom: -20,
					XTo:   -6,
					YFrom: -5,
					YTo:   5,
					ZFrom: -20,
					ZTo:   -6,
					Flip:  on,
				},
				// box left bottom edge
				"-5/5/-20/-6/-20/-6/on": {
					XFrom: -5,
					XTo:   5,
					YFrom: -20,
					YTo:   -6,
					ZFrom: -20,
					ZTo:   -6,
					Flip:  on,
				},
				// box bottom back left corner
				"-20/-6/-20/-6/-20/-6/on": {
					XFrom: -20,
					XTo:   -6,
					YFrom: -20,
					YTo:   -6,
					ZFrom: -20,
					ZTo:   -6,
					Flip:  on,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, overlapMap(tt.args.box, tt.args.otherBox),
				"overlapMap(%v, %v)", tt.args.box, tt.args.otherBox)
		})
	}
}
