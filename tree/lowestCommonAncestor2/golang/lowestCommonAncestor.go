package tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/*
给定一个二叉树, 找到该树中两个指定节点的最近公共祖先。

百度百科中最近公共祖先的定义为：“对于有根树 T 的两个结点 p、q，最近公共祖先表示为一个结点 x，满足 x 是 p、q 的祖先且 x 的深度尽可能大（一个节点也可以是它自己的祖先）。”

例如，给定如下二叉搜索树:  root = [6,2,8,0,4,7,9,null,null,3,5]



示例 1:

输入: root = [6,2,8,0,4,7,9,null,null,3,5], p = 2, q = 8
输出: 6
解释: 节点 2 和节点 8 的最近公共祖先是 6。

示例 2:

输入: root = [6,2,8,0,4,7,9,null,null,3,5], p = 2, q = 4
输出: 2
解释: 节点 2 和节点 4 的最近公共祖先是 2, 因为根据定义最近公共祖先节点可以为节点本身。



说明:

    所有节点的值都是唯一的。
    p、q 为不同节点且均存在于给定的二叉搜索树中。
*/

/**
 * Definition for TreeNode.
 * type TreeNode struct {
 *     Val int
 *     Left *ListNode
 *     Right *ListNode
 * }
 */

/*
当我们用递归去做这个题时不要被题目误导，应该要明确一点
这个函数的功能有三个：给定两个节点 p 和 q

    如果 p 和 q 都存在，则返回它们的公共祖先；
    如果只存在一个，则返回存在的一个；
    如果 p 和 q 都不存在，则返回NULL

本题说给定的两个节点都存在，那自然还是能用上面的函数来解决

具体思路：
（1） 如果当前结点 root 等于NULL，则直接返回NULL
（2） 如果 root 等于 p 或者 q ，那这棵树一定返回 p 或者 q
（3） 然后递归左右子树，因为是递归，使用函数后可认为左右子树已经算出结果，用 left 和 right 表示
（4） 此时若left为空，那最终结果只要看 right；若 right 为空，那最终结果只要看 left
（5） 如果 left 和 right 都非空，因为只给了 p 和 q 两个结点，都非空，说明一边一个，因此 root 是他们的最近公共祖先
（6） 如果 left 和 right 都为空，则返回空（其实已经包含在前面的情况中了）
*/
// 递归
//func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
//	if root == nil || root == p || root == q {
//		return root
//	}
//
//	left := lowestCommonAncestor(root.Left, p, q)
//	right := lowestCommonAncestor(root.Right, p, q)
//	if left == nil {
//		return right
//	}
//	if right == nil {
//		return left
//	}
//	return root
//}

/*
我们使用DFS搜索每一个节点的左右子树：
1、若子树上存在p和q的公共节点，返回此公共节点
2、若不存在公共节点，但是存在p或q任意一个节点，返回此节点
3、若不存在公共、p、q节点，则返回null。

那么，有以下几个结论：
0、若当前节点为null、p、q之一，直接返回当前节点
1、若左子树上存在公共节点（返回值非p、q），则函数返回值为左子树返回值，不需再遍历右子树
2、若左子树返回null，则函数返回值为右子树返回值
3、若左子树、右子树返回值均为非null，则肯定为一个p，一个q，则公共节点为当前节点
4、最后一种情况，即左子树返回非null，右子树返回null，则函数返回值为左子树返回值
*/
// 递归+剪枝
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil || root == p || root == q {
		return root
	}

	left := lowestCommonAncestor(root.Left, p, q)
	if left != nil && left != p && left != q {
		return left
	}

	right := lowestCommonAncestor(root.Right, p, q)
	if left != nil && right != nil {
		return root
	}

	if left != nil {
		return left
	}
	return right
}
