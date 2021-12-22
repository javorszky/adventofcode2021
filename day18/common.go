package day18

import (
	"io/ioutil"
	"log"
	"strconv"
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

func isLeaf(in *node) bool {
	return in.left == nil && in.right == nil
}

func isPair(in *node) bool {
	return in.left != nil && in.right != nil && isLeaf(in.left) && isLeaf(in.right)
}

func addNodes(left, right *node) *node {
	return &node{
		left:  left,
		right: right,
	}
}

func gatherNodesFromLeft(root *node) []*node {
	nodes := make([]*node, 0)

	var inOrder func(*node)

	inOrder = func(tree *node) {
		if tree != nil {
			inOrder(tree.left)

			if isLeaf(tree) {
				nodes = append(nodes, tree)
			}

			inOrder(tree.right)
		}
	}

	inOrder(root)

	return nodes
}

//
//func preOrder(tree *node) {
//	if tree != nil {
//		fmt.Printf("value of current node: %d\n", tree.value)
//		preOrder(tree.left)
//		preOrder(tree.right)
//	}
//}
//
//func postOrder(tree *node) {
//	if tree != nil {
//		postOrder(tree.left)
//		postOrder(tree.right)
//		fmt.Printf("value of current node: %d\n", tree.value)
//	}
//}

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

func parse(in string) *node {
	if len(in) == 1 {
		value, err := strconv.Atoi(in)
		if err != nil {
			log.Fatalf("string is 1 long, should be an int, is [%s], but encountered error: %s", in, err)
		}

		return &node{
			value: value,
			left:  nil,
			right: nil,
		}
	}

	// Trim the two ends of it. If the input is correct, they should be [].
	in = in[1 : len(in)-1]

	// find the middle comma, and parse the two sides
	var left strings.Builder

	var right strings.Builder

	writeLeft := true
	numSquares := 0

	for _, ch := range in {
		switch string(ch) {
		case "[":
			numSquares++

			if writeLeft {
				left.WriteString(string(ch))
			} else {
				right.WriteString(string(ch))
			}
		case "]":
			numSquares--

			if writeLeft {
				left.WriteString(string(ch))
			} else {
				right.WriteString(string(ch))
			}
		case ",":
			if numSquares == 0 {
				writeLeft = false
			} else {
				if writeLeft {
					left.WriteString(string(ch))
				} else {
					right.WriteString(string(ch))
				}
			}
		default:
			if writeLeft {
				left.WriteString(string(ch))
			} else {
				right.WriteString(string(ch))
			}
		}
	}

	return &node{
		value: 0,
		left:  parse(left.String()),
		right: parse(right.String()),
	}
}
