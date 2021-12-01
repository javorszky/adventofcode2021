package day01

import (
	"io"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCorrectness(t *testing.T) {
	tests := []struct {
		name     string
		filename string
		fn       func(testing.TB, string) int
		want     int
	}{
		{
			name:     "tests my solution with Daniel's input task 1",
			filename: "bench_input.txt",
			fn: func(t testing.TB, s string) int {
				input := getBenchInput(t, s)

				return task1(input)
			},
			want: 1462,
		},
		{
			name:     "tests my solution with Daniel's input task 2",
			filename: "bench_input.txt",
			fn: func(t testing.TB, s string) int {
				input := getBenchInput(t, s)

				return task2(input)
			},
			want: 1497,
		},
		{
			name:     "tests okra's solution with Daniel's input task 1",
			filename: "bench_input.txt",
			fn: func(t testing.TB, s string) int {
				input := okraBenchInput1(t, s)
				defer func() { _ = input.Close() }()

				return okraTask1(input)
			},
			want: 1462,
		},
		{
			name:     "tests okra's solution with Daniel's input task 2",
			filename: "bench_input.txt",
			fn: func(t testing.TB, s string) int {
				input := okraBenchInput2(t, s)

				return okraTask2(input)
			},
			want: 1497,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, test.want, test.fn(t, test.filename))
		})
	}
}

func Benchmark(b *testing.B) {
	benchmarks := []struct {
		name     string
		filename string
		fn       func([]int) int
	}{
		{
			name:     "task 1 bench with Daniel's input",
			filename: "bench_input.txt",
			fn:       task1,
		},
		{
			name:     "task 2 bench with Daniel's input",
			filename: "bench_input.txt",
			fn:       task2,
		},
	}
	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			inputs := getBenchInput(b, bm.filename)
			for i := 0; i < b.N; i++ {
				_ = bm.fn(inputs)
			}
		})
	}
}

func BenchmarkOkra(b *testing.B) {
	benchmarks := []struct {
		name     string
		filename string
		fn       func(reader io.Reader) int
	}{
		{
			name:     "okra's task 1 bench with Daniel's input",
			filename: "bench_input.txt",
			fn:       okraTask1,
		},
	}
	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			scn := okraBenchInput1(b, bm.filename)
			defer func() { _ = scn.Close() }()
			for i := 0; i < b.N; i++ {
				_ = bm.fn(scn)
			}
		})
	}
}

func BenchmarkOkra2(b *testing.B) {
	benchmarks := []struct {
		name     string
		filename string
		fn       func([]string) int
	}{
		{
			name:     "okra's task 2 bench with Daniel's input",
			filename: "bench_input.txt",
			fn:       okraTask2,
		},
	}
	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			scn := okraBenchInput2(b, bm.filename)
			for i := 0; i < b.N; i++ {
				_ = bm.fn(scn)
			}
		})
	}
}

func getBenchInput(b testing.TB, filename string) []int {
	b.Helper()

	return getInputs(filename)
}

func okraBenchInput1(b testing.TB, filename string) *os.File {
	b.Helper()

	cwd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	file, err := os.Open(filepath.Join(cwd, filename))
	if err != nil {
		log.Fatal(err)
	}

	return file
}

func okraBenchInput2(b testing.TB, filename string) []string {
	b.Helper()

	cwd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	reader, err := ioutil.ReadFile(filepath.Join(cwd, filename))
	if err != nil {
		log.Fatal(err)
	}

	values := strings.Split(strings.TrimRight(string(reader), "\n"), "\n")

	return values
}
