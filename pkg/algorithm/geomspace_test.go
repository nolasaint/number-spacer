package algorithm

import (
	"internal/asserts"
	"testing"
)

//
// tests
//

func TestGeomspaceNoElements(t *testing.T) {
	expected := []float64{}
	observed := Geomspace(0, 100, 0)

	asserts.Float64SliceEquals(expected, observed, t)
}

func TestGeomspaceAscending(t *testing.T) {
	expected := []float64{10, 17.78279, 31.62278, 56.23413, 100}
	observed := Geomspace(10, 100, 5)

	for i := 0; i < len(observed); i++ {
		observed[i] = toFixed(observed[i], 5)
	}

	asserts.Float64SliceEquals(expected, observed, t)
}

func TestGeomspaceDescending(t *testing.T) {
	expected := []float64{100, 56.23413, 31.62278, 17.78279, 10}
	observed := Geomspace(100, 10, 5)

	for i := 0; i < len(observed); i++ {
		observed[i] = toFixed(observed[i], 5)
	}

	asserts.Float64SliceEquals(expected, observed, t)
}
