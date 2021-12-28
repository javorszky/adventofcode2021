package day22

type cubespace map[int]flip

func (c *cubespace) applyInstructions(i instruction) {

}

func bla() {
	_ = []instruction{
		{xFrom: -20, xTo: -5, yFrom: -20, yTo: -5, zFrom: -20, zTo: -5, flip: 1},
		{xFrom: -5, xTo: 5, yFrom: -20, yTo: 5, zFrom: -20, zTo: -5, flip: 1},
		{xFrom: -20, xTo: 5, yFrom: -20, yTo: -5, zFrom: -5, zTo: 5, flip: 1},
		{xFrom: -20, xTo: -5, yFrom: -5, yTo: 5, zFrom: -20, zTo: 5, flip: 1},
		{xFrom: -5, xTo: 20, yFrom: -5, yTo: 20, zFrom: -5, zTo: 20, flip: 2},
	}

	_ = []instruction{
		{xFrom: -20, xTo: -5, yFrom: -20, yTo: -5, zFrom: -20, zTo: -5, flip: 1},
		{xFrom: -5, xTo: 5, yFrom: -20, yTo: 5, zFrom: -20, zTo: -5, flip: 1},
		{xFrom: -20, xTo: 5, yFrom: -20, yTo: -5, zFrom: -5, zTo: 5, flip: 1},
		{xFrom: -20, xTo: -5, yFrom: -5, yTo: 5, zFrom: -20, zTo: 5, flip: 1},
		{xFrom: -5, xTo: 20, yFrom: -5, yTo: 20, zFrom: -5, zTo: 20, flip: 2},
	}
}
