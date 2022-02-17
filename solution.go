package square

import (
	"math"
)

// Define custom int type to hold sides number and update CalcSquare
// signature by replacing #yourTypeNameHere#

// Define constants to represent 0, 3 and 4 sides.
//Test uses mnemos: SidesTriangle(==3), SidesSquare(==4), SidesCircle(==0)
// it's like:
// CalcSquare(10.0, SidesTriangle)
// CalcSquare(10.0, SidesSquare)
// CalcSquare(10.0, SidesCircle)
type MyTypeSides int

func CalcSquare(sideLen float64, sidesNum MyTypeSides) float64 {

	const (
		SidesTriangle = 3
		SidesSquare   = 4
		SidesCircle   = 0
	)

	var result float64

	if sidesNum == SidesTriangle {

		result = float64(sideLen) * 3

	} else if sidesNum == SidesSquare {
		result = sideLen * sideLen
	} else if sidesNum == SidesCircle {
		result = math.Pi * float64(sideLen*sideLen)
	}
	return result
}
