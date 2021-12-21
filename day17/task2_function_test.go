package day17

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_xSpeed(t *testing.T) {
	type args struct {
		initial int
		tick    int
	}

	tests := []struct {
		name  string
		args  args
		want  int
		want1 int
	}{
		{
			name: "16,0 produces 16,0",
			args: args{
				initial: 16,
				tick:    0,
			},
			want:  16,
			want1: 0,
		},
		{
			name: "16,1 produces 15,16",
			args: args{
				initial: 16,
				tick:    1,
			},
			want:  15,
			want1: 16,
		},
		{
			name: "16,2 produces 14,31",
			args: args{
				initial: 16,
				tick:    2,
			},
			want:  14,
			want1: 31,
		},
		{
			name: "16,3 produces 13,45",
			args: args{
				initial: 16,
				tick:    3,
			},
			want:  13,
			want1: 45,
		},
		{
			name: "16,16 produces 0,136",
			args: args{
				initial: 16,
				tick:    16,
			},
			want:  0,
			want1: 136,
		},
		{
			name: "16,17 produces 0,136",
			args: args{
				initial: 16,
				tick:    17,
			},
			want:  0,
			want1: 136,
		},
		{
			name: "16,5443 produces 0,136",
			args: args{
				initial: 16,
				tick:    5443,
			},
			want:  0,
			want1: 136,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := xSpeed(tt.args.initial, tt.args.tick)
			assert.Equalf(t, tt.want, got, "xSpeed(%v, %v)", tt.args.initial, tt.args.tick)
			assert.Equalf(t, tt.want1, got1, "xSpeed(%v, %v)", tt.args.initial, tt.args.tick)
		})
	}
}
