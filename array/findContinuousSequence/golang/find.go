package array

/*
输入一个正整数 target ，输出所有和为 target 的连续正整数序列（至少含有两个数）。

序列内的数字由小到大排列，不同序列按照首个数字从小到大排列。
*/

// 暴力求解：枚举每个正整数，以其为起点往后求和，找出所有和为target的序列，略

// 滑动窗口(双指针))
func FindContinuousSequence(target int) [][]int {
	l, r, sum := 1, 2, 0
	ret := make([][]int, 0)
	for l < r { //至少两个数，不取等
		sum = (l + r) * (r - l + 1) / 2
		switch {
		case sum < target:
			r++
		case sum > target: //l为起点不存在满足条件的序列
			l++
		default:
			s := make([]int, r-l+1)
			for k := 0; k < r-l+1; k++ {
				s[k] = k + l
			}
			ret = append(ret, s)
			l++
		}
	}

	return ret
}
