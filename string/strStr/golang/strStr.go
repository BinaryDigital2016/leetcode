package string

/*
实现 strStr() 函数。

给定一个 haystack 字符串和一个 needle 字符串，在 haystack 字符串中找出 needle 字符串出现的第一个位置 (从0开始)。如果不存在，则返回  -1。

示例 1:

输入: haystack = "hello", needle = "ll"
输出: 2

示例 2:

输入: haystack = "aaaaa", needle = "bba"
输出: -1

说明:

当 needle 是空字符串时，我们应当返回什么值呢？这是一个在面试中很好的问题。

对于本题而言，当 needle 是空字符串时我们应当返回 0 。这与C语言的 strstr() 以及 Java的 indexOf() 定义相符。
*/

// func strStr(haystack string, needle string) int {
//     return strings.Index(haystack, needle)
// }

func strStr(haystack string, needle string) int {
	if len(needle) <= 0 || haystack == needle {
		return 0
	}

	isFind := false
	i := 0
	// for ;i<len(haystack) && !isFind;i++{
	//     for j:=0;j<len(needle);j++{
	//         if haystack[i] != needle[j]{
	//             break
	//         } else {
	//             if j == len(needle)-1{
	//                 isFind = true
	//             }
	//         }
	//     }
	// }

	for ; i <= len(haystack)-len(needle); i++ {
		if haystack[i:i+len(needle)] == needle {
			isFind = true
			break
		}
	}

	if isFind {
		return i
	}
	return -1
}
