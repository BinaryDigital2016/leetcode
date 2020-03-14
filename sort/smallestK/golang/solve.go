package mysort

func SmallestK(arr []int, k int) []int {
	partition(arr, 0, len(arr)-1, k)
	return arr[:k]
}

func partition(arr []int, l, r, k int) {
	if l >= r {
		return
	}

	p := arr[l]
	ll, rr := l, r
	for ll < rr {
		for ll < rr && arr[rr] >= p {
			rr--
		}
		arr[ll] = arr[rr]

		for ll < rr && arr[ll] <= p {
			ll++
		}
		arr[rr] = arr[ll]
	}
	arr[ll] = p

	if k < ll-l+1 {
		partition(arr, l, ll, k)
	} else {
		partition(arr, ll+1, r, k-ll+l-1)
	}
}
