package shapes

import (
	"math"
	"testing"
)

const deltaTesting = 0.001

func equalDeltaFloat64(a, b, delta float64) bool {
	return math.Abs(b-a) < delta
}

func Test_solutionSquare(t *testing.T) {
	data := []struct {
		In       float64
		Expected float64
	}{
		{In: 0, Expected: 0},
		{In: 1, Expected: 1},
		{In: 4.4, Expected: 19.36},
		{In: 15.67, Expected: 245.5489},
	}
	for _, q := range data {
		got := CalcSquare(q.In, SidesSquare)
		if !equalDeltaFloat64(got, q.Expected, deltaTesting) {
			t.Logf("Square4 expected: %f, got: %f", q.Expected, got)
			t.Fail()
		}
	}
}

func Test_solutionTriangle(t *testing.T) {
	data := []struct {
		In       float64
		Expected float64
	}{
		{In: 0, Expected: 0},
		{In: 1, Expected: 0.433013},
		{In: 4.4, Expected: 8.383126},
		{In: 15.67, Expected: 106.325793},
	}
	for _, q := range data {
		got := CalcSquare(q.In, SidesTriangle)
		if !equalDeltaFloat64(got, q.Expected, deltaTesting) {
			t.Logf("Square4 expected: %f, got: %f", q.Expected, got)
			t.Fail()
		}
	}
}

func Test_solutionCircle(t *testing.T) {
	data := []struct {
		In       float64
		Expected float64
	}{
		{In: 0, Expected: 0},
		{In: 1, Expected: 3.141593},
		{In: 4.4, Expected: 60.821234},
		{In: 15.67, Expected: 771.414620},
	}
	for _, q := range data {
		got := CalcSquare(q.In, SidesCircle)
		if !equalDeltaFloat64(got, q.Expected, deltaTesting) {
			t.Logf("Square0 expected: %f, got: %f", q.Expected, got)
			t.Fail()
		}
	}
}