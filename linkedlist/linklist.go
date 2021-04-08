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
func HasCycle2(head *ListNode) bool {
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
