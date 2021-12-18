package day16

import "strings"

func task2(input string) int {
	reader := strings.NewReader(input)

	built := newBuilder(reader).build()

	return built.Value()
}
