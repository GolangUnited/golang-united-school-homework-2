package square
import "math"

// Define custom int type to hold sides number and update CalcSquare signature by replacing #yourTypeNameHere#

// Define constants to represent 0, 3 and 4 sides.  Test uses mnemos: SidesTriangle(==3), SidesSquare(==4), SidesCircle(==0)
// it's like:
// CalcSquare(10.0, SidesTriangle)
// CalcSquare(10.0, SidesSquare)
// CalcSquare(10.0, SidesCircle)
type sideNum int

const SidesTriangle sideNum = 3
const SidesSquare sideNum	= 4
const SidesCircle sideNum = 0

func CalcSquare(sideLen float64, sidesNum sideNum) float64 {
	if sidesNum == SidesTriangle {
		var h float64 = (sideLen*math.Sqrt(3))/2
		return 1.0/2.0 * h * sideLen
	} else if sidesNum == SidesSquare {
		return sideLen * sideLen
	}else {
		return math.Pi*sideLen*sideLen
	}
}

