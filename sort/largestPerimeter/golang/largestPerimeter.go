package sort

// 选择A中满足a+b>c的最大的a,b,c

//// 先排序，再比较
//func largestPerimeter(A []int) int {
//	tmp := sort.IntSlice(A)
//	sort.Sort(tmp)
//	for i := len(A) - 3; i >= 0; i-- {
//		if A[i]+A[i+1] > A[i+2] {
//			return A[i] + A[i+1] + A[i+2]
//		}
//	}
//
//	return 0
//
//}

// 排序的过程中比较，冒泡
func largestPerimeter(A []int) int {
	n := len(A)
	for i := 0; i < n; i++ {
		for j := 0; j < n-i-1; j++ {
			if A[j] > A[j+1] {
				A[j], A[j+1] = A[j+1], A[j]
			}
		}

		if i >= 2 {
			if A[n-i-1]+A[n-i] > A[n-i+1] {
				return A[n-i-1] + A[n-i] + A[n-i+1]
			}
		}
	}

	return 0
}
