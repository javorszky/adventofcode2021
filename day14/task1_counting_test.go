package day14

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_workCounter(t *testing.T) {
	type args struct {
		counter   map[string]int
		bestRules map[string][2]string
	}

	tests := []struct {
		name string
		args args
		want map[string]int
	}{
		{
			name: "counts a counter with one step",
			args: args{
				counter: parseTemplateIntoCounter("NNBHNB"),
				bestRules: getBestMap([]string{
					"CH -> B",
					"HH -> N",
					"CB -> H",
					"NH -> C",
					"HB -> C",
					"HC -> B",
					"HN -> C",
					"NN -> C",
					"BH -> H",
					"NC -> B",
					"NB -> B",
					"BN -> B",
					"BB -> N",
					"BC -> B",
					"CC -> N",
					"CN -> C",
				}),
			},
			want: parseTemplateIntoCounter("NCNBBHHCNBB"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, workCounter(tt.args.counter, tt.args.bestRules),
				"workCounter(%v, %v)", tt.args.counter, tt.args.bestRules)
		})
	}
}

func Test_parseTemplateIntoCounter(t *testing.T) {
	type args struct {
		template string
	}

	tests := []struct {
		name string
		args args
		want map[string]int
	}{
		{
			name: "creates a counter from template",
			args: args{template: "NVBBNV"},
			want: map[string]int{
				"NV": 2,
				"VB": 1,
				"BB": 1,
				"BN": 1,
			},
		},
		{
			name: "creates counter from NCNBBHHCNBB",
			args: args{
				template: "NCNBBHHCNBB",
			},
			want: map[string]int{
				"NC": 1,
				"CN": 2,
				"NB": 2,
				"BB": 2,
				"BH": 1,
				"HH": 1,
				"HC": 1,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, parseTemplateIntoCounter(tt.args.template),
				"parseTemplateIntoCounter(%v)", tt.args.template)
		})
	}
}

func Test_getPair(t *testing.T) {
	type args struct {
		in uint
	}

	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "works decoding",
			args: args{
				in: 0b0101011001001110,
			},
			want: "VN",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, getPair(tt.args.in), "getPair(%v)", tt.args.in)
		})
	}
}
