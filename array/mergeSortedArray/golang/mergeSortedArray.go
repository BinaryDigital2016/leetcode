func merge(nums1 []int, m int, nums2 []int, n int) {
	tmp := make([]int, len(nums1))
	copy(tmp, nums1) //append扩容会改变底层数组，所以不能直接修改nums1
	i, j, cnt := 0, 0, 0
	for i < m+cnt && j < n {
		if tmp[i] > nums2[j] {
			cnt++
			tmp = append(tmp[:i], append([]int{nums2[j]}, tmp[i:]...)...)
			j++
		}
		i++
	}

	if j < n {
		tmp = append(tmp[:m+cnt], nums2[j:]...)
	}
	copy(nums1, tmp)
}
