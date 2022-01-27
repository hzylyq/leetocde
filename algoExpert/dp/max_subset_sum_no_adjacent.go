package dp

// MaxSubsetSumNoAdjacent O(n) / O(n)
func MaxSubsetSumNoAdjacent(array []int) int {
	if len(array) == 0 {
		return 0
	}

	if len(array) == 1 {
		return array[0]
	}

	if len(array) == 2 {
		return max(array[0], array[1])
	}
	dp := make([]int, len(array))
	dp[0] = array[0]
	dp[1] = max(array[0], array[1])

	for i := 2; i < len(array); i++ {
		dp[i] = max(array[i]+dp[i-2], dp[i-1])
	}

	// Write your code here.
	return dp[len(array)]
}

// MaxSubsetSumNoAdjacent1 O(n) / O(1)
func MaxSubsetSumNoAdjacent1(array []int) int {
	if len(array) == 0 {
		return 0
	}

	if len(array) == 1 {
		return array[0]
	}

	if len(array) == 2 {
		return max(array[0], array[1])
	}

	first, second := array[0], max(array[0], array[1])

	var res int
	dp := make([]int, len(array))
	dp[0] = array[0]
	dp[1] = max(array[0], array[1])

	for i := 2; i < len(array); i++ {
		res = max(first + array[i], second)
		first = second
		second  = res
	}

	// Write your code here.
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
