package gomath

func Sum(x []float64) float64 {
	sum := float64(0)
	for i := 0; i < len(x); i++ {
		sum += x[i]
	}

	return sum
}

