package mysort

/*
在 O(n log n) 时间复杂度和常数级空间复杂度下，对链表进行排序。

示例 1:

输入: 4->2->1->3
输出: 1->2->3->4

示例 2:

输入: -1->5->3->4->0
输出: -1->0->3->4->5
*/

type ListNode struct {
	Val  int
	Next *ListNode
}

// 自底向上归并
func sortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	n := 0
	cur := head
	for cur != nil {
		n++
		cur = cur.Next
	}

	dummyHead := &ListNode{Val: 0, Next: head}

	for span := 1; span < n; span <<= 1 {
		cur = dummyHead.Next //不能是=head, cur后head标识的链表已被修改，dummyHead指向处理后的链表
		tail := dummyHead    //已合并链表的末尾
		for cur != nil {     //一个span循环
			left := cur
			right := cut(cur, span)
			cur = cut(right, span)

			tail.Next = merge(left, right)
			for tail.Next != nil {
				tail = tail.Next
			}
		}
	}
	return dummyHead.Next
}

// 断掉head的前size个元素，返回剩下的head
func cut(head *ListNode, size int) *ListNode {
	cur := head
	for cur != nil && size > 1 {
		cur = cur.Next
		size--
	}

	if cur == nil {
		return nil
	}

	next := cur.Next
	cur.Next = nil
	return next
}

func merge(l1 *ListNode, l2 *ListNode) *ListNode {
	dummyHead := &ListNode{Val: 0}
	cur := dummyHead
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			cur.Next = l1
			l1 = l1.Next
		} else {
			cur.Next = l2
			l2 = l2.Next
		}
		cur = cur.Next
	}

	if l1 != nil {
		cur.Next = l1
	}

	if l2 != nil {
		cur.Next = l2
	}

	return dummyHead.Next
}
