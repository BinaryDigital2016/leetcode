package tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/*
给定一个二叉树，它的每个结点都存放着一个整数值。

找出路径和等于给定数值的路径总数。

路径不需要从根节点开始，也不需要在叶子节点结束，但是路径方向必须是向下的（只能从父节点到子节点）。

二叉树不超过1000个节点，且节点数值范围是 [-1000000,1000000] 的整数。

示例：

root = [10,5,-3,3,2,null,11,3,-2,null,1], sum = 8

      10
     /  \
    5   -3
   / \    \
  3   2   11
 / \   \
3  -2   1

返回 3。和等于 8 的路径有:

1.  5 -> 3
2.  5 -> 2 -> 1
3.  -3 -> 11
*/

/*
根据题意，需要找出路径和等于给定值的路径总数，且起点不必从根结点开始，终点也不必在叶子结点，这意味着我们需要检查所有结点作为路径起点的情况（满足条件的路径的起点将可能是任意一个结点）。

换个角度，可以检查所有结点作为路径终点时的情况。即每遍历到一个结点，就以它为路径终点，向上回溯到根结点，同时在这个回溯的过程中不断累加顶点的和，相当于确定了终点，然后在向上回溯的过程中，把遇到的每个结点都尝试作为路径的起点，计算路径和，判断是否等于给定值，从而得到当前这个结点作为终点时的路径总数。

理解了上面这个过程，就可以得到递归代码了。

从根结点这棵二叉树开始，路径总数等于：根结点作为终点时的路径总数，加上左子树的路径总数，加上右子树的路径总数。
*/

func PathSum(root *TreeNode, sum int) int {
	s := make([]int, 1000)
	return findPathSum(root, s, sum, 0)
}

func findPathSum(root *TreeNode, s []int, sum int, layer int) int {
	if root == nil {
		return 0
	}
	cur, tmp := 0, 0
	s[layer] = root.Val
	for i := layer; i >= 0; i-- {
		tmp += s[i]
		if tmp == sum {
			cur++
		}
	}
	return cur + findPathSum(root.Left, s, sum, layer+1) + findPathSum(root.Right, s, sum, layer+1)
}
