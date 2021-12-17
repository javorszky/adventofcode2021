package day16

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_hexToBinString(t *testing.T) {
	type args struct {
		hexString string
	}

	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "replaces character 0 with 0000",
			args: args{
				hexString: "0",
			},
			want: "0000",
		},
		{
			name: "replaces character 1 with 0001",
			args: args{
				hexString: "1",
			},
			want: "0001",
		},
		{
			name: "replaces character 2 with 0010",
			args: args{
				hexString: "2",
			},
			want: "0010",
		},
		{
			name: "replaces character 3 with 0011",
			args: args{
				hexString: "3",
			},
			want: "0011",
		},
		{
			name: "replaces character 4 with 0100",
			args: args{
				hexString: "4",
			},
			want: "0100",
		},
		{
			name: "replaces character 5 with 0101",
			args: args{
				hexString: "5",
			},
			want: "0101",
		},
		{
			name: "replaces character 6 with 0110",
			args: args{
				hexString: "6",
			},
			want: "0110",
		},
		{
			name: "replaces character 7 with 0111",
			args: args{
				hexString: "7",
			},
			want: "0111",
		},
		{
			name: "replaces character 8 with 1000",
			args: args{
				hexString: "8",
			},
			want: "1000",
		},
		{
			name: "replaces character 9 with 1001",
			args: args{
				hexString: "9",
			},
			want: "1001",
		},
		{
			name: "replaces character A with 1010",
			args: args{
				hexString: "A",
			},
			want: "1010",
		},
		{
			name: "replaces character B with 1011",
			args: args{
				hexString: "B",
			},
			want: "1011",
		},
		{
			name: "replaces character C with 1100",
			args: args{
				hexString: "C",
			},
			want: "1100",
		},
		{
			name: "replaces character D with 1101",
			args: args{
				hexString: "D",
			},
			want: "1101",
		},
		{
			name: "replaces character E with 1110",
			args: args{
				hexString: "E",
			},
			want: "1110",
		},
		{
			name: "replaces character F with 1111",
			args: args{
				hexString: "F",
			},
			want: "1111",
		},
		{
			name: "replaces string 0123456789ABCDEF with its bin string",
			args: args{
				hexString: "0123456789ABCDEF",
			},
			want: "0000000100100011010001010110011110001001101010111100110111101111",
		},
		{
			name: "replaces multiple copies of string 0123456789ABCDEF with its bin string",
			args: args{
				hexString: "0123456789ABCDEF" +
					"0123456789ABCDEF" +
					"0123456789ABCDEF" +
					"0123456789ABCDEF",
			},
			want: "0000000100100011010001010110011110001001101010111100110111101111" +
				"0000000100100011010001010110011110001001101010111100110111101111" +
				"0000000100100011010001010110011110001001101010111100110111101111" +
				"0000000100100011010001010110011110001001101010111100110111101111",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, hexToBinString(tt.args.hexString))
		})
	}
}
