package list

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	cur := head.Next
	pre := head
	for cur != nil {
		if cur.Val == pre.Val {
			pre.Next = cur.Next
			cur = cur.Next
			continue
		}
		pre = cur
		cur = cur.Next
	}

	return head
}
