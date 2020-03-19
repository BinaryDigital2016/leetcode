package tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/*
给定一个二叉树，检查它是否是镜像对称的。

例如，二叉树 [1,2,2,3,4,4,3] 是对称的。

    1
   / \
  2   2
 / \ / \
3  4 4  3

但是下面这个 [1,2,2,null,3,null,3] 则不是镜像对称的:

    1
   / \
  2   2
   \   \
   3    3
*/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// 递归
// func isSymmetric(root *TreeNode) bool {
//     if root == nil {
//         return true
//     }
//     return isMirror(root.Left, root.Right)
// }

// func isMirror(l *TreeNode, r *TreeNode) bool{
//     if l == nil && r == nil {
//         return true
//     }
//     if l == nil || r == nil {
//         return false
//     }

//     return (l.Val == r.Val) && isMirror(l.Left, r.Right) && isMirror(l.Right, r.Left)
// }

// 非递归
func isSymmetric(root *TreeNode) bool {
	q := Queue{}
	q.Append(root)
	q.Append(root)
	for !q.Empty() {
		a := q.Pop()
		b := q.Pop()
		if a == nil && b == nil {
			continue
		}
		if a == nil || b == nil {
			return false
		}
		if a.Val != b.Val {
			return false
		}
		q.Append(a.Left)
		q.Append(b.Right)
		q.Append(a.Right)
		q.Append(b.Left)
	}
	return true
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
	if len(q.elem) > 0 {
		t := q.elem[0]
		q.elem = q.elem[1:]
		return t
	}
	return nil
}

func (q *Queue) Empty() bool {
	return len(q.elem) == 0
}
