package square


import (
	"math"
)

type SidesNumber int

const (
	SidesCircle   SidesNumber = 0
	SidesTriangle SidesNumber = 3
	SidesSquare   SidesNumber = 4
)

func CalcSquare(sideLen float64, sidesNum SidesNumber) float64 {
	switch sidesNum {
	case SidesTriangle:
		return (sideLen * sideLen * math.Sqrt(3) / 4)
	case SidesSquare:
		return (sideLen * sideLen)
	case SidesCircle:
		return (math.Pi * sideLen * sideLen)
	default:
		return (0)
	}
}
