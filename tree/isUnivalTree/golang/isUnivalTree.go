package tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isUnivalTree(root *TreeNode) bool {
	if root == nil || root.Left == nil && root.Right == nil {
		return true
	}

	return dfs(root, root.Val)
}

func dfs(root *TreeNode, x int) bool {
	if root == nil {
		return true
	}

	if root.Val != x {
		return false
	}

	return dfs(root.Left, x) && dfs(root.Right, x)
}
