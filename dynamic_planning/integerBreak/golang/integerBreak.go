package dp

/*
给定一个正整数 n，将其拆分为至少两个正整数的和，并使这些整数的乘积最大化。 返回你可以获得的最大乘积。

示例 1:

输入: 2
输出: 1
解释: 2 = 1 + 1, 1 × 1 = 1。

示例 2:

输入: 10
输出: 36
解释: 10 = 3 + 3 + 4, 3 × 3 × 4 = 36。
*/

func integerBreak(n int) int {
	if n <= 2 {
		return 1
	}
	dp := make([]int, n+1)
	for i := 3; i <= n; i++ {
		for j := 1; j < i; j++ {
			dp[i] = max(dp[i], max(j*(i-j), max(dp[j]*(i-j), j*dp[i-j]))) //j*(i-j)表示将i拆成i和i-j两段 j*dp[i-j]表示将在i中拆出j，剩下的i-j接着拆
		}
	}
	return dp[n]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
