package day21

func task1(player1, player2 int) int {
	p1Ring := assembledTask1()
	for p1Ring.value() != player1 {
		p1Ring = p1Ring.next()
	}

	p2Ring := assembledTask1()
	for p2Ring.value() != player2 {
		p2Ring = p2Ring.next()
	}

	p1Score, p2Score, die, c := 0, 0, 0, 0

	for {
		// this is one round
		p1Step := 0

		for i := 0; i < 3; i++ {
			p1d := (die % 100) + 1
			p1Step += p1d
			die++
			c++
		}

		p1Ring, _ = p1Ring.rotateBy(p1Step)

		p1Score += p1Ring.value()

		if p1Score >= 1000 {
			break
		}

		p2Step := 0

		for i := 0; i < 3; i++ {
			p2d := (die % 100) + 1
			p2Step += p2d
			die++
			c++
		}

		p2Ring, _ = p2Ring.rotateBy(p2Step)

		p2Score += p2Ring.value()

		if p2Score >= 1000 {
			break
		}
	}

	if p1Score < p2Score {
		return p1Score * c
	}

	return p2Score * c
}
