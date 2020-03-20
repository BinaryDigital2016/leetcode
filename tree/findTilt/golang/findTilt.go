package tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func findTilt(root *TreeNode) int {
	ret := 0
	sum(root, &ret)
	return ret
}

func sum(root *TreeNode, tilt *int) int {
	if root == nil {
		return 0
	}

	leftSum := sum(root.Left, tilt)
	rightSum := sum(root.Right, tilt)
	*tilt += abs(leftSum, rightSum)
	return leftSum + rightSum + root.Val
}

func abs(a, b int) int {
	if a > b {
		return a - b
	}
	return b - a
}
