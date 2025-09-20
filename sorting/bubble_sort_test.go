package sorting

import "testing"

func TestBubbleSort(t *testing.T) {
	tests := []struct {
		input    []int
		expected []int
	}{
		{[]int{5, 1, 4, 2, 8}, []int{1, 2, 4, 5, 8}},
		{[]int{3, 2, 1}, []int{1, 2, 3}},
		{[]int{}, []int{}},
	}

	for _, test := range tests {
		result := BubbleSort(append([]int{}, test.input...))
		for i := range result {
			if result[i] != test.expected[i] {
				t.Errorf("BubbleSort(%v) = %v, expected %v", test.input, result, test.expected)
			}
		}
	}
}
