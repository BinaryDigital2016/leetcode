package bt

import (
	"strconv"
	"strings"
)

/*
给定一个只包含数字的字符串，复原它并返回所有可能的 IP 地址格式。

示例:

输入: "25525511135"
输出: ["255.255.11.135", "255.255.111.35"]
*/
func restoreIpAddresses(s string) []string {
	ret := make([]string, 0)
	if len(s) < 4 || len(s) > 12 {
		return ret
	}
	backtrace(s, 0, 4, "", &ret)
	return ret
}

func backtrace(s string, start, remain int, trace string, res *[]string) {
	n := len(s)
	if remain == 0 {
		if start == n {
			*res = append(*res, trace[:len(trace)-1])
		}
		return
	}

	for j := 1; j < 4; j++ {
		if start+j > n || //越界
			n-start > remain*3 { //剩余位数过多
			break
		}
		sub := s[start : start+j]
		if len(sub) > 1 && strings.HasPrefix(sub, "0") { //0开头
			continue
		}

		subI, err := strconv.Atoi(sub)
		if err != nil { //非数字
			break
		}
		if subI > 255 {
			break
		}
		l := len(trace)
		trace = trace + sub + "."
		backtrace(s, start+j, remain-1, trace, res)
		trace = trace[:l]
	}
}
