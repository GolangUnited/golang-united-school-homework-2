# Square task 

How to:
---
* Clone the repo  
* run `go mod init somename`
* run `go mod tidy`
* Edit `solution.go` 
	* it contains correct package name 
	* follow comments placeholder

Tasks:
---
Implement function to calculate square of an equilateral figurine following rules:
* `func CalcSquare(sideLen float64, sidesNum intCustomType) float64` 
* `CalcSquare` func must return correct square for:
   * equilateral triangle(3 sides),
   * square(4 sides) 
   * circle(0 sides) (count sideLen as radius)
   * if any other `sideNum` param is passed, return 0
* built-in Pi constant must be used to bypass the test
