package day06

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_parseFishSplitAtoi(t *testing.T) {
	s := "3,4,3,1,2"
	w := []int{3, 4, 3, 1, 2}

	tests := []struct {
		name string
		in   string
		want []int
		fn   func(string) []int
	}{
		{
			name: "split atoi",
			in:   s,
			want: w,
			fn:   parseFishSplitAtoi,
		},
		{
			name: "walk atoi",
			in:   s,
			want: w,
			fn:   parseFishWalkAtoi,
		},
		{
			name: "for atoi",
			in:   s,
			want: w,
			fn:   parseFishForAtoi,
		},
		{
			name: "split map",
			in:   s,
			want: w,
			fn:   parseFishSplitMap,
		},
		{
			name: "walk map",
			in:   s,
			want: w,
			fn:   parseFishWalkMap,
		},
		{
			name: "for map",
			in:   s,
			want: w,
			fn:   parseFishForMap,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.fn(tt.in))
		})
	}
}

func Benchmark(b *testing.B) {
	input := benchInput(b, "input.txt")

	benchmarks := []struct {
		name string
		fn   func(string) []int
	}{
		{
			name: "split atoi",
			fn:   parseFishSplitAtoi,
		},
		{
			name: "walk atoi",
			fn:   parseFishWalkAtoi,
		},
		{
			name: "for atoi",
			fn:   parseFishForAtoi,
		},
		{
			name: "split map",
			fn:   parseFishSplitMap,
		},
		{
			name: "walk map",
			fn:   parseFishWalkMap,
		},
		{
			name: "for map",
			fn:   parseFishForMap,
		},
	}
	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				bm.fn(input)
			}
		})
	}
}

func benchInput(b testing.TB, filename string) string {
	b.Helper()

	return getInputs(filename)
}
