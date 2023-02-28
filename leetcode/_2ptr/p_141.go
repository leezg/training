package _2ptr

func hasCycle(head *ListNode) bool {
	if head == nil {
		return false
	}
	count := map[*ListNode]int{}
	for {
		if _, ok := count[head]; ok {
			return true
		}
		count[head] = 1
		if head.Next == nil {
			break
		}
		head = head.Next
	}
	return false
}

//快慢指针
//func hasCycle(head *ListNode) bool {
//	if head == nil || head.Next == nil {
//		return false
//	}
//	slow, fast := head, head.Next
//	for fast != slow {
//		if fast == nil || fast.Next == nil {
//			return false
//		}
//		slow = slow.Next
//		fast = fast.Next.Next
//	}
//	return true
//}
