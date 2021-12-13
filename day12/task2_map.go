package day12

import (
	"strings"
	"sync"
)

func task2Map(input []string) int {
	nodes := parseIntoNodeMap(input)
	smolCaves := make([]string, 0)

	for n := range nodes {
		if strings.ToLower(n) == n && n != startName && n != endName {
			smolCaves = append(smolCaves, n)
		}
	}

	paths := make([][]string, 0, 10000)
	for _, smolCave := range smolCaves {
		paths = append(paths, walkNodeMap(startName, nodes, []string{}, containsTwice(smolCave))...)
	}

	uniquePaths := make(map[string]struct{})
	for _, p := range paths {
		uniquePaths[strings.Join(p, ",")] = struct{}{}
	}

	return len(uniquePaths)
}

func task2MapConcurrent(input []string) int {
	nodes := parseIntoNodeMap(input)
	smolCaves := make([]string, 0)

	for n := range nodes {
		if strings.ToLower(n) == n && n != startName && n != endName {
			smolCaves = append(smolCaves, n)
		}
	}

	paths := make([][]string, 0, 10000)
	wg := sync.WaitGroup{}
	queue := make(chan [][]string, 1)

	for _, smolCave := range smolCaves {
		wg.Add(1)

		go func(smc string) {
			queue <- walkNodeMap(startName, nodes, []string{}, containsTwice(smc))
		}(smolCave)
	}

	go func() {
		for m := range queue {
			paths = append(paths, m...)

			wg.Done()
		}
	}()

	wg.Wait()

	uniquePaths := make(map[string]struct{})
	for _, p := range paths {
		uniquePaths[strings.Join(p, ",")] = struct{}{}
	}

	return len(uniquePaths)
}
