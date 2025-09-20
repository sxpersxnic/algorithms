package sorting

// QuickSort sorts a slice of integers using the Quick Sort algorithm.
// It selects a pivot and partitions the array around it recursively.
//
// Time Complexity: O(n log n) on average, O(n^2) in the worst case
// Space Complexity: O(log n) due to recursion
func QuickSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	pivot := arr[len(arr)/2]
	left := []int{}
	right := []int{}
	mid := []int{}

	for _, val := range arr {
		if val < pivot {
			left = append(left, val)
		} else if val > pivot {
			right = append(right, val)
		} else {
			mid = append(mid, val)
		}
	}

	left = QuickSort(left)
	right = QuickSort(right)

	return append(append(left, mid...), right...)
}