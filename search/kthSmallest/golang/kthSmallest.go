package search

/*
给定一个 n x n 矩阵，其中每行和每列元素均按升序排序，找到矩阵中第k小的元素。
请注意，它是排序后的第 k 小元素，而不是第 k 个不同的元素。



示例:

matrix = [
   [ 1,  5,  9],
   [10, 11, 13],
   [12, 13, 15]
],
k = 8,

返回 13。



提示：
你可以假设 k 的值永远是有效的, 1 ≤ k ≤ n2 。
*/

/*
解析详见：
https://leetcode-cn.com/problems/kth-smallest-element-in-a-sorted-matrix/solution/er-fen-chao-ji-jian-dan-by-jacksu1024/
*/
func kthSmallest(matrix [][]int, k int) int {
	row := len(matrix)
	col := len(matrix[0])
	left := matrix[0][0]
	right := matrix[row-1][col-1]

	for left < right {
		mid := left + (right-left)/2
		cnt := countNotBiggerThan(matrix, mid, row, col)

		if cnt >= k {
			right = mid
		} else {
			left = mid + 1
		}
	}

	return left
}

func countNotBiggerThan(matrix [][]int, target int, row, col int) int {
	i := row - 1
	j := 0
	cnt := 0
	for i >= 0 && j < col {
		if matrix[i][j] <= target {
			cnt += i + 1
			j++
		} else {
			i--
		}
	}
	return cnt
}
