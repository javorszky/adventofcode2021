package day15

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_shiftValue(t *testing.T) {
	type args struct {
		v  int
		by int
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "shifts 5 by 1",
			args: args{
				v:  5,
				by: 1,
			},
			want: 6,
		},
		{
			name: "shifts 5 by 3",
			args: args{
				v:  5,
				by: 3,
			},
			want: 8,
		},
		{
			name: "shifts 5 by 5",
			args: args{
				v:  5,
				by: 5,
			},
			want: 1,
		},
		{
			name: "shifts 5 by 7",
			args: args{
				v:  5,
				by: 7,
			},
			want: 3,
		},
		{
			name: "shifts 5 by 9",
			args: args{
				v:  5,
				by: 9,
			},
			want: 5,
		},
		{
			name: "shifts 5 by 18",
			args: args{
				v:  5,
				by: 18,
			},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, shiftValue(tt.args.v, tt.args.by), "shiftValue(%v, %v)", tt.args.v, tt.args.by)
		})
	}
}

func Test_makeMapCopy(t *testing.T) {
	type args struct {
		field        map[int]int
		register     int
		smolRegister int
		shiftX       int
		shiftY       int
	}

	tests := []struct {
		name string
		args args
		want map[int]int
	}{
		{
			name: "copies a map with the big registers",
			args: args{
				field: map[int]int{
					0:     1,
					1000:  2,
					1:     3,
					99099: 4,
					23045: 5,
				},
				register:     t2register,
				smolRegister: t1register,
				shiftX:       3,
				shiftY:       2,
			},
			want: map[int]int{
				3200:   6,
				4200:   7,
				3201:   8,
				102299: 9,
				26245:  1,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, makeMapCopy(
				tt.args.field,
				tt.args.register,
				tt.args.smolRegister,
				tt.args.shiftX,
				tt.args.shiftY,
			),
				"makeMapCopy(%v, %v, %v, %v, %v)",
				tt.args.field,
				tt.args.register,
				tt.args.smolRegister,
				tt.args.shiftX,
				tt.args.shiftY)
		})
	}
}
