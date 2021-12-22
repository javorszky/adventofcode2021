package day18

func task2(input []string) int {
	max := 0

	for i, vs := range input {
		for j, ws := range input {
			if i == j {
				continue
			}

			v := parse(vs)
			w := parse(ws)
			wv := addNodes(v, w)
			wvMag := wv.Magnitude()

			if max < wvMag {
				max = wvMag
			}
		}
	}

	return max
}
