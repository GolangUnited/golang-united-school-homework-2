package square

import "math"

type myType int

const (
	SidesCircle   myType = 0
	SidesTriangle myType = 3
	SidesSquare   myType = 4
)

func CalcSquare(sideLen float64, sidesNum myType) (res float64) {

	if sidesNum == SidesTriangle {
		res = (math.Pow(sideLen, 2.0) * math.Sqrt(3.0)) / 4
	} else if sidesNum == SidesSquare {
		res = math.Pow(sideLen, 2.0)
	} else if sidesNum == SidesCircle {
		res = math.Pi * math.Pow(sideLen, 2.0)
	}
	return
}
