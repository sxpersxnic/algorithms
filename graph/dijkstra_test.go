package graph

import (
	"reflect"
	"testing"
)

func TestDijkstra(t *testing.T) {
	adj := map[int]map[int]int{
		0: {1: 4, 2: 1},
		1: {3: 1},
		2: {1: 2, 3: 5},
		3: {},
	}
	expected := map[int]int{
		0: 0,
		1: 3,
		2: 1,
		3: 4,
	}
	result := Dijkstra(adj, 0)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Dijkstra failed: expected %v, got %v", expected, result)
	}
}