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
