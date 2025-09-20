package graph

import (
	"testing"
	"reflect"
)

func TestBFS(t *testing.T) {
	adj := map[int][]int{
		0: {1, 2},
		1: {0, 3, 4},
		2: {0, 4},
		3: {1, 5},
		4: {1, 2, 5},
		5: {3, 4},
	}
	expected := []int{0, 1, 2, 3, 4, 5}
	result := BFS(adj, 0)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("BFS failed: expected %v, got %v", expected, result)
	}
}