package tries

type SuffixTrie map[byte]SuffixTrie

func NewSuffixTrie() SuffixTrie {
	trie := SuffixTrie{}
	return trie
}

func (trie SuffixTrie) PopulateSuffixTrieFrom(str string) {
	for i := range str {
		node := trie
		for j := i; j < len(str); j++ {
			letter := str[j]
			if _, ok := node[letter]; !ok {
				node[letter] = NewSuffixTrie()
			}
			node = node[letter]
		}
		node['*'] = nil
	}
}

func (trie SuffixTrie) Contains(str string) bool {
	node := trie
	for i := 0; i < len(str); i++ {
		letter := str[i]
		if _, ok := node[letter]; !ok {
			return false
		}

		node = node[letter]
	}
	_, ok := node['*']
	return ok
}
