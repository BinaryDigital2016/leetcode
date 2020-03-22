package tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/*
给定一个有相同值的二叉搜索树（BST），找出 BST 中的所有众数（出现频率最高的元素）。

假定 BST 有如下定义：

    结点左子树中所含结点的值小于等于当前结点的值
    结点右子树中所含结点的值大于等于当前结点的值
    左子树和右子树都是二叉搜索树

例如：
给定 BST [1,null,2,2],

   1
    \
     2
    /
   2

返回[2].

提示：如果众数超过1个，不需考虑输出顺序
*/

func findMode(root *TreeNode) []int {
	ret := make([]int, 0)
	if root == nil {
		return ret
	}
	var pre *TreeNode
	maxCount, curCount := 0, 1
	inorder(root, &pre, &maxCount, &curCount, &ret)
	return ret
}

// pre在递归过程中需要被修改，所以得是指针传递
func inorder(root *TreeNode, pre **TreeNode, maxCount, curCount *int, ret *[]int) {
	if root == nil {
		return
	}

	inorder(root.Left, pre, maxCount, curCount, ret)
	if *pre != nil {
		if root.Val == (*pre).Val {
			*curCount++
		} else {
			*curCount = 1
		}
	}
	if *curCount == *maxCount {
		*ret = append(*ret, root.Val)
	} else if *curCount > *maxCount {
		*ret = []int{root.Val}
		*maxCount = *curCount
	}
	*pre = root
	inorder(root.Right, pre, maxCount, curCount, ret)
}
