package dp

/*
给你一个整数数组 nums ，请你找出数组中乘积最大的连续子数组（该子数组中至少包含一个数字）。



示例 1:

输入: [2,3,-2,4]
输出: 6
解释: 子数组 [2,3] 有最大乘积 6。

示例 2:

输入: [-2,0,-1]
输出: 0
解释: 结果不能为 2, 因为 [-2,-1] 不是子数组。
*/

func maxProduct(nums []int) int {
	imax, imin := 1, 1
	ret := -1 << 31
	for i := 0; i < len(nums); i++ {
		if nums[i] < 0 {
			imax, imin = imin, imax //当前值为负数，交换最大和最小
		}
		imax = max(imax*nums[i], nums[i])
		imin = min(imin*nums[i], nums[i])
		ret = max(imax, ret)
	}
	return ret
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
