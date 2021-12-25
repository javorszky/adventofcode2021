package day19

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_collapseMultipleRotations(t *testing.T) {
	pos := position{
		x: 10,
		y: 20,
		z: 30,
	}

	type args struct {
		pos       position
		rotations []int
	}

	tests := []struct {
		name    string
		args    args
		want    int
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name: "collapses 1-1",
			args: args{
				pos:       pos,
				rotations: []int{1, 1},
			},
			want:    2,
			wantErr: assert.NoError,
		},
		{
			name: "collapses 1-1",
			args: args{
				pos:       pos,
				rotations: []int{1, 1},
			},
			want:    2,
			wantErr: assert.NoError,
		},
		{
			name: "collapses 1-1",
			args: args{
				pos:       pos,
				rotations: []int{1, 1},
			},
			want:    2,
			wantErr: assert.NoError,
		},
		{
			name: "collapses 1-1",
			args: args{
				pos:       pos,
				rotations: []int{1, 1},
			},
			want:    2,
			wantErr: assert.NoError,
		},
		{
			name: "collapses 1-1",
			args: args{
				pos:       pos,
				rotations: []int{1, 1},
			},
			want:    2,
			wantErr: assert.NoError,
		},
		{
			name: "collapses 1-1",
			args: args{
				pos:       pos,
				rotations: []int{1, 1},
			},
			want:    2,
			wantErr: assert.NoError,
		},
		{
			name: "collapses 1-1",
			args: args{
				pos:       pos,
				rotations: []int{1, 1},
			},
			want:    2,
			wantErr: assert.NoError,
		},
		{
			name: "collapses 1-1",
			args: args{
				pos:       pos,
				rotations: []int{1, 1},
			},
			want:    2,
			wantErr: assert.NoError,
		},
		{
			name: "collapses 1-1",
			args: args{
				pos:       pos,
				rotations: []int{1, 1},
			},
			want:    2,
			wantErr: assert.NoError,
		},
		{
			name: "collapses 1-1",
			args: args{
				pos:       pos,
				rotations: []int{1, 1},
			},
			want:    2,
			wantErr: assert.NoError,
		},
		{
			name: "collapses 1-1",
			args: args{
				pos:       pos,
				rotations: []int{1, 1},
			},
			want:    2,
			wantErr: assert.NoError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := collapseMultipleRotations(tt.args.pos, tt.args.rotations)
			if !tt.wantErr(t, err, fmt.Sprintf("collapseMultipleRotations(%v, %v)", tt.args.pos, tt.args.rotations)) {
				return
			}
			assert.Equalf(t, tt.want, got, "collapseMultipleRotations(%v, %v)", tt.args.pos, tt.args.rotations)
		})
	}
}
