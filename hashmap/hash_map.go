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
