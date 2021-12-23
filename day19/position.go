package day19

import (
	"fmt"
	"math"
)

type position struct {
	x, y, z int
}

func (p position) String() string {
	return fmt.Sprintf("%d, %d, %d", p.x, p.y, p.z)
}

//
//const (
//	xPzP int = iota
//	xPyP
//	xPzM
//	xPyM
//
//	xMzP
//	xMyP
//	xMzM
//	xMyM
//
//	yPzP
//	yPxP
//	yPzM
//	yPxM
//
//	yMzP
//	yMxP
//	yMzM
//	yMxM
//
//	zPxP
//	zPyP
//	zPxM
//	zPyM
//
//	zMxP
//	zMyP
//	zMxM
//	zMyM
//)

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

func distance(p1, p2 position) float64 {
	x := float64(p1.x - p2.x)
	y := float64(p1.y - p2.y)
	z := float64(p1.z - p2.z)

	return math.Sqrt(x*x + y*y + z*z)
}

// rotations
// z 90
// z 180
// z 270
// x 90
// x 180
// x 270
// y 90
// y 180
// y 270
