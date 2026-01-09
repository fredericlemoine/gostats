package gostats

import (
	"math"
)

// Generate from exponential distribution
func (gsr *GoStatRand) Exp(lambda float64) float64 {
	exp := gsr.rand.Float64()
	exp = -math.Log(1-exp) / lambda
	return (exp)
}
