package square

import "math"

type SidesCount int

const (
	SidesCircle   SidesCount = 0
	SidesTriangle SidesCount = 3
	SidesSquare   SidesCount = 4
)

func CalcSquare(sideLen float64, sidesNum SidesCount) float64 {
	if sidesNum == SidesCircle {
		return math.Pi * sideLen * sideLen
	}

	if sidesNum == SidesTriangle {
		return (math.Sqrt(3) / 4) * sideLen * sideLen
	}

	if sidesNum == SidesSquare {
		return sideLen * sideLen
	}

	return 0
}
