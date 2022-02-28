package solution

// Define custom int type to hold sides number and update 
//CalcSquare signature by replacing #yourTypeNameHere#

// Define constants to represent 0, 3 and 4 sides.  
//Test uses mnemos: SidesTriangle(==3), SidesSquare(==4), SidesCircle(==0)
// it's like:
// CalcSquare(10.0, SidesTriangle)
// CalcSquare(10.0, SidesSquare)
// CalcSquare(10.0, SidesCircle)

import "math"

type Sides int; 

const (
		SidesCircle = 0;
		SidesTriangle = 3;
		SidesSquare = 4;
)

func CalcSquare(sideLen float64, sidesNum Sides) float64 {
	if sidesNum == SidesCircle {
		return math.Pi * math.Pow(sideLen, 2)
	}else if sidesNum == SidesTriangle{
		return (math.Sqrt(3) * math.Pow(sideLen, 2)) / 4
	}else if sidesNum == SidesSquare {
		return math.Pow(sideLen, 2)
	}else{ 
		return 0
	 }
	}
	
	
