package bt

/*
有重复字符串的排列组合。编写一种方法，计算某字符串的所有排列组合。

示例1:

 输入：S = "qqe"
 输出：["eqq","qeq","qqe"]

示例2:

 输入：S = "ab"
 输出：["ab", "ba"]
*/

func permutation(S string) []string {
	mret := make(map[string]struct{}, 0)
	selected := make([]int8, len(S))
	backtrack(S, []byte{}, selected, &mret)
	ret := make([]string, 0)
	for k := range mret {
		ret = append(ret, k)
	}
	return ret
}

func backtrack(s string, path []byte, selected []int8, res *map[string]struct{}) {
	if len(path) == len(s) {
		(*res)[string(path)] = struct{}{}
		return
	}

	for i := 0; i < len(s); i++ {
		if selected[i] == 1 {
			continue
		}
		selected[i] = 1
		n := len(path)
		path = append(path, s[i])
		backtrack(s, copy(path), selected, res)
		path = path[:n]
		selected[i] = 0
	}
}

func copy(a []byte) []byte {
	b := make([]byte, len(a))
	for i := 0; i < len(a); i++ {
		b[i] = a[i]
	}
	return b
}
