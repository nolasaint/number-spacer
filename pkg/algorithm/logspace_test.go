package algorithm

import (
	"math"
	"reflect"
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

func assertFloat64SliceEquals(expected []float64, observed []float64, t *testing.T) {
	if !reflect.DeepEqual(expected, observed) {
		t.Fatalf("\nexpected: %v\nbut got:  %v", expected, observed)
	}
}

//
// tests
//

func TestLogspaceNoElements(t *testing.T) {
	expected := []float64{}
	observed := Logspace(10., 100., 0)

	assertFloat64SliceEquals(expected, observed, t)
}

func TestLogspaceAscending(t *testing.T) {
	expected := []float64{100, 215.443469, 464.158883, 1000}
	observed := Logspace(2., 3., 4)

	for i := 0; i < len(observed); i++ {
		observed[i] = toFixed(observed[i], 6)
	}

	assertFloat64SliceEquals(expected, observed, t)
}

func TestLogspaceDescending(t *testing.T) {
	expected := []float64{1000, 464.158883, 215.443469, 100}
	observed := Logspace(3., 2., 4)

	for i := 0; i < len(observed); i++ {
		observed[i] = toFixed(observed[i], 6)
	}

	assertFloat64SliceEquals(expected, observed, t)
}
