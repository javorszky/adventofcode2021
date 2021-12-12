package day12

func task1(input []string) int {
	nodes := parseIntoNodes(input)
	paths := walkNodes(nodes[startName], []string{}, contains)

	return len(paths)
}

func walkNodes(currentNode *node, currentPath []string, cntns func([]string, string) bool) [][]string {
	paths := make([][]string, 0)
	name := currentNode.name

	// if we're on a small node, and we've seen this one before, return an empty path without the current visited
	// elements. This means we've reached a dead end as we've gone back to a small node that we've seen before.
	if currentNode.smol && cntns(currentPath, name) {
		return paths
	}

	// Add the allowTwice of the current node to the path slice.
	currentPath = append(currentPath, name)

	// If we're at the endName node, we have reached the end of our journey.
	if name == endName {
		return append(paths, currentPath)
	}

	for _, n := range currentNode.lynx {
		v := make([]string, len(currentPath))
		copy(v, currentPath)
		paths = append(paths, walkNodes(n, v, cntns)...)
	}

	return paths
}

func contains(path []string, element string) bool {
	for _, p := range path {
		if p == element {
			return true
		}
	}

	return false
}
