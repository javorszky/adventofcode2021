package day14

import (
	"log"
)

func task1Counting(template string, rules []string) int {
	counter := parseTemplateIntoCounter(template)
	bestRules := getBestMap(rules)

	for i := 0; i < 10; i++ {
		counter = workCounter(counter, bestRules)
	}

	return countElements(template, counter)
}

func parseTemplateIntoCounter(template string) map[string]int {
	polymer := make(map[string]int)

	for i := 0; i < len(template)-1; i++ {
		polymer[codePointToLetter[uint(template[i])]+codePointToLetter[uint(template[i+1])]]++
	}

	return polymer
}

func workCounter(counter map[string]int, bestRules map[string][2]string) map[string]int {
	newCounter := make(map[string]int)

	for k, v := range counter {
		newKeys, ok := bestRules[k]
		if !ok {
			log.Fatalf("Could not find key %s in best rules. This should not have happened", k)
		}

		newCounter[newKeys[0]] += v
		newCounter[newKeys[1]] += v
	}

	return newCounter
}

func getPair(in uint) string {
	return codePointToLetter[in>>8] + codePointToLetter[in&(1<<9-1)]
}

func countElements(template string, counter map[string]int) int {
	first := string(template[0])
	last := string(template[len(template)-1])
	counts := make(map[string]int)
	most := 0
	least := 1<<63 - 1

	for k, v := range counter {
		counts[string(k[0])] += v
		counts[string(k[1])] += v
	}

	for k, v := range counts {
		counts[k] = v / 2

		if k == first || k == last {
			counts[k]++
		}

		if counts[k] > most {
			most = counts[k]
		}

		if counts[k] < least {
			least = counts[k]
		}
	}

	return most - least
}
