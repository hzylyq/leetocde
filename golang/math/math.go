package math

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

// 9. 回文数
func isPalindrome(x int) bool {
	// 负数
	if x < 0 || (x%10 == 0 && x != 0) {
		return false
	}

	revertNum := 0
	for x > revertNum {
		revertNum = revertNum*10 + x%10
		x /= 10
	}

	return x == revertNum || x == revertNum/10
}

// 69. x 的平方根
func mySqrt(x int) int {
	if x == 0 {
		return x
	}
	ans := int(math.Exp(0.5 * math.Log(float64(x))))
	if (ans+1)*(ans+1) <= x {
		return ans + 1
	}
	return ans
}

func mySqrt2(x int) int {

}

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

// Candy 135. 分发糖果
func Candy(rating []int) int {
	// n := len(rating)
	// // rating[i-1] < rate[i], x[i-1] < x[i]
	// left := make([]int, n)

	// // 左规则
	// for i, val := range rating {
	// 	if i > 0 && val > rating[i-1] {
	// 		left[i] = left[i-1] + 1
	// 	} else {
	// 		left[i] = 0
	// 	}
	// }
	// right := 0
	// var res int

	// for i := n - 1; i >= 0; i-- {
	// 	rating[i+1]
	// 	{
	// 		right++
	// 	} else {
	// 		right = 1
	// 	}
	// 	res += max(left[i], right)
	// }
	// return res
	return 0
}

// 357. 统计各位数字都不同的数字个数
func countNumbersWithUniqueDigits(n int) int {
	if n == 0 {
		return 1
	}
	if n == 1 {
		return 10
	}

	ans, cur := 10, 9
	for i := 0; i < n-1; i++ {
		cur *= 9 - i
		ans += cur
	}

	return ans

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

// 537. 复数乘法
func complexNumberMultiply(num1 string, num2 string) string {
	real1, imag1 := parseComplexNumber(num1)
	real2, imag2 := parseComplexNumber(num2)
	return fmt.Sprintf("%d+%di", real1*real2-imag1*imag2, real1*imag2+real2*imag1)
}

// 553. 最优除法
func optimalDivision(nums []int) string {
	n := len(nums)
	if n == 1 {
		return strconv.Itoa(nums[0])
	}
	if n == 2 {
		return fmt.Sprintf("%d/%d", nums[0], nums[1])
	}

	var ans string
	ans += fmt.Sprintf("%d/(%d", nums[0], nums[1])
	for _, num := range nums[2:] {
		ans += fmt.Sprintf("/%d", num)
	}
	ans += ")"
	return ans
}

func parseComplexNumber(num string) (real, imag int) {
	i := strings.Index(num, "+")
	real, _ = strconv.Atoi(num[:i])
	imag, _ = strconv.Atoi(num[i+1 : len(num)-1])
	return
}

// 1238. 循环码排列
func circularPermutation(n int, start int) []int {
	ans := make([]int, 1<<n)
	for i := range ans {
		ans[i] = (i >> 1) ^ i ^ start
	}

	return ans
}

// 1401. 圆和矩形是否有重叠
// func checkOverlap(radius int, xCenter int, yCenter int, x1 int, y1 int, x2 int, y2 int) bool {
// 	var dist float64
// 	if xCenter
//
// 	return false
// }

// 1716. 计算力扣银行的钱
// Hercy 想要为购买第一辆车存钱。他 每天 都往力扣银行里存钱。
// 最开始，他在周一的时候存入 1 块钱。从周二到周日，他每天都比前一天多存入 1 块钱。在接下来每一个周一，他都会比 前一个周一 多存入 1 块钱。
// 给你 n ，请你返回在第 n 天结束的时候他在力扣银行总共存了多少块钱。
func totalMoney(n int) int {
	return 0
}

// 2485. 找出中枢整数 数学公式
func pivotInteger(n int) int {
	t := (n*n + n) / 2
	res := int(math.Sqrt(float64(t)))
	if res*res == t {
		return res
	}

	return -1
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
