package tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/*
二叉搜索树原地转成链表
*/

func convertBiNode(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}

	head := new(TreeNode)
	inorder(root, head)
	return head.Right
}

func inorder(root *TreeNode, pre *TreeNode) *TreeNode {
	if root != nil {
		pre = inorder(root.Left, pre)
		root.Left = nil
		pre.Right = root
		pre = root
		pre = inorder(root.Right, pre)
	}
	return pre
}
