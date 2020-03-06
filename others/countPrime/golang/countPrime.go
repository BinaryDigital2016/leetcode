package others

/*
统计所有小于非负整数 n 的质数的数量。
*/

//暴力法忽略

// 筛选法
func countPrimes(n int) int {
	flag := make([]byte, n) //初始默认都是质数,可以用位图代替
	count := 0
	for i := 2; i < n; i++ {
		if flag[i] == 0 { //flag[i]是质数
			count++
			//for j:=i;j<n;j+=i{ //flag[i]的倍数都不是质数
			for j := i * i; j < n; j += i { //flag[i]的倍数都不是质数
				flag[j] = 1
			}
		}
	}
	return count
}
