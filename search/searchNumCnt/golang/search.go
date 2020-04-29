package search

/*
统计一个数字在排序数组中出现的次数。



示例 1:

输入: nums = [5,7,7,8,8,10], target = 8
输出: 2

示例 2:

输入: nums = [5,7,7,8,8,10], target = 6
输出: 0



限制：

0 <= 数组长度 <= 50000
*/

func search(nums []int, target int) int {
	count := 0
	binarySearch(nums, target, 0, len(nums)-1, &count)
	return count
}

func binarySearch(nums []int, target, left, right int, count *int) {
	if left <= right {
		if nums[left] > target || nums[right] < target {
			return
		}

		mid := left + (right-left)/2
		if target == nums[mid] {
			(*count)++
			binarySearch(nums, target, left, mid-1, count)
			binarySearch(nums, target, mid+1, right, count)
		} else if target > nums[mid] {
			binarySearch(nums, target, mid+1, right, count)
		} else {
			binarySearch(nums, target, left, mid-1, count)
		}
	}
}
