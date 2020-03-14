package mysort

// 按对角线冒泡排序
func DiagonalSort(mat [][]int) [][]int {
	if len(mat) == 0 || len(mat[0]) == 0 {
		return mat
	}

	row := len(mat)
	col := len(mat[0])
	for i := 0; i < row-1; i++ {
		for j := 0; j < row-i-1; j++ {
			for k := 0; k < col-1; k++ {
				if mat[j][k] > mat[j+1][k+1] {
					mat[j][k], mat[j+1][k+1] = mat[j+1][k+1], mat[j][k]
				}
			}
		}
	}
	return mat
}
