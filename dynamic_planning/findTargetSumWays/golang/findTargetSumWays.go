package dp

/*
给定一个非负整数数组，a1, a2, ..., an, 和一个目标数，S。现在你有两个符号 + 和 -。对于数组中的任意一个整数，你都可以从 + 或 -中选择一个符号添加在前面。

返回可以使最终数组和为目标数 S 的所有添加符号的方法数。

示例 1:

输入: nums: [1, 1, 1, 1, 1], S: 3
输出: 5
解释:

-1+1+1+1+1 = 3
+1-1+1+1+1 = 3
+1+1-1+1+1 = 3
+1+1+1-1+1 = 3
+1+1+1+1-1 = 3

一共有5种方法让最终目标和为3。

注意:

    数组非空，且长度不会超过20。
    初始的数组的和不会超过1000。
    保证返回的最终结果能被32位整数存下。
*/

func findTargetSumWays(nums []int, S int) int {
	n := len(nums)
	sum := 0
	for i := 0; i < n; i++ {
		sum += nums[i]
	}
	if S > sum || S < -sum {
		return 0
	}

	dp := make([][]int, n) //dp[i][j]表示nums[0,...,i]组成和为j-sum的方法数

	dp[0] = make([]int, 2*sum+1) //防止-S越界
	dp[0][nums[0]+sum] = 1
	dp[0][-nums[0]+sum] += 1 //nums[0]==0有两种组成方法，+0和-0
	for i := 1; i < n; i++ {
		if len(dp[i]) == 0 {
			dp[i] = make([]int, 2*sum+1)
		}
		for j := 0; j < 2*sum+1; j++ {
			l, r := 0, 0
			if j-nums[i] >= 0 { //防止越界
				l = j - nums[i]
			}
			if j+nums[i] < 2*sum+1 {
				r = j + nums[i]
			}
			dp[i][j] = dp[i-1][l] + dp[i-1][r]
		}
	}

	return dp[n-1][S+sum]
}
