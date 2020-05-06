package main

/*
给定一个正整数 num，编写一个函数，如果 num 是一个完全平方数，则返回 True，否则返回 False。

说明：不要使用任何内置的库函数，如  sqrt。

示例 1：

输入：16
输出：True

示例 2：

输入：14
输出：Falsess
*/

func isPerfectSquare(num int) bool {
	if num < 2 {
		return true
	}

	left, right := 0, num/2
	for left <= right {
		m := left + (right-left)/2
		if m == num/m { //防止m*m溢出
			if num%m == 0 {
				return true
			} else {
				right = m - 1
			}
		} else if m < num/m {
			left = m + 1
		} else {
			right = m - 1
		}
	}
	return false
}
