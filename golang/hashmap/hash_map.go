package hashmap

type MyHashMap struct {
	m map[int]int
}

/** Initialize your data structure here. */
func Constructor() MyHashMap {
	return MyHashMap{
		m: make(map[int]int),
	}
}

/** value will always be non-negative. */
func (this *MyHashMap) Put(key int, value int) {
	this.m[key] = value
}

/** Returns the value to which the specified key is mapped, or -1 if this map contains no mapping for the key */
func (this *MyHashMap) Get(key int) int {
	if val, ok := this.m[key]; ok {
		return val
	}
	return -1
}

/** Removes the mapping of the specified value key if this map contains a mapping for the key */
func (this *MyHashMap) Remove(key int) {
	delete(this.m, key)
}

// 1.两数之和
func twoSum(nums []int, target int) []int {
	hashMap := make(map[int]int)

	for i, num := range nums {
		if j, ok := hashMap[target-num]; ok {
			return []int{i, j}
		}

		hashMap[num] = i
	}

	return nil
}

// 13. 罗马数字转整数
func romanToInt(s string) int {
	resMap := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}

	n := len(s)
	var res int
	for i, ch := range s {
		val := resMap[byte(ch)]

		if i < n-1 && val < resMap[byte(s[i+1])] {
			res -= val
		} else {
			res += val
		}
	}
	return res
}

/**
 * Your MyHashMap object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Put(key,value);
 * param_2 := obj.Get(key);
 * obj.Remove(key);
 */

// 128. 最长连续序列
func longestConsecutive(nums []int) int {
	numMap := make(map[int]bool, len(nums))
	for i := 0; i < len(nums); i++ {
		numMap[nums[i]] = true
	}

	var longStreak int
	for num := range numMap {
		if !numMap[num-1] {
			currentNum := num
			currentStreak := 1
			for numMap[currentNum+1] {
				currentNum++
				currentStreak++
			}

			if currentStreak > longStreak {
				longStreak = currentStreak
			}
		}
	}

	return longStreak
}

// 205. 同构字符串
func isIsomorphic(s string, t string) bool {
	s1Map := make(map[byte]byte, len(s))
	s2Map := make(map[byte]byte, len(t))

	for i := range s {
		x, y := s[i], t[i]

		if s1Map[x] > 0 && s1Map[x] != y || s2Map[y] > 0 && s2Map[y] != x {
			return false
		}

		s1Map[x] = y
		s2Map[y] = x
	}

	return true
}

// 219. 存在重复元素 II
// 给定一个整数数组和一个整数 k，判断数组中是否存在两个不同的索引 i 和 j，使得 nums [i] = nums [j]，并且 i 和 j 的差的 绝对值 至多为 k。
func containsNearbyDuplicate(nums []int, k int) bool {
	indexDiff := make(map[int]int, len(nums))
	for i := 0; i < len(nums); i++ {
		if val, ok := indexDiff[nums[i]]; ok && i-val <= k {
			return true
		}
		indexDiff[nums[i]] = i
	}
	return false
}

// 383. 赎金信
func canConstruct(ransomNote string, magazine string) bool {
	if len(ransomNote) > len(magazine) {
		return false
	}

	cnt := [26]int{}
	for _, ch := range magazine {
		cnt[ch-'a']++
	}

	for _, ch := range ransomNote {
		cnt[ch-'a']--
		if cnt[ch-'a'] < 0 {
			return false
		}
	}
	return true
}

// 1365. 有多少小于当前数字的数字
// 暴力破解
func SmallerNumbersThanCurrent(nums []int) []int {
	res := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		for j := 0; j < len(nums); j++ {
			if nums[i] > nums[j] && j != i {
				res[i]++
			}
		}
	}
	return res
}

// 2325. 解密消息
func decodeMessage(key string, message string) string {
	var keyMap = map[rune]byte{}

	curr := byte('a')

	for _, c := range key {
		if _, ok := keyMap[c]; ok || c == ' ' {
			continue
		}

		keyMap[c] = curr
		curr++
	}

	res := []byte(message)
	for i, c := range message {
		if c != ' ' {
			res[i] = keyMap[c]
		}
	}

	return string(res)
}
