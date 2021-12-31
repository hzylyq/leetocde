package math

import (
	"math"
)

func BitwiseComplement(N int) int {
	if N == 0 {
		return 1
	}

	var arr []int
	for N != 0 {
		arr = append(arr, N%2)
		N /= 2
	}
	var resArr []int
	for _, val := range arr {
		if val == 0 {
			resArr = append(resArr, 1)
		} else {
			resArr = append(resArr, 0)
		}
	}

	var res int
	for idx, val := range resArr {
		if val == 1 {
			res += int(math.Pow(2, float64(idx)))
		}
	}

	return res
}

// 135. 分发糖果
func Candy(rating []int) int {
	n := len(rating)
	// rating[i-1] < rate[i], x[i-1] < x[i]
	left := make([]int, n)

	// 左规则
	for i, val := range rating {
		if i > 0 && val > rating[i-1] {
			left[i] = left[i-1] + 1
		} else {
			left[i] = 0
		}
	}
	right := 0
	var res int
	for i := n - 1; i >= 0; i-- {
		if i < n-1 && rating[i] > rating[i+1] {
			right++
		} else {
			right = 1
		}
		res += max(left[i], right)
	}
	return res
}

// 507. 完美数
// 对于一个 正整数，如果它和除了它自身以外的所有 正因子 之和相等，我们称它为 「完美数」。
// 给定一个整数n，如果是完美数，返回 true，否则返回 false
func checkPerfectNumber(num int) bool {
	if num == 1 {
		return false
	}

	sum := 1
	for d := 2; d*d <= num; d++ {
		if num%d == 0 {
			sum += d
			if d*d < num {
				sum += num / d
			}
		}
	}

	return sum == num
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
