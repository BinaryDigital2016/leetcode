package tree

/*
将一个按照升序排列的有序数组，转换为一棵高度平衡二叉搜索树。

本题中，一个高度平衡二叉树是指一个二叉树每个节点 的左右两个子树的高度差的绝对值不超过 1。

示例:

给定有序数组: [-10,-3,0,5,9],

一个可能的答案是：[0,-3,9,-10,null,5]，它可以表示下面这个高度平衡二叉搜索树：

      0
     / \
   -3   9
   /   /
 -10  5
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 二叉搜索树的中序遍历为有序数组，该题为二叉搜索树的逆过程，答案不唯一
func SortedArrayToBST(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}

	p := len(nums) / 2
	root := &TreeNode{Val: nums[p]}
	root.Left = SortedArrayToBST(nums[:p])
	root.Right = SortedArrayToBST(nums[p+1:])

	return root

}
