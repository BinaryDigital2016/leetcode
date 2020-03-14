package mysort

func HeapSort(s []int) {
	// 建堆
	for idx := len(s) / 2; idx >= 0; idx-- {
		adjustHeap(s, idx)
	}

	end := len(s) - 1
	for end >= 0 {
		s[0], s[end] = s[end], s[0]
		adjustHeap(s[:end], 0)
		end--
	}

}

func adjustHeap(s []int, idx int) {
	parent := idx
	child := idx*2 + 1
	for child < len(s) {
		if child+1 < len(s) && s[child+1] > s[child] {
			child++
		}

		if s[child] <= s[parent] {
			break
		}

		s[child], s[parent] = s[parent], s[child]
		parent = child
		child = parent*2 + 1
	}
}
