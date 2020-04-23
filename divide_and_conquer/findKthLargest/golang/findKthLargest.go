package dc

func findKthLargest(nums []int, k int) int {
	qSort(nums, 0, len(nums)-1, k)
	return nums[k-1]
}

func qSort(nums []int, l, r, k int) {
	if l >= r {
		return
	}
	index := partition(nums, l, r, k)
	if index-l+1 == k {
		return
	} else if index-l+1 > k {
		qSort(nums, l, index, k)
	} else {
		qSort(nums, index+1, r, k-index+l-1)
	}
}

func partition(nums []int, l, r, k int) int {
	pivot := nums[l]
	i, j := l, r
	for i < j {
		for i < j && nums[j] <= pivot {
			j--
		}
		nums[i] = nums[j]

		for i < j && nums[i] >= pivot {
			i++
		}
		nums[j] = nums[i]
	}
	nums[i] = pivot
	return i
}
