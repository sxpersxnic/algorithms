package sorting

import "testing"

func TestQuickSort(t *testing.T) {
	tests := []struct {
		input    []int
		expected []int
	}{
		{[]int{10, 7, 8, 9, 1, 5}, []int{1, 5, 7, 8, 9, 10}},
		{[]int{3, 3, 2, 1}, []int{1, 2, 3, 3}},
		{[]int{}, []int{}},
	}

	for _, test := range tests {
		result := QuickSort(test.input)
		for i := range result {
			if result[i] != test.expected[i] {
				t.Errorf("QuickSort(%v) = %v, expected %v", test.input, result, test.expected)
			}
		}
	}
}
