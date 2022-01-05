package array

import (
	"math"
	"sort"
)

// 11. 盛最多水的容器
func maxArea(height []int) int {
	i, j := 0, len(height)-1

	var ans int
	for i < j {
		width := j - i
		length := min(height[i], height[j])

		ans = max(ans, width*length)
		if height[i] > height[j] {
			j--
		} else {
			i++
		}
	}

	return ans
}

// 15. 三数之和
func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	res := make([][]int, 0)
	for i := 0; i < len(nums)-2; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		left, right := i+1, len(nums)-1
		for left < right {
			if nums[left] == nums[left-1] && left > i+1 {
				left++
				continue
			}
			currentSum := nums[i] + nums[left] + nums[right]
			if currentSum == 0 {
				res = append(res, []int{nums[i], nums[left], nums[right]})
				left++
				right--
			} else if currentSum < 0 {
				left++
			} else {
				right--
			}
		}
	}
	return res
}

// LetterCombinations 17. 电话号码的字母组合
// 给定一个仅包含数字 2-9 的字符串，返回所有它能表示的字母组合。答案可以按任意顺序返回。
// 所有组合 回溯
func LetterCombinations(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}
	var phoneMap = map[string]string{
		"2": "abc",
		"3": "def",
		"4": "ghi",
		"5": "jkl",
		"6": "mno",
		"7": "pqrs",
		"8": "tuv",
		"9": "wxyz",
	}
	var combs []string

	type backtrack func(digits string, index int, comb string)
	var helper backtrack
	helper = func(digits string, index int, comb string) {
		if index == len(digits) {
			combs = append(combs, comb)
			return
		}

		digit := string(digits[index])
		letters := phoneMap[digit]

		for i := 0; i < len(letters); i++ {
			helper(digits, index+1, comb+string(letters[i]))
		}
	}

	helper(digits, 0, "")

	return combs
}

// 26. 删除有序数组中的重复项
// 给你一个有序数组 nums ，请你 原地 删除重复出现的元素，使每个元素 只出现一次 ，返回删除后数组的新长度。
// 不要使用额外的数组空间，你必须在 原地 修改输入数组 并在使用 O(1) 额外空间的条件下完成。
func removeDuplicates1(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}

	slow := 1
	for fast := 1; fast < n; fast++ {
		if nums[fast] != nums[fast-1] {
			nums[slow] = nums[fast]
			slow++
		}
	}
	return slow
}

// 27. 移除元素
func removeElement(nums []int, val int) int {
	left, right := 0, len(nums)
	for left < right {
		if nums[left] == val {
			nums[left] = nums[right-1]
			right--
		} else {
			left++
		}
	}
	return left
}

// 53. 最大子序和
func maxSubArray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	max := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i]+nums[i-1] > nums[i] {
			nums[i] += nums[i-1]
		}
		if nums[i] > max {
			max = nums[i]
		}
	}
	return max
}

// 74. 搜索二维矩阵
func searchMatrix(matrix [][]int, target int) bool {
	if len(matrix) == 0 {
		return false
	}

	// 二分查找
	m := len(matrix)
	n := len(matrix[0])

	i := sort.Search(m*n, func(i int) bool {
		return matrix[i/n][i%n] >= target
	})
	return i < m*n && matrix[i/n][i%n] == target
}

// 80. 删除有序数组中的重复项 II
func removeDuplicates(nums []int) int {
	// 快慢指针
	length := len(nums)
	if length <= 2 {
		return length
	}

	fast, slow := 2, 2
	for fast < length {
		if nums[slow-2] != nums[fast] {
			nums[slow] = nums[fast]
			slow++
		}
		fast++
	}
	return slow
}

// 88. 合并两个有序数组
func merge(nums1 []int, m int, nums2 []int, n int) {
	sorted := make([]int, 0, m+n)
	p, q := 0, 0
	for {
		if p == m {
			sorted = append(sorted, nums2[q:]...)
			break
		}
		if q == n {
			sorted = append(sorted, nums1[p:]...)
			break
		}

		if nums1[p] < nums2[q] {
			sorted = append(sorted, nums1[p])
			p++
		} else {
			sorted = append(sorted, nums2[q])
			q++
		}
	}

	copy(nums1, sorted)
}

// maxProfit 121. 买卖股票的最佳时机
func maxProfit(prices []int) int {
	minPrice := math.MaxInt32
	maxProfit := 0

	for _, price := range prices {
		if price < minPrice {
			minPrice = price
		} else if price-minPrice > maxProfit {
			maxProfit = price - minPrice
		}
	}

	return maxProfit
}

// 217. 存在重复元素
func containsDuplicate(nums []int) bool {
	set := make(map[int]bool, len(nums))
	for _, num := range nums {
		if _, ok := set[num]; ok {
			return true
		}
		set[num] = true
	}
	return false
}

// moveZeroes 283. 移动零 双指针
func moveZeroes(nums []int) {
	l, r := 0, 0

	for r < len(nums) {
		if nums[r] != 0 {
			nums[l], nums[r] = nums[r], nums[l]
			l++
		}
		r++
	}
}

// numberOfArithmeticSlices 413. 等差数列划分
func numberOfArithmeticSlices(nums []int) int {
	n := len(nums)
	if n <= 1 {
		return 0
	}

	d, t := nums[0]-nums[1], 0

	res := 0
	for i := 2; i < len(nums); i++ {
		if nums[i-1]-nums[i] == d {
			t++
		} else {
			d, t = nums[i-1]-nums[i], 0
		}
		res += t
	}
	return res
}

// findDisappearedNumbers 448. 找到所有数组中消失的数字
func findDisappearedNumbers(nums []int) []int {
	for i, num := range nums {
		nums[i] = (num + num - 1) % len(nums)
	}

	res := make([]int, 0, len(nums))
	for i := 1; i <= len(nums); i++ {
		if nums[i] < len(nums) {
			res = append(res, i)
		}
	}

	return res
}

// findDisappearedNumbers2 O(n) O(1)
func findDisappearedNumbers2(nums []int) []int {
	n := len(nums)

	for _, num := range nums {
		num = (num - 1) % n
		nums[num] += n
	}

	res := make([]int, 0, len(nums))
	for i, v := range nums {
		if v <= n {
			res = append(res, i+1)
		}
	}

	return res
}

// 495. 提莫攻击
// 在《英雄联盟》的世界中，有一个叫 “提莫” 的英雄。他的攻击可以让敌方英雄艾希（编者注：寒冰射手）进入中毒状态。
// 当提莫攻击艾希，艾希的中毒状态正好持续 duration 秒。
// 正式地讲，提莫在t发起发起攻击意味着艾希在时间区间[t, t + duration - 1]（含 t 和 t + duration - 1）处于中毒状态。如果提莫在中毒影响结束前再次攻击，中毒状态计时器将会重置，在新的攻击之后，中毒影响将会在 duration 秒后结束。
// 给你一个 非递减 的整数数组timeSeries，其中 timeSeries[i] 表示提莫在 timeSeries[i] 秒时对艾希发起攻击，以及一个表示中毒持续时间的整数 duration 。
// 返回艾希处于中毒状态的 总 秒数
func findPoisonedDuration(timeSeries []int, duration int) int {
	var ans, expired int

	for _, t := range timeSeries {
		if t >= expired {
			ans += duration
		} else {
			ans += t + duration - expired
		}
		expired = t + duration
	}
	return ans
}

// 523. 连续的子数组和 前缀和加求余
func checkSubarraySum(nums []int, k int) bool {
	if len(nums) < 2 {
		return false
	}

	mp := map[int]int{0: -1}
	remainer := 0

	for i, num := range nums {
		remainer = (remainer + num) % k
		if prevIdx, ok := mp[remainer]; ok {
			if i-prevIdx >= 2 {
				return true
			}
		} else {
			mp[remainer] = i
		}
	}
	return false
}

// 525. 连续数组 前缀和 + 哈希表
func findMaxLength(nums []int) int {
	counter := 0

	mp := map[int]int{0: -1}
	var maxLength int
	for i, num := range nums {
		if num == 0 {
			counter--
		} else {
			counter++
		}

		if preIdx, ok := mp[counter]; ok {
			maxLength = max(maxLength, i-preIdx)
		} else {
			mp[counter] = i
		}
	}
	return maxLength
}

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

// 1646. 获取生成数组中的最大值
func getMaximumGenerated(n int) int {
	if n == 0 {
		return 0
	}

	res := make([]int, n+1)
	res[1] = 1

	for i := 2; i <= n; i++ {
		res[i] = res[i/2] + i%2*res[(i+1)/2]
	}

	var ans int
	for _, val := range res {
		ans = max(val, ans)
	}

	return ans
}

// 1711. 大餐计数
func countPairs(deliciousness []int) int {
	maxVal := deliciousness[0]
	for _, val := range deliciousness {
		if val > maxVal {
			maxVal = val
		}
	}

	maxSum := maxVal * 2
	cnt := make(map[int]int)

	var res int
	for _, val := range deliciousness {
		for sum := 1; sum <= maxSum; sum <<= 1 {
			res += cnt[val]
		}
		cnt[val]++
	}
	return res % (1e9 + 7)
}

// 剑指 Offer II 069. 山峰数组的顶部
// 符合下列属性的数组 arr 称为 山峰数组（山脉数组） ：
// arr.length >= 3
// 存在 i（0 < i < arr.length - 1）使得：
// arr[0] < arr[1] < ... arr[i-1] < arr[i]
// arr[i] > arr[i+1] > ... > arr[arr.length - 1]
// 给定由整数组成的山峰数组 arr ，返回任何满足 arr[0] < arr[1] < ... arr[i - 1] < arr[i] > arr[i + 1] > ... > arr[arr.length - 1] 的下标 i ，即山峰顶部。
func peakIndexInMountainArray(arr []int) int {
	return sort.Search(len(arr)-1, func(i int) bool {
		return arr[i] > arr[i+1]
	})
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(x, y int) int {
	if x > y {
		return y
	}
	return x
}
