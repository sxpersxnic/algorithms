package dynamic_programming

// Fibonacci computes the n-th Fibonacci.
// Time Complexity: O(n)
// Space Complexity: O(n), can be optimized to O(1) if only two variables are kept.
func Fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	return Fibonacci(n-1) + Fibonacci(n-2)
}
