package day19

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

type position struct {
	x, y, z int
}

func (p position) String() string {
	return fmt.Sprintf("%d, %d, %d", p.x, p.y, p.z)
}

func (p position) rotations() [24]position {
	return [24]position{
		// Double checked
		{x: p.x, y: p.y, z: p.z},   // xpzp, self
		{x: p.x, y: -p.z, z: p.y},  // xpyp
		{x: p.x, y: -p.y, z: -p.z}, // xpzm (z minus also flips y by 180)
		{x: p.x, y: p.z, z: -p.y},  // xpym

		// Double checked
		{x: -p.x, y: -p.y, z: p.z},  // xmzp
		{x: -p.x, y: -p.z, z: -p.y}, // xmym
		{x: -p.x, y: p.y, z: -p.z},  // xmzm
		{x: -p.x, y: p.z, z: p.y},   // xmyp

		// Double checked
		{x: -p.y, y: p.x, z: p.z},  // ypzp
		{x: -p.z, y: p.x, z: -p.y}, // ypxm
		{x: p.y, y: p.x, z: -p.z},  // ypzm
		{x: p.z, y: p.x, z: p.y},   // ypxp

		// Double checked
		{x: p.y, y: -p.x, z: p.z},   // ymzp
		{x: -p.z, y: -p.x, z: p.y},  // ymxm
		{x: -p.y, y: -p.x, z: -p.z}, // ymzm
		{x: p.z, y: -p.x, z: -p.y},  // ymxp

		// Double checked
		{x: -p.z, y: p.y, z: p.x},  // zpxm
		{x: -p.y, y: -p.z, z: p.x}, // zpym
		{x: p.z, y: -p.y, z: p.x},  // zpxp
		{x: p.y, y: p.z, z: p.x},   // zpyp

		// Double checked
		{x: p.z, y: p.y, z: -p.x},   // zmxp
		{x: -p.y, y: p.z, z: -p.x},  // zmyp
		{x: -p.z, y: -p.y, z: -p.x}, // zmxm
		{x: p.y, y: -p.z, z: -p.x},  // zmym
	}
}

func distance(p1, p2 position) int {
	x := p1.x - p2.x
	y := p1.y - p2.y
	z := p1.z - p2.z

	return x*x + y*y + z*z
}

func parseBeacon(s string) position {
	parts := strings.Split(s, ",")
	if len(parts) != 3 {
		log.Fatalf("tryng to parse beacon, expecting three numbers, got %d: %s", len(parts), s)
	}

	numbers := make([]int, 3)

	for j, p := range parts {
		n, err := strconv.Atoi(p)
		if err != nil {
			log.Fatalf("tried to convert string [%s] to int, but failed in beacon: %s", p, err)
		}

		numbers[j] = n
	}

	return position{
		x: numbers[0],
		y: numbers[1],
		z: numbers[2],
	}
}

func shiftPositionBy(shiftThis, by position) position {
	return position{
		x: shiftThis.x - by.x,
		y: shiftThis.y - by.y,
		z: shiftThis.z - by.z,
	}
}
