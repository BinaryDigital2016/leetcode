package tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/*
给定一个二叉搜索树，同时给定最小边界L 和最大边界 R。通过修剪二叉搜索树，使得所有节点的值在[L, R]中 (R>=L) 。你可能需要改变树的根节点，所以结果应当返回修剪好的二叉搜索树的新的根节点。

示例 1:

输入:
    1
   / \
  0   2

  L = 1
  R = 2

输出:
    1
      \
       2

示例 2:

输入:
    3
   / \
  0   4
   \
    2
   /
  1

  L = 1
  R = 3

输出:
      3
     /
   2
  /
 1
*/

/*
当 node.val > R，那么修剪后的二叉树必定出现在节点的左边。
类似地，当 node.val < L，那么修剪后的二叉树出现在节点的右边。
否则，我们将会修剪树的两边。
*/

func trimBST(root *TreeNode, L int, R int) *TreeNode {
	if root == nil {
		return root
	}

	if root.Val > R {
		return trimBST(root.Left, L, R)
	}
	if root.Val < L {
		return trimBST(root.Right, L, R)
	}
	root.Left = trimBST(root.Left, L, R)
	root.Right = trimBST(root.Right, L, R)
	return root
}
