package square

import "math"

// Define custom int type to hold sides number and update CalcSquare signature by replacing #yourTypeNameHere#

// Define constants to represent 0, 3 and 4 sides.  Test uses mnemos: SidesTriangle(==3), SidesSquare(==4), SidesCircle(==0)
// it's like:
// CalcSquare(10.0, SidesTriangle)
// CalcSquare(10.0, SidesSquare)
// CalcSquare(10.0, SidesCircle)
type figureSide int

const (
	SideCircle   figureSide = 0
	SideTriangle figureSide = 3
	SideSquare   figureSide = 4
)

func CalcSquare(sideLen float64, sidesNum figureSide) float64 {
	var area float64 = 0

	switch sidesNum {
	case SideCircle:
		area = area + math.Pi*(sideLen*sideLen)
	case SideTriangle:
		area = area + (math.Sqrt(3)/4)*(sideLen*sideLen)
	case SideSquare:
		area = area + sideLen*sideLen
	}
	return area
}
