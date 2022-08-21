package populator

import (
	"math"
	"math/rand"
)

func GeneratePoissonInterval(lambda float64) float64 {
	return -1000 / lambda * math.Log(rand.Float64())
}
