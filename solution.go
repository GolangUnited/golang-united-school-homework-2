package square

import "math"

type sides int

func CalcSquare(sideLen float64, sidesNum sides) (res float64) {

	const SidesTriangle = 3
	const SidesSquare = 4
	const SidesCircle = 0

	switch sidesNum {
	case SidesTriangle:
		return math.Sqrt(3*sideLen*sideLen) / 4
	case SidesSquare:
		return sideLen * sideLen
	case SidesCircle:
		return math.Pi * sideLen * sideLen
	default:
		return 0
	}
}
