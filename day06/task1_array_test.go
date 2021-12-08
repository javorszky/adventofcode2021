package day06

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_parseIntoSlice(t *testing.T) {
	tests := []struct {
		name string
		in   []int
		want [7]int
	}{
		{
			name: "converts example slice into array",
			in:   []int{0, 1, 0, 5, 6, 0, 1, 2, 2, 3},
			want: [7]int{3, 2, 2, 1, 0, 1, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, parseIntoSlice(tt.in), "parseIntoSlice(%v)", tt.in)
		})
	}
}

func Test_tickArray(t *testing.T) {
	type args struct {
		in      []int
		forDays int
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "calculates example for 18 days",
			args: args{
				in:      []int{3, 4, 3, 1, 2},
				forDays: 18,
			},
			want: 26,
		},
		{
			name: "calculates example for 80 days",
			args: args{
				in:      []int{3, 4, 3, 1, 2},
				forDays: 80,
			},
			want: 5934,
		},
		{
			name: "calculates example for 256 days",
			args: args{
				in:      []int{3, 4, 3, 1, 2},
				forDays: 256,
			},
			want: 26984457539,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, tickArray(parseIntoSlice(tt.args.in), tt.args.forDays),
				"tickArray(%v, %v)", tt.args.in, tt.args.forDays)
		})
	}
}
