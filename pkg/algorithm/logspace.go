package algorithm

import "math"

// create a list of n base10 logarithmically-spaced numbers between 10**start and 10**stop
// if max > min, list will be ascending, else it will be descending
func Logspace(start float64, stop float64, n uint32) []float64 {
	numbers := make([]float64, n)
	for i := 0; i < int(n); i++ {
		exponent := start + (float64(i) * (stop - start) / (float64(n) - 1))
		numbers[int(i)] = math.Pow(10., exponent)
	}
	return numbers
}
