package dp

/*
给两个整数数组 A 和 B ，返回两个数组中公共的、长度最长的子数组的长度。

示例 1:

输入:
A: [1,2,3,2,1]
B: [3,2,1,4,7]
输出: 3
解释:
长度最长的公共子数组是 [3, 2, 1]。
*/

func findLength(A []int, B []int) int {
	m := len(A)
	n := len(B)
	if m == 0 || n == 0 {
		return 0
	}
	dp := make([][]int, m) //dp[i][j]表示A[0,...,i]与B[0,...,j]最长公共子数组长度
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
		if A[i] == B[0] {
			dp[i][0] = 1
		}
	}
	for i := 0; i < n; i++ {
		if A[0] == B[i] {
			dp[0][i] = 1
		}
	}
	ret := 0
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if A[i] == B[j] {
				dp[i][j] = dp[i-1][j-1] + 1
			}
			if ret < dp[i][j] {
				ret = dp[i][j]
			}
		}

	}
	return ret
}
