package array

import (
	"fmt"
	"math"
	"sort"
	"strconv"
	"strings"
	"unicode"
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

func threeSum2(nums []int) [][]int {
	target := 0

	sort.Ints(nums)
	n := len(nums)

	res := make([][]int, 0)

	for first := 0; first < n; first++ {
		if first > 0 && nums[first] == nums[first-1] {
			continue
		}

		third := n - 1
		for second := first + 1; second < n; second++ {
			if second > first+1 && nums[second] == nums[second-1] {
				continue
			}

			for second < third && nums[first]+nums[second]+nums[third] > target {
				third--
			}

			if second == third {
				break
			}

			if nums[first]+nums[second]+nums[third] == target {
				res = append(res, []int{nums[first], nums[second], nums[third]})
			}
		}
	}

	return res
}

// 26. 删除有序数组中的重复项
// 给你一个有序数组 nums ，请你 原地 删除重复出现的元素，使每个元素 只出现一次 ，返回删除后数组的新长度。
// 不要使用额外的数组空间，你必须在 原地 修改输入数组 并在使用 O(1) 额外空间的条件下完成。
func removeDuplicates1(nums []int) int {
	var slow int
	for i := 0; i < len(nums); i++ {
		if nums[i] != nums[slow] {
			slow++
			nums[slow] = nums[i]
		}
	}
	return slow + 1
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

// 31. 下一个排列
func nextPermutation(nums []int) {
	n := len(nums)
	i := n - 2
	for i >= 0 && nums[i] >= nums[i+1] {
		i--
	}

	for i >= 0 {
		j := n - 1
		for j >= 0 && nums[i] >= nums[j] {
			j--
		}
		nums[i], nums[j] = nums[j], nums[i]
	}
	reverse(nums[i+1:])
}

// 56. 合并区间
func mergeSpace(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	// for _, item := range intervals {

	// }

	return intervals

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

func removeDuplicateN(nums []int, k int) int {
	var u int
	for _, v := range nums {
		if u < k || v != nums[u-k] {
			nums[u] = v
			u++
		}
	}
	return u
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

// 125. 验证回文串
func isPalindrome(s string) bool {
	var sgood string
	for i := 0; i < len(s); i++ {
		if isalnum(s[i]) {
			sgood += string(s[i])
		}
	}

	sgood = strings.ToLower(sgood)

	for n := 0; n < len(sgood)/2; n++ {
		if sgood[n] != sgood[len(sgood)-1-n] {
			return false
		}
	}

	return true
}

func isalnum(ch byte) bool {
	return (ch >= '0' && ch <= '9') || (ch >= 'a' && ch <= 'z') || (ch >= 'A' && ch <= 'Z')
}

// 134. 加油站
func canCompleteCircuit(gas []int, cost []int) int {
	// for i, n := 0, len(gas); i < n; {
	// 	sumOfGas, sumOfCost, cnt := 0, 0, 0
	// 	for cnt < n {

	// 	}
	// }
	return 0
}

// 167. 两数之和 II - 输入有序数组
// 给定一个已按照 非递减顺序排列 的整数数组 numbers ，请你从数组中找出两个数满足相加之和等于目标数 target 。
// 函数应该以长度为 2 的整数数组的形式返回这两个数的下标值。numbers 的下标 从 1 开始计数 ，所以答案数组应当满足 1 <= answer[0] < answer[1] <= numbers.length 。
// 你可以假设每个输入 只对应唯一的答案 ，而且你 不可以 重复使用相同的元素。
func twoSum(numbers []int, target int) []int {
	i, j := 0, len(numbers)-1
	var res []int
	for i < j {
		if numbers[i]+numbers[j] > target {
			j--
		} else if numbers[i]+numbers[j] == target {
			res = []int{i + 1, j + 1}
			break
		} else {
			i++
		}
	}

	return res
}

// 169. 多数元素
func majorityElement(nums []int) int {
	var (
		count     int
		candidate int
	)

	for _, v := range nums {
		if count == 0 {
			candidate = v
		}
		if candidate == v {
			count += 1
		} else {
			count -= 1
		}
	}

	return candidate
}

// 189. 轮转数组
func reverse(a []int) {
	for i, n := 0, len(a); i < n/2; i++ {
		a[i], a[n-1-i] = a[n-1-i], a[i]
	}
}

func rotate(nums []int, k int) {
	k %= len(nums)
	reverse(nums)
	reverse(nums[:k])
	reverse(nums[k:])
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

// 228. 汇总区间
// 给定一个无重复元素的有序整数数组 nums 。
// 返回 恰好覆盖数组中所有数字 的 最小有序 区间范围列表。也就是说，nums 的每个元素都恰好被某个区间范围所覆盖，并且不存在属于某个范围但不属于 nums 的数字 x 。
// 列表中的每个区间范围 [a,b] 应该按如下格式输出：
// "a->b" ，如果 a != b
// "a" ，如果 a == b
func summaryRanges(nums []int) []string {
	ans := make([]string, 0, len(nums))
	for i := 0; i < len(nums); {
		left := i
		for i++; i < len(nums) && nums[i]-nums[i-1] == 1; i++ {
		}
		s := strconv.Itoa(nums[left])
		if left < i-1 {
			s += "->" + strconv.Itoa(nums[i-1])
		}
		ans = append(ans, s)
	}
	return ans
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

// 392. 判断子序列
func isSubsequence(s string, t string) bool {
	n, m := len(s), len(t)

	i, j := 0, 0
	for i < n && j < m {
		if s[i] == t[j] {
			i++
		}
		j++
	}

	return i == n
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

// 724. 寻找数组的中心下标
func pivotIndex(nums []int) int {
	var total int

	for _, num := range nums {
		total += num
	}

	var helf int
	for i := 0; i < len(nums); i++ {
		if helf*2 == total-nums[i] {
			return i
		}
		helf += nums[i]
	}

	return -1
}

// 844. 比较含退格的字符串
func backspaceCompare(s string, t string) bool {
	skipS, skipT := 0, 0

	i, j := len(s)-1, len(t)-1
	for i >= 0 || j >= 0 {
		for i >= 0 {
			if s[i] == '#' {
				skipS++
				i--
			} else if skipS > 0 {
				i--
				skipS--
			} else {
				break
			}
		}

		for j >= 0 {
			if t[j] == '#' {
				skipT++
				j--
			} else if skipT > 0 {
				j--
				skipT--
			} else {
				break
			}
		}

		if i >= 0 && j >= 0 {
			if s[i] != t[j] {
				return false
			}
		} else if i >= 0 || j >= 0 {
			return false
		}

		i--
		j--
	}

	return true
}

// 917. 仅仅反转字母
// 给你一个字符串 s ，根据下述规则反转字符串：

// 所有非英文字母保留在原有位置。
// 所有英文字母（小写或大写）位置反转。
// 返回反转后的 s 。
func reverseOnlyLetters(s string) string {
	left, right := 0, len(s)-1
	ans := []byte(s)
	for left < right {
		if !unicode.IsLetter(rune(s[left])) {
			left++
			continue
		}
		if !unicode.IsLetter(rune(s[right])) {
			right--
			continue
		}

		ans[left], ans[right] = ans[right], ans[left]
		left++
		right--
	}

	return string(ans)
}

// 977. 有序数组的平方
func sortedSquares(nums []int) []int {
	res := make([]int, len(nums))
	i, j := 0, len(nums)-1
	for pos := len(nums) - 1; pos >= 0; pos-- {
		v, w := nums[i]*nums[i], nums[j]*nums[j]
		if v > w {
			res[pos] = v
			i++
		} else {
			res[pos] = w
			j--
		}
	}

	return res
}

// 986. 区间列表的交集
func intervalIntersection(firstList [][]int, secondList [][]int) [][]int {
	ans := make([][]int, 0)
	var i, j int

	for i < len(firstList) && j < len(secondList) {
		lo := max(firstList[i][0], secondList[j][0])
		hi := min(firstList[i][1], secondList[j][1])

		if lo < hi {
			ans = append(ans, []int{lo, hi})
		}

		if firstList[i][1] < secondList[j][1] {
			i++
		} else {
			j++
		}
	}

	return ans
}

// 1051. 高度检查器
func heightChecker(heights []int) int {
	var ans int

	sorted := append([]int{}, heights...)

	sort.Ints(sorted)

	for i, v := range heights {
		if v != sorted[i] {
			ans++
		}
	}

	return ans
}

// 计数排序
func heightChecker2(heights []int) int {
	// var ans int
	//
	// sorted := append([]int{}, heights...)
	//
	// sort.Ints(sorted)
	//
	// for i, v := range heights {
	//	if v != sorted[i] {
	//		ans++
	//	}
	// }
	//
	// return ans
	return 0
}

// 1480. Running Sum of 1d Array
func runningSum(nums []int) []int {
	for i := 1; i < len(nums); i++ {
		nums[i] += nums[i-1]
	}

	return nums
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

// 1706. 球会落何处
func findBall(grid [][]int) []int {
	n := len(grid[0])
	ans := make([]int, n)

	for j := range ans {
		col := j
		for _, row := range grid {
			dir := row[col]
			col += dir
			if col < 0 || col == n || row[col] != dir {
				col = -1
				break
			}
		}

		ans[j] = col
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

// 1331. 数组序号转换
func arrayRankTransform(arr []int) []int {
	a := append([]int{}, arr...)

	sort.Ints(a)

	res := make([]int, 0, len(arr))
	rankMap := make(map[int]int)
	for _, item := range a {
		if _, ok := rankMap[item]; !ok {
			rankMap[item] = len(rankMap) + 1
		}
	}

	for _, item := range arr {
		res = append(res, rankMap[item])
	}

	return res
}

// 1470. 重新排列数组
func shuffle(nums []int, n int) []int {
	res := make([]int, n*2)
	for i, num := range nums[:n] {
		res[i*2] = num
		res[i*2+1] = nums[n+i]
	}
	return res
}

// 2042. 检查句子中的数字是否递增
func areNumbersAscending(s string) bool {
	var pre, i int

	for i < len(s) {
		if unicode.IsDigit(rune(s[i])) {
			cur := 0
			for i < len(s) && unicode.IsDigit(rune(s[i])) {
				cur = cur*10 + int(s[i]-'0')
				i++
			}

			if pre > cur {
				return false
			}
			pre = cur

		} else {
			i++
		}
	}

	return true
}

func numberOfPairs(nums []int) []int {
	var res int

	cnt := map[int]bool{}
	for _, num := range nums {
		cnt[num] = !cnt[num]
		if !cnt[num] {
			res++
		}
	}

	return []int{res, len(nums) - res*2}
}

// 2352. 相等行列对
func equalPairs(grid [][]int) int {
	n := len(grid)
	cntMap := make(map[string]int, n)
	for _, row := range grid {
		cntMap[fmt.Sprint(row)]++
	}

	var res int
	for j := 0; j < n; j++ {
		var arr []int
		for i := 0; i < n; i++ {
			arr = append(arr, grid[i][j])
		}
		if val, ok := cntMap[fmt.Sprint(arr)]; ok {
			res += val
		}
	}

	return res
}

// 2357. 使数组中所有元素都等于零
func minimumOperations(nums []int) int {
	var res int

	boolMap := make(map[int]bool, len(nums))
	for _, num := range nums {
		if num != 0 && !boolMap[num] {
			res++
		}
		boolMap[num] = true
	}

	return res
}

// 2383. 赢得比赛需要的最少训练时长
func minNumberOfHours(initialEnergy int, initialExperience int, energy []int, experience []int) int {
	sum := 0
	for _, v := range energy {
		sum += v
	}

	var res int
	if initialEnergy <= sum {
		res = sum + 1 - initialEnergy
	}

	for _, v := range experience {
		if initialExperience <= v {
			res += (v + 1) - initialExperience
			initialExperience += (v + 1) - initialExperience + v
		} else {
			initialExperience += v
		}
	}

	return res
}

// 2460. 对数组执行操作
func applyOperations(nums []int) []int {
	var j int
	for i := 0; i < len(nums); i++ {
		if i+1 < len(nums) && nums[i] == nums[i+1] {
			nums[i] *= 2
			nums[i+1] = 0
		}
		if nums[i] != 0 {
			nums[i], nums[j] = nums[j], nums[i]
			j++
		}
	}

	return nums
}

// 2475. 数组中不等三元组的数目
func unequalTriplets(nums []int) int {
	countMap := make(map[int]int, len(nums))
	for _, num := range nums {
		countMap[num]++
	}

	res, n, t := 0, len(nums), 0
	for _, v := range countMap {
		res, t = res+t*v*(n-t-v), t+v
	}
	return res
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
