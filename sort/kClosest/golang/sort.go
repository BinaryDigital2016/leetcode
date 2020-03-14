package mysort

// 本质为求k小元素

//排序取前k个,略

//快排变形
func KClosest(points [][]int, K int) [][]int {
	partition(points, 0, len(points)-1, K)
	return points[0:K]
}

func partition(points [][]int, le, ri, K int) {
	if le >= ri {
		return
	}
	p := points[le]
	pivot := dist(p)
	l, r := le, ri
	for l < r {
		for l < r && dist(points[r]) >= pivot {
			r--
		}
		points[l] = points[r]
		for l < r && dist(points[l]) <= pivot {
			l++
		}
		points[r] = points[l]
	}
	points[l] = p

	//l==r
	if K < l-le+1 { //K位于左半部分
		partition(points, le, l, K)
	} else {
		partition(points, l+1, ri, K-l+le-1)
	}
}

func dist(s []int) int {
	return s[0]*s[0] + s[1]*s[1]
}
