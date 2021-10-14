package linkedlist

// This is an input struct. Do not edit.
type ListNode struct {
	Val  int
	Next *ListNode
}

func RemoveDuplicatesFromLinkedList(linkedList *ListNode) *ListNode {
	// Write your code here.
	currentNode := linkedList
	for currentNode != nil {
		nextNode := currentNode.Next
		for nextNode != nil && currentNode.Val == nextNode.Val {
			nextNode = nextNode.Next
		}

		currentNode.Next = nextNode
		currentNode = nextNode
	}
	return linkedList
}

// 2. 两数相加
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	carry := 0
	var head *ListNode
	var tail *ListNode
	for l1 != nil || l2 != nil {
		n1, n2 := 0, 0
		if l1 != nil {
			n1 = l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			n2 = l2.Val
			l2 = l2.Next
		}
		sum := n1 + n2 + carry
		sum, carry = sum%10, sum/10
		if head == nil {
			head = &ListNode{
				Val: sum,
			}
			tail = head
		} else {
			tail.Next = &ListNode{
				Val: sum,
			}
			tail = tail.Next
		}
	}
	if carry > 0 {
		tail.Next = &ListNode{Val: carry}
	}
	return head
}

// 19. 删除链表的倒数第 N 个结点
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummy := &ListNode{0, head}
	first, second := head, &ListNode{0, head}

	for i := 0; i < n; i++ {
		first = first.Next
	}
	for ; first != nil; first = first.Next {
		second = second.Next
	}
	second.Next = second.Next.Next
	return dummy.Next
}

// 206. 反转链表
func reverseList(head *ListNode) *ListNode {
	var prev *ListNode
	curr := head
	for curr != nil {
		next := curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}
	return prev
}

// 141. 环形链表
// hash表
func hasCycle(head *ListNode) bool {
	listNodeMap := make(map[*ListNode]struct{}, 0)
	for head != nil {
		if _, ok := listNodeMap[head]; ok {
			return true
		}

		listNodeMap[head] = struct{}{}
		head = head.Next
	}
	return false
}

// LRUCache 146. LRU 缓存机制
// 运用你所掌握的数据结构，设计和实现一个  LRU (最近最少使用) 缓存机制 。
// 实现 LRUCache 类：
//
// LRUCache(int capacity) 以正整数作为容量 capacity 初始化 LRU 缓存
// int get(int key) 如果关键字 key 存在于缓存中，则返回关键字的值，否则返回 -1 。
// void put(int key, int value) 如果关键字已经存在，则变更其数据值；如果关键字不存在，则插入该组「关键字-值」。当缓存容量达到上限时，它应该在写入新数据之前删除最久未使用的数据值，从而为新的数据值留出空间。
//
//
// 进阶：你是否可以在 O(1) 时间复杂度内完成这两种操作？
type LRUCache struct {
	capacity int
	size     int
	cache    map[int]*dLinkedList
	head     *dLinkedList
	tail     *dLinkedList
}

type dLinkedList struct {
	key int
	val int

	prev *dLinkedList
	next *dLinkedList
}

func NewDLinkedList(key, value int) *dLinkedList {
	return &dLinkedList{
		key: key,
		val: value,
	}
}

func Constructor(capacity int) LRUCache {
	l := LRUCache{
		cache:    make(map[int]*dLinkedList, capacity),
		capacity: capacity,
		head:     NewDLinkedList(0, 0),
		tail:     NewDLinkedList(0, 0),
	}
	l.head.next = l.tail
	l.tail.prev = l.head
	return l
}

func (c *LRUCache) Get(key int) int {
	node, ok := c.cache[key]
	if !ok {
		return -1
	}

	c.moveToHead(node)
	return node.val
}

func (c *LRUCache) Put(key int, value int) {
	node, ok := c.cache[key]
	if ok {
		node.val = value
		c.moveToHead(node)
	} else {
		node = NewDLinkedList(key, value)
		c.cache[key] = node
		c.addToHead(node)
		c.size++
		if c.size > c.capacity {
			tail := c.removeTail()
			delete(c.cache, tail.key)
			c.size--
		}
	}
}

func (c *LRUCache) addToHead(node *dLinkedList) {
	node.prev = c.head
	node.next = c.head.next
	c.head.next.prev = node
	c.head.next = node
}

func (c *LRUCache) removeNode(node *dLinkedList) {
	node.prev.next = node.next
	node.next.prev = node.prev
}

func (c *LRUCache) moveToHead(node *dLinkedList) {
	c.removeNode(node)
	c.addToHead(node)
}

func (c *LRUCache) removeTail() *dLinkedList {
	node := c.tail.prev
	c.removeNode(node)
	return node
}

// 快慢指针
func hasCycle2(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}

	slow, fast := head, head.Next
	for fast != slow {
		if fast == nil || fast.Next == nil {
			return false
		}
		fast = fast.Next.Next
		slow = slow.Next
	}
	return true
}

// 160. 相交链表
// 快慢指针
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	pA := headA
	pB := headB

	for pA != pB {
		if pA == nil {
			pA = headB
		} else {
			pA = pA.Next
		}
		if pB == nil {
			pB = headA
		} else {
			pB = pB.Next
		}
	}
	return pA
}

// 234. 回文链表
func isPalindrome(head *ListNode) bool {
	arr := make([]int, 0)

	for ; head != nil; head = head.Next {
		arr = append(arr, head.Val)
	}

	for i := 0; i < len(arr)/2; i++ {
		if arr[i] != arr[len(arr)-i-1] {
			return false
		}
	}
	return true
}
