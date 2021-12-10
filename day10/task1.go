package day10

func task1(input []string) int {
	r := getChunkReplacer()
	or := getOpenerReplacer()
	acc := 0

	for _, line := range input {
		for {
			l := len(line)
			line = r.Replace(line)

			if l == len(line) {
				for {
					k := len(line)
					line = or.Replace(line)

					if k == len(line) {
						break
					}
				}

				break
			}
		}

		if line == `` {
			continue
		}

		acc += invalidScore[line[0]]
	}

	return acc
}
