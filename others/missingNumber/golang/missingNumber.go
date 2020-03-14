package others

// // 暴力求解
// func MissingNumber(nums []int) int {
//     n := len(nums)
//     f := make([]int8,n+1)
//     for _,v := range nums{
//         f[v] = 1
//     }

//     for k, v:=range f{
//         if v == 0{
//             return k
//         }
//     }

//     return 0
// }

// // 异或
// func MissingNumber(nums []int) int {
//     f := 0
//     n := len(nums)
//     for i:=1; i <= n;i++{
//         f ^= i
//     }

//     for i:=0;i<n;i++{
//         f ^= nums[i]
//     }

//     return f
// }

// 求和
func MissingNumber(nums []int) int {
	sum := 0
	for k, v := range nums {
		sum += k + 1
		sum -= v
	}
	return sum
}
