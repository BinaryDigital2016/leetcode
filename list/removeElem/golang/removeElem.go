package list

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeElements(head *ListNode, val int) *ListNode {
	for head != nil {
		if head.Val == val {
			head = head.Next
		} else {
			break
		}
	}

	if head != nil {
		pre := head
		cur := head.Next
		for cur != nil {
			if cur.Val == val {
				pre.Next = cur.Next
			} else {
				pre = cur
			}
			cur = cur.Next
		}
	}

	return head
}
