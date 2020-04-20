package bt

/*
给定一个无重复元素的数组 candidates 和一个目标数 target ，找出 candidates 中所有可以使数字和为 target 的组合。

candidates 中的数字可以无限制重复被选取。

说明：

    所有数字（包括 target）都是正整数。
    解集不能包含重复的组合。

示例 1:

输入: candidates = [2,3,6,7], target = 7,
所求解集为:
[
  [7],
  [2,2,3]
]

示例 2:

输入: candidates = [2,3,5], target = 8,
所求解集为:
[
  [2,2,2,2],
  [2,3,3],
  [3,5]
]
*/

func combinationSum(candidates []int, target int) [][]int {
	ret := make([][]int, 0)
	backtrace(target, candidates, 0, []int{}, &ret)
	return ret
}

func backtrace(target int, candidates []int, start int, trace []int, res *[][]int) {
	if target == 0 {
		*res = append(*res, trace)
		return
	}

	for i := start; i < len(candidates); i++ { //不往前 防止重复组合
		if target < candidates[i] {
			continue
		}
		n := len(trace)
		trace = append(trace, candidates[i])
		backtrace(target-candidates[i], candidates, i, copy(trace), res)
		trace = trace[:n]
	}
}

func copy(a []int) []int {
	b := make([]int, len(a))
	for k, v := range a {
		b[k] = v
	}
	return b
}
