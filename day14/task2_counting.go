package day14

func task2Counting(template string, rules []string) int {
	counter := parseTemplateIntoCounter(template)
	bestRules := getBestMap(rules)

	for i := 0; i < 40; i++ {
		counter = workCounter(counter, bestRules)
	}

	return countElements(template, counter)
}
