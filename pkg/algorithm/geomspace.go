package algorithm

import "math"

// create a list of n logarithmically-spaced numbers between start and stop
// similar to Logspace, except outputs are between [start, stop], not [base**start, base**stop]
// if start > stop, list will be ascending, else it will be descending
func Geomspace(start float64, stop float64, n uint32) []float64 {
	logStart := math.Log10(start)
	logStop := math.Log10(stop)
	numbers := make([]float64, n)

	for i := 0; i < int(n); i++ {
		exponent := logStart + (float64(i) * (logStop - logStart) / (float64(n) - 1))
		numbers[int(i)] = math.Pow(10, exponent)
	}

	return numbers
}
