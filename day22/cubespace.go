package day22

import "fmt"

type cubespace map[string]instruction

func (c cubespace) applyInstructions(i instruction) {
	if c.Length() == 0 {
		c[i.String()] = i

		return
	}

	for _, b := range c {
		bla := filterOffs(overlapAndMerge(b, i))
		fmt.Printf("bla after instructions\n\n%#v\n", bla)
	}
}

func (c cubespace) Length() int {
	return len(c)
}
