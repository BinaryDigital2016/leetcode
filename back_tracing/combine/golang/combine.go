package bt

/*
给定两个整数 n 和 k，返回 1 ... n 中所有可能的 k 个数的组合。

示例:

输入: n = 4, k = 2
输出:
[
  [2,4],
  [3,4],
  [2,3],
  [1,2],
  [1,3],
  [1,4],
]
*/

func combine(n int, k int) [][]int {
	ret := make([][]int, 0)
	backtrace(n, k, []int{}, &ret)
	return ret
}

func backtrace(n, k int, trace []int, res *[][]int) {
	if len(trace) == k {
		*res = append(*res, trace)
		return
	}

	for i := 0; i < n; i++ {
		l := len(trace)
		if !valid(trace, i+1) {
			continue
		}
		trace = append(trace, i+1)
		backtrace(n, k, copy(trace), res)
		trace = trace[0:l]
	}
}

func valid(a []int, b int) bool {
	for _, v := range a {
		if v == b {
			return false
		}
		if v > b {
			return false
		}
	}
	return true
}

func copy(a []int) []int {
	b := make([]int, len(a))
	for i := 0; i < len(a); i++ {
		b[i] = a[i]
	}
	return b
}
