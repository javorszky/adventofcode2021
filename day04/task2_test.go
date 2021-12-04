package day04

import "testing"

func Benchmark_Task2(b *testing.B) {
	fileData := benchInput(b, "input.txt")
	for i := 0; i < b.N; i++ {
		_, _ = task2(fileData)
	}
}
