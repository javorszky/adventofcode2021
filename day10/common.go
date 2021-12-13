package day10

import (
	"io/ioutil"
	"strings"

	"github.com/javorszky/adventofcode2021/util"
)

const (
	braces  = `()`
	squares = `[]`
	curlies = `{}`
	angles  = `<>`
	empty   = ``
)

// ): 3 points.
// ]: 57 points.
// }: 1197 points.
// >: 25137 points.
var invalidScore = map[uint8]int{
	0x29: 3,     // )
	0x5d: 57,    // ]
	0x7d: 1197,  // }
	0x3e: 25137, // >
}

// getInputs reads the input.txt file and returns them as a slice of strings for each row.
func getInputs(fn string) []string {
	data, err := ioutil.ReadFile(fn)
	if err != nil {
		panic(err)
	}

	return strings.Split(strings.TrimRight(string(data), util.NewLine), util.NewLine)
}

func getChunkReplacer() *strings.Replacer {
	return strings.NewReplacer(braces, empty,
		squares, empty,
		curlies, empty,
		angles, empty,
	)
}

func getOpenerReplacer() *strings.Replacer {
	return strings.NewReplacer(`(`, ``,
		`[`, ``,
		`{`, ``,
		`<`, ``,
	)
}
