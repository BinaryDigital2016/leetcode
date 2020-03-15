package tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func RangeSumBST(root *TreeNode, L int, R int) int {
	if root == nil || L > R {
		return 0
	}

	if root.Val < L {
		return RangeSumBST(root.Right, L, R)
	} else if root.Val > R {
		return RangeSumBST(root.Left, L, R)
	} else {
		return root.Val + RangeSumBST(root.Right, L, R) + RangeSumBST(root.Left, L, R)
	}
}
