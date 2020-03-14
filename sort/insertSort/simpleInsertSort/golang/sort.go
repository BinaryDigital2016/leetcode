package mysort

func SimpleInsertSort(s []int) {
	for i := 1; i < len(s); i++ {
		end := i - 1
		tmp := s[i]
		for end >= 0 {
			if s[end] > tmp {
				s[end+1] = s[end]
				end--
			} else {
				break
			}
		}
		s[end+1] = tmp
	}
}
