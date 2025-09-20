package dynamic_programming

func Knapsack(weights []int, values []int, capacity int) int {
	n := len(weights)
	if n == 0 || capacity == 0 {
		return 0
	}

	dp := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([]int, capacity+1)
	}

	for i := 1; i <= n; i++ {
		for w := 1; w <= capacity; w++ {
			if weights[i-1] <= w {
				dp[i][w] = max(values[i-1]+dp[i-1][w-weights[i-1]], dp[i-1][w])
			} else {
				dp[i][w] = dp[i-1][w]
			}
		}
	}
	return dp[n][capacity]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}