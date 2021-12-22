package day18

import (
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/javorszky/adventofcode2021/util"
)

const (
	notLeafError nodeError = "not a Leaf, can't split"
	valueTooLow  nodeError = "value too low to split"
)

type nodeError string

func (n nodeError) Error() string {
	return string(n)
}

// getInputs reads the input.txt file and returns them as a slice of strings for each row.
func getInputs(fn string) []string {
	data, err := ioutil.ReadFile(fn)
	if err != nil {
		panic(err)
	}

	return strings.Split(strings.TrimRight(string(data), util.NewLine), util.NewLine)
}

type node struct {
	value int
	left  *node
	right *node
}

func inOrder(tree *node) {
	if tree != nil {
		inOrder(tree.left)
		fmt.Printf("value of current node: %d\n", tree.value)
		inOrder(tree.right)
	}
}

func preOrder(tree *node) {
	if tree != nil {
		fmt.Printf("value of current node: %d\n", tree.value)
		preOrder(tree.left)
		preOrder(tree.right)
	}
}

func postOrder(tree *node) {
	if tree != nil {
		postOrder(tree.left)
		postOrder(tree.right)
		fmt.Printf("value of current node: %d\n", tree.value)
	}
}

func Split(n *node) error {
	if n.left != nil || n.right != nil {
		return notLeafError
	}

	if n.value < 10 {
		return valueTooLow
	}

	l := n.value / 2
	r := n.value - l

	leftNode := &node{
		value: l,
		left:  nil,
		right: nil,
	}

	rightNode := &node{
		value: r,
		left:  nil,
		right: nil,
	}

	n.value = 0
	n.left = leftNode
	n.right = rightNode

	return nil
}
