package day21

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
