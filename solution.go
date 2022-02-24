package square

import "math"

// Define custom int type to hold sides number and update CalcSquare signature by replacing #yourTypeNameHere#

// Define constants to represent 0, 3 and 4 sides.  Test uses mnemos: SidesTriangle(==3), SidesSquare(==4), SidesCircle(==0)
// it's like:
// CalcSquare(10.0, SidesTriangle)
// CalcSquare(10.0, SidesSquare)
// CalcSquare(10.0, SidesCircle)

const (
	SidesCircle SidesNumber =0
	SidesTriangle SidesNumber =3
	SidesSquare SidesNumber =4
)

func CalcSquare(sideLen float64, sidesNum SidesNumber) float64 {

	switch sidesNum{
	case SidesTriangle:
		return (sideLen*sideLen*math.Sqrt(3)/4)
	
	case SidesSquare:
		return (sideLen*sideLen)

	case SidesCircle:
		return (math.Pi*sideLen*sideLen)
		
	default:
		return (0)

}
