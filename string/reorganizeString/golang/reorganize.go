package string

/*
给定一个字符串S，检查是否能重新排布其中的字母，使得两相邻的字符不同。

若可行，输出任意可行的结果。若不可行，返回空字符串。

示例 1:

输入: S = "aab"
输出: "aba"

示例 2:

输入: S = "aaab"
输出: ""

注意:

    S 只包含小写字母并且长度在[1, 500]区间内。
*/
func ReorganizeString(S string) string {
	n := len(S)
	arr := make([]int, 26)
	b := []byte(S)
	for _, v := range b {
		arr[v-byte('a')]++
		if arr[v-byte('a')] > (n+1)/2 {
			return ""
		}
	}

	bRet := make([]byte, n)
	even, odd := 0, 1
	for i := 0; i < 26; i++ {
		for arr[i] > 0 && arr[i] < n/2+1 && odd < n {
			bRet[odd] = byte(i + 'a')
			arr[i]--
			odd += 2
		}
		for arr[i] > 0 {
			bRet[even] = byte(i + 'a')
			arr[i]--
			even += 2
		}
	}
	return string(bRet)
}
