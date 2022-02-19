package square

import "math"

// Define custom int type to hold sides number and update CalcSquare signature by replacing #yourTypeNameHere#

// Define constants to represent 0, 3 and 4 sides.  Test uses mnemos: SidesTriangle(==3), SidesSquare(==4), SidesCircle(==0)
// it's like:
// CalcSquare(10.0, SidesTriangle)
// CalcSquare(10.0, SidesSquare)
// CalcSquare(10.0, SidesCircle)
type MyInt int

var area float64

const (
	SidesCircle   MyInt = 0
	SidesTriangle MyInt = 3
	SidesSquare   MyInt = 4
)

func CalcSquare(sideLen float64, sidesNum MyInt) float64 {
	if sidesNum == SidesCircle {
		area = math.Pi * (sideLen * sideLen)
	} else if SidesTriangle == 3 {
		area = (math.Sqrt(3) / 4) * (sideLen * sideLen)
	} else {
		area = sideLen * sideLen
	}

	return area
}
