package search

/*
稀疏数组搜索。有个排好序的字符串数组，其中散布着一些空字符串，编写一种方法，找出给定字符串的位置。

示例1:

 输入: words = ["at", "", "", "", "ball", "", "", "car", "", "","dad", "", ""], s = "ta"
 输出：-1
 说明: 不存在返回-1。

示例2:

 输入：words = ["at", "", "", "", "ball", "", "", "car", "", "","dad", "", ""], s = "ball"
 输出：4
*/

func findString(words []string, s string) int {
	ret := -1
	search(words, s, 0, len(words)-1, &ret)
	return ret
}

func search(words []string, s string, l, r int, res *int) {
	if l <= r {
		m := l + (r-l)/2
		tm := m
		for m < r && words[m] == "" {
			m++
		}

		if words[m] == s {
			*res = m
			return
		}

		if words[m] < s {
			search(words, s, m+1, r, res)
			return
		}

		search(words, s, l, tm-1, res)
	}
}
