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
