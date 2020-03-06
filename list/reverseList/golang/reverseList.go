package list

type ListNode struct {
	Val  int
	Next *ListNode
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

// 为啥超时 一脸懵逼
// func reverseList(head *ListNode) *ListNode {
// 	if head == nil || head.Next == nil {
// 		return head
// 	}
// 	t := &ListNode{
// 		Val:-1,
// 		Next:head,
// 	}
// 	cur := head.Next
// 	for cur != nil{
// 		n := cur.Next
// 		cur.Next = t.Next
// 		t.Next = cur
// 		cur = n
// 	}

// 	return t.Next
// }

// func reverseList(head *ListNode) *ListNode {
//     var pre *ListNode
// 	cur := head
// 	for cur != nil{
// 		n := cur.Next
// 		cur.Next = pre
// 		pre = cur
// 		cur = n
// 	}

// 	return pre
// }

// 递归版
func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	cur := reverseList(head.Next) //最后一个节点
	head.Next.Next = head         //下一个节点指向自己，形成环
	head.Next = nil               //断掉环

	return cur
}
