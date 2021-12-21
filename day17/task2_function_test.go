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

func Test_ySpeed(t *testing.T) {
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
			name: "15, 0 returns 15,0",
			args: args{
				initial: 15,
				tick:    0,
			},
			want:  15,
			want1: 0,
		},
		{
			name: "15, 1 returns 14,15",
			args: args{
				initial: 15,
				tick:    1,
			},
			want:  14,
			want1: 15,
		},
		{
			name: "15, 15 returns 0,120",
			args: args{
				initial: 15,
				tick:    15,
			},
			want:  0,
			want1: 120,
		},
		{
			name: "15, 16 returns -1,120",
			args: args{
				initial: 15,
				tick:    16,
			},
			want:  -1,
			want1: 120,
		},
		{
			name: "15, 17 returns -2,119",
			args: args{
				initial: 15,
				tick:    17,
			},
			want:  -2,
			want1: 119,
		},
		{
			name: "15, 18 returns -3,117",
			args: args{
				initial: 15,
				tick:    18,
			},
			want:  -3,
			want1: 117,
		},
		{
			name: "15, 30 returns -15,15",
			args: args{
				initial: 15,
				tick:    30,
			},
			want:  -15,
			want1: 15,
		},
		{
			name: "15, 31 returns -16,0",
			args: args{
				initial: 15,
				tick:    31,
			},
			want:  -16,
			want1: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := ySpeed(tt.args.initial, tt.args.tick)
			assert.Equalf(t, tt.want, got, "ySpeed(%v, %v)", tt.args.initial, tt.args.tick)
			assert.Equalf(t, tt.want1, got1, "ySpeed(%v, %v)", tt.args.initial, tt.args.tick)
		})
	}
}

func Test_xFunc(t *testing.T) {
	type args struct {
		initial int
		xMin    int
		xMax    int
	}

	tests := []struct {
		name string
		args args
		want map[int]int
	}{
		{
			name: "will never reach it",
			args: args{
				initial: 4,
				xMin:    12,
				xMax:    19,
			},
			want: nil,
		},
		{
			name: "will overshoot it",
			args: args{
				initial: 20,
				xMin:    12,
				xMax:    19,
			},
			want: nil,
		},
		{
			name: "will skip over it",
			args: args{
				initial: 11,
				xMin:    12,
				xMax:    19,
			},
			want: map[int]int{},
		},
		{
			name: "produces hits, but scrubs inside target area",
			args: args{
				initial: 5,
				xMin:    12,
				xMax:    19,
			},
			want: map[int]int{
				3: 5,
				4: 5,
				5: 5,
			},
		},
		{
			name: "produces hits, but scrubs outside target area",
			args: args{
				initial: 7,
				xMin:    12,
				xMax:    19,
			},
			want: map[int]int{
				2: 7,
				3: 7,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, xFunc(tt.args.initial, tt.args.xMin, tt.args.xMax),
				"xFunc(%v, %v, %v)", tt.args.initial, tt.args.xMin, tt.args.xMax)
		})
	}
}
