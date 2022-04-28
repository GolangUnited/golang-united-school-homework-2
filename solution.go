package square

import (
	"math"
)

const (
	SidesCircle   SidesNumber = 0
	SidesTriangle SidesNumber = 3
	SidesSquare   SidesNumber = 4
)

type SidesNumber int

func CalcSquare(sideLen float64, sidesNum SidesNumber) float64 {

	square := float64(0)

	if sidesNum == SidesTriangle {
		square = sideLen * sideLen / 4 * math.Sqrt(3)
	} else if sidesNum == SidesSquare {
		square = math.Pow(sideLen, 2)
	} else if sidesNum == SidesCircle {
		square = sideLen * math.Pi
	} else {
		print("Error. Not calculated because of incorrect data.")
	}

	return square
}
