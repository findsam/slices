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

func TestReduce(t *testing.T) {
	s := []int{1, 2, 3, 4}
	expected := 10

	result := Reduce(s, 0, func(acc int, i int) int {
		return acc + i
	})

	if result != expected {
		t.Errorf("Expected %d, got %d", expected, result)
	}
}

func TestFilter(t *testing.T) {
	s := []int{1, 2, 3, 4, 5}
	expected := []int{2, 4}

	got := Filter(s, func(n int) bool {
		return n%2 == 0
	})

	if len(got) != len(expected) {
		t.Fatalf("got %v, want %v", got, expected)
	}
}
