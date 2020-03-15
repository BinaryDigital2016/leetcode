package tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func kthLargest(root *TreeNode, k int) int {
	s := make([]int,0)
	dfs(root, &s)
	return s[k-1]
}

func dfs(root *TreeNode, s *[]int){
	if root == nil {
		return
	}
	dfs(root.Right,s)
	*s = append(*s, root.Val)
	dfs(root.Left,s)
}
