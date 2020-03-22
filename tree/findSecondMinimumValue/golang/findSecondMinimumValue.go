package tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/*
给定一个非空特殊的二叉树，每个节点都是正数，并且每个节点的子节点数量只能为 2 或 0。如果一个节点有两个子节点的话，那么这个节点的值不大于它的子节点的值。

给出这样的一个二叉树，你需要输出所有节点中的第二小的值。如果第二小的值不存在的话，输出 -1 。

示例 1:

输入:
    2
   / \
  2   5
     / \
    5   7

输出: 5
说明: 最小的值是 2 ，第二小的值是 5 。

示例 2:

输入:
    2
   / \
  2   2

输出: -1
说明: 最小的值是 2, 但是不存在第二小的值。
*/

// func findSecondMinimumValue(root *TreeNode) int {
//     if root == nil {
//         return -1
//     }
//     first,second := math.MaxUint32,math.MaxUint32
//     count := 1 //多少个不同的数
//     preorder(root, &first,&second,&count)
//     if count == 1{
//         return -1
//     }
//     return second
// }

// func preorder(root *TreeNode, first,second,count *int) {
//     if root == nil {
//         return
//     }
//     if root.Val<*first{
//         *second = *first
//         *first=root.Val
//     } else if root.Val>*first&&root.Val<=*second{
//         *second = root.Val
//         *count++
//     }
//     preorder(root.Left, first,second,count)
//     preorder(root.Right, first,second,count)
// }

func findSecondMinimumValue(root *TreeNode) int {
	if root == nil {
		return -1
	}
	return preorder(root, root.Val)
}

func preorder(root *TreeNode, min int) int {
	if root == nil {
		return -1
	}
	if root.Val > min {
		return root.Val //子节点一定大于当前节点，当前节点一定是第二小的
	}
	l := preorder(root.Left, min)
	r := preorder(root.Right, min)
	if l == -1 {
		return r
	}
	if r == -1 {
		return l
	}
	return findMin(l, r)
}

func findMin(a, b int) int {
	if a < b {
		return a
	}
	return b
}
