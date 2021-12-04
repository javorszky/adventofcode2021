package day04

import (
	"testing"
)

func Benchmark_Task1(b *testing.B) {
	fileData := benchInput(b, "input.txt")
	for i := 0; i < b.N; i++ {
		_, _ = task1(fileData)
	}
}

func benchInput(b testing.TB, filename string) []string {
	b.Helper()

	return getInputs(filename)
}
