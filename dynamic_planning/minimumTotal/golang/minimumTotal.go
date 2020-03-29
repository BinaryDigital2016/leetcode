package dp

/*
给定一个三角形，找出自顶向下的最小路径和。每一步只能移动到下一行中相邻的结点上。

例如，给定三角形：

[
     [2],
    [3,4],
   [6,5,7],
  [4,1,8,3]
]

自顶向下的最小路径和为 11（即，2 + 3 + 5 + 1 = 11）。
*/

// func minimumTotal(triangle [][]int) int {
//     row := len(triangle)
//     if row == 0 {
//         return 0
//     }
//     dp := make([][]int,row) //dp[i][j]:到达第i行第j列的做小路径和
//     dp[0] = make([]int,1)
//     dp[0][0]=triangle[0][0]
//     for i:=1;i<row;i++{
//         if len(dp[i])==0{
//             dp[i]=make([]int,i+1)
//         }
//         for j:=0;j<=i;j++{
//             if j==0 {
//                 dp[i][j]=dp[i-1][j]+triangle[i][j]
//             } else if j==i{
//                 dp[i][j]=dp[i-1][j-1]+triangle[i][j]
//             } else {
//                 dp[i][j]=min(dp[i-1][j],dp[i-1][j-1])+triangle[i][j]
//             }
//         }
//     }
//     fmt.Println(dp)
//     ret := dp[row-1][0]
//     for i:=1;i<row;i++{
//         ret = min(ret, dp[row-1][i])
//     }
//     return ret
// }

func minimumTotal(triangle [][]int) int {
	row := len(triangle)
	if row == 0 {
		return 0
	}
	dp := make([]int, row) //dp[i]:到达当前行j列的最小路径
	dp[0] = triangle[0][0]
	first, second := 0, 0
	for i := 1; i < row; i++ {
		for j := 0; j <= i; j++ {
			second = dp[j]
			if j == 0 {
				dp[j] = second + triangle[i][j]
			} else if j == i {
				dp[j] = first + triangle[i][j]
			} else {
				dp[j] = min(first, second) + triangle[i][j]
			}
			first = second
		}
	}
	ret := dp[0]
	for i := 1; i < row; i++ {
		ret = min(ret, dp[i])
	}
	return ret
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
