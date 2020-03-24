package tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func kthLargest(root *TreeNode, k int) int {
	s := make([]int, 0)
	dfs(root, &s)
	return s[k-1]
}

func dfs(root *TreeNode, s *[]int) {
	if root == nil {
		return
	}
	dfs(root.Right, s)
	*s = append(*s, root.Val)
	dfs(root.Left, s)
}

// 迭代
//class Solution:
//	def kthSmallest(self, root, k):
//		stack = []
//		while True:
//			while root:
//				stack.append(root)
//				root = root.left
//			root = stack.pop()
//			k -= 1
//			if not k:
//				return root.val
//			root = root.right
