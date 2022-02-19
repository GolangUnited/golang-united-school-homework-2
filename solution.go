package square

import "math"

// Define custom int type to hold sides number and update CalcSquare signature by replacing #yourTypeNameHere#

// Define constants to represent 0, 3 and 4 sides.  Test uses mnemos: SidesTriangle(==3), SidesSquare(==4), SidesCircle(==0)
// it's like:
// CalcSquare(10.0, SidesTriangle)
// CalcSquare(10.0, SidesSquare)
// CalcSquare(10.0, SidesCircle)
type myInt int

var area float64

const (
	SidesCircle   myInt = 0
	SidesTriangle myInt = 3
	SidesSquare   myInt = 4
)

func CalcSquare(sideLen float64, sidesNum myInt) float64 {
	if sidesNum == SidesCircle {
		area = math.Pi * (sideLen * sideLen)
	} else if sidesNum == SidesTriangle {
		area = (math.Sqrt(3) / 4) * (sideLen * sideLen)
	} else if sidesNum == SidesSquare {
		area = sideLen * sideLen
	}
	return area
}
