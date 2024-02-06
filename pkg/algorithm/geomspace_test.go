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

//TODO:finish writing tests
