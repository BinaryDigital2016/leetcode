package list

type ListNode struct {
	Val  int
	Next *ListNode
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}
	p := headA
	q := headB
	for i := 0; i < 2; i++ {
		for p != nil && q != nil {
			p = p.Next
			q = q.Next
		}
		if p == nil {
			p = headB
		} else {
			q = headA
		}
	}

	for {
		if p == q {
			return p
		}
		p = p.Next
		q = q.Next
	}
}
