package algorithm

import (
	"internal/asserts"
	"math"
	"testing"
)

//
// helpers
//

func round(num float64) int {
	return int(num + math.Copysign(0.5, num))
}

func toFixed(num float64, precision int) float64 {
	output := math.Pow(10, float64(precision))
	return float64(round(num*output)) / output
}

//
// tests
//

func TestLogspaceNoElements(t *testing.T) {
	expected := []float64{}
	observed := Logspace(10., 100., 0, 10.)

	asserts.Float64SliceEquals(expected, observed, t)
}

func TestLogspaceAscending(t *testing.T) {
	expected := []float64{100, 215.443469, 464.158883, 1000}
	observed := Logspace(2., 3., 4, 10.)

	for i := 0; i < len(observed); i++ {
		observed[i] = toFixed(observed[i], 6)
	}

	asserts.Float64SliceEquals(expected, observed, t)
}

func TestLogspaceDescending(t *testing.T) {
	expected := []float64{1000, 464.158883, 215.443469, 100}
	observed := Logspace(3., 2., 4, 10.)

	for i := 0; i < len(observed); i++ {
		observed[i] = toFixed(observed[i], 6)
	}

	asserts.Float64SliceEquals(expected, observed, t)
}
