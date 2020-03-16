package tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/*
给定一个二叉树，返回其节点值自底向上的层次遍历。 （即按从叶子节点所在层到根节点所在的层，逐层从左向右遍历）

例如：
给定二叉树 [3,9,20,null,null,15,7],

    3
   / \
  9  20
    /  \
   15   7

返回其自底向上的层次遍历为：

[
  [15,7],
  [9,20],
  [3]
]
*/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func levelOrderBottom(root *TreeNode) [][]int {
	ret := make([][]int, 0)
	if root == nil {
		return nil
	}

	q := Queue{}
	ret = append(ret, []int{root.Val})
	q.Append(root)
	for !q.Empty() {
		t := make([]int, 0)
		n := q.Size()
		for i := 0; i < n; i++ { //每次循环开始，队列中的节点都是同一层的节点
			cur := q.Pop()
			if cur.Left == nil && cur.Right == nil {
				continue
			}
			if cur.Left != nil {
				t = append(t, cur.Left.Val)
				q.Append(cur.Left)
			}
			if cur.Right != nil {
				t = append(t, cur.Right.Val)
				q.Append(cur.Right)
			}
		}
		if len(t) > 0 {
			//ret = append([][]int{t}, ret...) //往前插入，效率低，内存大
			ret = append(ret, t) //往后插入，最后逆置
		}
	}
	Reverse(ret)
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

func Reverse(s [][]int) {
	if len(s) == 0 {
		return
	}
	i, j := 0, len(s)-1
	for i <= j {
		s[i], s[j] = s[j], s[i]
		i++
		j--
	}
}
