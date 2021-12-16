package day15

import "sort"

func task1Dijkstra(input []string) int {
	field := makeMapMap(input)

	lastNode := walkDijkstra(field)

	return lastNode.Cost()
}

func walkDijkstra(field map[int]map[int]int) *pathNode {
	var n *pathNode

	lastOne := (len(field)-1)*10000 + (len(field) - 1)
	pf := newPathFinder()
	pf.AddElement(0, 0, 0)

	for {
		n = pf.Next()

		if n == nil || n.Designation() == lastOne {
			break
		}

		for _, nay := range getNeighbours(n, field) {
			pf.AddElement(nay.col, nay.row, nay.cost)
		}
	}

	return n
}

type neighbour struct {
	col, row, cost int
}

func getNeighbours(n *pathNode, field map[int]map[int]int) []neighbour {
	col := n.Designation() / 10000
	row := n.Designation() - col*10000
	neighbours := make([]neighbour, 0)

	// check left
	if leftRisk, ok := field[row][col-1]; ok {
		neighbours = append(neighbours, neighbour{
			col:  col - 1,
			row:  row,
			cost: n.Cost() + leftRisk,
		})
	}

	// check right
	if rightRisk, ok := field[row][col+1]; ok {
		neighbours = append(neighbours, neighbour{
			col:  col + 1,
			row:  row,
			cost: n.Cost() + rightRisk,
		})
	}

	// check up
	if upRisk, ok := field[row-1][col]; ok {
		neighbours = append(neighbours, neighbour{
			col:  col,
			row:  row - 1,
			cost: n.Cost() + upRisk,
		})
	}

	// check down
	if downRisk, ok := field[row+1][col]; ok {
		neighbours = append(neighbours, neighbour{
			col:  col,
			row:  row + 1,
			cost: n.Cost() + downRisk,
		})
	}

	return neighbours
}

// pathNode is a simple struct containing a designation, which is an int derived from the row, and col coordinates of
// the field, and a cost to reach that field.
type pathNode struct {
	designation int
	cost        int
}

// SetCost updates the cost, if the incoming cost is lower, otherwise silently ignores it.
func (n *pathNode) SetCost(c int) *pathNode {
	if c < n.cost {
		n.cost = c
	}

	return n
}

// Designation returns the int that we use as a key for the path nodes. It's 10,000 * col coordinate (x, horizontal)
// + row coordinate (y, vertical). An element on 45,91 would have a designation of 450,091.
func (n *pathNode) Designation() int {
	return n.designation
}

// Cost returns the current cost of the node.
func (n *pathNode) Cost() int {
	return n.cost
}

// newPathNode is a constructor. It needs col (x, horizontal coord), row (y, vertical coord), and cost. Returns a
// pointer to the new pathNode.
func newPathNode(col, row, cost int) *pathNode {
	return &pathNode{
		designation: col*10000 + row,
		cost:        cost,
	}
}

// nodeSlice is a custom slice type []*pathNode that implements sort.Interface to make sorting easy.
type nodeSlice []*pathNode

// Len returns the number of elements in the slice.
func (ns nodeSlice) Len() int {
	return len(ns)
}

// Less determines if element with index i should come before element with index j. In reality, it compares the cost of
// the pathNodes at the two indexes, and returns true if the cost of i is smaller.
func (ns nodeSlice) Less(i, j int) bool {
	return ns[i].cost < ns[j].cost
}

func (ns nodeSlice) Swap(i, j int) {
	ns[i], ns[j] = ns[j], ns[i]
}

// newPathFinder will return a pointer to a new pathFinder struct with all fields initialized and ready to use.
func newPathFinder() *pathFinder {
	return &pathFinder{
		queue:           make(nodeSlice, 0),
		elementsInQueue: make(map[int]*pathNode),
		visitedElements: make(map[int]struct{}),
	}
}

// pathFinder is the struct that has a queue where the first element has the next lowest cost, an elementsInQueue to
// hold keys for which elements are in the queue, and a visitedElements map to keep track of nodes we've visited.
type pathFinder struct {
	queue           nodeSlice
	elementsInQueue map[int]*pathNode
	visitedElements map[int]struct{}
}

// AddElement takes a pair of coordinates and a cost and returns the pointer to the pathFinder object to provide a
// fluent API for it.
//
// - If the element we're adding has been visited, we ignore it.
// - If the element is already in the queue, and
//     - the cost is smaller than what we have in the queue, we update the cost, re-sort the queue, and return pointer
//       to pathFinder
//     - the cost is the same or larger than what we already have in the queue, we ignore it, and return the pointer to
//       pathFinder
// - If the element is not in the queue already, we add it to the queue, and the elements in the queue map, and re-sort
//   the queue.
func (p *pathFinder) AddElement(col, row, cost int) *pathFinder {
	node := newPathNode(col, row, cost)

	if _, ok := p.visitedElements[node.Designation()]; ok {
		return p
	}

	if stored, ok := p.elementsInQueue[node.Designation()]; ok {
		stored.SetCost(cost)
	} else {
		p.queue = append(p.queue, node)
		p.elementsInQueue[node.Designation()] = node
	}

	sort.Sort(p.queue)

	return p
}

// Next returns the next element in the queue with the lowest cost that we haven't visited yet, and shifts it off the
// beginning of the queue.
func (p *pathFinder) Next() *pathNode {
	if p.queue.Len() == 0 {
		return nil
	}

	pn := p.queue[0]
	p.queue = p.queue[1:]

	p.visitedElements[pn.Designation()] = struct{}{}
	delete(p.elementsInQueue, pn.Designation())

	return pn
}

func (p *pathFinder) Left() int {
	return p.queue.Len()
}
