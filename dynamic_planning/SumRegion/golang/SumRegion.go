package dp

/*
给定一个二维矩阵，计算其子矩形范围内元素的总和，该子矩阵的左上角为 (row1, col1) ，右下角为 (row2, col2)。

Range Sum Query 2D
上图子矩阵左上角 (row1, col1) = (2, 1) ，右下角(row2, col2) = (4, 3)，该子矩形内元素的总和为 8。

示例:

给定 matrix = [
  [3, 0, 1, 4, 2],
  [5, 6, 3, 2, 1],
  [1, 2, 0, 1, 5],
  [4, 1, 0, 1, 7],
  [1, 0, 3, 0, 5]
]

sumRegion(2, 1, 4, 3) -> 8
sumRegion(1, 1, 2, 2) -> 11
sumRegion(1, 2, 2, 4) -> 12

说明:

    你可以假设矩阵不可变。
    会多次调用 sumRegion 方法。
    你可以假设 row1 ≤ row2 且 col1 ≤ col2。
*/

type NumMatrix struct {
	dp [][]int //左上角和
}

func Constructor(matrix [][]int) NumMatrix {
	row := len(matrix)
	if row == 0 {
		return NumMatrix{}
	}
	col := len(matrix[0])
	if col == 0 {
		return NumMatrix{}
	}
	dp := make([][]int, row+1)
	for i := 0; i < row; i++ {
		if len(dp[i]) == 0 {
			dp[i] = make([]int, col+1)
		}
		if len(dp[i+1]) == 0 {
			dp[i+1] = make([]int, col+1)
		}
		for j := 0; j < col; j++ {
			dp[i+1][j+1] = dp[i+1][j] + dp[i][j+1] - dp[i][j] + matrix[i][j]
		}
	}
	return NumMatrix{dp}
}

func (this *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {
	if len(this.dp) == 0 {
		return 0
	}
	return this.dp[row2+1][col2+1] - this.dp[row2+1][col1] - this.dp[row1][col2+1] + this.dp[row1][col1]
}

/**
 * Your NumMatrix object will be instantiated and called as such:
 * obj := Constructor(matrix);
 * param_1 := obj.SumRegion(row1,col1,row2,col2);
 */
