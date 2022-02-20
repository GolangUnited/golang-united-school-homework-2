package square

import "math"

// Define custom int type to hold sides number and update CalcSquare signature by replacing #yourTypeNameHere#
type sides int

// Define constants to represent 0, 3 and 4 sides.  Test uses mnemos: SidesTriangle(==3), SidesSquare(==4), SidesCircle(==0)
const (
	SidesTriangle = sides(3)
	SidesSquare   = sides(4)
	SidesCircle   = sides(0)
	Pi            = 3.1415
)

// it's like:
// CalcSquare(10.0, SidesTriangle)
// CalcSquare(10.0, SidesSquare)
// CalcSquare(10.0, SidesCircle)

func CalcSquare(sideLen float64, sidesNum sides) float64 {
	if sidesNum == SidesTriangle {
		return sideLen * sideLen * sideLen
	}
	if sidesNum == SidesSquare {
		return math.Sqrt(3) / 4.0 * sideLen * sideLen
	}
	if sidesNum == SidesCircle {
		return Pi * sideLen * sideLen
	}

	return 0
}
