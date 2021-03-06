package others

/*
判断是否是回文数
*/

// func isPalindrome(x int) bool {
//     if x < 0 || (x%10==0 && x!= 0){
//         return false
//     }
//     origin := x
//     result := 0
//     for x != 0{
//         b := x%10
//         x = x/10
//         result = result*10 + b //溢出必然不是回文数
//     }
//     return result == origin
// }

func isPalindrome(x int) bool {
	// 个位是0且x!=0返回false
	if x < 0 || (x%10 == 0 && x != 0) {
		return false
	}

	result := 0
	for x > result { //翻转一半
		b := x % 10
		x = x / 10
		result = result*10 + b
	}
	// 位数为奇数，去掉result的个位
	return result == x || result/10 == x
}
