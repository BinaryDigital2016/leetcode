package tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/*
给定一个二叉树，找到最长的路径，这个路径中的每个节点具有相同值。 这条路径可以经过也可以不经过根节点。

注意：两个节点之间的路径长度由它们之间的边数表示。

示例 1:

输入:

              5
             / \
            4   5
           / \   \
          1   1   5

输出:

2

示例 2:

输入:

              1
             / \
            4   5
           / \   \
          4   4   5

输出:

2
*/

func longestUnivaluePath(root *TreeNode) int {
	ret := 0
	help(root, &ret)
	return ret
}

// 求root节点左右子树最大路径
func help(root *TreeNode, max *int) int {
	if root == nil {
		return 0
	}
	l := help(root.Left, max)
	r := help(root.Right, max)
	if root.Left != nil && root.Left.Val == root.Val {
		l++
	} else { //不能省略，不相等时置零
		l = 0
	}
	if root.Right != nil && root.Right.Val == root.Val {
		r++
	} else {
		r = 0
	}
	*max = findMax(*max, l+r)
	return findMax(l, r)
}

func findMax(a, b int) int {
	if a > b {
		return a
	}
	return b
}
