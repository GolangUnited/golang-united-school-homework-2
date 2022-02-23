package square

import "math"

// Define custom int type to hold sides number and update CalcSquare signature by replacing #yourTypeNameHere#

// Define constants to represent 0, 3 and 4 sides.  Test uses mnemos: SidesTriangle(==3), SidesSquare(==4), SidesCircle(==0)
// it's like:
// CalcSquare(10.0, SidesTriangle)
// CalcSquare(10.0, SidesSquare)
// CalcSquare(10.0, SidesCircle)

type sideNumber int

const (
	sideCircle    sideNumber = 0
	sideTriangle  sideNumber = 3
	sideRectangle sideNumber = 4
)

func CalcSquare(sideLen float64, sidesNum sideNumber) float64 {
	sideLenSquare := sideLen * sideLen
	switch sidesNum {
	case sideCircle:
		return sideLenSquare * math.Pi
	case sideTriangle:
		return sideLenSquare * math.Sqrt(3) / 4
	case sideRectangle:
		return sideLenSquare
	}
	return 0
}
