package mysort

func SimpleSelectSort(s []int) {
	begin, end := 0, len(s)-1
	for begin < end {
		min, max := begin, end
		for i := begin; i <= end; i++ {
			if s[i] > s[max] {
				max = i
			}
			if s[i] < s[min] {
				min = i
			}
		}
		s[begin], s[min] = s[min], s[begin]
		if max == begin {
			//begin已被交换值min处
			max = min
		}
		s[end], s[max] = s[max], s[end]
		begin++
		end--
	}

}
