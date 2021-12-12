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
			name: "task 2 smol example",
			args: args{fn: "example_input_small.txt"},
			f:    task2,
			want: 36,
		},
		{
			name: "task 2 example",
			args: args{fn: "example_input.txt"},
			f:    task2,
			want: 103,
		},
		{
			name: "task 2 slightly larger example",
			args: args{fn: "example_input_large.txt"},
			f:    task2,
			want: 3509,
		},
		{
			name: "task 2 concurrently smol example",
			args: args{fn: "example_input_small.txt"},
			f:    task2Concurrent,
			want: 36,
		},
		{
			name: "task 2 concurrently example",
			args: args{fn: "example_input.txt"},
			f:    task2Concurrent,
			want: 103,
		},
		{
			name: "task 2 concurrently slightly larger example",
			args: args{fn: "example_input_large.txt"},
			f:    task2Concurrent,
			want: 3509,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, task2(getInputs(tt.args.fn)), "task2(getInputs(%s))", tt.args.fn)
		})
	}
}

func Benchmark_Task2(b *testing.B) {
	benchmarks := []struct {
		name string
		fn   string
		f    func([]string) int
	}{
		{
			name: "linear smallex",
			fn:   "example_input_small.txt",
			f:    task2,
		},
		{
			name: "linear ex",
			fn:   "example_input.txt",
			f:    task2,
		},
		{
			name: "linear largex",
			fn:   "example_input_large.txt",
			f:    task2,
		},
		{
			name: "linear actual",
			fn:   "input.txt",
			f:    task2,
		},
		{
			name: "concurrent smallex",
			fn:   "example_input_small.txt",
			f:    task2Concurrent,
		},
		{
			name: "concurrent ex",
			fn:   "example_input.txt",
			f:    task2Concurrent,
		},
		{
			name: "concurrent largex",
			fn:   "example_input_large.txt",
			f:    task2Concurrent,
		},
		{
			name: "concurrent actual",
			fn:   "input.txt",
			f:    task2Concurrent,
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
