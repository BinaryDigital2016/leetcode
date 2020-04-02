package dp

/*
给定一个字符串s，找到其中最长的回文子序列。可以假设s的最大长度为1000。

示例 1:
输入:

"bbbab"

输出:

4

一个可能的最长回文子序列为 "bbbb"。

示例 2:
输入:

"cbbd"

输出:

2

一个可能的最长回文子序列为 "bb"。
*/

func longestPalindromeSubseq(s string) int {
	n := len(s)
	if n == 0 {
		return n
	}
	dp := make([][]int, n) //dp[i][j]表示s[i,...,j]最长回文子序列的长度
	for i := 0; i < n; i++ {
		dp[i] = make([]int, n)
		dp[i][i] = 1
	}
	for i := n - 1; i >= 0; i-- {
		for j := i + 1; j < n; j++ {
			if s[i] == s[j] {
				dp[i][j] = dp[i+1][j-1] + 2
			} else {
				dp[i][j] = max(dp[i+1][j], dp[i][j-1])
			}
		}
	}
	return dp[0][n-1]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
