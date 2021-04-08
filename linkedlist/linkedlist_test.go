package linkedlist

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRemoveDuplicatesFromLinkedList(t *testing.T) {

}

func TestHasCycle2(t *testing.T) {
	list := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
		},
	}

	assert.Equal(t, false, HasCycle2(list))
}
