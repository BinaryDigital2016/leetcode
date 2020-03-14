package mysort

import "sort"

// 排序
func WiggleSort(nums []int) {
	newNums := sort.IntSlice(nums)
	sort.Sort(newNums)
	n := len(newNums)
	n1 := n / 2
	if n%2 == 1 {
		n1++
	}
	pre := make([]int, n1)
	last := make([]int, n-n1)
	copy(last, newNums[n1:])
	copy(pre, newNums[:n1])
	j, k := n1-1, n-n1-1
	for i := 0; i < n; i++ {
		if i%2 == 0 {
			newNums[i] = pre[j]
			j--
		} else {
			newNums[i] = last[k]
			k--
		}
	}
}
