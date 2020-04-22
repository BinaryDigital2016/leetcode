package bt

//N皇后问题

func solveNQueens(n int) [][]string {
	trace := make([][]byte, n)
	for i := 0; i < n; i++ {
		trace[i] = make([]byte, n)
		for j := 0; j < n; j++ {
			trace[i][j] = '.'
		}
	}

	ret := make([][]string, 0)
	backtrace(trace, 0, n, &ret)
	return ret
}

func backtrace(trace [][]byte, row, n int, res *[][]string) {
	if row == n {
		*res = append(*res, byte2string(trace))
		return
	}

	for col := 0; col < n; col++ {
		if !isValid(trace, row, col, n) {
			continue
		}
		trace[row][col] = 'Q'
		backtrace(trace, row+1, n, res)
		trace[row][col] = '.'
	}
}

func isValid(trace [][]byte, row, col, n int) bool {
	for i := 0; i < n; i++ { //列是否冲突
		if trace[i][col] == 'Q' {
			return false
		}
	}

	i := row - 1
	j := col - 1
	for i >= 0 && j >= 0 { //左上对角线是否冲突
		if trace[i][j] == 'Q' {
			return false
		}
		i--
		j--
	}

	i = row - 1
	j = col + 1
	for i >= 0 && j < n { //右上对角线是否冲突
		if trace[i][j] == 'Q' {
			return false
		}
		i--
		j++
	}

	return true
}

func byte2string(b [][]byte) []string {
	s := make([]string, len(b))
	for k, v := range b {
		s[k] = string(v)
	}
	return s
}
