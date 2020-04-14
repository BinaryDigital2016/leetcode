package dp

/*
给定一个字符串S，通过将字符串S中的每个字母转变大小写，我们可以获得一个新的字符串。返回所有可能得到的字符串集合。

示例:
输入: S = "a1b2"
输出: ["a1b2", "a1B2", "A1b2", "A1B2"]

输入: S = "3z4"
输出: ["3z4", "3Z4"]

输入: S = "12345"
输出: ["12345"]

注意：

    S 的长度不超过12。
    S 仅由数字和字母组成。
*/

func letterCasePermutation(S string) []string {
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
	if isLetter(s[start]) { //如果当前字符是字母，改变字母大小写后再处理一次
		b[start] = change(s, start)
		dfs(s, start+1, b, res)
	}
}

func isLetter(c byte) bool {
	if (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z') {
		return true
	}
	return false
}

func change(s string, i int) byte {
	b := []byte(s)
	if b[i] >= 'a' && b[i] <= 'z' {
		b[i] -= 32
	} else {
		b[i] += 32
	}
	return b[i]
}
