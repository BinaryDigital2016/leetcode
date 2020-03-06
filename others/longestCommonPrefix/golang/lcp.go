package others

import "fmt"

/*
编写一个函数来查找字符串数组中的最长公共前缀。

如果不存在公共前缀，返回空字符串 ""。

示例 1:

输入: ["flower","flow","flight"]
输出: "fl"

示例 2:

输入: ["dog","racecar","car"]
输出: ""
解释: 输入不存在公共前缀。

说明:

所有输入只包含小写字母 a-z 。
*/

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	result := ""
	s0 := strs[0]
	for k, i := range s0 {
		si := string(i)
		for _, s := range strs {
			// 判断长度，否则会panic，若s0使用长度最短str则不用判断
			if len(s) <= k || s0[k] != s[k] {
				return result
			}
		}
		result = fmt.Sprintf("%s%s", result, si)
	}

	return result
}
