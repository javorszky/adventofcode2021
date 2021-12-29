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

	c.Collapse()
}

func (c *cubespace) Length() int {
	return len(*c)
}

func (c *cubespace) Collapse() {
	overlaps := make([]instruction, c.Length())
	counter := 0

	for _, v := range *c {
		overlaps[counter] = v
		counter++
	}

	merged := mergeInstructionSlice(overlaps)
	mergedMap := make(map[string]instruction)

	for _, i := range merged {
		mergedMap[i.String()] = i
	}

	*c = mergedMap
}

func (c *cubespace) Lights() int {
	acc := 0

	for _, i := range *c {
		acc += i.Lights()
	}

	return acc
}

func mergeInstructionSlice(overlaps []instruction) []instruction {
	checked := make(map[string]instruction)

	for {
		merges := map[string]instruction{}

		for i, overlapBox := range overlaps {
			for _, overlapOtherBox := range overlaps[i+1:] {
				_, ok := checked[overlapBox.String()]
				_, ok2 := checked[overlapOtherBox.String()]

				if ok || ok2 {
					continue
				}

				m := mergeBoxes(overlapBox, overlapOtherBox)

				if len(m) == 1 {
					merges[m[0].String()] = m[0]
					checked[overlapBox.String()] = overlapBox
					checked[overlapOtherBox.String()] = overlapOtherBox
				}
			}
		}

		if len(merges) == 0 {
			break
		}

		newOverlaps := make(map[string]instruction)

		for _, _o := range overlaps {
			if _, ok := checked[_o.String()]; !ok {
				newOverlaps[_o.String()] = _o
			}
		}

		newOverlaps = mergeMap(merges, newOverlaps)
		checked = map[string]instruction{}
		overlaps = make([]instruction, len(newOverlaps))
		i := 0

		for _, v := range newOverlaps {
			overlaps[i] = v
			i++
		}
	}

	return overlaps
}
