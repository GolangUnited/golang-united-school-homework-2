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
	SidesCircle   figureSide = 0
	SidesTriangle figureSide = 3
	SidesSquare   figureSide = 4
)

func CalcSquare(sideLen float64, sidesNum figureSide) float64 {
	var area float64 = 0

	switch sidesNum {
	case SidesCircle:
		area = area + math.Pi*(sideLen*sideLen)
	case SidesTriangle:
		area = area + (math.Sqrt(3)/4)*(sideLen*sideLen)
	case SidesSquare:
		area = area + sideLen*sideLen
	}
	return area
}
