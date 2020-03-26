package dp

/*
给定一个非负整数 num。对于 0 ≤ i ≤ num 范围中的每个数字 i ，计算其二进制数中的 1 的数目并将它们作为数组返回。

示例 1:

输入: 2
输出: [0,1,1]

示例 2:

输入: 5
输出: [0,1,1,2,1,2]

进阶:

    给出时间复杂度为O(n*sizeof(integer))的解答非常容易。但你可以在线性时间O(n)内用一趟扫描做到吗？
    要求算法的空间复杂度为O(n)。
*/

// func countBits(num int) []int {
//     dp := make([]int,num+1)
//     for i:=1;i<num+1;i++{
//         dp[i] = dp[i&(i-1)]+1
//     }
//     return dp
// }

/*
观察x 和 x′=x/2的关系：

x=(1001011101)2=(605)10

x′=(100101110)2=(302)10

可以发现 x′与 x 只有一位不同，这是因为x′ 可以看做 x 移除最低有效位的结果。

这样，我们就有了下面的状态转移函数：

P(x)=P(x/2)+(xmod  2)
*/
func countBits(num int) []int {
	dp := make([]int, num+1)
	for i := 1; i < num+1; i++ {
		dp[i] = dp[i/2] + (i & 1)
	}
	return dp
}
