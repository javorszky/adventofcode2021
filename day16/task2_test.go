package day16

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_task2(t *testing.T) {
	type args struct {
		input string
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "example 1: C200B40A82",
			args: args{input: hexToBinString("C200B40A82")},
			want: 3,
		},
		{
			name: "example 2: 04005AC33890",
			args: args{input: hexToBinString("04005AC33890")},
			want: 54,
		},
		{
			name: "example 3: 880086C3E88112",
			args: args{input: hexToBinString("880086C3E88112")},
			want: 7,
		},
		{
			name: "example 4: CE00C43D881120",
			args: args{input: hexToBinString("CE00C43D881120")},
			want: 9,
		},
		{
			name: "example 5: D8005AC2A8F0",
			args: args{input: hexToBinString("D8005AC2A8F0")},
			want: 1,
		},
		{
			name: "example 6: F600BC2D8F",
			args: args{input: hexToBinString("F600BC2D8F")},
			want: 0,
		},
		{
			name: "example 7: 9C005AC2F8F0",
			args: args{input: hexToBinString("9C005AC2F8F0")},
			want: 0,
		},
		{
			name: "example 8: 9C0141080250320F1802104A08",
			args: args{input: hexToBinString("9C0141080250320F1802104A08")},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, task2(tt.args.input), "task2(%v)", tt.args.input)
		})
	}
}
