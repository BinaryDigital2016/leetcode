package mysort

import (
	"sort"
	"strconv"
	"strings"
)

/*
给定一组非负整数，重新排列它们的顺序使之组成一个最大的整数。

示例 1:

输入: [10,2]
输出: 210

示例 2:

输入: [3,30,34,5,9]
输出: 9534330

说明: 输出结果可能非常大，所以你需要返回一个字符串而不是整数。
*/

func largestNumber(nums []int) string {
	strSlice := make([]string, len(nums))
	allZero := true
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			allZero = false
		}
		strSlice[i] = strconv.Itoa(nums[i])
	}

	if allZero {
		return "0"
	}

	myStrSlc := MyStringSlice(strSlice)
	sort.Sort(myStrSlc)
	return strings.Join(myStrSlc, "")
}

type MyStringSlice []string

func (m MyStringSlice) Len() int {
	return len(m)
}

func (m MyStringSlice) Swap(i, j int) {
	m[i], m[j] = m[j], m[i]
}

// 逆序排列
func (m MyStringSlice) Less(i, j int) bool {
	s1 := m[i] + m[j]
	s2 := m[j] + m[i]
	return strings.Compare(s1, s2) == 1
}
