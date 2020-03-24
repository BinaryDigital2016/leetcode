package tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/*
给定一个二叉树，在树的最后一行找到最左边的值。

示例 1:

输入:

    2
   / \
  1   3

输出:
1



示例 2:

输入:

        1
       / \
      2   3
     /   / \
    4   5   6
       /
      7

输出:
7



注意: 您可以假设树（即给定的根节点）不为 NULL。
*/

func findBottomLeftValue(root *TreeNode) int {
	ret := 0
	maxLevel := 0
	help(root, 1, &ret, &maxLevel)
	return ret
}

func help(root *TreeNode, curLevel int, ret, maxLevel *int) {
	if root == nil {
		return
	}
	help(root.Left, curLevel+1, ret, maxLevel)
	if curLevel > *maxLevel {
		*maxLevel = curLevel
		*ret = root.Val
	}
	help(root.Right, curLevel+1, ret, maxLevel)
}
