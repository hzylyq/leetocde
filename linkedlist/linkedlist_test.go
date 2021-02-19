package linkedlist

import "testing"

func TestRemoveDuplicatesFromLinkedList(t *testing.T) {
	list := &LinkedList{Value:1,
		Next:&LinkedList{
		Value:1,
		Next:&LinkedList{
			Value: 1,
			Next:  &LinkedList{
				Value: 1,
				Next:  &LinkedList{
					Value: 3,
					Next:  &LinkedList{
						Value: 4,
						Next:  &LinkedList{
							Value: 4,
							Next:  &LinkedList{
								Value: 4,
								Next:  &LinkedList{
									Value: 5,
									Next:  &LinkedList{
										Value: 6,
										Next:  &LinkedList{
											Value: 6,
											Next:  nil,
										},
									},
								},
							},
						},
					},
				},
			},
		},
		}}
	RemoveDuplicatesFromLinkedList(list)
	t.Log(list)

}
