package priorityqueue

import (
	"container/heap"
	"sort"
)

// 692. 前K个高频单词
func topKFrequent(words []string, k int) []string {
	wordCountMap := make(map[string]int)
	for _, word := range words {
		wordCountMap[word]++
	}
	uniqueWords := make([]string, 0, len(wordCountMap))
	for word := range wordCountMap {
		uniqueWords = append(uniqueWords, word)
	}
	sort.Slice(uniqueWords, func(i, j int) bool {
		s, t := uniqueWords[i], uniqueWords[j]
		return wordCountMap[s] > wordCountMap[t] || wordCountMap[s] == wordCountMap[t] && s < t
	})
	return uniqueWords[:k]
}

type pair struct {
	w string
	c int
}

type hp []pair

func (h hp) Len() int {
	return len(h)
}

func (h hp) Less(i, j int) bool {
	s, t := h[i], h[j]
	return s.c < t.c || s.c == t.c && s.w > t.w
}

func (h hp) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *hp) Push(v interface{}) {
	*h = append(*h, v.(pair))
}

func (h *hp) Pop() interface{} {
	a := *h
	v := a[len(a)-1]
	*h = a[:len(a)-1]
	return v
}

func topKFrequent2(words []string, k int) []string {
	cnt := map[string]int{}
	for _, w := range words {
		cnt[w] ++
	}
	
	h := &hp{}
	for w, c := range cnt {
		heap.Push(h, pair{w, c})
		if h.Len() > k {
			heap.Pop(h)
		}
	}
	ans := make([]string, k)
	for i:= k-1; i>=0;i-- {
		ans[i] = heap.Pop(h).(pair).w
	}
	return ans
}
