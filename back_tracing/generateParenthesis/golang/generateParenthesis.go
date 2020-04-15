package bt

/*
括号。设计一种算法，打印n对括号的所有合法的（例如，开闭一一对应）组合。

说明：解集不能包含重复的子集。

例如，给出 n = 3，生成结果为：

[
  "((()))",
  "(()())",
  "(())()",
  "()(())",
  "()()()"
]*/

/*
左边第一个一定是左括号（根节点），然后往下递归遍历时，记录左括号和右括号数，左括号可以直接加（前提是还能放右括号），
放右括号前提是左括号数大于当前右括号数且还能放右括号。
这里面之所以可以遍历完所有情况是因为在每一个选择时都做了分支（就是所谓的解空间树），叶结点就是所有的结果
*/
func generateParenthesis(n int) []string {
	if n <= 0 {
		return []string{}
	}
	ret := make([]string, 0)
	dfs("(", 1, 0, n, &ret)
	return ret
}

func dfs(s string, lnum, rnum, n int, res *[]string) {
	if lnum+rnum == 2*n {
		*res = append(*res, s)
		return
	}

	if lnum < n {
		dfs(s+"(", lnum+1, rnum, n, res)
	}
	if rnum < n && lnum > rnum {
		dfs(s+")", lnum, rnum+1, n, res)
	}
}
