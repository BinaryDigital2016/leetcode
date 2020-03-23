package tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/*
给你一棵二叉树，请你返回满足以下条件的所有节点的值之和：

    该节点的祖父节点的值为偶数。（一个节点的祖父节点是指该节点的父节点的父节点。）

如果不存在祖父节点值为偶数的节点，那么返回 0 。
*/

func sumEvenGrandparent(root *TreeNode) int {
	if root == nil {
		return 0
	}
	ret := 0
	gp, p := 1, 1
	help(root, gp, p, &ret)
	return ret
}

func help(root *TreeNode, gp, p int, sum *int) {
	if root == nil {
		return
	}
	if gp%2 == 0 {
		*sum += root.Val
	}
	help(root.Left, p, root.Val, sum)
	help(root.Right, p, root.Val, sum)
}
