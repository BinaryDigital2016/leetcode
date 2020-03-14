package mysort

func MergeSort(s []int) {
	tmp := make([]int, len(s))
	doMergeSort(s, 0, len(s)-1, tmp)
}

func doMergeSort(s []int, l, r int, tmp []int) {
	if l < r {
		m := l + (r-l)>>1
		doMergeSort(s, l, m, tmp)
		doMergeSort(s, m+1, r, tmp)
		merge(s, l, m, r, tmp)
	}
}

func merge(s []int, l, m, r int, tmp []int) {
	b1, e1 := l, m
	b2, e2 := m+1, r
	k := l
	for b1 <= e1 && b2 <= e2 {
		if s[b1] < s[b2] {
			tmp[k] = s[b1]
			b1++
		} else {
			tmp[k] = s[b2]
			b2++
		}
		k++
	}

	for b1 <= e1 {
		tmp[k] = s[b1]
		b1++
		k++
	}

	for b2 <= e2 {
		tmp[k] = s[b2]
		b2++
		k++
	}

	k = l
	for l <= r {
		s[l] = tmp[k]
		l++
		k++
	}
}
