package string

import "fmt" /*
给定两个二进制字符串，返回他们的和（用二进制表示）。

输入为非空字符串且只包含数字 1 和 0。

示例 1:

输入: a = "11", b = "1"
输出: "100"

示例 2:

输入: a = "1010", b = "1011"
输出: "10101"
*/
func AddBinary(a string, b string) string {
	var carry uint8 = '0'
	n := 0
	n1, n2 := len(a), len(b)
	result := ""
	if n1 > n2 {
		for j := 0; j < n1-n2; j++ {
			b = fmt.Sprintf("%s%s", "0", b)
		}
		n = n1
	} else {
		for j := 0; j < n2-n1; j++ {
			a = fmt.Sprintf("%s%s", "0", a)
		}
		n = n2
	}
	for i := n - 1; i >= 0; i-- {
		cur := a[i] + b[i] - '0' + carry - '0'
		if cur >= '2' {
			carry = '1'
			result = fmt.Sprintf("%d%s", cur-'2', result)
		} else {
			carry = '0'
			result = fmt.Sprintf("%c%s", cur, result)
		}
	}
	if carry == '1' {
		result = fmt.Sprintf("%s%s", "1", result)
	}
	return result
}
