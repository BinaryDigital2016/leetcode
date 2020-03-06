package list

type ListNode struct {
	Val  int
	Next *ListNode
}

// // map+遍历
// func hasCycle(head *ListNode) bool {
//     m := make(map[*ListNode]struct{})
//     q := head
//     for q != nil {
//         if _,ok:=m[q];ok{
//             return true
//         }
//         m[q]=struct{}{}
//         q = q.Next
//     }
//     return false
// }

// 快慢指针
func hasCycle(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}
	s := head
	q := head.Next
	for q != nil && q.Next != nil {
		if s == q {
			return true
		}
		s = s.Next
		q = q.Next.Next
	}
	return false
}
