package day04

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Benchmark_Task1s(b *testing.B) {
	benchmarks := []struct {
		name string
		fn   func([]string) (int, error)
	}{
		{
			name: "task1",
			fn:   task1,
		},
		{
			name: "task1BoardPlay",
			fn:   task1BoardPlay,
		},
		{
			name: "task1BoardPlayConcurrent",
			fn:   task1BoardPlayConcurrent,
		},
	}
	for _, bm := range benchmarks {
		fileData := benchInput(b, "input.txt")
		b.Run(bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_, _ = bm.fn(fileData)
			}
		})
	}
}

func benchInput(b testing.TB, filename string) []string {
	b.Helper()

	return getInputs(filename)
}

func Test_task1BoardPlay(t *testing.T) {
	tests := []struct {
		name     string
		filename string
		want     int
		wantErr  assert.ErrorAssertionFunc
	}{
		{
			name:     "verifies that task1BoardPlay, task1BoardPlayConcurrent and task1 come to the same result",
			filename: "input.txt",
			want:     33348,
			wantErr:  assert.NoError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fileData := benchInput(t, tt.filename)

			bpGot, bpErr := task1BoardPlay(benchInput(t, tt.filename))
			tt.wantErr(t, bpErr)

			bpGotCC, bpErrCC := task1BoardPlayConcurrent(benchInput(t, tt.filename))
			tt.wantErr(t, bpErrCC)

			t1Got, t1Err := task1(fileData)
			tt.wantErr(t, t1Err)

			assert.Equal(t, tt.want, bpGot)
			assert.Equal(t, bpGot, t1Got)
			assert.Equal(t, bpGot, bpGotCC)
		})
	}
}
