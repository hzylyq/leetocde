package search

// 33. 搜索旋转排序数组
func search(nums []int, target int) int {
	length := len(nums)
	if length == 0 {
		return -1
	}

	var (
		l = 0
		r = len(nums) - 1
	)

	for l <= r {
		mid := l + (r-l)/2
		if nums[mid] == target {
			return mid
		}

		if nums[0] <= nums[mid] {
			if nums[0] <= target && target <= nums[mid] {
				r = mid - 1
			} else {
				l = mid + 1
			}
		} else {
			if nums[mid] <= target && target <= nums[length-1] {
				l = mid + 1
			} else {
				r = mid - 1
			}
		}
	}

	return -1
}
