package sort

import "strings"

/*
给定一个字符串和一个字符串字典，找到字典里面最长的字符串，该字符串可以通过删除给定字符串的某些字符来得到。如果答案不止一个，返回长度最长且字典顺序最小的字符串。如果答案不存在，则返回空字符串。

示例 1:

输入:
s = "abpcplea", d = ["ale","apple","monkey","plea"]

输出:
"apple"

示例 2:

输入:
s = "abpcplea", d = ["a","b","c"]

输出:
"a"

说明:

    所有输入的字符串只包含小写字母。
    字典的大小不会超过 1000。
    所有输入的字符串长度不会超过 1000。
*/

// // 排序后匹配
// func findLongestWord(s string, d []string) string {
//     t :=MySliceString(d)
//     sort.Sort(t)
//     for _,v := range d {
//         if isSubString(s,v){
//             return v
//         }
//     }
//     return ""
// }

// func isSubString(s string, sub string) bool {
//     j:=0
//     for i:=0;i<len(s)&&j<len(sub);i++{
//         if sub[j] == s[i] {
//             j++
//         }
//     }
//     return j == len(sub)
// }

// type MySliceString []string
// func (m MySliceString) Len() int {
//     return len(m)
// }

// func (m MySliceString) Swap(i,j int) {
//     m[i],m[j]=m[j],m[i]
// }

// func (m MySliceString) Less(i,j int) bool {
//     if len(m[i]) < len(m[j]) {
//         return false
//     } else if len(m[i]) > len(m[j]) {
//         return true
//     } else {
//         return strings.Compare(m[i],m[j]) == -1
//     }
// }

// 不排序
func findLongestWord(s string, d []string) string {
	ret := ""
	for _, v := range d {
		if isSubString(s, v) && isMatcher(v, ret) {
			ret = v
		}
	}
	return ret
}

func isSubString(s string, sub string) bool {
	j := 0
	for i := 0; i < len(s) && j < len(sub); i++ {
		if sub[j] == s[i] {
			j++
		}
	}
	return j == len(sub)
}

func isMatcher(s1, s2 string) bool {
	if len(s1) < len(s2) {
		return false
	} else if len(s1) > len(s2) {
		return true
	} else {
		return strings.Compare(s1, s2) == -1
	}
}
