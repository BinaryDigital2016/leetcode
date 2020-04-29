package search

/*
魔术索引。 在数组A[0...n-1]中，有所谓的魔术索引，满足条件A[i] = i。给定一个有序整数数组，编写一种方法找出魔术索引，若有的话，在数组A中找出一个魔术索引，如果没有，则返回-1。若有多个魔术索引，返回索引值最小的一个。

示例1:

 输入：nums = [0, 2, 3, 4, 5]
 输出：0
 说明: 0下标的元素为0

示例2:

 输入：nums = [1, 1, 1]
 输出：1
*/

func findMagicIndex(nums []int) int {
	ret := -1
	search(nums, 0, len(nums)-1, &ret)
	return ret
}

func search(nums []int, l, r int, res *int) {
	// 剪枝
	// if *res != -1 && l >= *res{
	//     return
	// }
	if l <= r {
		m := l + (r-l)/2
		if nums[m] == m {
			if *res == -1 {
				*res = m
			} else {
				if m < *res {
					*res = m
				}
			}
			search(nums, l, m-1, res)
		} else {
			search(nums, m+1, r, res)
			search(nums, l, m-1, res)
		}
	}
}
