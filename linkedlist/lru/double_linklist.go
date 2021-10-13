package lru

type doubleLinkList struct {
	key int
	val int

	prev *doubleLinkList
	next *doubleLinkList
}

func (l *doubleLinkList) moveToFront() {

}


