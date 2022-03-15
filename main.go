package main

import (
	square "github.com/golang-united-school-homework-2/square"
)

func main()  {
	a := square.CalcSquare(9.0, square.SidesTriangle)
	b := square.CalcSquare(9.0, square.SidesSquare)
	c := square.CalcSquare(9.0, square.SidesCircle)
	println(a)
	println(b)
	println(c)
}
