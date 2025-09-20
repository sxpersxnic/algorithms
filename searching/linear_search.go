package searching

// LinearSearch searches for a target in an integer slice using a sequential scan.
// Returns the index of the target if found, otherwise -1.
//
// Time Complexity: O(n)
// Space Complexity: O(1)
func LinearSearch(arr []int, target int) int {
	for i, val := range arr {
		if val == target {
			return i
		}
	}
	return -1
}