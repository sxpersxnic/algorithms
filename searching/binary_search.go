package searching

// BinarySearch searches for a target in a sorted integer slice using the binary search algorithm.
// Returns the index of the target if found, otherwise -1.
//
// Time Complexity: O(log n)
// Space Complexity: O(1)
func BinarySearch(arr []int, target int) int {
	low, high := 0, len(arr)-1
	for low <= high {
		mid := low + (high-low)/2
		if arr[mid] == target {
			return mid
		} else if arr[mid] < target {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return -1
}