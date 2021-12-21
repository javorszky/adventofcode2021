package day17

import (
	"fmt"
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
		name    string
		args    args
		want    map[int]int
		want1   int
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name: "will never reach it",
			args: args{
				initial: 4,
				xMin:    12,
				xMax:    19,
			},
			want:    nil,
			want1:   0,
			wantErr: assert.Error,
		},
		{
			name: "will overshoot it",
			args: args{
				initial: 20,
				xMin:    12,
				xMax:    19,
			},
			want:    nil,
			want1:   0,
			wantErr: assert.Error,
		},
		{
			name: "will skip over it",
			args: args{
				initial: 11,
				xMin:    12,
				xMax:    19,
			},
			want:    map[int]int{},
			want1:   9,
			wantErr: assert.NoError,
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
			want1:   0,
			wantErr: assert.NoError,
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
			want1:   3,
			wantErr: assert.NoError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			hits, newSpeed, err := xFunc(tt.args.initial, tt.args.xMin, tt.args.xMax)

			if !tt.wantErr(t, err, fmt.Sprintf("xFunc(%v, %v, %v)", tt.args.initial, tt.args.xMin, tt.args.xMax)) {
				return
			}
			assert.Equalf(t, tt.want, hits, "xFunc(%v, %v, %v)", tt.args.initial, tt.args.xMin, tt.args.xMax)
			assert.Equalf(t, tt.want1, newSpeed, "xFunc(%v, %v, %v)", tt.args.initial, tt.args.xMin, tt.args.xMax)
		})
	}
}

func Test_yFunc(t *testing.T) {
	type args struct {
		initial int
		yMin    int
		yMax    int
	}

	tests := []struct {
		name      string
		args      args
		want      map[int]int
		want1     int
		wantError assert.ErrorAssertionFunc
	}{
		{
			name: "overshoots it",
			args: args{
				initial: -130,
				yMin:    -128,
				yMax:    -83,
			},
			want:      nil,
			want1:     -130,
			wantError: assert.Error,
		},
		{
			name: "skips over it",
			args: args{
				initial: -82,
				yMin:    -128,
				yMax:    -83,
			},
			want:      map[int]int{},
			want1:     -84,
			wantError: assert.NoError,
		},
		{
			name: "gets hits",
			args: args{
				initial: -4,
				yMin:    -128,
				yMax:    -83,
			},
			want: map[int]int{
				10: -4,
				11: -4,
				12: -4,
			},
			want1:     -17,
			wantError: assert.NoError,
		},
		{
			name: "gets hits but positive launch",
			args: args{
				initial: 6,
				yMin:    -128,
				yMax:    -83,
			},
			want: map[int]int{
				21: 6,
				22: 6,
				23: 6,
			},
			want1:     -18,
			wantError: assert.NoError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, gotErr := yFunc(tt.args.initial, tt.args.yMin, tt.args.yMax)
			if !tt.wantError(t, gotErr, fmt.Sprintf("yFunc(%v, %v, %v)", tt.args.initial, tt.args.yMin, tt.args.yMax)) {
				return
			}
			assert.Equalf(t, tt.want, got, "yFunc(%v, %v, %v)", tt.args.initial, tt.args.yMin, tt.args.yMax)
			assert.Equalf(t, tt.want1, got1, "yFunc(%v, %v, %v)", tt.args.initial, tt.args.yMin, tt.args.yMax)
		})
	}
}
