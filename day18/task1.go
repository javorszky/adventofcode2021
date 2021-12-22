package day18

func task1(input []string) int {
	return addNodesInOrder(input).Magnitude()
}

func addNodesInOrder(forest []string) *node {
	tree := parse(forest[0])
	reduce(tree)

	for _, newTree := range forest[1:] {
		tree = addNodes(tree, parse(newTree))

		reduce(tree)
	}

	return tree
}

func reduce(tree *node) {
	cycleHash := tree.String()
	explosionHash := tree.String()

	for {
		for {
			runExplosion(tree)

			if explosionHash == tree.String() {
				break
			}

			explosionHash = tree.String()
		}

		runSplit(tree)

		if cycleHash == tree.String() {
			break
		}

		cycleHash = tree.String()
	}
}
