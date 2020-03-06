package search

/*
给定一个排序数组和一个目标值，在数组中找到目标值，并返回其索引。如果目标值不存在于数组中，返回它将会被按顺序插入的位置。

你可以假设数组中无重复元素。
示例 1:

输入: [1,3,5,6], 5
输出: 2

示例 2:

输入: [1,3,5,6], 2
输出: 1

示例 3:

输入: [1,3,5,6], 7
输出: 4

示例 4:

输入: [1,3,5,6], 0
输出: 0
*/
// func searchInsert(nums []int, target int) int {
//     pre := -1
//     for k,v := range nums{
//         if v < target {
//             pre = k
//             continue
//         }
//         if v == target{
//             return k
//         }
//         return pre+1
//     }

//     return len(nums)
// }

//

func searchInsert(nums []int, target int) int {
	i, j := 0, len(nums)
	for i < j {
		m := (i + j) / 2
		if nums[m] == target {
			return m
		}
		if nums[m] < target {
			i++
		} else {
			j--
		}
	}
	return i
}
