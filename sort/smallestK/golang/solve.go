package mysort

//func SmallestK(arr []int, k int) []int {
//	partition(arr, 0, len(arr)-1, k)
//	return arr[:k]
//}
//
//func partition(arr []int, l, r, k int) {
//	if l >= r {
//		return
//	}
//
//	p := arr[l]
//	ll, rr := l, r
//	for ll < rr {
//		for ll < rr && arr[rr] >= p {
//			rr--
//		}
//		arr[ll] = arr[rr]
//
//		for ll < rr && arr[ll] <= p {
//			ll++
//		}
//		arr[rr] = arr[ll]
//	}
//	arr[ll] = p
//
//	if k < ll-l+1 {
//		partition(arr, l, ll, k)
//	} else {
//		partition(arr, ll+1, r, k-ll+l-1)
//	}
//}

func getLeastNumbers(arr []int, k int) []int {
	if k >= len(arr) {
		return arr
	}

	if k <= 0 {
		return []int{}
	}

	qsort(arr, 0, len(arr)-1, k)
	return arr[:k]
}

func qsort(arr []int, l, r, k int) {
	if l < r {
		index := partition(arr, l, r)
		if index-l+1 == k {
			return
		} else if index-l+1 > k {
			qsort(arr, l, index, k)
		} else {
			qsort(arr, index+1, r, k-index+l-1)
		}
	}
}

func partition(arr []int, l, r int) int {
	pivot := arr[l]
	i, j := l, r
	for i < j {
		for i < j && arr[j] >= pivot {
			j--
		}
		arr[i] = arr[j]

		for i < j && arr[i] <= pivot {
			i++
		}
		arr[j] = arr[i]
	}
	arr[i] = pivot

	return i
}
