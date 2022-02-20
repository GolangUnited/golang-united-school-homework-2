package square

import "math"

// Define custom int type to hold sides number and update CalcSquare signature by replacing #yourTypeNameHere#

// Define constants to represent 0, 3 and 4 sides.  Test uses mnemos: SidesTriangle(==3), SidesSquare(==4), SidesCircle(==0)
// it's like:
// CalcSquare(10.0, SidesTriangle)
// CalcSquare(10.0, SidesSquare)
// CalcSquare(10.0, SidesCircle)

type myType int

func CalcSquare(sideLen float64, sidesNum myType) (res float64) {

	const SidesCircle = 0
	const SidesTriangle myType = 3
	const SidesSquare myType = 4

	if sidesNum == SidesTriangle {
		res = (math.Pow(sideLen, 2.0) * math.Sqrt(3.0)) / 4
	} else if sidesNum == SidesSquare {
		res = math.Pow(sideLen, 2.0)
	} else if sidesNum == SidesCircle {
		res = math.Pi * math.Pow(sideLen, 2.0)
	}
	return
}
