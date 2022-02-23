package square

import (
	"math"
)

type figurSide int

// Define custom int type to hold sides number and update CalcSquare signature by replacing #yourTypeNameHere#

// Define constants to represent 0, 3 and 4 sides.  Test uses mnemos: SidesTriangle(==3), SidesSquare(==4), SidesCircle(==0)
// it's like:
// CalcSquare(10.0, SidesTriangle)
// CalcSquare(10.0, SidesSquare)
// CalcSquare(10.0, SidesCircle)

const (
	SidesTriangle = 3
	SidesSquare   = 4
	SidesCircle   = 0
)

func CalcSquare(sideLen float64, sidesNum figurSide) float64 {
	switch sidesNum {
	case SidesCircle:
		return float64(math.Pi * math.Pow(sideLen, 2))
	case SidesTriangle:
		return float64((math.Pow(sideLen, 2) * math.Sqrt(3)) / 4)
	case SidesSquare:
		return float64(math.Pow(sideLen, 2))
	}
	return float64(0)
}
