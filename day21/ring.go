package day21

import "log"

type node struct {
	nextNode     *node
	previousNode *node
	nodeValue    int
}

func (n *node) next() *node {
	return n.nextNode
}

func (n *node) previous() *node {
	return n.previousNode
}

func (n *node) value() int {
	return n.nodeValue
}

func assembledTask1() *node {
	var first *node

	var previous *node

	for i := 1; i <= 10; i++ {
		n := &node{
			previousNode: previous,
			nextNode:     nil,
			nodeValue:    i,
		}

		if first == nil {
			previous = n
			first = n

			continue
		}

		previous.nextNode = n
		previous = n
	}

	first.previousNode = previous
	previous.nextNode = first

	return first
}

func (n *node) rotateTo(v int) *node {
	if v < 1 || v > 10 {
		log.Fatalf("can't rotate a ring with values [1-10] to %d", v)
	}

	rotate := n

	for rotate.value() != v {
		rotate = rotate.next()
	}

	return rotate
}

func (n *node) rotateBy(v int) (*node, int) {
	ring := n
	for i := 0; i < v%10; i++ {
		ring = ring.next()
	}

	return ring, ring.value()
}
