package complex

// type node struct {
// 	key   string
// 	count int
// }

// type AllOne struct {
// 	nodes map[string]*node
// }

// func Constructor() AllOne {
// 	return AllOne{
// 		nodes: make(map[string]*list.Element),
// 	}
// }

// func (l *AllOne) Inc(key string) {
// 	if node, ok := l.nodes[key]; ok {
// 		node.count++
// 		l.nodes[key] = node
// 		return
// 	}

// 	l.nodes[key] = &node{
// 		key:   key,
// 		count: 1,
// 	}
// }

// func (l *AllOne) Dec(key string) {
// 	if node, ok := l.nodes[key]; ok {
// 		node.count--
// 		l.nodes[key] = node
// 		return
// 	}
// }

// func (l *AllOne) GetMaxKey() string {

// }

// func (l *AllOne) GetMinKey() string {

// }

/**
 * Your AllOne object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Inc(key);
 * obj.Dec(key);
 * param_3 := obj.GetMaxKey();
 * param_4 := obj.GetMinKey();
 */

type LRUCache struct {
	size     int
	capacity int
	
	cache      map[int]*DLinkNode
	head, tail *DLinkNode
}

type DLinkNode struct {
	key, value int
	prev, next *DLinkNode
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		size:     0,
		capacity: capacity,
		cache:    make(map[int]*DLinkNode, capacity),
		head:     nil,
		tail:     nil,
	}
}

func (this *LRUCache) Get(key int) int {
	value, ok := this.cache[key]
	if !ok {
		return 0
	}
	
	return value.value
}

func (this *LRUCache) Put(key int, value int) {

}
