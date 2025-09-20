package sorting

// BubbleSort sorts a slice of integers using the Bubble Sort algorithm.
// It repeatedly swaps adjacent elements if they are in the wrong order.
//
// Time Complexity: O(n^2)
// Space Complexity: O(1)
func BubbleSort(arr []int) []int {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	return arr
}