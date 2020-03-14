package mysort

/*
每次都把当前区间最大的数移动到当前区间的最后面
假设当前区间为[0, i]，该区间内最大的数的下标为j
那么我们可以通过两次翻转达到目的：
1，j + 1，将该数移动到区间头部
2，i + 1，将该数移动到区间尾部

作者：da-li-wang
链接：https://leetcode-cn.com/problems/pancake-sorting/solution/c-jian-dan-shu-xue-fa-by-da-li-wang/
*/
func PancakeSort(A []int) []int {
	ret := make([]int, 0, len(A))
	for i := len(A) - 1; i > 0; i-- {
		j := maxIndex(A[:i+1])
		if j > 0 {
			reverse(A[:j+1])
			ret = append(ret, j+1)
		}
		reverse(A[:i+1])
		ret = append(ret, i+1)
	}
	return ret
}

func maxIndex(s []int) int {
	if len(s) == 0 {
		return 0
	}
	max := 0
	for k, v := range s {
		if v > s[max] {
			max = k
		}
	}
	return max
}

func reverse(s []int) {
	if len(s) == 0 {
		return
	}

	i, j := 0, len(s)-1
	for i < j {
		s[i], s[j] = s[j], s[i]
		i++
		j--
	}
}
