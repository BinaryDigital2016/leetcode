package dp

/*
如果一个数列至少有三个元素，并且任意两个相邻元素之差相同，则称该数列为等差数列。

例如，以下数列为等差数列:

1, 3, 5, 7, 9
7, 7, 7, 7
3, -1, -5, -9

以下数列不是等差数列。

1, 1, 2, 5, 7



数组 A 包含 N 个数，且索引从0开始。数组 A 的一个子数组划分为数组 (P, Q)，P 与 Q 是整数且满足 0<=P<Q<N 。

如果满足以下条件，则称子数组(P, Q)为等差数组：

元素 A[P], A[p + 1], ..., A[Q - 1], A[Q] 是等差的。并且 P + 1 < Q 。

函数要返回数组 A 中所有为等差数组的子数组个数。



示例:

A = [1, 2, 3, 4]

返回: 3, A 中有三个子等差数组: [1, 2, 3], [2, 3, 4] 以及自身 [1, 2, 3, 4]。
*/

// func numberOfArithmeticSlices(A []int) int {
//     count := 0
//     for i:=0;i<len(A)-2;i++{ //最少三个数
//         d := A[i+1]-A[i]
//         for j:=i+2;j<len(A);j++{ //j为当前i的最右边元素下标
//             if d != A[j]-A[j-1]{
//                 break
//             }
//             count++
//         }
//     }
//     return count
// }

// //dp
// func numberOfArithmeticSlices(A []int) int {
//     dp := make([]int,len(A)) //dp[i]表示以下标i结束的子序列包含的等差数据列个数
//     count := 0
//     for i:=2;i<len(A);i++{ //最少三个数
//         if A[i]-A[i-1]==A[i-1]-A[i-2] { //增加当前元素后差值仍相等，数量加1
//             dp[i] = 1+dp[i-1]
//             count += dp[i] //count为dp中所有元素的和
//         }
//     }
//     return count
// }

//计数
func numberOfArithmeticSlices(A []int) int {
	curCount := 0 //当前子序列等差数列个数
	count := 0
	for i := 2; i < len(A); i++ { //最少三个数
		if A[i]-A[i-1] == A[i-1]-A[i-2] { //增加当前元素后差值仍相等，数量加1
			curCount += 1
			count += curCount
		} else {
			curCount = 0 //不是等差，从0计数
		}
	}
	return count
}
