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

// 21. 合并两个有序链表
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil {
		return list2
	} else if list2 == nil {
		return list1
	} else if list1.Val > list2.Val {
		list2.Next = mergeTwoLists(list1, list2.Next)
		return list2
	} else {
		list1.Next = mergeTwoLists(list2, list1.Next)
		return list1
	}
}

func mergeTwoLists2(list1 *ListNode, list2 *ListNode) *ListNode {
	var curr *ListNode

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

// 142. 环形链表 II
// 给定一个链表，返回链表开始入环的第一个节点。 如果链表无环，则返回 null。
// 为了表示给定链表中的环，我们使用整数 pos 来表示链表尾连接到链表中的位置（索引从 0 开始）。 如果 pos 是 -1，则在该链表中没有环。注意，pos 仅仅是用于标识环的情况，并不会作为参数传递到函数中。
// 说明：不允许修改给定的链表。
// 进阶：你是否可以使用 O(1) 空间解决此题？
func detectCycle(head *ListNode) *ListNode {
	slow, fast := head, head
	for fast != nil {
		slow = slow.Next
		if fast.Next == nil {
			return nil
		}
		fast = fast.Next.Next
		if fast == slow {
			p := head
			for p != slow {
				p = p.Next
				slow = slow.Next
			}
			return p
		}
	}

	return nil
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

// 148. 排序链表
// 给你链表的头结点 head ，请将其按 升序 排列并返回 排序后的链表 。
// 进阶：你可以在 O(n log n) 时间复杂度和常数级空间复杂度下，对链表进行排序吗？
func sortList(head *ListNode) *ListNode {
	return sort(head, nil)
}

func sort(head, tail *ListNode) *ListNode {
	if head == nil {
		return head
	}

	if head.Next == tail {
		head.Next = nil
		return head
	}

	slow, fast := head, head
	for fast != tail {
		slow = slow.Next
		fast = fast.Next
		if fast != tail {
			fast = fast.Next
		}
	}

	mid := slow
	return mergeList(sort(head, mid), sort(mid, tail))
}

func mergeList(head1, head2 *ListNode) *ListNode {
	dummyHead := &ListNode{}
	temp, temp1, temp2 := dummyHead, head1, head2
	for temp1 != nil && temp2 != nil {
		if temp1.Val <= temp2.Val {
			temp.Next = temp1
			temp1 = temp1.Next
		} else {
			temp.Next = temp2
			temp2 = temp2.Next
		}

		temp = temp.Next
	}

	if temp1 != nil {
		temp.Next = temp1
	} else if temp2 != nil {
		temp.Next = temp2
	}

	return dummyHead.Next
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

// 237. 删除链表中的节点
// 请编写一个函数，用于删除单链表中某个特定节点。在设计函数时需要注意，你无法访问链表的头节点 head，只能直接访问要被删除的节点。
// 题目数据保证需要删除的节点不是末尾节点。
func deleteNode(node *ListNode) {
	node.Val = node.Next.Val
	node.Next = node.Next.Next
}

// 876. 链表的中间结点
func middleNode(head *ListNode) *ListNode {
	slow, fast := head, head

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	return slow
}
