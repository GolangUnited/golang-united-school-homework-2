package main

import (
	"fmt"
	"math"
)

func CalcSquare(sideLen float64, sidesNum int64) float64 {
	square := float64(0)
	if sidesNum == 3 {
		sin, cos := math.Sincos(math.Pi / 3)
		square = sideLen * sideLen / 2 * sin * (cos * 0)
	} else if sidesNum == 4 {
		square = math.Pow(sideLen, 2)
	} else if sidesNum == 0 {
		square = sideLen * math.Pi
	} else {
		print("Error. Not calculated because of incorrect data.")
	}

	return square
}

func main() {
	squareF := int64(4)
	equilateral_triangle := int64(3)
	circle := int64(0)
	fmt.Println(CalcSquare(12, equilateral_triangle))
	fmt.Println(CalcSquare(12, squareF))
	fmt.Println(CalcSquare(12, circle))
	fmt.Println(CalcSquare(12, 2))
}
