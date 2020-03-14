package mysort

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// // 遍历后排序
// func GetAllElements(root1 *TreeNode, root2 *TreeNode) []int {
//     ret := make([]int,0)
//     traverse(root1, &ret)
//     traverse(root2, &ret)
//     tmp := sort.IntSlice(ret)
//     sort.Sort(tmp)
//     return ret
// }

// func traverse(root *TreeNode, ret *[]int){
//     if root == nil {
//         return
//     }

//     traverse(root.Left, ret)
//     *ret = append(*ret, root.Val)
//     traverse(root.Right, ret)
// }

// 遍历后合并列表
func GetAllElements(root1 *TreeNode, root2 *TreeNode) []int {
	s1 := make([]int, 0)
	s2 := make([]int, 0)
	traverse(root1, &s1)
	traverse(root2, &s2)
	ret := make([]int, 0, len(s1)+len(s2))
	i, j := 0, 0
	for i < len(s1) && j < len(s2) {
		if s1[i] < s2[j] {
			ret = append(ret, s1[i])
			i++
		} else {
			ret = append(ret, s2[j])
			j++
		}
	}
	if i < len(s1) {
		ret = append(ret, s1[i:]...)
	}
	if j < len(s2) {
		ret = append(ret, s2[j:]...)
	}
	return ret
}

func traverse(root *TreeNode, ret *[]int) {
	if root == nil {
		return
	}

	traverse(root.Left, ret)
	*ret = append(*ret, root.Val)
	traverse(root.Right, ret)
}
