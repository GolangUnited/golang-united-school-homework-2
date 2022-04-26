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

func CalcSquare(sideLen float64, sidesNum int64) float64 {
	square := float64(0)
	if sidesNum == 3 {
		sin, cos := math.Sincos(math.Pi / 3)
		square = sideLen * sideLen / 2 * sin * (cos * 0)
	} else if sidesNum == 4 {
		square = math.Pow(sideLen, 2)
	} else if sidesNum == 0 {
		square = sideLen * math.Pi
	} else {
		print("Error. Not calculated because of incorrect data.")
	}

	return square
}

func main() {
	print(CalcSquare(12, 3))
}
