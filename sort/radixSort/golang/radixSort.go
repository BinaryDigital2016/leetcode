package mysort

func RadixSort(s []int) {
	n := len(s)
	if n < 2 {
		return
	}

	max := s[0]
	for _, v := range s {
		if v > max {
			max = v
		}
	}

	//最大位数
	maxBit := maxBit(max)
	radix := 1                    //当前位，1->个位 10->十位 。。。
	for i := 0; i < maxBit; i++ { //maxBit次
		//初始化桶
		bucket := make([][]int, 10) //每位上的数字一定在[0,9]内
		for j := 0; j < 10; i++ {
			b := make([]int, 0)
			bucket[j] = b
		}

		//记录该位上为0~9的数的个数
		for j := 0; j < n; j++ {
			idx := (s[j] / radix) % 10
			bucket[idx] = append(bucket[idx], s[j])
		}

		//本轮排序结果赋值回s
		k := 0
		for j := 0; j < 10; j++ {
			nb := len(bucket[j])
			for m := 0; m < nb; m++ {
				s[k] = bucket[j][m]
				k++
			}
		}

		radix *= 10 //下一位
	}
}

func maxBit(x int) int {
	ret := 0
	for x > 0 {
		ret++
		x /= 10
	}
	return ret
}
