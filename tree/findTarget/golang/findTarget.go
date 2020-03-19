package tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func findTarget(root *TreeNode, k int) bool {
	if root == nil {
		return false
	}
	m := make(map[int]struct{})
	return find(root, k, m)
}

func find(root *TreeNode, k int, m map[int]struct{}) bool {
	if root == nil {
		return false
	}
	if _, ok := m[k-root.Val]; ok {
		return true
	}
	m[root.Val] = struct{}{}
	return find(root.Left, k, m) || find(root.Right, k, m)
}
