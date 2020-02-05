/*
给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串，判断字符串是否有效。

有效字符串需满足：

    左括号必须用相同类型的右括号闭合。
    左括号必须以正确的顺序闭合。

注意空字符串可被认为是有效字符串。

示例 1:

输入: "()"
输出: true

示例 2:

输入: "()[]{}"
输出: true

示例 3:

输入: "(]"
输出: false

示例 4:

输入: "([)]"
输出: false

示例 5:

输入: "{[]}"
输出: true
*/

func isValid(s string) bool {
	stack := make([]uint8, 0)
	for k, _ := range s {
		if isPush(s[k]) {
			stack = append(stack, s[k])
		} else {
			sLen := len(stack)
			if sLen <= 0 || !isMatch(stack[sLen-1], s[k]) {
				return false
			}
			stack = stack[:len(stack)-1]
		}
	}

	if len(stack) != 0 {
		return false
	}

	return true
}

func isPush(item uint8) bool {
	switch string(item) {
	case "(", "[", "{":
		return true
	default:
		return false
	}
}

func isMatch(i, j uint8) bool {
	switch string(i) {
	case "(":
		return string(j) == ")"
	case "[":
		return string(j) == "]"
	case "{":
		return string(j) == "}"
	default:
		return false
	}
}
