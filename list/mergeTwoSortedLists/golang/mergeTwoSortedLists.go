// // 递归
// func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
//     //如果有一条链是nil，直接返回另外一条链
//     if l1 == nil {
//         return l2
//     }
//     if l2 == nil {
//         return l1
//     }
//     // 定义一个结果节点
//     var res *ListNode
//     // 当l1节点的值大于l2节点的值，那么res指向l2的节点，从l2开始遍历，反之从l1开始
//     if l1.Val >= l2.Val {
//         res = l2
//         res.Next = mergeTwoLists(l1, l2.Next)
//     } else {
//         res = l1
//         res.Next = mergeTwoLists(l1.Next, l2)
//     }
//     return res
// }

// 非递归
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
    //如果有一条链是nil，直接返回另外一条链
    if l1 == nil {
        return l2
    }
    if l2 == nil {
        return l1
    }

    result := l1
	r := result

    if l1.Val > l2.Val {
		result = l2
		l2 = l2.Next
		r = result
	} else {
		l1 = l1.Next
		r = result
	}

    for l1 != nil && l2 != nil {
        if l1.Val > l2.Val {
            result.Next = l2
            l2 = l2.Next
            result = result.Next
        } else {
            result.Next = l1
            l1 = l1.Next
            result = result.Next
        }
    }

    if l1 == nil && l2 != nil{
		result.Next = l2
	}

	if l1 != nil && l2 == nil{
		result.Next = l1
	}
 
	return r

}
