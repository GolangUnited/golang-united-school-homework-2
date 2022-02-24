package square

import "math"

// Define custom int type to hold sides number and update CalcSquare signature by replacing #yourTypeNameHere#

// Define constants to represent 0, 3 and 4 sides.  Test uses mnemos: SidesTriangle(==3), SidesSquare(==4), SidesCircle(==0)
// it's like:
// CalcSquare(10.0, SidesTriangle)
// CalcSquare(10.0, SidesSquare)
// CalcSquare(10.0, SidesCircle)
type sidesType int64

const (
	SidesCircle   int64 = 0
	SidesTriangle int64 = 3
	SidesSquare   int64 = 4
)

func CalcSquare(sideLen float64, sidesNum sidesType) float64 {
	switch sidesNum {
	case sidesType(SidesSquare):
		return (math.Pow(sideLen, 2))
	case sidesType(SidesTriangle):
		return (math.Pow(sideLen, 2) * math.Sqrt(3) / 4)
	case sidesType(SidesCircle):
		return (math.Pow(sideLen, 2) * math.Pi)
	default:
		return (0)
	}
}
