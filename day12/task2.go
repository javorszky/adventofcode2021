package day12

import (
	"strings"
)

func task2(input []string) int {
	nodes := parseIntoNodes(input)
	smolCaves := make([]string, 0)

	for _, n := range nodes {
		if n.smol && n.name != "start" && n.name != "end" {
			smolCaves = append(smolCaves, n.name)
		}
	}

	paths := make([][]string, 0, 10000)
	for _, smolCave := range smolCaves {
		paths = append(paths, walkNodes(nodes["start"], []string{}, containsTwice(smolCave))...)
	}

	uniquePaths := make(map[string]struct{})
	for _, p := range paths {
		uniquePaths[strings.Join(p, ",")] = struct{}{}
	}

	return len(uniquePaths)
}

func containsTwice(name string) func([]string, string) bool {
	return func(paths []string, element string) bool {
		twiceCounter := 1
		// For each element in the segment.
		for _, p := range paths {
			// If the current element in the iteration is the same as the one passed to the function.
			if p == element {
				// But if it's the one we allowed a twice counter.
				if p == name {
					// And it hasn't seen it twice yet.
					if twiceCounter > 0 {
						// Then decrement the counter.
						twiceCounter--

						continue
					} else {
						// Otherwise yell "yah, we've been here twice".
						return true
					}
				}

				return true
			}
		}

		return false
	}
}
