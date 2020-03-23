package tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/*
给你一棵二叉树，请你返回层数最深的叶子节点的和。



示例：

输入：root = [1,2,3,4,5,null,6,7,null,null,null,null,8]
输出：15
*/

func deepestLeavesSum(root *TreeNode) int {
	if root == nil {
		return 0
	}
	q := Queue{make([]*TreeNode, 0)}
	q.Append(root)
	ret := 0
	for !q.Empty() {
		ret = 0
		n := q.Size()
		for i := 0; i < n; i++ {
			cur := q.Pop()
			ret += cur.Val
			if cur.Left != nil {
				q.Append(cur.Left)
			}
			if cur.Right != nil {
				q.Append(cur.Right)
			}
		}
	}
	return ret
}

type Queue struct {
	elem []*TreeNode
}

func (q *Queue) Append(e *TreeNode) {
	if q.elem == nil {
		q.elem = make([]*TreeNode, 0)
	}
	q.elem = append(q.elem, e)
}

func (q *Queue) Pop() *TreeNode {
	if q.elem == nil || len(q.elem) == 0 {
		return nil
	}

	t := q.elem[0]
	q.elem = q.elem[1:]
	return t
}

func (q *Queue) Empty() bool {
	return len(q.elem) == 0
}

func (q *Queue) Size() int {
	return len(q.elem)
}
