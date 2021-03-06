package dp

/*
给定一个无序的整数数组，找到其中最长上升子序列的长度。

示例:

输入: [10,9,2,5,3,7,101,18]
输出: 4
解释: 最长的上升子序列是 [2,3,7,101]，它的长度是 4。

说明:

    可能会有多种最长上升子序列的组合，你只需要输出对应的长度即可。
    你算法的时间复杂度应该为 O(n2) 。
*/

func lengthOfLIS(nums []int) int {
	n := len(nums)
	if n < 2 {
		return n
	}
	dp := make([]int, n) //dp[i]表示以nums[i]结尾的最长上升子序列的长度
	for i := 0; i < n; i++ {
		dp[i] = 1 //最初的最长上升子序列长度为1
	}

	for i := 1; i < n; i++ {
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] { //当前值更大
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
	}

	ret := 0
	for i := 0; i < n; i++ {
		if dp[i] > ret {
			ret = dp[i]
		}
	}
	return ret
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
