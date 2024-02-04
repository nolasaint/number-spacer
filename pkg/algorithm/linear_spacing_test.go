package algorithm

import (
	"reflect"
	"testing"
)

//
// helpers
//

func fail(expected []int, observed []int, t *testing.T) {
	t.Fatalf("\nexpected: %v\nbut got:  %v", expected, observed)
}

//
// tests
//

func TestLinearSpacingNoElements(t *testing.T) {
	expected := []int{}
	observed := LinearSpacing(0, 0, 100)

	if !reflect.DeepEqual(expected, observed) {
		fail(expected, observed, t)
	}
}

func TestLinearSpacingEvenNumElements(t *testing.T) {
	expected := []int{10, 20, 30, 40}
	observed := LinearSpacing(4, 10, 40)

	if !reflect.DeepEqual(expected, observed) {
		fail(expected, observed, t)
	}
}

func TestLinearSpacingOddNumElements(t *testing.T) {
	expected := []int{100, 200, 300, 400, 500}
	observed := LinearSpacing(5, 100, 500)

	if !reflect.DeepEqual(expected, observed) {
		fail(expected, observed, t)
	}
}

func TestLinearSpacingRounding(t *testing.T) {
	expected := []int{500, 666, 833, 1000}
	observed := LinearSpacing(4, 500, 1000)

	if !reflect.DeepEqual(expected, observed) {
		fail(expected, observed, t)
	}
}

func testLinearSpacingTooFewElementsForRange(t *testing.T) {
	expected := []int{1, 1, 1, 2}
	observed := LinearSpacing(4, 1, 2)

	if !reflect.DeepEqual(expected, observed) {
		fail(expected, observed, t)
	}
}

func TestLinearSpacingMinAndMaxEqual(t *testing.T) {
	expected := []int{420, 420, 420}
	observed := LinearSpacing(3, 420, 420)

	if !reflect.DeepEqual(expected, observed) {
		fail(expected, observed, t)
	}
}

func TestLinearSpacingMinGreaterThanMax(t *testing.T) {
	expected := []int{40, 30, 20, 10}
	observed := LinearSpacing(4, 40, 10)

	if !reflect.DeepEqual(expected, observed) {
		fail(expected, observed, t)
	}
}
