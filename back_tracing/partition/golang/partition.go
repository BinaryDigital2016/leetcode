package bt

/*
给定一个字符串 s，将 s 分割成一些子串，使每个子串都是回文串。

返回 s 所有可能的分割方案。

示例:

输入: "aab"
输出:
[
  ["aa","b"],
  ["a","a","b"]
]
*/

/*
aabb
先考虑在第 1 个位置切割，a | abb
把 a 加入到结果中 [a]

然后考虑 abb
先考虑在第 1 个位置切割，a | bb
把 a  加入到结果中 [a a]

然后考虑 bb
先考虑在第 1 个位置切割，b | b
把 b 加入到结果中 [a a b]

然后考虑 b
先考虑在第 1 个位置切割，b |
把 b 加入到结果中 [a a b b]

然后考虑空串
把结果加到最终结果中 [[a a b b]]

回溯到上一层
考虑 bb
考虑在第 2 个位置切割，bb |
把 bb 加入到结果中 [a a bb]

然后考虑 空串
把结果加到最终结果中 [[a a b b] [a a bb]]

然后继续回溯
*/

func partition(s string) [][]string {
	ret := make([][]string, 0)
	backtrace(s, 0, []string{}, &ret)
	return ret
}

func backtrace(s string, start int, trace []string, res *[][]string) {
	if start == len(s) {
		*res = append(*res, trace)
		return
	}

	for i := start; i < len(s); i++ {
		if !valid(s[start : i+1]) {
			continue
		}
		n := len(trace)
		trace = append(trace, s[start:i+1])
		backtrace(s, i+1, copy(trace), res)
		trace = trace[:n]
	}
}

func copy(a []string) []string {
	b := make([]string, len(a))
	for k, v := range a {
		b[k] = v
	}
	return b
}

func valid(s string) bool {
	i, j := 0, len(s)-1
	for i <= j {
		if s[i] != s[j] {
			return false
		}
		i++
		j--
	}
	return true
}
