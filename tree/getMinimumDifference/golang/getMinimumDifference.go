package tree

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/*
给你一棵所有节点为非负值的二叉搜索树，请你计算树中任意两节点的差的绝对值的最小值。



示例：

输入：

   1
    \
     3
    /
   2

输出：
1
*/

func getMinimumDifference(root *TreeNode) int {
	if root == nil {
		return -1
	}
	min := math.MaxInt32
	pre := -1
	inorder(root, &pre, &min)
	return min
}

func inorder(root *TreeNode, pre, min *int) *int {
	if root != nil {
		pre = inorder(root.Left, pre, min)
		if *pre != -1 {
			*min = minV(*min, root.Val-*pre)
		}
		*pre = root.Val
		pre = inorder(root.Right, pre, min)
	}
	return pre
}

func minV(a, b int) int {
	if a < b {
		return a
	}
	return b
}
