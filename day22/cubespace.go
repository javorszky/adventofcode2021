package day22

type cubespace map[string]instruction

func (c *cubespace) applyInstructions(i instruction) {
	if c.Length() == 0 {
		(*c)[i.String()] = i

		return
	}

	merge := make(map[string]instruction)

	for _, b := range *c {
		merge = mergeMap(filterOffs(overlapAndMerge(b, i)), merge)
	}

	*c = merge
}

func (c *cubespace) Length() int {
	return len(*c)
}
