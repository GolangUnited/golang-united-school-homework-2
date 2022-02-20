package square

import "math"

// Define custom int type to hold sides number and update CalcSquare signature by replacing #yourTypeNameHere#

// Define constants to represent 0, 3 and 4 sides.  Test uses mnemos: SidesTriangle(==3), SidesSquare(==4), SidesCircle(==0)
// it's like:
// CalcSquare(10.0, SidesTriangle)
// CalcSquare(10.0, SidesSquare)
// CalcSquare(10.0, SidesCircle)

type SidesTriangle int
type SidesSquare int
type SidesCircle int

const triangle SidesTriangle = 3
const square SidesSquare = 4
const circle SidesCircle = 0

func CalcSquare(sideLen float64, sidesNum int) float64 {
	var square float64
	switch sidesNum {
	case 3:
		square = float64((sideLen * sideLen * math.Sqrt(3)) / 4)
	case 4:
		square = float64(sideLen * sideLen)
	case 0:
		square = float64(math.Pi * sideLen * sideLen)
	}
	return square
}
