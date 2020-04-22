package bt

/*
给定一个二维网格和一个单词，找出该单词是否存在于网格中。

单词必须按照字母顺序，通过相邻的单元格内的字母构成，其中“相邻”单元格是那些水平相邻或垂直相邻的单元格。同一个单元格内的字母不允许被重复使用。



示例:

board =
[
  ['A','B','C','E'],
  ['S','F','C','S'],
  ['A','D','E','E']
]

给定 word = "ABCCED", 返回 true
给定 word = "SEE", 返回 true
给定 word = "ABCB", 返回 false



提示：

    board 和 word 中只包含大写和小写英文字母。
    1 <= board.length <= 200
    1 <= board[i].length <= 200
    1 <= word.length <= 10^3
*/

func exist(board [][]byte, word string) bool {
	m := len(board)
	if m == 0 {
		return word == ""
	}
	n := len(board[0])

	visited := make([][]bool, m)
	for i := 0; i < m; i++ {
		visited[i] = make([]bool, n)
	}

	direction := [][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if backtrace(board, 0, i, j, visited, direction, word) {
				return true
			}
		}
	}
	return false
}

func backtrace(board [][]byte, start, x, y int, visited [][]bool, direction [][2]int, word string) bool {
	if start == len(word)-1 {
		return board[x][y] == word[start]
	}

	if board[x][y] == word[start] { //当前匹配，继续搜索
		visited[x][y] = true
		m := len(board)
		n := len(board[0])
		for _, v := range direction {
			nx := x + v[0]
			ny := y + v[1]
			if nx >= 0 && nx < m && ny >= 0 && ny < n &&
				!visited[nx][ny] &&
				backtrace(board, start+1, nx, ny, visited, direction, word) {
				return true
			}
		}
		visited[x][y] = false
	}
	return false
}
