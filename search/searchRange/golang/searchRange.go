package search

/*
给定一个按照升序排列的整数数组 nums，和一个目标值 target。找出给定目标值在数组中的开始位置和结束位置。

你的算法时间复杂度必须是 O(log n) 级别。

如果数组中不存在目标值，返回 [-1, -1]。

示例 1:

输入: nums = [5,7,7,8,8,10], target = 8
输出: [3,4]

示例 2:

输入: nums = [5,7,7,8,8,10], target = 6
输出: [-1,-1]
*/

func searchRange(nums []int, target int) []int {
	n := len(nums)
	if n == 0 {
		return []int{-1, -1}
	}
	left := left_bound(nums, target, n)
	right := right_bound(nums, target, n)
	return []int{left, right}
}

func left_bound(nums []int, target, n int) int {
	left, right := 0, n-1
	for left <= right {
		mid := left + (right-left)/2
		if target > nums[mid] {
			left = mid + 1
		} else if target < nums[mid] {
			right = mid - 1
		} else {
			right = mid - 1 //必须-1，否则无法跳出循环。如果只有mid一个值等于target,那么最终left会等于right+1，也就是mid，不会被漏掉
		}
	}

	//1.target大于所有值，left一直右移到n;
	//2.target介于数组最值之间，但是不存在于数组中
	if left >= n || nums[left] != target { //没找到
		return -1
	}
	return left
}

func right_bound(nums []int, target, n int) int {
	left, right := 0, n-1
	for left <= right {
		mid := left + (right-left)/2
		if target > nums[mid] {
			left = mid + 1
		} else if target < nums[mid] {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}

	if right < 0 || nums[right] != target { //没找到
		return -1
	}
	return right
}
