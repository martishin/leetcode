package reverse_linked_list_206

func reverseListRecursive(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	var newHead = reverseListRecursive(head.Next)
	head.Next.Next = head
	head.Next = nil

	return newHead
}

func TestRecursive() {
	head := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 4,
					Next: &ListNode{
						Val:  5,
						Next: nil,
					},
				},
			},
		},
	}

	newHead := reverseListRecursive(head)
	printList(newHead)
}
