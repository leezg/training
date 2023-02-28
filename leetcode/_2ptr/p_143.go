package _2ptr

func departList(head *ListNode) *ListNode {
	slow, fast := head, head
	for {
		if fast.Next == nil {
			tmp := slow.Next
			slow.Next = nil
			return tmp
		}
		if fast.Next.Next == nil {
			tmp := slow.Next
			slow.Next = nil
			return tmp
		}
		slow = slow.Next
		fast = fast.Next.Next
	}
}

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

func reorderList(head *ListNode) {
	mid := departList(head)
	mid = reverseList(mid)
	curr := head
	for mid != nil {
		cNext := curr.Next
		mNext := mid.Next
		curr.Next = mid
		mid.Next = cNext
		curr = cNext
		mid = mNext
	}
}
