package day19

import "fmt"

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
		{x: p.x, y: p.y, z: p.z},   // xpzp, self
		{x: p.x, y: p.z, z: -p.y},  // xpym
		{x: p.x, y: -p.y, z: -p.z}, // xpzm (z minus also flips y by 180)
		{x: p.x, y: p.z, z: -p.y},  // xpyp

		{x: -p.x, y: -p.y, z: p.z},  // xmzp
		{x: -p.x, y: -p.z, z: -p.y}, // xmym
		{x: -p.x, y: p.y, z: -p.z},  // xmzm
		{x: -p.x, y: p.z, z: p.y},   // xmyp

		{x: p.y, y: -p.x, z: p.z},  // ypzp
		{x: p.y, y: -p.z, z: -p.x}, // ypxm
		{x: p.y, y: p.x, z: -p.z},  // ypzm
		{x: p.y, y: p.z, z: p.x},   // ypxp

		{x: -p.y, y: p.x, z: p.z},   // ymzp
		{x: -p.y, y: -p.z, z: p.x},  // ymxm
		{x: -p.y, y: -p.x, z: -p.z}, // ymzm
		{x: -p.z, y: p.z, z: -p.x},  // ymxp

		{x: p.z, y: p.y, z: -p.x},  // zpxm
		{x: p.z, y: p.x, z: p.y},   //zpym
		{x: p.z, y: -p.y, z: p.x},  // zpxp
		{x: p.z, y: -p.x, z: -p.y}, // zpyp

		{x: -p.z, y: p.y, z: p.x},   // zmxp
		{x: -p.z, y: -p.x, z: p.y},  // zmym
		{x: -p.z, y: -p.y, z: -p.x}, // zmxm
		{x: -p.z, y: p.x, z: -p.y},  // zmyp
	}
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
