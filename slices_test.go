package slices

import (
	"reflect"
	"testing"
)
func TestMap(t *testing.T) {
	s := []int{1, 2, 3}
	expected := []int{2, 4, 6}

	result := Map(s, func(i int) int {
		return i * 2
	})

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}