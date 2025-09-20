package sorting

import "testing"

func TestMergeSort(t *testing.T) {
	tests := []struct {
		input    []int
		expected []int
	}{
		{[]int{38, 27, 43, 3, 9, 82, 10}, []int{3, 9, 10, 27, 38, 43, 82}},
		{[]int{1}, []int{1}},
		{[]int{}, []int{}},
	}

	for _, test := range tests {
		result := MergeSort(test.input)
		for i := range result {
			if result[i] != test.expected[i] {
				t.Errorf("MergeSort(%v) = %v, expected %v", test.input, result, test.expected)
			}
		}
	}
}
