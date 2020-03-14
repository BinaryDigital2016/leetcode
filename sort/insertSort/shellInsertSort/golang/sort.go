package mysort

func ShellInsertSort(s []int) {
	gap := len(s) >> 1
	for gap >= 1 {
		for i := 1; i < len(s); i += gap {
			end := i - gap
			tmp := s[i]
			for end >= 0 {
				if s[end] > tmp {
					s[end+1] = s[end]
					end -= gap
				} else {
					break
				}
			}
			s[end+gap] = tmp
		}
		gap >>= 1
	}

}
