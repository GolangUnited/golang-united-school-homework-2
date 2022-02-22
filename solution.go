package square

import "math"

// Define custom int type to hold sides number and update CalcSquare signature by replacing #yourTypeNameHere#

// Define constants to represent 0, 3 and 4 sides.  Test uses mnemos: SidesTriangle(==3), SidesSquare(==4), SidesCircle(==0)
// it's like:
// CalcSquare(10.0, SidesTriangle)
// CalcSquare(10.0, SidesSquare)
// CalcSquare(10.0, SidesCircle)

const (
	SidesCircle   = 0
	SidesTriangle = 3
	SidesSquare   = 4
	Pi            = 3.141592
)

func CalcSquare(sideLen float64, sidesNum int) float64 {
	if sideLen <= 0 {
		return 0
	}
	switch sidesNum {
	case SidesSquare:
		return AreaSquare(sideLen)
	case SidesTriangle:
		return AreaTriangle(sideLen)
	case SidesCircle:
		return AreaCircle(sideLen)
	default:
		return 0
	}
}

func AreaSquare(sideLen float64) float64 {
	return math.Pow(sideLen, 2)
}

func AreaTriangle(sideLen float64) float64 {
	return (math.Sqrt(3) / 4) * math.Pow(sideLen, 2)
}

func AreaCircle(sideLen float64) float64 {
	return Pi * math.Pow(sideLen, 2)
}
