package day14

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_getBestMap(t *testing.T) {
	type args struct {
		rules []string
	}

	tests := []struct {
		name string
		args args
		want map[uint][2]uint
	}{
		{
			name: "parses a ruleset into the best map",
			args: args{
				rules: []string{
					"NN -> B",
					"NB -> V",
					"BV -> N",
				},
			},
			want: map[uint][2]uint{
				0b0100111001001110: {0b0100111001000010, 0b0100001001001110},
				0b0100111001000010: {0b0100111001010110, 0b0101011001000010},
				0b0100001001010110: {0b0100001001001110, 0b0100111001010110},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, getBestMap(tt.args.rules), "getBestMap(%v)", tt.args.rules)
		})
	}
}
