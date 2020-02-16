/*
给定一个非负整数 numRows，生成杨辉三角的前 numRows 行。

在杨辉三角中，每个数是它左上方和右上方的数的和。

示例:

输入: 5
输出:
[
     [1],
    [1,1],
   [1,2,1],
  [1,3,3,1],
 [1,4,6,4,1]
]
*/

// func generate(numRows int) [][]int {
// 	if numRows == 0 {
// 		return [][]int{}
// 	}

// 	cur := []int{1}
// 	ret := [][]int{cur}
// 	i := 1
// 	for i < numRows {
// 		cur = append(cur, 0)
// 		for j := len(cur) - 1; j > 0; j-- {
// 			cur[j] = cur[j] + cur[j-1]
// 		}
// 		tmp := make([]int, len(cur))
// 		copy(tmp, cur)
// 		ret = append(ret, tmp)
// 		i++
// 	}
// 	return ret
// }

// 下一行由上一行数组左右各加一个0得到两个数组，然后对应相加
// 如 由0121 1210对应相加1331
func generate(numRows int) [][]int {
	if numRows == 0 {
		return [][]int{}
	}

	cur := []int{1}
	ret := [][]int{cur}
	i := 1
	for i < numRows {
		cur = append(cur, 0)
		for j := len(cur) - 1; j > 0; j-- {
			cur[j] = cur[j] + cur[j-1]
		}
		tmp := make([]int, len(cur))
		copy(tmp, cur)
		ret = append(ret, tmp)
		i++
	}
	return ret
}
