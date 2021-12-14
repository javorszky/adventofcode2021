package day14

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_parsePolymer(t *testing.T) {
	type args struct {
		template string
	}

	tests := []struct {
		name string
		args args
		want []uint
	}{
		{
			name: "parses example template into polymer slice",
			args: args{template: "NNCB"},
			want: []uint{
				0b01001110,
				0b01001110,
				0b01000011,
				0b01000010,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := parsePolymer(tt.args.template); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parsePolymer() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_parseBetterRules(t *testing.T) {
	type args struct {
		rules []string
	}

	tests := []struct {
		name string
		args args
		want map[uint]uint
	}{
		{
			name: "parses better rules into a map",
			args: args{rules: []string{
				"NN -> C",
				"NC -> B",
				"CB -> H",
			}},
			want: map[uint]uint{
				0b0100111001001110: 0b01000011,
				0b0100111001000011: 0b01000010,
				0b0100001101000010: 0b01001000,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := parseBetterRules(tt.args.rules); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parseBetterRules() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_work(t *testing.T) {
	type args struct {
		polymer     func() []uint
		betterRules func() map[uint]uint
	}

	tests := []struct {
		name string
		args args
		want func() []uint
	}{
		{
			name: "example work start -> step 1",
			args: args{
				polymer: func() []uint {
					return parsePolymer("NNCB")
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
			want: func() []uint {
				return parsePolymer("NCNBCHB")
			},
		},
		{
			name: "example work step 1 -> step 2",
			args: args{
				polymer: func() []uint {
					return parsePolymer("NNCB")
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
			want: func() []uint {
				return parsePolymer("NCNBCHB")
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want(), work(tt.args.polymer(), tt.args.betterRules()))
		})
	}
}

func Test_task1(t *testing.T) {
	type args struct {
		template string
		rules    []string
	}

	tests := []struct {
		name string
		args args
		fn   func(string, []string) int
		want int
	}{
		{
			name: "solves example slice",
			args: args{
				template: "NNCB",
				rules: []string{
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
				},
			},
			fn:   task1,
			want: 1588,
		},
		{
			name: "solves example linked list",
			args: args{
				template: "NNCB",
				rules: []string{
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
				},
			},
			fn:   task1LinkedList,
			want: 1588,
		},
		{
			name: "solves example counter",
			args: args{
				template: "NNCB",
				rules: []string{
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
				},
			},
			fn:   task1Counting,
			want: 1588,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, tt.fn(tt.args.template, tt.args.rules), "task1(%v, %v)", tt.args.template, tt.args.rules)
		})
	}
}

func Benchmark_Tasks(b *testing.B) {
	benchmarks := []struct {
		name     string
		fn       func(string, []string) int
		filename string
	}{
		{
			name:     "task1 slice example",
			fn:       task1,
			filename: "example_input.txt",
		},
		{
			name:     "task1 linked list example",
			fn:       task1LinkedList,
			filename: "example_input.txt",
		},
	}

	for _, bm := range benchmarks {
		template, rules := benchInput(b, bm.filename)
		b.Run(bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				bm.fn(template, rules)
			}
		})
	}
}

func benchInput(tb testing.TB, filename string) (string, []string) {
	tb.Helper()

	return getInputs(filename)
}
