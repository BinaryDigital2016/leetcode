package mysort

import "sort"

func BucketSort(s []int) {
	n := len(s)
	if n < 2 {
		return
	}
	min, max := s[0], s[0]
	for _, v := range s {
		if v > max {
			max = v
		}
		if v < min {
			min = v
		}
	}
	//数组值都一样
	if max == min {
		return
	}

	bucketCnt := (max-min)/4 + 1
	//初始化桶
	bucket := make([][]int, 0)
	for i := 0; i < bucketCnt; i++ {
		b := make([]int, 0)
		bucket = append(bucket, b)
	}

	//入桶
	for _, v := range s {
		idx := mappingFunc(v, max, min)
		bucket[idx] = append(bucket[idx], v)
	}

	//排序桶、输出
	k := 0
	for i := 0; i < bucketCnt; i++ {
		sort.Sort(sort.IntSlice(bucket[i])) //排序方法任意
		//mysort.BubbleSort(bucket[i])
		for j := 0; j < len(bucket[i]); j++ {
			s[k] = bucket[i][j]
			k++
		}
	}
}

// 映射函数，x的单调函数
func mappingFunc(x, max, min int) int {
	return (x - min) * 4 / (max - min)
}
