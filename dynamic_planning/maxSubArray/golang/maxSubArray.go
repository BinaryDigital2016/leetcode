package dp

func maxSubArray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	ret := nums[0]
	sum := nums[0]
	for i := 1; i < len(nums); i++ {
		sum = max(sum+nums[i], nums[i])
		ret = max(ret, sum)
	}
	return ret
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
