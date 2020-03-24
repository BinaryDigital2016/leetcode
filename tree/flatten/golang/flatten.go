package tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/*
给定一个二叉树，原地将它展开为链表。

例如，给定二叉树

    1
   / \
  2   5
 / \   \
3   4   6

将其展开为：

1
 \
  2
   \
    3
     \
      4
       \
        5
         \
          6
*/

func flatten(root *TreeNode) {
	var pre *TreeNode
	help(root, &pre)
}

func help(root *TreeNode, pre **TreeNode) {
	if root != nil {
		help(root.Right, pre)
		help(root.Left, pre)
		root.Right = *pre
		root.Left = nil
		*pre = root
	}
}
