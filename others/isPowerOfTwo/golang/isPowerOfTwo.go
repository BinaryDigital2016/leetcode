package others

// n&(n-1)将n最右边的1置0
// n(-n)获取n中最右边的1
func IsPowerOfTwo(n int) bool {
	// if n <= 0{
	//     return false
	// }
	return n > 0 && n&(n-1) == 0
}
