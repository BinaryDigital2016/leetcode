package array

import "sort"

/*
给定两个整数数组，请交换一对数值（每个数组中取一个数值），使得两个数组所有元素的和相等。

返回一个数组，第一个元素是第一个数组中要交换的元素，第二个元素是第二个数组中要交换的元素。若有多个答案，返回任意一个均可。若无满足条件的数值，返回空数组。

示例:

输入: array1 = [4, 1, 2, 1, 1, 2], array2 = [3, 6, 3, 3]
输出: [1, 3]

示例:

输入: array1 = [1, 2, 3], array2 = [4, 5, 6]
输出: []

提示：

    1 <= array1.length, array2.length <= 100000
*/

// 超时
// func findSwapValues(array1 []int, array2 []int) []int {
//     sum1 := sumList(array1)
//     sum2 := sumList(array2)
//     diff := sum1 - sum2
//     for _,v:=range array1{
//         for _,vv:=range array2{
//             if (v-vv)*2==diff{
//                 return []int{v,vv}
//             }
//         }
//     }
//     return []int{}
// }

// func sumList(l []int) int{
//     sum := 0
//     for _,v:=range l{
//         sum += v
//     }
//     return sum
// }

func FindSwapValues(array1 []int, array2 []int) []int {
	sum1 := sumList(array1)
	sum2 := sumList(array2)
	diff := sum1 - sum2
	if diff%2 != 0 {
		return []int{}
	}
	diff >>= 1
	sort.Sort(sort.IntSlice(array1))
	sort.Sort(sort.IntSlice(array2))
	i, j := 0, 0
	for i < len(array1) && j < len(array2) {
		if array1[i]-array2[j] == diff {
			return []int{array1[i], array2[j]}
		} else if array1[i]-array2[j] < diff {
			i++
		} else {
			j++
		}
	}
	return []int{}
}

func sumList(l []int) int {
	sum := 0
	for _, v := range l {
		sum += v
	}
	return sum
}
