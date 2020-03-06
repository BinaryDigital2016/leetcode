/*
对链表进行插入排序
*/
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func insertionSortList(head *ListNode) *ListNode {
    if head == nil || head.Next == nil {
        return head
    }

    pre := head
    cur := head.Next
    for cur != nil{
        if cur.Val < pre.Val{
            tmp := cur
            cur = cur.Next
            pre.Next = cur //断掉cur， pre链接cur.Next

            var pre2 *ListNode //插入位置前节点
            cur2 := head //插入位置后节点
            for cur2 != pre && tmp.Val > cur2.Val{
                pre2 = cur2
                cur2 = cur2.Next
            }

            if pre2 == nil{
                // 插入链表头
                tmp.Next = head
                head = tmp
            } else {
                pre2.Next = tmp //插入
                tmp.Next = cur2
            }
            
        } else {
            pre = cur
            cur = cur.Next
        }
        
    }

    return head
}
