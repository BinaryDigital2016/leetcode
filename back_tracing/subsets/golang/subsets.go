package bt

/*
给定一组不含重复元素的整数数组 nums，返回该数组所有可能的子集（幂集）。

说明：解集不能包含重复的子集。

示例:

输入: nums = [1,2,3]
输出:
[
  [3],
  [1],
  [2],
  [1,2,3],
  [1,3],
  [2,3],
  [1,2],
  []
]
*/

func subsets(nums []int) [][]int {
	ret := make([][]int, 0)
	backtrack(nums, 0, []int{}, &ret)
	return ret
}

func backtrack(nums []int, start int, track []int, res *[][]int) {
	*res = append(*res, track)
	for i := start; i < len(nums); i++ {
		track = append(track, nums[i])
		backtrack(nums, i+1, copy(track), res)
		track = track[:len(track)-1]
	}
}

func copy(a []int) []int {
	b := make([]int, len(a))
	for i := 0; i < len(a); i++ {
		b[i] = a[i]
	}
	return b
}
