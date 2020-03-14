package dp

// 斐波那契递归
// func climbStairs(n int) int {
//     if n < 2{
//         return 1
//     }
//     return climbStairs(n-1)+climbStairs(n-2)
// }

// 斐波那契非递归
// func ClimbStairs(n int) int {
//     if n < 2{
//         return 1
//     }
//     if n==2{
//         return 2
//     }

//     i,j:= 1,2
//     ret :=0
//     for k:= 3;k<=n;k++{
//         ret = i+j
//         i=j
//         j=ret
//     }

//     return ret
// }

// 动态规划
func ClimbStairs(n int) int {
	if n < 2 {
		return 1
	}
	if n == 2 {
		return 2
	}

	ret := make([]int, 0)
	ret = append(ret, []int{0, 1, 2}...) //0占位
	for k := 3; k <= n; k++ {
		ret = append(ret, ret[k-1]+ret[k-2])
	}

	return ret[n]
}
