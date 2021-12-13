package day13

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_makePaper(t *testing.T) {
	type args struct {
		dots []string
	}

	tests := []struct {
		name string
		args args
		want map[uint]uint
	}{
		{
			name: "parses small input into paper",
			args: args{dots: []string{
				"543,332",
				"1300,1300",
				"0,0",
				"392,66",
			}},
			want: map[uint]uint{
				0b0100001111100101001100: 1, // 543,332
				0b1010001010010100010100: 1, // 1300,1300
				0b0000000000000000000000: 1, // 0,0
				0b0011000100000001000010: 1, // 392,66
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, makePaper(tt.args.dots))
		})
	}
}
