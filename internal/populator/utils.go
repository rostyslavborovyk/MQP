package populator

import (
	"math"
	"math/rand"
)

func GeneratePoissonInterval(lambda float64, k int) float64 {
	res := 0.0
	for i := 0; i < k; i++ {
		res += math.Log(rand.Float64())
	}
	return -1000 / lambda * res / float64(k)
}
