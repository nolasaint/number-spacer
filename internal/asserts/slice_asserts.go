package asserts

import (
	"reflect"
	"testing"
)

func Float64SliceEquals(expected []float64, observed []float64, t *testing.T) {
	if !reflect.DeepEqual(expected, observed) {
		t.Fatalf("\nexpected: %v\nbut got:  %v", expected, observed)
	}
}

func IntSliceEquals(expected []int, observed []int, t *testing.T) {
	if !reflect.DeepEqual(expected, observed) {
		t.Fatalf("\nexpected: %v\nbut got:  %v", expected, observed)
	}
}
