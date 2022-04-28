package square

import (
	"math"
)

func CalcSquare(sideLen float64, sidesNum int) float64 {
	SidesSquare := int(4)

	SidesTriangle := int(3)

	SidesCircle := int(0)

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
