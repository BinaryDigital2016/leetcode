package mysort

//s := []int{1, 42, 23, 233, 33, 1, 2, 5, 7, 8, 3}
//s要有明确的范围
func CountSort(s []int) {
	t := make([]int, 233+1)
	for _, v := range s {
		t[v]++
	}
	j := 0
	for i := 0; i < len(t); i++ {
		for t[i] > 0 {
			s[j] = i
			t[i]--
			j++
		}
	}
}
