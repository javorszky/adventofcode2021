package day22

type cubespace map[string]instruction

type step struct {
	Existing  map[string]instruction `json:"Existing,omitempty"`
	Incoming  map[string]instruction `json:"Incoming"`
	Result    map[string]instruction `json:"Result,omitempty"`
	Collapsed map[string]instruction `json:"Collapsed,omitempty"`
}

func (c *cubespace) applyInstructions(i instruction) step {
	if c.Length() == 0 {
		(*c)[i.String()] = i

		return step{
			Existing: map[string]instruction{},
			Incoming: map[string]instruction{
				i.String(): i,
			},
			Result: map[string]instruction{
				i.String(): i,
			},
			Collapsed: map[string]instruction{
				i.String(): i,
			},
		}
	}

	unaffected,
		affected,
		existing :=
		make(map[string]instruction),
		make(map[string]instruction),
		make(map[string]instruction)

	for k, v := range *c {
		existing[k] = v
		_, err := findOverlapBox(v, i)

		if err == nil {
			affected[k] = v

			continue
		}

		unaffected[k] = v
	}

	if len(affected) == 0 {
		if i.Flip == off {
			// nothing needs doing, return
			return step{
				Existing: existing,
				Incoming: map[string]instruction{
					i.String(): i,
				},
				Result:    existing,
				Collapsed: existing,
			}
		}

		*c = mergeMap(map[string]instruction{i.String(): i}, unaffected)

		result := make(map[string]instruction)
		for k, v := range *c {
			result[k] = v
		}

		c.Collapse()

		collapsed := make(map[string]instruction)
		for k, v := range *c {
			collapsed[k] = v
		}

		return step{
			Existing: existing,
			Incoming: map[string]instruction{
				i.String(): i,
			},
			Result:    result,
			Collapsed: collapsed,
		}
	}

	merge := make(map[string]instruction)

	for _, b := range affected {
		merge = mergeMap(filterOffs(overlapAndMerge(b, i)), merge)
	}

	*c = mergeMap(unaffected, merge)

	result := make(map[string]instruction)
	for k, v := range *c {
		result[k] = v
	}

	c.Collapse()

	collapsed := make(map[string]instruction)
	for k, v := range *c {
		collapsed[k] = v
	}

	return step{
		Existing: existing,
		Incoming: map[string]instruction{
			i.String(): i,
		},
		Result:    result,
		Collapsed: collapsed,
	}
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
