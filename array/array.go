package array

// 1480. Running Sum of 1d Array
func runningSum(nums []int) []int {
	res := make([]int, len(nums))
	var total int
	for k, v := range nums {
		res[k] = total + v
		total += v
	}
	return res
}
