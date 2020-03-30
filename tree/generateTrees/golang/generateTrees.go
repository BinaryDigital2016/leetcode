package tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 递归
func generateTrees(n int) []*TreeNode {
	if n <= 0 {
		return []*TreeNode{}
	}
	return generate(1, n)
}

func generate(start, end int) []*TreeNode {
	curRet := make([]*TreeNode, 0)
	if start > end {
		curRet = append(curRet, nil)
		return curRet
	}
	for i := start; i <= end; i++ { //每个点作为根节点
		leftList := generate(start, i-1) //左边的为左子树，递归
		rightList := generate(i+1, end)  //右边的为右子树，递归
		for _, v := range leftList {     //组合左右子树所有可能的结果
			for _, vv := range rightList {
				curRoot := TreeNode{Val: i}
				curRoot.Left = v
				curRoot.Right = vv
				curRet = append(curRet, &curRoot)
			}
		}
	}
	return curRet
}
