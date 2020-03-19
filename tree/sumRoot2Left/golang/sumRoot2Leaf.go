package tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/*
给出一棵二叉树，其上每个结点的值都是 0 或 1 。每一条从根到叶的路径都代表一个从最高有效位开始的二进制数。例如，如果路径为 0 -> 1 -> 1 -> 0 -> 1，那么它表示二进制数 01101，也就是 13 。

对树上的每一片叶子，我们都要找出从根到该叶子的路径所表示的数字。

示例：

输入：[1,0,1,0,1,0,1]
输出：22
解释：(100) + (101) + (110) + (111) = 4 + 5 + 6 + 7 = 22
*/

func SumRootToLeaf(root *TreeNode) int {
	sum := 0
	ret := 0
	sumRootToLeafDFS(root, sum, &ret)
	return ret
}

func sumRootToLeafDFS(root *TreeNode, sum int, ret *int) {
	if root == nil {
		return
	}

	sum = (sum << 1) + root.Val
	if root.Left == nil && root.Right == nil {
		*ret += sum
		return
	}
	sumRootToLeafDFS(root.Left, sum, ret)
	sumRootToLeafDFS(root.Right, sum, ret)
}
