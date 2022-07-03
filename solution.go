package square

// Define custom int type to hold sides number and update CalcSquare signature by replacing #yourTypeNameHere#

// Define constants to represent 0, 3 and 4 sides.  Test uses mnemos: SidesTriangle(==3), SidesSquare(==4), SidesCircle(==0)
// it's like:
// CalcSquare(10.0, SidesTriangle)
// CalcSquare(10.0, SidesSquare)
// CalcSquare(10.0, SidesCircle)

import (
	"math"
)

const (
	SidesTriangle CustomIntType = 3
	SidesSquare   CustomIntType = 4
	SidesCircle   CustomIntType = 0
	pi                          = math.Pi
)

type CustomIntType int

func CalcSquare(sideLen float64, sidesNum CustomIntType) float64 {
	var result float64
	switch sidesNum {
	case 4:
		result = sideLen * sideLen
	case 3:
		result = ((sideLen * sideLen) * math.Sqrt(3)) / 4
	case 0:
		result = pi * (sideLen * sideLen)
	default:
		result = 0.0
	}
	return result
}
