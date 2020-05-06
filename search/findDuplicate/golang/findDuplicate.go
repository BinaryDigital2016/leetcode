package search

/*
给定一个包含 n + 1 个整数的数组 nums，其数字都在 1 到 n 之间（包括 1 和 n），可知至少存在一个重复的整数。假设只有一个重复的整数，找出这个重复的数。

示例 1:

输入: [1,3,4,2,2]
输出: 2

示例 2:

输入: [3,1,3,4,2]
输出: 3

说明：

    不能更改原数组（假设数组是只读的）。
    只能使用额外的 O(1) 的空间。
    时间复杂度小于 O(n2) 。
    数组中只有一个重复的数字，但它可能不止重复出现一次。
*/

func findDuplicate(nums []int) int {
	size := len(nums)
	left, right := 1, size-1 //1,2,...,n
	for left < right {
		mid := left + (right-left)/2

		cnt := 0
		for _, v := range nums {
			if v <= mid {
				cnt++
			}
		}

		// 根据抽屉原理，小于等于 4 的数的个数如果严格大于 4 个，
		// 此时重复元素一定出现在 [1, 4] 区间里
		if cnt > mid {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return left
}
