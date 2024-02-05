package algorithm

import (
	"reflect"
	"testing"
)

//
// helpers
//

func assertIntSliceEquals(expected []int, observed []int, t *testing.T) {
	if !reflect.DeepEqual(expected, observed) {
		t.Fatalf("\nexpected: %v\nbut got:  %v", expected, observed)
	}
}

//
// tests
//

func TestLinspaceNoElements(t *testing.T) {
	expected := []int{}
	observed := Linspace(0, 100, 0)

	assertIntSliceEquals(expected, observed, t)
}

func TestLinspaceEvenNumElements(t *testing.T) {
	expected := []int{10, 20, 30, 40}
	observed := Linspace(10, 40, 4)

	assertIntSliceEquals(expected, observed, t)
}

func TestLinspaceOddNumElements(t *testing.T) {
	expected := []int{100, 200, 300, 400, 500}
	observed := Linspace(100, 500, 5)

	assertIntSliceEquals(expected, observed, t)
}

func TestLinspaceRounding(t *testing.T) {
	expected := []int{500, 666, 833, 1000}
	observed := Linspace(500, 1000, 4)

	assertIntSliceEquals(expected, observed, t)
}

func testLinspaceTooFewElementsForRange(t *testing.T) {
	expected := []int{1, 1, 1, 2}
	observed := Linspace(1, 2, 4)

	assertIntSliceEquals(expected, observed, t)
}

func TestLinspaceMinAndMaxEqual(t *testing.T) {
	expected := []int{420, 420, 420}
	observed := Linspace(420, 420, 3)

	assertIntSliceEquals(expected, observed, t)
}

func TestLinspaceMinGreaterThanMax(t *testing.T) {
	expected := []int{40, 30, 20, 10}
	observed := Linspace(40, 10, 4)

	assertIntSliceEquals(expected, observed, t)
}
