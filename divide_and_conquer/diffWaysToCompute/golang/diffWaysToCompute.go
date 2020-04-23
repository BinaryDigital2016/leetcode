package dc

import "strconv"

/*
给定一个含有数字和运算符的字符串，为表达式添加括号，改变其运算优先级以求出不同的结果。你需要给出所有可能的组合的结果。有效的运算符号包含 +, - 以及 * 。

示例 1:

输入: "2-1-1"
输出: [0, 2]
解释:
((2-1)-1) = 0
(2-(1-1)) = 2

示例 2:

输入: "2*3-4*5"
输出: [-34, -14, -10, -10, 10]
解释:
(2*(3-(4*5))) = -34
((2*3)-(4*5)) = -14
((2*(3-4))*5) = -10
(2*((3-4)*5)) = -10
(((2*3)-4)*5) = 10
*/

/*
分治算法三步走：

    分解：按运算符分成左右两部分，分别求解
    解决：实现一个递归函数，输入算式，返回算式解
    合并：根据运算符合并左右两部分的解，得出最终解
*/
func diffWaysToCompute(input string) []int {
	inputI, err := strconv.Atoi(input)
	if err == nil { //全是数字
		return []int{inputI}
	}

	ret := make([]int, 0)
	for k, v := range input {
		if v == '+' || v == '-' || v == '*' {
			left := diffWaysToCompute(input[:k])
			right := diffWaysToCompute(input[k+1:])
			for _, vl := range left {
				for _, vr := range right {
					switch v {
					case '+':
						ret = append(ret, vl+vr)
					case '-':
						ret = append(ret, vl-vr)
					case '*':
						ret = append(ret, vl*vr)
					}
				}
			}
		}
	}
	return ret
}
