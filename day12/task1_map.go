package day12

import "strings"

func task1Map(input []string) int {
	nodes := parseIntoNodes(input)
	paths := walkNodes(nodes[startName], []string{}, contains)

	return len(paths)
}

func parseIntoNodeMap(input []string) map[string][]string {
	m := make(map[string][]string)

	for _, line := range input {
		parts := strings.Split(line, "-")
		if _, ok := m[parts[0]]; !ok {
			m[parts[0]] = make([]string, 0)
		}

		if _, ok := m[parts[1]]; !ok {
			m[parts[1]] = make([]string, 0)
		}

		m[parts[0]] = append(m[parts[0]], parts[1])
		m[parts[1]] = append(m[parts[1]], parts[0])
	}

	return m
}

func walkNodeMap(
	currentNode string,
	allNodes map[string][]string,
	currentPath []string,
	cntns func([]string, string) bool,
) [][]string {
	paths := make([][]string, 0)
	smol := strings.ToLower(currentNode) == currentNode
	// if we're on a small node, and we've seen this one before, return an empty path without the current visited
	// elements. This means we've reached a dead end as we've gone back to a small node that we've seen before.
	if smol && cntns(currentPath, currentNode) {
		return paths
	}

	// Add the allowTwice of the current node to the path slice.
	currentPath = append(currentPath, currentNode)

	// If we're at the endName node, we have reached the end of our journey.
	if currentNode == endName {
		return append(paths, currentPath)
	}

	for _, n := range allNodes[currentNode] {
		v := make([]string, len(currentPath))
		copy(v, currentPath)
		paths = append(paths, walkNodeMap(n, allNodes, v, cntns)...)
	}

	return paths
}
