package day18

import (
	"fmt"
	"io/ioutil"
	"log"
	"regexp"
	"strconv"
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
	value int
	left  *node
	right *node
}

func (n *node) String() string {
	if (n.left == nil && n.right != nil) || (n.left != nil && n.right == nil) {
		log.Fatalf("String: found a node where one of the branches has a thing, the other does not. " +
			"It should not have happened")
	}

	if n.left == nil && n.right == nil {
		return fmt.Sprintf("%d", n.value)
	}

	return fmt.Sprintf("[%s,%s]", n.left.String(), n.right.String())
}

func (n *node) Magnitude() int {
	if (n.left == nil && n.right != nil) || (n.left != nil && n.right == nil) {
		log.Fatalf("Magnitude: found a node where one of the branches has a thing, the other does not. " +
			"It should not have happened")
	}

	if n.left == nil && n.right == nil {
		return n.value
	}

	return 3*n.left.Magnitude() + 2*n.right.Magnitude()
}

func isLeaf(in *node) bool {
	return in.left == nil && in.right == nil
}

func isPair(in *node) bool {
	return in.left != nil && in.right != nil && isLeaf(in.left) && isLeaf(in.right)
}

func addNodes(left, right *node) *node {
	addedNode := &node{
		left:  left,
		right: right,
	}

	reduce(addedNode)

	return addedNode
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

func gatherNodesAtTiers(root *node) map[int][]*node {
	nodes := make(map[int][]*node)

	var inOrder func(*node, int)

	inOrder = func(tree *node, depth int) {
		if tree != nil {
			if nodes[depth] == nil {
				nodes[depth] = make([]*node, 0)
			}

			inOrder(tree.left, depth+1)

			nodes[depth] = append(nodes[depth], tree)

			inOrder(tree.right, depth+1)
		}
	}

	inOrder(root, 0)

	return nodes
}

// runExplosion will run one set of explosion on the tree.
//
// "If any pair is nested inside four pairs, the leftmost such pair explodes." - day 18 task 1.
func runExplosion(root *node) {
	tiers := gatherNodesAtTiers(root)
	tierFour, ok := tiers[4]

	if !ok {
		// nothing to do
		return
	}

	var explodeThis *node

	for _, n := range tierFour {
		if isPair(n) {
			explodeThis = n

			break
		}
	}

	if explodeThis == nil {
		// there were no pairs on level 4, nothing to do
		return
	}

	leftToRight := gatherNodesFromLeft(root)
	left := explodeThis.left
	right := explodeThis.right

	var (
		leafToLeft  *node
		leafToRight *node
	)

	for i, leaf := range leftToRight {
		if leaf == left {
			if i > 0 {
				leafToLeft = leftToRight[i-1]
			}
		}

		if leaf == right {
			if i < len(leftToRight)-1 {
				leafToRight = leftToRight[i+1]
			}
		}
	}

	if leafToRight != nil {
		leafToRight.value += right.value
	}

	if leafToLeft != nil {
		leafToLeft.value += left.value
	}

	explodeThis.value = 0
	explodeThis.left = nil
	explodeThis.right = nil
}

func runSplit(root *node) {
	var splitThis *node

	for _, leaf := range gatherNodesFromLeft(root) {
		if leaf.value >= 10 {
			splitThis = leaf

			break
		}
	}

	if splitThis == nil {
		return
	}

	left := splitThis.value / 2
	right := splitThis.value - left

	leftNode := &node{value: left}
	rightNode := &node{value: right}

	splitThis.value = 0
	splitThis.left = leftNode
	splitThis.right = rightNode
}

var reOnlyDigits = regexp.MustCompile(`^\d+$`)

func parse(in string) *node {
	if reOnlyDigits.MatchString(in) {
		value, err := strconv.Atoi(in)
		if err != nil {
			log.Fatalf("string is only a number, should be an int, is [%s], but encountered error: %s", in, err)
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
