package base

import (
	"math"
	"sort"
	"strconv"
)

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

// 78. 子集
// 给你一个整数数组 nums ，数组中的元素 互不相同。返回该数组所有可能的子集（幂集）。
// 解集不能包含重复的子集。你可以按任意顺序返回解集。
func subsets(nums []int) [][]int {
	n := len(nums)
	ans := make([][]int, 0)
	for i := 0; i < 1<<n; i++ {
		res := make([]int, 0)
		for k, val := range nums {
			if i>>k&1 > 0 {
				res = append(res, val)
			}
		}
		ans = append(ans, append([]int(nil), res...))
	}
	return ans
}

// 231. 2的幂
func isPowerOfTwo(n int) bool {
	if n == 0 {
		return false
	}

	return n&(n-1) == 0
}

// countBits 338. 比特位计数
func countBits(n int) []int {
	var res []int
	for i := 0; i <= n; i++ {
		res = append(res, onesCount(i))
	}
	return res
}

func onesCount(x int) (ones int) {
	for ; x > 0; x &= x - 1 {
		ones++
	}
	return
}

func IsPowerOfFour(n int) bool {
	return n > 0 && ((n & (n - 1)) > 0) && n%3 == 1
}

// 412. Fizz Buzz
// 给你一个整数 n ，找出从 1 到 n 各个整数的 Fizz Buzz 表示，并用字符串数组 answer（下标从 1 开始）返回结果，其中：
// answer[i] == "FizzBuzz" 如果 i 同时是 3 和 5 的倍数。
// answer[i] == "Fizz" 如果 i 是 3 的倍数。
// answer[i] == "Buzz" 如果 i 是 5 的倍数。
// answer[i] == i 如果上述条件全不满足。
func fizzBuzz(n int) []string {
	var res []string
	for i := 1; i <= n; i++ {
		if i%3 == 0 && i%5 == 0 {
			res = append(res, "FizzBuzz")
		} else if i%3 == 0 {
			res = append(res, "Fizz")
		} else if i%5== 0 {
			res = append(res, "Buzz")
		} else {
			res = append(res, strconv.Itoa(i))
		}
	}
	return res
}

// hammingDistance 461. 汉明距离
func hammingDistance(x int, y int) int {
	var res int
	for s := x ^ y; s > 0; s >>= 1 {
		res += s & 1
	}

	return res
}

// tribonacci 1137. 第 N 个泰波那契数
func tribonacci(n int) int {
	if n == 0 {
		return 0
	}

	if n <= 2 {
		return 1
	}

	first, second, third, s := 0, 0, 1, 1
	for i := 3; i <= n; i++ {
		first = second
		second = third
		third = s
		s = first + second + third
	}
	return s
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
