package base

import (
	"math"
	"sort"
)

// O(n^2) time
// O(1) space
func TwoNumberSum(array []int, target int) []int {
	// Write your code here.
	for i, val1 := range array {
		for j := i + 1; j < len(array); j++ {
			val2 := array[j]
			if val1+val2 == target {
				return []int{val1, val2}
			}
		}
	}
	return nil
}

func TwoNumberSum2(array []int, target int) []int {
	numMap := make(map[int]bool)
	// Write your code here.
	for _, val1 := range array {
		val2 := target - val1
		if _, ok := numMap[val2]; ok {
			return []int{val1, val2}
		}
		numMap[val1] = true
	}
	return nil
}

func TwoNumberSum3(array []int, target int) []int {
	sort.Ints(array)
	left, right := 0, len(array)-1
	for left < right {
		sum := array[left] + array[right]
		if sum == target {
			return []int{array[left], array[right]}
		} else if sum < target {
			left += 1
		} else {
			right -= 1
		}
	}
	return []int{}
}

func IsValidSubsequence(array []int, sequence []int) bool {
	// Write your code here.
	var seqIndex, arrIndex int
	for arrIndex < len(array) && seqIndex < len(sequence) {
		if array[arrIndex] == sequence[seqIndex] {
			seqIndex++
			arrIndex++
		} else {
			arrIndex++
		}
	}

	return seqIndex == len(sequence)
}

func IsValidSubsequence2(array []int, sequence []int) bool {
	// Write your code here.
	var seqIndex int
	for _, val := range array {
		if seqIndex == len(sequence) {
			break
		}
		if val == sequence[seqIndex] {
			seqIndex++
		}
	}

	return seqIndex == len(sequence)
}

func BinarySearch(array []int, target int) int {
	// Write your code here.
	low, high := 0, len(array)-1
	for low < high {
		mid := low + (high-low)/2
		numMid := array[mid]
		if target == numMid {
			return mid
		}

		if target < numMid {
			high = mid
		} else {
			low = mid
		}
	}

	return -1
}

func NonConstructibleChange(coins []int) int {
	// Write your code here.
	sort.Ints(coins)
	currentChangeCreated := 0
	for _, coin := range coins {
		if coin > currentChangeCreated+1 {
			return currentChangeCreated + 1
		}
		currentChangeCreated += coin
	}

	return currentChangeCreated + 1
}

func GetNthFib(n int) int {
	// Write your code here.
	if n < 0 {
		return -1
	}

	if n == 1 {
		return 0
	}
	if n == 2 {
		return 1
	}

	return GetNthFib(n-1) + GetNthFib(n-2)
}

func GetNthFib2(n int) int {
	// Write your code here.
	memoize := map[int]int{
		1: 0,
		2: 1,
	}

	return helper(n, memoize)
}

func helper(n int, memoize map[int]int) int {
	if val, ok := memoize[n]; ok {
		return val
	}
	memoize[n] = helper(n-1, memoize) + helper(n-2, memoize)
	return memoize[n]
}

func GetNthFib3(n int) int {
	// Write your code here.
	lastTwo := []int{0, 1}
	counter := 3
	for counter <= n {
		nextFib := lastTwo[0] + lastTwo[1]
		lastTwo = []int{lastTwo[1], nextFib}
		counter++
	}
	if n > 1 {
		return lastTwo[1]
	}
	return lastTwo[0]
}

func FindThreeLargestNumbers(array []int) []int {
	// Write your code here.
	three := []int{math.MinInt32, math.MinInt32, math.MinInt32}
	for _, num := range array {
		updateLargest(three, num)
	}
	return three
}

func updateLargest(three []int, num int) {
	if num > three[2] {
		shiftAndUpdate(three, num, 2)
	} else if num > three[1] {
		shiftAndUpdate(three, num, 1)
	} else if num > three[0] {
		shiftAndUpdate(three, num, 0)
	}
}

func shiftAndUpdate(array []int, num int, idx int) {
	for i := 0; i < idx+1; i++ {
		if i == idx {
			array[i] = num
		} else {
			array[i] = array[i+1]
		}
	}
}

func IsPalindrome(str string) bool {
	// Write your code here.
	for i := 0; i <= len(str)/2; i++ {
		if str[i] == str[len(str)-1-i] {
			continue
		}
		return false
	}
	return true
}

func SortedSquaredArray(array []int) []int {
	sort.Ints(array)
	res := make([]int, 0)
	for _, val := range array {
		res = append(res, val*val)
	}
	return res
}

func SortedSquaredArray2(array []int) []int {
	res := make([]int, 0)

	smallValIdx := 0
	largerValIdx := len(array) - 1

	for idx := len(array) - 1; idx >= 0; idx-- {
		smallVal := array[smallValIdx]
		largeVal := array[largerValIdx]

		if abs(smallVal) > abs(largeVal) {
			res[idx] = smallVal * smallVal
			smallValIdx++
		} else {
			res[idx] = largeVal * largeVal
			largerValIdx--
		}
	}

	return res
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

const HOME_TEAM_WON = 1

func TournamentWinner(competitions [][]string, results []int) string {
	var currentBestTeam string
	points := make(map[string]int, 0)
	for idx, competition := range competitions {
		result := results[idx]
		homeTeam, awayTeam := competition[0], competition[1]
		winner := awayTeam
		if result == HOME_TEAM_WON {
			winner = homeTeam
		}
		points[winner] += 3
		if points[winner] > points[currentBestTeam] {
			currentBestTeam = winner
		}

	}

	return currentBestTeam
}

// 4. 寻找两个正序数组的中位数
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	totalLength := len(nums1) + len(nums2)
	if totalLength%2 == 1 {
		midIdx := totalLength / 2
		return float64(getKthElement(nums1, nums2, midIdx+1))
	} else {
		midIdx1, midIdx2 := totalLength/2-1, totalLength/2
		return float64(getKthElement(nums1, nums2, midIdx1+1) + getKthElement(nums1, nums2, midIdx2+1)/2)
	}
}

func getKthElement(nums1, nums2 []int, k int) int {
	return 0
}

// 7. 整数反转
func reverse(x int) int {
	var rev int
	for x != 0 {
		pop := x % 10
		x /= 10
		if rev > math.MaxInt32 || rev < math.MinInt32 {
			return 0
		}
		rev = rev*10 + pop
	}
	return rev
}

// 231. 2的幂
func isPowerOfTwo(n int) bool {
	if n == 0 {
		return false
	}

	return n&(n-1) == 0
}

func IsPowerOfFour(n int) bool {
	return n > 0 && ((n & (n - 1)) > 0) && n%3 == 1
}

// hammingDistance 461. 汉明距离
func hammingDistance(x int, y int) int {
	var res int
	for s := x ^ y; s > 0; s >>= 1 {
		res += s & 1
	}

	return res
}

// 1442. 形成两个异或相等数组的三元组数目
func countTriplets(arr []int) int {

	return 0
}

// 1720. 解码异或后的数组
func decode(encoded []int, first int) []int {
	res := make([]int, len(encoded)+1)
	res[0] = first
	for i, e := range encoded {
		res[i+1] = res[i] ^ e
	}
	return res
}
