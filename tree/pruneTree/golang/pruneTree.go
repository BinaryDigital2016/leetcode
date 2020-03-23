package tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/*
给定二叉树根结点 root ，此外树的每个结点的值要么是 0，要么是 1。

返回移除了所有不包含 1 的子树的原二叉树。

( 节点 X 的子树为 X 本身，以及所有 X 的后代。)

示例1:
输入: [1,null,0,0,1]
输出: [1,null,0,null,1]
*/

func pruneTree(root *TreeNode) *TreeNode {
	if !containOne(root) {
		return nil
	}
	prune(root.Left, root, true)
	prune(root.Right, root, false)
	return root
}
func prune(root, parent *TreeNode, isLeft bool) {
	if root == nil {
		return
	}
	if !containOne(root) {
		if isLeft {
			parent.Left = nil
		} else {
			parent.Right = nil
		}
	}
	prune(root.Left, root, true)
	prune(root.Right, root, false)
}

func containOne(root *TreeNode) bool {
	if root == nil {
		return false
	}
	if root.Val == 1 {
		return true
	}
	return containOne(root.Left) || containOne(root.Right)
}
