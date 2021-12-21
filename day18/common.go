package day18

import (
	"io/ioutil"
	"strings"

	"github.com/javorszky/adventofcode2021/util"
)

// getInputs reads the input.txt file and returns them as a slice of strings for each row.
func getInputs(fn string) []string {
	data, err := ioutil.ReadFile(fn)
	if err != nil {
		panic(err)
	}

	return strings.Split(strings.TrimRight(string(data), util.NewLine), util.NewLine)
}

type node struct {
	parent *node
	value  int
	left   *node
	right  *node
	depth  int
}

func (n *node) Left() *node {
	return n.left
}

func (n *node) Right() *node {
	return n.right
}

func (n *node) Self() int {
	return n.value
}

func (n *node) Parent() *node {
	return n.parent
}

func (n *node) Depth() int {
	return n.depth
}
