package day06

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_spawnsOn(t *testing.T) {
	type args struct {
		day     int
		current int
		until   int
		cycle   int
	}

	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "finds spawn days for 5 now, 14 days in the future",
			args: args{
				day:     0,
				current: 5,
				until:   14,
				cycle:   defaultCycle,
			},
			want: []int{6, 13},
		},
		{
			name: "finds spawn days when it falls on the last day",
			args: args{
				day:     0,
				current: 5,
				until:   6,
				cycle:   defaultCycle,
			},
			want: []int{6},
		},
		{
			name: "finds spawn days when it starts with a new spawn",
			args: args{
				day:     5,
				current: 8,
				until:   21,
				cycle:   defaultCycle,
			},
			want: []int{14, 21},
		},
		{
			name: "finds first example spawn days",
			args: args{
				day:     0,
				current: 3,
				until:   18,
				cycle:   defaultCycle,
			},
			want: []int{4, 11, 18},
		},
		{
			name: "finds first example normalized spawn days",
			args: args{
				day:     -5,
				current: 8,
				until:   18,
				cycle:   defaultCycle,
			},
			want: []int{4, 11, 18},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, spawnsOn(tt.args.day, tt.args.current, tt.args.until, tt.args.cycle),
				"spawnsOn(%v, %v, %v, %v, %v)",
				tt.args.day,
				tt.args.current,
				tt.args.until,
				tt.args.cycle)
		})
	}
}

func Test_getSpawnDays(t *testing.T) {
	tests := []struct {
		name string
		in   []int
		want []int
	}{
		{
			name: "normalizes ages to birth dates made up",
			in:   []int{1, 1, 2, 1, 5, 2, 3, 1, 6},
			want: []int{-7, -7, -6, -7, -3, -6, -5, -7, -2},
		},
		{
			name: "normalizes ages to birth dates example",
			in:   []int{3, 4, 3, 1, 2},
			want: []int{-5, -4, -5, -7, -6},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, getSpawnDays(tt.in), "getSpawnDays(%v)", tt.in)
		})
	}
}

func Test_calculateAllSpawns(t *testing.T) {
	tests := []struct {
		name string
		in   []int
		days int
		fn   func([]int, int) int
		want int
	}{
		{
			name: "calculates example recursively 18 days",
			in:   getSpawnDays([]int{3, 4, 3, 1, 2}),
			days: 18,
			fn:   calculateAllSpawns,
			want: 26,
		},
		{
			name: "calculates example recursively 80 days",
			in:   getSpawnDays([]int{3, 4, 3, 1, 2}),
			days: 80,
			fn:   calculateAllSpawns,
			want: 5934,
		},
		{
			name: "calculates example ticks 18",
			in:   []int{3, 4, 3, 1, 2},
			days: 18,
			fn:   calcState,
			want: 26,
		},
		{
			name: "calculates example ticks 80",
			in:   []int{3, 4, 3, 1, 2},
			days: 80,
			fn:   calcState,
			want: 5934,
		},
		//{
		//	name: "calculates example recursively",
		//	in:   []int{-5, -4, -5, -7, -6},
		//	days: 256,
		//	want: 26984457539,
		//},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, tt.fn(tt.in, tt.days), "calculateAllSpawns(%v, %v)", tt.in, tt.days)
		})
	}
}

func Benchmark_Task1(b *testing.B) {
	input := benchInput(b, "input.txt")
	for i := 0; i < b.N; i++ {
		task1(input)
	}
}
