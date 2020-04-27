package others

func intersection(nums1 []int, nums2 []int) []int {
	if len(nums1) == 0 || len(nums2) == 0 {
		return []int{}
	}
	m := make(map[int]int8)
	for _, v := range nums1 {
		c, ok := m[v]
		if !ok || c == 0 {
			m[v] = 0
		}
	}

	for _, v := range nums2 {
		_, ok := m[v]
		if ok {
			m[v]++
		}
	}

	ret := make([]int, 0)
	for k, v := range m {
		if v > 0 {
			ret = append(ret, k)
		}
	}

	return ret
}
