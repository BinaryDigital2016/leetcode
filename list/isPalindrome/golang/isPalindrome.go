package list

type ListNode struct {
	Val  int
	Next *ListNode
}

// 1.链表保存到数组中，再判断，略

// 2.链表逆置后再判断，略

// 3.链表后半部分逆置后再判断
func isPalindrome(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return true
	}
	slow := head
	fast := head
	// slow指向中间节点
	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	// 后半部分节点逆置,pre指向逆置后的第一个节点
	var pre *ListNode
	cur := slow.Next
	for cur != nil {
		next := cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}

	for pre != nil && head != nil {
		if pre.Val != head.Val {
			return false
		}
		pre = pre.Next
		head = head.Next
	}
	return true
}
