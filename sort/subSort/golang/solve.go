package mysort

/*
给定一个整数数组，编写一个函数，找出索引m和n，只要将索引区间[m,n]的元素排好序，整个数组就是有序的。注意：n-m尽量最小，也就是说，找出符合条件的最短序列。函数返回值为[m,n]，若不存在这样的m和n（例如整个数组是有序的），请返回[-1,-1]。

示例：

输入： [1,2,4,7,10,11,7,12,6,7,16,18,19]
输出： [3,9]

提示：

    0 <= len(array) <= 1000000
*/

/*
进行两次遍历，一次从左到右找出N，一次从右到左找出M
（1）从左到右找出N
如果当前元素小于之前的最大元素则说明当前元素应处于[M N]无序序列中而且当前元素是当前最大下标的无序元素所以更新N为当前元素下标。
（2）从右到左找出M
如果当前元素大于之前的最小元素则说明当前元素应处于[M N]无序序列中而且当前元素是当前最小下标的无序元素所以更新M为当前元素下标
*/

func SubSort(array []int) []int {
	size := len(array)
	if size < 2 {
		return []int{-1, -1}
	}
	max := array[0]
	n := 0
	//找乱序的最大下标
	for i := 1; i < size; i++ {
		if array[i] >= max {
			max = array[i]
		} else {
			n = i
		}
	}

	min := array[size-1]
	m := len(array) - 1
	//找乱序的最大下标
	for i := size - 2; i >= 0; i-- {
		if array[i] <= min {
			min = array[i]
		} else {
			m = i
		}
	}

	if m >= n {
		return []int{-1, -1}
	}

	return []int{m, n}
}
