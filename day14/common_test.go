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
		want map[string][2]string
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
			want: map[string][2]string{
				"NN": {"NB", "BN"},
				"NB": {"NV", "VB"},
				"BV": {"BN", "NV"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, getBestMap(tt.args.rules), "getBestMap(%v)", tt.args.rules)
		})
	}
}
