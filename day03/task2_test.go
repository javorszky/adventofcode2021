package day03

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_filters(t *testing.T) {
	l := []uint{
		0b00100,
		0b11110,
		0b10110,
		0b10111,
		0b10101,
		0b01111,
		0b00111,
		0b11100,
		0b10000,
		0b11001,
		0b00010,
		0b01010,
	}

	type args struct {
		list     []uint
		position uint
		fn       func(int, int) bool
	}

	tests := []struct {
		name string
		args args
		want []uint
	}{
		{
			name: "filters for most common values in least significant bit (ends in 0)",
			args: args{
				list:     l,
				position: 0b1,
				fn:       reduceForOxygen,
			},
			want: []uint{
				0b00100,
				0b11110,
				0b10110,
				0b11100,
				0b10000,
				0b00010,
				0b01010,
			},
		},
		{
			name: "filters for most common values in third bit (1)",
			args: args{
				list:     l,
				position: 0b100,
				fn:       reduceForOxygen,
			},
			want: []uint{
				0b00100,
				0b11110,
				0b10110,
				0b10111,
				0b10101,
				0b01111,
				0b00111,
				0b11100,
			},
		},
		{
			name: "filters for least common values in least significant bit (ends in 1)",
			args: args{
				list:     l,
				position: 0b1,
				fn:       reduceForCO2,
			},
			want: []uint{
				0b10111,
				0b10101,
				0b01111,
				0b00111,
				0b11001,
			},
		},
		{
			name: "filters for lest common values in third bit (0)",
			args: args{
				list:     l,
				position: 0b100,
				fn:       reduceForCO2,
			},
			want: []uint{
				0b10000,
				0b11001,
				0b00010,
				0b01010,
			},
		},
		{
			name: "filters for most common values, but they are equal, returns list with ones",
			args: args{
				list: []uint{
					0b0001,
					0b1000,
					0b0101,
					0b1111,
					0b0010,
					0b0110,
				},
				position: 0b1,
				fn:       reduceForOxygen,
			},
			want: []uint{
				0b0001,
				0b0101,
				0b1111,
			},
		},
		{
			name: "filters for least common values, but they are equal, returns list with zeroes",
			args: args{
				list: []uint{
					0b0001,
					0b1000,
					0b0101,
					0b1111,
					0b0010,
					0b0110,
				},
				position: 0b1,
				fn:       reduceForCO2,
			},
			want: []uint{
				0b1000,
				0b0010,
				0b0110,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(
				t, tt.want,
				filterList(tt.args.list, tt.args.position, tt.args.fn),
				"filterMostCommon(%v, %v)", tt.args.list, tt.args.position)
		})
	}
}

func Test_reduceList(t *testing.T) {
	l := []uint{
		0b00100,
		0b11110,
		0b10110,
		0b10111,
		0b10101,
		0b01111,
		0b00111,
		0b11100,
		0b10000,
		0b11001,
		0b00010,
		0b01010,
	}

	type args struct {
		input   []uint
		width   int
		reducer func(int, int) bool
	}

	tests := []struct {
		name    string
		args    args
		want    uint
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name: "grabs value for oxygen from original list",
			args: args{
				input:   l,
				width:   5,
				reducer: reduceForOxygen,
			},
			want:    23,
			wantErr: assert.NoError,
		},
		{
			name: "grabs value for co2 from original list",
			args: args{
				input:   l,
				width:   5,
				reducer: reduceForCO2,
			},
			want:    10,
			wantErr: assert.NoError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := reduceList(tt.args.input, tt.args.width, tt.args.reducer)
			if !tt.wantErr(t, err, fmt.Sprintf("reduceList(%v, %v, reducer)", tt.args.input, tt.args.width)) {
				return
			}
			assert.Equalf(t, tt.want, got, "reduceList(%v, %v, reducer)", tt.args.input, tt.args.width)
		})
	}
}
