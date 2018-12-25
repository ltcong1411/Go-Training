package circle

import (
	"math"
	"mymath"
)

func Area(r float32) float32 {
	// return r * r * math.Pi
	return mymath.Mul(r, r) * math.Pi
}

func Per(r float32) float32 {
	// return r * 2 * math.Pi
	return mymath.Mul(r, math.Pi) * 2
}
