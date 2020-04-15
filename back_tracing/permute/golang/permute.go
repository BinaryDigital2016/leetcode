package bt

/*
给定一个 没有重复 数字的序列，返回其所有可能的全排列。

示例:

输入: [1,2,3]
输出:
[
  [1,2,3],
  [1,3,2],
  [2,1,3],
  [2,3,1],
  [3,1,2],
  [3,2,1]
]
*/

func permute(nums []int) [][]int {
	ret := make([][]int, 0)
	backtrack(nums, []int{}, &ret)
	return ret
}

func backtrack(nums, path []int, res *[][]int) {
	if len(path) == len(nums) {
		*res = append(*res, path)
		return
	}

	for i := 0; i < len(nums); i++ {
		if contains(path, nums[i]) {
			continue
		}
		n := len(path)
		path = append(path, nums[i])
		backtrack(nums, copy(path), res)
		path = path[:n]
	}
}

func contains(a []int, b int) bool {
	for _, v := range a {
		if v == b {
			return true
		}
	}
	return false
}

func copy(a []int) []int {
	b := make([]int, len(a))
	for i := 0; i < len(a); i++ {
		b[i] = a[i]
	}
	return b
}
