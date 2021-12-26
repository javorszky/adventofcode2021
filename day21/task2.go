package day21

func task2(p1, p2 int) int {
	w := simulateMultiverses(p1, p2)

	if w.p1 > w.p2 {
		return w.p1
	}

	return w.p2
}

func calculatePossibilities() map[int]int {
	poss := make(map[int]int)

	for i := 1; i <= 3; i++ {
		for j := 1; j <= 3; j++ {
			for k := 1; k <= 3; k++ {
				poss[i+j+k]++
			}
		}
	}

	return poss
}

type wins struct {
	p1, p2 int
}

type universePossibilities struct {
	u Universe
	p int
}

func simulateMultiverses(p1, p2 int) wins {
	rollPossibilities := calculatePossibilities()

	universe := Universe{
		p1Step:  p1,
		p1Score: 0,
		p2Step:  p2,
		p2Score: 0,
	}

	w := wins{}
	possibilityMap := map[string]universePossibilities{
		universe.String(): {u: universe, p: 1},
	}
	ring := assembledTask1()

	c := 0

	for {
		if len(possibilityMap) == 0 {
			break
		}

		c++

		localMap := copyMap(possibilityMap)

		for encoded, possibilities := range localMap {
			delete(possibilityMap, encoded)

			for rollValue, frequency := range rollPossibilities {
				// Create separate copies for each branch.
				p1Ring := ring.rotateTo(possibilities.u.p1Step)
				p1Score := possibilities.u.p1Score

				rollRing := p1Ring.rotateBy(rollValue)
				p1Score += rollRing.value()

				if p1Score > 21 {
					w.p1 += possibilities.p * frequency

					continue
				}

				for rollValue2, frequency2 := range rollPossibilities {
					// Create separate copies for each branch
					p2Ring := ring.rotateTo(possibilities.u.p2Step)
					p2Score := possibilities.u.p2Score

					rollRing2 := p2Ring.rotateBy(rollValue2)
					p2Score += rollRing2.value()

					if p2Score > 21 {
						w.p2 += possibilities.p * frequency2 * frequency

						continue
					}

					newUni := Universe{
						p1Step:  rollRing.value(),
						p1Score: p1Score + rollRing.value(),
						p2Step:  rollRing2.value(),
						p2Score: p2Score + rollRing2.value(),
					}

					possibilityMap[newUni.String()] = struct {
						u Universe
						p int
					}{u: newUni, p: frequency * frequency2}
				}
			}
		}
	}

	return w
}

func copyMap(copyThis map[string]universePossibilities) map[string]universePossibilities {
	localCopy := make(map[string]universePossibilities)
	for k, v := range copyThis {
		localCopy[k] = v
	}

	return localCopy
}
