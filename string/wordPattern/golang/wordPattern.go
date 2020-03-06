package string

/*
给定一种规律 pattern 和一个字符串 str ，判断 str 是否遵循相同的规律。

这里的 遵循 指完全匹配，例如， pattern 里的每个字母和字符串 str 中的每个非空单词之间存在着双向连接的对应规律。

示例1:

输入: pattern = "abba", str = "dog cat cat dog"
输出: true

示例 2:

输入:pattern = "abba", str = "dog cat cat fish"
输出: false

示例 3:

输入: pattern = "aaaa", str = "dog cat cat dog"
输出: false

示例 4:

输入: pattern = "abba", str = "dog dog dog dog"
输出: false

说明:
你可以假设 pattern 只包含小写字母， str 包含了由单个空格分隔的小写字母。
*/import "strings"

func wordPattern(pattern string, str string) bool {
	strs := strings.Split(str, " ")
	if len(pattern) != len(strs) {
		return false
	}

	n := len(pattern)
	m := make(map[uint8]string)
	mm := make(map[string]uint8)
	for i := 0; i < n; i++ {
		if s, ok := m[pattern[i]]; !ok {
			if _, ok2 := mm[strs[i]]; ok2 { //是否被其他pattern占用
				return false
			}
			m[pattern[i]] = strs[i]
			mm[strs[i]] = pattern[i]
		} else {
			if s != strs[i] {
				return false
			}
		}
	}
	return true
}
