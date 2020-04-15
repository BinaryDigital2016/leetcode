package bt

/*
无重复字符串的排列组合。编写一种方法，计算某字符串的所有排列组合，字符串每个字符均不相同。

示例1:

 输入：S = "qwe"
 输出：["qwe", "qew", "wqe", "weq", "ewq", "eqw"]

示例2:

 输入：S = "ab"
 输出：["ab", "ba"]

提示:

    字符都是英文字母。
    字符串长度在[1, 9]之间。
*/

func permutation(S string) []string {
	n := len(S)
	ret := make([]string, 0)
	if n == 0 {
		return ret
	}
	b := make([]byte, n)
	dfs(S, 0, b, &ret)
	return ret
}

func dfs(s string, start int, b []byte, res *[]string) {
	if start == len(s) {
		*res = append(*res, string(b))
		return
	}
	b[start] = s[start]
	dfs(s, start+1, b, res)
	for i := 0; i < start; i++ {
		dfs(s, start+1, change(b, i, start), res)
	}
}

func change(b []byte, i, j int) []byte {
	cb := make([]byte, len(b))
	for k := 0; k < len(b); k++ {
		if k == i {
			cb[k] = b[j]
		} else if k == j {
			cb[k] = b[i]
		} else {
			cb[k] = b[k]
		}
	}
	return cb
}
