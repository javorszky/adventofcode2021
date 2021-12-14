package day14

import (
	"fmt"
	"sort"

	"github.com/javorszky/adventofcode2021/util"
)

func task2(template string, rules []string) int {
	betterRules := parseBetterRules(rules)
	polymer := parseTemplateLinkedList(template)

	fmt.Printf("step 0%s%v%s%s", util.NewLine, getDistributions(polymer), util.NewLine, util.NewLine)

	for j := 0; j < 10; j++ {
		workLinkedList(polymer, betterRules)

		fmt.Printf("step %d%s%v%s%s", j+1, util.NewLine, getDistributions(polymer), util.NewLine, util.NewLine)
	}

	counts := map[uint]int{}
	walker := polymer
	most := 0
	least := 0

	for {
		counts[walker.Point()]++
		walker = walker.Next()

		if walker == nil {
			break
		}

		least++
	}

	for _, v := range counts {
		if v > most {
			most = v
		}

		if v < least {
			least = v
		}
	}

	return most - least
}

func getDistributions(polymer *polymerElement) []string {
	counts := map[uint]int{}
	walker := polymer

	for {
		counts[walker.Point()]++
		walker = walker.Next()

		if walker == nil {
			break
		}
	}

	countsString := make(map[string]int)
	forSorting := make([]int, 0)
	sorted := make([]string, 0)
	reverse := make(map[int][]string)

	for k, v := range counts {
		countsString[codePointToLetter[k]] = v
		forSorting = append(forSorting, v)

		if reverse[v] == nil {
			reverse[v] = make([]string, 0)
		}

		reverse[v] = append(reverse[v], codePointToLetter[k])
	}

	sort.Ints(forSorting)

	for i := 0; i < len(forSorting)/2; i++ {
		forSorting[i], forSorting[len(forSorting)-1-i] = forSorting[len(forSorting)-1-i], forSorting[i]
	}

	prevHz := 0
	for _, hz := range forSorting {
		if hz == prevHz {
			continue
		}

		prevHz = hz

		for _, letter := range reverse[hz] {
			sorted = append(sorted, fmt.Sprintf("%s: %d", letter, hz))
		}
	}

	return sorted
}
