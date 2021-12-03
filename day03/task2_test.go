package day03

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_reduceList(t *testing.T) {
	l := []uint{
		0b00100,
		0b11110,
		0b10110,
		0b10111,
		0b10101,
		0b01111,
		0b00111,
		0b11100,
		0b10000,
		0b11001,
		0b00010,
		0b01010,
	}

	type args struct {
		input   []uint
		width   int
		reducer func(int, int) bool
	}

	tests := []struct {
		name    string
		args    args
		want    uint
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name: "grabs value for oxygen from original list",
			args: args{
				input:   l,
				width:   5,
				reducer: reduceForOxygen,
			},
			want:    23,
			wantErr: assert.NoError,
		},
		{
			name: "grabs value for co2 from original list",
			args: args{
				input:   l,
				width:   5,
				reducer: reduceForCO2,
			},
			want:    10,
			wantErr: assert.NoError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := reduceList(tt.args.input, tt.args.width, tt.args.reducer)
			if !tt.wantErr(t, err, fmt.Sprintf("reduceList(%v, %v, reducer)", tt.args.input, tt.args.width)) {
				return
			}
			assert.Equalf(t, tt.want, got, "reduceList(%v, %v, reducer)", tt.args.input, tt.args.width)
		})
	}
}

func Benchmark_Task2_ReduceList(b *testing.B) {
	benchmarks := []struct {
		name    string
		input   func(testing.TB) ([]uint, int)
		reducer func(int, int) bool
	}{
		{
			name: "example list for co2",
			input: func(b testing.TB) ([]uint, int) {
				b.Helper()

				return []uint{
					0b00100,
					0b11110,
					0b10110,
					0b10111,
					0b10101,
					0b01111,
					0b00111,
					0b11100,
					0b10000,
					0b11001,
					0b00010,
					0b01010,
				}, 5
			},
			reducer: reduceForCO2,
		},
		{
			name: "example list for o2",
			input: func(b testing.TB) ([]uint, int) {
				b.Helper()

				return []uint{
					0b00100,
					0b11110,
					0b10110,
					0b10111,
					0b10101,
					0b01111,
					0b00111,
					0b11100,
					0b10000,
					0b11001,
					0b00010,
					0b01010,
				}, 5
			},
			reducer: reduceForOxygen,
		},
		{
			name: "actual input for co2",
			input: func(b testing.TB) ([]uint, int) {
				b.Helper()

				return getInputs(getLines("input.txt"))
			},
			reducer: reduceForCO2,
		},
		{
			name: "actual input for o2",
			input: func(b testing.TB) ([]uint, int) {
				b.Helper()

				return getInputs(getLines("input.txt"))
			},
			reducer: reduceForOxygen,
		},
	}
	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			list, w := bm.input(b)
			for i := 0; i < b.N; i++ {
				_, _ = reduceList(list, w, bm.reducer)
			}
		})
	}
}

func Benchmark_Task2_e2e(b *testing.B) {
	benchmarks := []struct {
		name  string
		input func(tb testing.TB) ([]uint, int)
	}{
		{
			name: "total task 2, small input",
			input: func(tb testing.TB) ([]uint, int) {
				tb.Helper()

				return getInputs(getLines("smallinput.txt"))
			},
		},
		{
			name: "total task 2, full input",
			input: func(tb testing.TB) ([]uint, int) {
				tb.Helper()

				return getInputs(getLines("input.txt"))
			},
		},
	}
	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			in, w := bm.input(b)
			for i := 0; i < b.N; i++ {
				_ = task2(in, w)
			}
		})
	}
}
