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
		name      string
		args      args
		wantHits  map[int]int
		wantSpeed int
		wantDist  int
		wantErr   assert.ErrorAssertionFunc
	}{
		{
			name: "will never reach it",
			args: args{
				initial: 4,
				xMin:    12,
				xMax:    19,
			},
			wantHits:  nil,
			wantSpeed: 0,
			wantDist:  10,
			wantErr:   assert.Error,
		},
		{
			name: "will overshoot it",
			args: args{
				initial: 20,
				xMin:    12,
				xMax:    19,
			},
			wantHits:  nil,
			wantSpeed: 0,
			wantDist:  20,
			wantErr:   assert.Error,
		},
		{
			name: "will skip over it",
			args: args{
				initial: 11,
				xMin:    12,
				xMax:    19,
			},
			wantHits:  map[int]int{},
			wantSpeed: 9,
			wantDist:  21,
			wantErr:   assert.NoError,
		},
		{
			name: "produces hits, but scrubs inside target area",
			args: args{
				initial: 5,
				xMin:    12,
				xMax:    19,
			},
			wantHits: map[int]int{
				3: 5,
				4: 5,
				5: 5,
			},
			wantSpeed: 0,
			wantDist:  15,
			wantErr:   assert.NoError,
		},
		{
			name: "produces hits, but scrubs outside target area",
			args: args{
				initial: 6,
				xMin:    12,
				xMax:    20,
			},
			wantHits: map[int]int{
				3: 6,
				4: 6,
				5: 6,
			},
			wantSpeed: 0,
			wantDist:  21,
			wantErr:   assert.NoError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			hits, newSpeed, dist, err := xFunc(tt.args.initial, tt.args.xMin, tt.args.xMax)

			if !tt.wantErr(t, err, fmt.Sprintf("xFunc(%v, %v, %v)", tt.args.initial, tt.args.xMin, tt.args.xMax)) {
				return
			}
			assert.Equalf(t, tt.wantHits, hits, "xFunc(%v, %v, %v)", tt.args.initial, tt.args.xMin, tt.args.xMax)
			assert.Equalf(t, tt.wantSpeed, newSpeed, "xFunc(%v, %v, %v)", tt.args.initial, tt.args.xMin, tt.args.xMax)
			assert.Equalf(t, tt.wantDist, dist, "xFunc(%v, %v, %v)", tt.args.initial, tt.args.xMin, tt.args.xMax)
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

func Test_mergeMaps(t *testing.T) {
	type args struct {
		mergeTo   map[int]map[int]struct{}
		mergeThis map[int]int
	}

	tests := []struct {
		name string
		args args
		want map[int]map[int]struct{}
	}{
		{
			name: "merges maps where no new exist yet",
			args: args{
				mergeTo: map[int]map[int]struct{}{
					1: {
						2: {},
					},
				},
				mergeThis: map[int]int{
					3: 4,
				},
			},
			want: map[int]map[int]struct{}{
				1: {
					2: {},
				},
				3: {
					4: {},
				},
			},
		},
		{
			name: "merges maps where some exist already",
			args: args{
				mergeTo: map[int]map[int]struct{}{
					1: {
						2: {},
					},
					3: {
						4: {},
					},
				},
				mergeThis: map[int]int{
					3: 4,
					4: 5,
				},
			},
			want: map[int]map[int]struct{}{
				1: {
					2: {},
				},
				3: {
					4: {},
				},
				4: {
					5: {},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, mergeMaps(tt.args.mergeTo, tt.args.mergeThis),
				"mergeMaps(%v, %v)", tt.args.mergeTo, tt.args.mergeThis)
		})
	}
}

func Test_firingSolutionsFuncs(t *testing.T) {
	type args struct {
		xMin int
		xMax int
		yMin int
		yMax int
	}

	tests := []struct {
		name string
		args args
		want map[string]struct{}
	}{
		{
			name: "produces firing solutions for example",
			args: args{
				xMin: 20,
				xMax: 30,
				yMin: -10,
				yMax: -5,
			},
			want: map[string]struct{}{
				"6, 0":    {},
				"6, 1":    {},
				"6, 2":    {},
				"6, 3":    {},
				"6, 4":    {},
				"6, 5":    {},
				"6, 6":    {},
				"6, 7":    {},
				"6, 8":    {},
				"6, 9":    {},
				"7, -1":   {},
				"7, 0":    {},
				"7, 1":    {},
				"7, 2":    {},
				"7, 3":    {},
				"7, 4":    {},
				"7, 5":    {},
				"7, 6":    {},
				"7, 7":    {},
				"7, 8":    {},
				"7, 9":    {},
				"8, -2":   {},
				"8, -1":   {},
				"8, 0":    {},
				"8, 1":    {},
				"9, -2":   {},
				"9, -1":   {},
				"9, 0":    {},
				"10, -2":  {},
				"10, -1":  {},
				"11, -4":  {},
				"11, -3":  {},
				"11, -2":  {},
				"11, -1":  {},
				"12, -4":  {},
				"12, -3":  {},
				"12, -2":  {},
				"13, -4":  {},
				"13, -3":  {},
				"13, -2":  {},
				"14, -4":  {},
				"14, -3":  {},
				"14, -2":  {},
				"15, -4":  {},
				"15, -3":  {},
				"15, -2":  {},
				"20, -10": {},
				"20, -9":  {},
				"20, -8":  {},
				"20, -7":  {},
				"20, -6":  {},
				"20, -5":  {},
				"21, -10": {},
				"21, -9":  {},
				"21, -7":  {},
				"21, -8":  {},
				"21, -6":  {},
				"21, -5":  {},
				"22, -10": {},
				"22, -9":  {},
				"22, -8":  {},
				"22, -7":  {},
				"22, -6":  {},
				"22, -5":  {},
				"23, -10": {},
				"23, -9":  {},
				"23, -8":  {},
				"23, -7":  {},
				"23, -6":  {},
				"23, -5":  {},
				"24, -10": {},
				"24, -9":  {},
				"24, -8":  {},
				"24, -7":  {},
				"24, -6":  {},
				"24, -5":  {},
				"25, -10": {},
				"25, -9":  {},
				"25, -8":  {},
				"25, -7":  {},
				"25, -6":  {},
				"25, -5":  {},
				"26, -10": {},
				"26, -9":  {},
				"26, -8":  {},
				"26, -7":  {},
				"26, -6":  {},
				"26, -5":  {},
				"27, -10": {},
				"27, -9":  {},
				"27, -8":  {},
				"27, -7":  {},
				"27, -6":  {},
				"27, -5":  {},
				"28, -10": {},
				"28, -9":  {},
				"28, -8":  {},
				"28, -7":  {},
				"28, -6":  {},
				"28, -5":  {},
				"29, -10": {},
				"29, -9":  {},
				"29, -8":  {},
				"29, -7":  {},
				"29, -6":  {},
				"29, -5":  {},
				"30, -10": {},
				"30, -9":  {},
				"30, -8":  {},
				"30, -7":  {},
				"30, -6":  {},
				"30, -5":  {},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, firingSolutionsFuncs(tt.args.xMin, tt.args.xMax, tt.args.yMin, tt.args.yMax),
				"firingSolutionsFuncs(%v, %v, %v, %v)", tt.args.xMin, tt.args.xMax, tt.args.yMin, tt.args.yMax)
		})
	}
}
