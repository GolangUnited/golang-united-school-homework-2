package square

import (
	"math"
	//"reflect"
)

// Define custom int type to hold sides number and update CalcSquare signature by replacing #yourTypeNameHere#

// Define constants to represent 0, 3 and 4 sides.  Test uses mnemos: SidesTriangle(==3), SidesSquare(==4), SidesCircle(==0)
// it's like:
//CalcSquare(10.0, SidesTriangle)
// CalcSquare(10.0, SidesSquare)
// CalcSquare(10.0, SidesCircle)

const SidesCircle int64 = 0
const SidesTriangle int64 = 3
const SidesSquare int64 = 4

func CalcSquare(sideLen float64, sidesNum int64) float64 {
	if sidesNum < 0 || sidesNum == 1 {
		return 0
	}
	if sidesNum == 0 {

		return math.Pi * math.Pow(sideLen, 2)
	}

	return (sideLen * sideLen) / 4.0 * float64(sidesNum) * (math.Cos(float64(math.Pi*2.0/float64(sidesNum*2))) / math.Sin(float64(math.Pi*2/float64(sidesNum*2))))
}
