package square

import (
	"math"
)

func CalcSquare(sideLen float64, sidesNum int) float64 {
	square := float64(0)
	if sidesNum == 3 {
		square = sideLen * sideLen / 4 * math.Sqrt(3)
	} else if sidesNum == 4 {
		square = math.Pow(sideLen, 2)
	} else if sidesNum == 0 {
		square = sideLen * math.Pi
	} else {
		print("Error. Not calculated because of incorrect data.")
	}

	return square
}
