package string

/*
给定一个仅包含大小写字母和空格 ' ' 的字符串 s，返回其最后一个单词的长度。

如果字符串从左向右滚动显示，那么最后一个单词就是最后出现的单词。

如果不存在最后一个单词，请返回 0 。

说明：一个单词是指仅由字母组成、不包含任何空格的 最大子字符串。



示例:

输入: "Hello World"
输出: 5
*/
// func lengthOfLastWord(s string) int {
//     i := strings.LastIndex(s," ")
//     if i < 0{
//         return len(s)
//     }
//     if i == len(s) - 1{ //空格在最后，去掉空格后重新计算，也可以使用strings.TrimRight()去掉后面空格
//         return lengthOfLastWord(s[:len(s)-1])
//     }
//     return len(s)-i-1
// }

func lengthOfLastWord(s string) int {
	n := len(s)
	ret := 0
	for i := n - 1; i >= 0; i-- {
		if string(s[i]) == " " {
			if ret == 0 {
				continue
			}
			break
		}
		ret += 1
	}
	return ret
}
