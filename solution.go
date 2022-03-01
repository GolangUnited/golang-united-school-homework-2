package square

import (
	"math"
)

// Define custom int type to hold sides number and update CalcSquare signature by replacing #yourTypeNameHere#
type sidesType int

// Define constants to represent 0, 3 and 4 sides.  Test uses mnemos: SidesTriangle(==3), SidesSquare(==4), SidesCircle(==0)
// it's like:
// CalcSquare(10.0, SidesTriangle)
// CalcSquare(10.0, SidesSquare)
// CalcSquare(10.0, SidesCircle)
const (
	SidesTriangle sidesType = 3
	SidesSquare   sidesType = 4
	SidesCircle   sidesType = 0
)

func CalcSquare(sideLen float64, sidesNum sidesType) float64 {
	var square float64
	switch sidesNum {
	case 3:
		h := math.Sqrt(math.Pow(sideLen, 2) - math.Pow(sideLen/2, 2))
		square = h * sideLen / 2
	case 4:
		square = sideLen * sideLen
	case 0:
		square = math.Pi * math.Pow(sideLen, 2)
	default:
		square = 0
	}
	return square
}
