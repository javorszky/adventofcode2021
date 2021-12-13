package day12

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_containsTwice(t *testing.T) {
	type args struct {
		allowTwice string
		list       []string
		el         string
	}

	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "contains twice says no for a thing we have one of",
			args: args{
				allowTwice: "ak",
				list:       []string{"start", "ak", "bo", "end"},
				el:         "ak",
			},
			want: false,
		},
		{
			name: "contains twice says true for a thing we have two of",
			args: args{
				allowTwice: "ak",
				list:       []string{"start", "ak", "bo", "ak", "end"},
				el:         "ak",
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, containsTwice(tt.args.allowTwice)(tt.args.list, tt.args.el),
				"containsTwice(%v)", tt.args.allowTwice)
		})
	}
}

func Test_task2(t *testing.T) {
	type args struct {
		fn string
	}

	tests := []struct {
		name string
		args args
		f    func([]string) int
		want int
	}{
		{
			name: "node smol example",
			args: args{fn: "example_input_small.txt"},
			f:    task2,
			want: 36,
		},
		{
			name: "node example",
			args: args{fn: "example_input.txt"},
			f:    task2,
			want: 103,
		},
		{
			name: "node slightly larger example",
			args: args{fn: "example_input_large.txt"},
			f:    task2,
			want: 3509,
		},
		{
			name: "map smol example",
			args: args{fn: "example_input_small.txt"},
			f:    task2Map,
			want: 36,
		},
		{
			name: "map example",
			args: args{fn: "example_input.txt"},
			f:    task2Map,
			want: 103,
		},
		{
			name: "map slightly larger example",
			args: args{fn: "example_input_large.txt"},
			f:    task2Map,
			want: 3509,
		},
		{
			name: "node concurrently smol example",
			args: args{fn: "example_input_small.txt"},
			f:    task2Concurrent,
			want: 36,
		},
		{
			name: "node concurrently example",
			args: args{fn: "example_input.txt"},
			f:    task2Concurrent,
			want: 103,
		},
		{
			name: "node concurrently slightly larger example",
			args: args{fn: "example_input_large.txt"},
			f:    task2Concurrent,
			want: 3509,
		},
		{
			name: "map concurrently smol example",
			args: args{fn: "example_input_small.txt"},
			f:    task2MapConcurrent,
			want: 36,
		},
		{
			name: "map concurrently example",
			args: args{fn: "example_input.txt"},
			f:    task2MapConcurrent,
			want: 103,
		},
		{
			name: "map concurrently slightly larger example",
			args: args{fn: "example_input_large.txt"},
			f:    task2MapConcurrent,
			want: 3509,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, task2(getInputs(tt.args.fn)), "task2(getInputs(%s))", tt.args.fn)
		})
	}
}

func Benchmark_Tasks(b *testing.B) {
	benchmarks := []struct {
		name string
		fn   string
		f    func([]string) int
	}{
		{
			name: "1: nodes smallex",
			fn:   "example_input_small.txt",
			f:    task1,
		},
		{
			name: "1: nodes ex",
			fn:   "example_input.txt",
			f:    task1,
		},
		{
			name: "1: nodes largex",
			fn:   "example_input_large.txt",
			f:    task1,
		},
		{
			name: "1: nodes actual",
			fn:   "input.txt",
			f:    task1,
		},
		{
			name: "1: map smallex",
			fn:   "example_input_small.txt",
			f:    task1Map,
		},
		{
			name: "1: map ex",
			fn:   "example_input.txt",
			f:    task1Map,
		},
		{
			name: "1: map largex",
			fn:   "example_input_large.txt",
			f:    task1Map,
		},
		{
			name: "1: map actual",
			fn:   "input.txt",
			f:    task1Map,
		},
		{
			name: "2: nodes smallex",
			fn:   "example_input_small.txt",
			f:    task2,
		},
		{
			name: "2: nodes ex",
			fn:   "example_input.txt",
			f:    task2,
		},
		{
			name: "2: nodes largex",
			fn:   "example_input_large.txt",
			f:    task2,
		},
		{
			name: "2: nodes actual",
			fn:   "input.txt",
			f:    task2,
		},
		{
			name: "2: nodes concurrent smallex",
			fn:   "example_input_small.txt",
			f:    task2Concurrent,
		},
		{
			name: "2: nodes concurrent ex",
			fn:   "example_input.txt",
			f:    task2Concurrent,
		},
		{
			name: "2: nodes concurrent largex",
			fn:   "example_input_large.txt",
			f:    task2Concurrent,
		},
		{
			name: "2: nodes concurrent actual",
			fn:   "input.txt",
			f:    task2Concurrent,
		},
		{
			name: "2: map smallex",
			fn:   "example_input_small.txt",
			f:    task2Map,
		},
		{
			name: "2: map ex",
			fn:   "example_input.txt",
			f:    task2Map,
		},
		{
			name: "2: map largex",
			fn:   "example_input_large.txt",
			f:    task2Map,
		},
		{
			name: "2: map actual",
			fn:   "input.txt",
			f:    task2Map,
		},
		{
			name: "2: map concurrent smallex",
			fn:   "example_input_small.txt",
			f:    task2MapConcurrent,
		},
		{
			name: "2: map concurrent ex",
			fn:   "example_input.txt",
			f:    task2MapConcurrent,
		},
		{
			name: "2: map concurrent largex",
			fn:   "example_input_large.txt",
			f:    task2MapConcurrent,
		},
		{
			name: "2: map concurrent actual",
			fn:   "input.txt",
			f:    task2MapConcurrent,
		},
	}
	for _, bm := range benchmarks {
		input := benchInput(b, bm.fn)
		b.Run(bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				bm.f(input)
			}
		})
	}
}

func benchInput(b testing.TB, fn string) []string {
	b.Helper()

	return getInputs(fn)
}
