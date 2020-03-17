package tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/*
给定一个二叉搜索树（Binary Search Tree），把它转换成为累加树（Greater Tree)，使得每个节点的值是原来的节点值加上所有大于它的节点值之和。



例如：

输入: 原始二叉搜索树:
              5
            /   \
           2     13

输出: 转换为累加树:
             18
            /   \
          20     13
*/

/*
二叉树的中序遍历是升序的，那么把中序遍历的左中右顺序改为右中左，则这样的遍历是降序的。
把二叉树的所有结点看作一个降序序列，降序遍历每个数，每次都把上一个数累加到下一个数中即可。
*/
func convertBST(root *TreeNode) *TreeNode {
	sum := 0
	convertBST2(root, &sum)
	return root
}

func convertBST2(root *TreeNode, sum *int) {
	if root != nil {
		convertBST2(root.Right, sum)
		*sum += root.Val
		root.Val = *sum
		convertBST2(root.Left, sum)
	}
}
