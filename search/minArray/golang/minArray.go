package search

/*
把一个数组最开始的若干个元素搬到数组的末尾，我们称之为数组的旋转。输入一个递增排序的数组的一个旋转，输出旋转数组的最小元素。例如，数组 [3,4,5,1,2] 为 [1,2,3,4,5] 的一个旋转，该数组的最小值为1。

示例 1：

输入：[3,4,5,1,2]
输出：1

示例 2：

输入：[2,2,2,0,1]
输出：0
*/

/*
https://leetcode-cn.com/problems/xuan-zhuan-shu-zu-de-zui-xiao-shu-zi-lcof/solution/mian-shi-ti-11-xuan-zhuan-shu-zu-de-zui-xiao-shu-3/
*/

func minArray(numbers []int) int {
	i, j := 0, len(numbers)-1
	if len(numbers) == 0 {
		return 0
	}

	for i < j {
		m := i + (j-i)/2
		if numbers[m] > numbers[j] {
			i = m + 1
		} else if numbers[m] < numbers[j] {
			j = m
		} else {
			j -= 1
		}
	}

	return numbers[i]
}
