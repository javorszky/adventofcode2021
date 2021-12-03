package day03

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_task1(t *testing.T) {
	tests := []struct {
		name        string
		input       []uint
		stringInput []string
		want        uint
	}{
		{
			name: "correctly generates product",
			stringInput: []string{
				"00100",
				"11110",
				"10110",
				"10111",
				"10101",
				"01111",
				"00111",
				"11100",
				"10000",
				"11001",
				"00010",
				"01010",
			},
			want: 0b10110 * 0b01001,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, task1(getInputs(tt.stringInput)))
			assert.Equal(t, tt.want, task1Strings(tt.stringInput))
		})
	}
}

func Benchmark_Task1_Uint(b *testing.B) {
	benchmarks := []struct {
		name     string
		filename string
	}{
		{
			name:     "uint with small filename",
			filename: "smallinput.txt",
		},
		{
			name:     "uint with regular input",
			filename: "input.txt",
		},
	}
	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			in, w := getBenchInputs(b, bm.filename)
			for i := 0; i < b.N; i++ {
				_ = task1(in, w)
			}
		})
	}
}

func getBenchInputs(b testing.TB, filename string) ([]uint, int) {
	b.Helper()

	return getInputs(getLines(filename))
}

func Benchmark_Task1_Strings(b *testing.B) {
	benchmarks := []struct {
		name     string
		filename string
	}{
		{
			name:     "string with small filename",
			filename: "smallinput.txt",
		},
		{
			name:     "string with regular input",
			filename: "input.txt",
		},
	}
	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			lines := getBenchInputsStrings(b, bm.filename)
			for i := 0; i < b.N; i++ {
				_ = task1Strings(lines)
			}
		})
	}
}

func getBenchInputsStrings(b testing.TB, filename string) []string {
	b.Helper()

	return getLines(filename)
}
