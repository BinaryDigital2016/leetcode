package bt

/*
找出所有相加之和为 n 的 k 个数的组合。组合中只允许含有 1 - 9 的正整数，并且每种组合中不存在重复的数字。

说明：

    所有数字都是正整数。
    解集不能包含重复的组合。

示例 1:

输入: k = 3, n = 7
输出: [[1,2,4]]

示例 2:

输入: k = 3, n = 9
输出: [[1,2,6], [1,3,5], [2,3,4]]
*/

func combinationSum3(k int, n int) [][]int {
	ret := make([][]int, 0)
	backtrace(k, n, []int{}, &ret)
	return ret
}

func backtrace(k, n int, trace []int, res *[][]int) {
	if len(trace) == k && n == 0 {
		*res = append(*res, trace)
		return
	}

	for i := 1; i < 10; i++ {
		if !valid(trace, i) || n < i {
			continue
		}
		l := len(trace)
		trace = append(trace, i)
		ctrace := make([]int, len(trace))
		copy(ctrace, trace)
		backtrace(k, n-i, ctrace, res)
		trace = trace[:l]
	}
}

func valid(a []int, b int) bool {
	for _, v := range a {
		if v == b { //已在列表中
			return false
		}
		if v > b {
			return false
		}
	}
	return true
}

// func copy(a []int) []int {
// 	b := make([]int, len(a))
// 	for i := 0; i < len(a); i++ {
// 		b[i] = a[i]
// 	}
// 	return b
// }

//func combinationSum3(k int, n int) [][]int {
//	ret := make([][]int, 0)
//	backtrace(k, n, []int{}, &ret)
//	return ret
//}
//
//func backtrace(k, n int, trace []int, res *[][]int) {
//	if len(trace) == k && n == 0 {
//		*res = append(*res, trace)
//		return
//	}
//
//	for i := 1; i < 10; i++ {
//		if !valid(trace, i) || n < i {
//			continue
//		}
//		l := len(trace)
//		trace = append(trace, i)
//		backtrace(k, n-i, copy(trace), res)
//		trace = trace[:l]
//	}
//}
//
//func valid(a []int, b int) bool {
//	for _, v := range a {
//		if v == b { //已在列表中
//			return false
//		}
//		if v > b {
//			return false
//		}
//	}
//	return true
//}
//
//func copy(a []int) []int {
//	b := make([]int, len(a))
//	for i := 0; i < len(a); i++ {
//		b[i] = a[i]
//	}
//	return b
//}
