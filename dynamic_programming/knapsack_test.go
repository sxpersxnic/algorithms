package dynamic_programming

import "testing"

func TestKnapsack(t *testing.T) {
	tests := []struct {
		weights  []int
		values   []int
		capacity int
		expected int
	}{
		{[]int{1, 2, 3}, []int{10, 20, 30}, 5, 50},
		{[]int{2, 3, 4, 5}, []int{3, 4, 5, 6}, 5, 7},
		{[]int{3, 4, 5}, []int{30, 50, 60}, 8, 90},
	}

	for _, test := range tests {
		result := Knapsack(test.weights, test.values, test.capacity)
		if result != test.expected {
			t.Errorf("Knapsack(%v, %v, %d) = %d, expected %d",
				test.weights, test.values, test.capacity, result, test.expected)
		}
	}
}
