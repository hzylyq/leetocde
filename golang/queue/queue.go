package queue

// MyCircularDeque 641. 设计循环双端队列
type MyCircularDeque struct {
	front, rear int
	element     []int
}

func Constructor(k int) MyCircularDeque {
	return MyCircularDeque{
		element: make([]int, k+1),
	}
}

func (q *MyCircularDeque) InsertFront(value int) bool {
	if q.IsFull() {
		return false
	}

	q.front = (q.front - 1 + len(q.element)) % len(q.element)
	q.element[q.front] = value
	return true
}

func (q *MyCircularDeque) InsertLast(value int) bool {
	if q.IsFull() {
		return false
	}

	q.element[q.rear] = value
	q.rear = (q.rear + 1) % len(q.element)
	return true
}

func (q *MyCircularDeque) DeleteFront() bool {
	if q.IsEmpty() {
		return false
	}

	q.front = (q.front + 1) % len(q.element)
	return true
}

func (q *MyCircularDeque) DeleteLast() bool {
	if q.IsEmpty() {
		return false
	}

	q.rear = (q.rear - 1 + len(q.element)) % len(q.element)
	return true
}

func (q *MyCircularDeque) GetFront() int {
	if q.IsEmpty() {
		return -1
	}

	return q.element[q.front]
}

func (q *MyCircularDeque) GetRear() int {
	if q.IsEmpty() {
		return -1
	}

	return q.element[(q.rear-1+len(q.element))%len(q.element)]
}

func (q *MyCircularDeque) IsEmpty() bool {
	return q.front == q.rear
}

func (q *MyCircularDeque) IsFull() bool {
	return (q.rear+1)%len(q.element) == q.front
}
