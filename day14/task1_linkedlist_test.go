package day14

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_parseTemplateLinkedList(t *testing.T) {
	type args struct {
		template string
	}

	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "parses example template into a linked list",
			args: args{
				template: "NNCB",
			},
			want: "NNCB",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, drain(parseTemplateLinkedList(tt.args.template)), "parseTemplateLinkedList(%v)")
		})
	}
}

func Test_workLinkedList(t *testing.T) {
	type args struct {
		start       func() *polymerElement
		betterRules func() map[uint]uint
	}

	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "works example start -> step 1",
			args: args{
				start: func() *polymerElement {
					return parseTemplateLinkedList("NNCB")
				},
				betterRules: func() map[uint]uint {
					return parseBetterRules([]string{
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
					})
				},
			},
			want: "NCNBCHB",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, drain(workLinkedList(tt.args.start(), tt.args.betterRules())),
				"workLinkedList(%v, %v)", tt.args.start, tt.args.betterRules)
		})
	}
}
