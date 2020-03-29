package dp

/*
给定一个包含非负整数的 m x n 网格，请找出一条从左上角到右下角的路径，使得路径上的数字总和为最小。

说明：每次只能向下或者向右移动一步。

示例:

输入:
[
  [1,3,1],
  [1,5,1],
  [4,2,1]
]
输出: 7
解释: 因为路径 1→3→1→1→1 的总和最小。
*/

func minPathSum(grid [][]int) int {
	row := len(grid)
	if row == 0 {
		return 0
	}
	col := len(grid[0])

	dp := make([][]int, row)
	//初始化边界
	dp[0] = make([]int, col)
	dp[0][0] = grid[0][0]
	for i := 1; i < row; i++ {
		if len(dp[i]) == 0 {
			dp[i] = make([]int, col)
		}
		dp[i][0] = dp[i-1][0] + grid[i][0]
	}
	for i := 1; i < col; i++ {
		dp[0][i] = dp[0][i-1] + grid[0][i]
	}

	for i := 1; i < row; i++ {
		for j := 1; j < col; j++ {
			dp[i][j] = min(dp[i-1][j], dp[i][j-1]) + grid[i][j]
		}
	}
	return dp[row-1][col-1]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
