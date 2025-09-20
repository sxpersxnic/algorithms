package searching

import "testing"

func TestLinearSearch(t *testing.T) {
	tests := []struct {
		arr			[]int
		target		int
		expected	int
	}{
		{[]int{1, 2, 3, 4, 5}, 3, 2},
		{[]int{10, 20, 30, 40, 50}, 25, -1},
		{[]int{}, 1, -1},
		{[]int{5}, 5, 0},
		{[]int{5}, 10, -1},
	}

	for _, test := range tests {
		result := LinearSearch(test.arr, test.target)
		if result != test.expected {
			t.Errorf("LinearSearch(%v, %d) = %d; want %d", test.arr, test.target, result, test.expected)
		}
	}
}