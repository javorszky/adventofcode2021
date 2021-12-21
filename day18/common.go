package day18

import (
	"errors"
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

func leaf(value int, parent *node) *node {
	depth := 0
	if parent != nil {
		depth = parent.Depth() + 1
	}

	return &node{
		parent: parent,
		value:  value,
		left:   nil,
		right:  nil,
		depth:  depth,
	}
}

func branch(left, right, parent *node) (*node, error) {
	depth := 0
	if parent != nil {
		depth = parent.Depth() + 1
	}

	if left == nil || right == nil {
		return nil, errors.New("can't create branch with nil left / right")
	}

	thisNode := &node{
		parent: parent,
		value:  0,
		left:   left,
		right:  right,
		depth:  depth,
	}

	left.parent = thisNode
	left.depth = depth + 1
	right.parent = thisNode
	right.depth = depth + 1

	return thisNode, nil
}
