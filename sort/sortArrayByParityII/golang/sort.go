package sort

/*
给定一个非负整数数组 A， A 中一半整数是奇数，一半整数是偶数。

对数组进行排序，以便当 A[i] 为奇数时，i 也是奇数；当 A[i] 为偶数时， i 也是偶数。

你可以返回任何满足上述条件的数组作为答案。



示例：

输入：[4,2,5,7]
输出：[4,5,2,7]
解释：[4,7,2,5]，[2,5,4,7]，[2,7,4,5] 也会被接受。



提示：

    2 <= A.length <= 20000
    A.length % 2 == 0
    0 <= A[i] <= 1000
*/
// // 两次遍历
// func sortArrayByParityII(A []int) []int {
//     ret := make([]int, len(A))
//     i := 0
//     for j:=0;j<len(A);j++{
//         if A[j]%2 == 0{
//             ret[i] = A[j]
//             i+=2
//         }
//     }

//     i = 1
//     for j:=0;j<len(A);j++{
//         if A[j]%2 == 1{
//             ret[i] = A[j]
//             i+=2
//         }
//     }

//     return ret
// }

// 双指针，找到不符合规则的偶数位，交换下一个不符合规则的奇数位
func sortArrayByParityII(A []int) []int {
	//i := 0
	j := 1
	for i := 0; i < len(A); i += 2 {
		if A[i]%2 == 1 {
			for ; j < len(A); j += 2 {
				if A[j]%2 == 0 {
					A[i], A[j] = A[j], A[i]
					break
				}
			}
		}
	}

	return A
}
