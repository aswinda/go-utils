package gomath

func Sum(x []float64) float64 {
	sum := float64(0)
	for i := 0; i < len(x); i++ {
		sum += x[i]
	}

	return sum
}

func Average(x []float64) float64 {
	sum := Sum(x)
	n := len(x)

	avg := sum/n

	return avg
}

func Min(x []float64) float64 {
	min := x[0]

	for i := 1; i < len(x); i++ {
		if min < x[i]
			min := x[i]
	}

	return min
}
