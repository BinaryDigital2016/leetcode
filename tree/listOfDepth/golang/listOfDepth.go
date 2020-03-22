package tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/*
给定一棵二叉树，设计一个算法，创建含有某一深度上所有节点的链表（比如，若一棵树的深度为 D，则会创建出 D 个链表）。返回一个包含所有深度的链表的数组。



示例：

输入：[1,2,3,4,5,null,7,8]

        1
       /  \
      2    3
     / \    \
    4   5    7
   /
  8

输出：[[1],[2,3],[4,5,7],[8]]
*/

func listOfDepth(tree *TreeNode) []*ListNode {
	ret := make([]*ListNode, 0)
	if tree == nil {
		return nil
	}
	q := Queue{make([]*TreeNode, 0)}
	q.Push(tree)
	for !q.Empty() {
		n := q.Size()
		dummy := new(ListNode)
		cur := dummy
		for i := 0; i < n; i++ {
			t := q.Pop()
			cur.Next = &ListNode{Val: t.Val}
			cur = cur.Next
			if t.Left != nil {
				q.Push(t.Left)
			}
			if t.Right != nil {
				q.Push(t.Right)
			}
		}
		ret = append(ret, dummy.Next)
	}
	return ret
}

type Queue struct {
	queue []*TreeNode
}

func (q *Queue) Push(x *TreeNode) {
	if q.queue == nil {
		q.queue = make([]*TreeNode, 0)
	}
	q.queue = append(q.queue, x)
}

func (q *Queue) Pop() *TreeNode {
	if len(q.queue) <= 0 {
		return nil
	}

	t := q.queue[0]
	if len(q.queue) == 1 {
		q.queue = nil
	} else {
		q.queue = q.queue[1:]
	}

	return t
}

func (q *Queue) Empty() bool {
	return q.queue == nil || len(q.queue) == 0
}

func (q *Queue) Size() int {
	if q.queue == nil {
		return 0
	}
	return len(q.queue)
}
