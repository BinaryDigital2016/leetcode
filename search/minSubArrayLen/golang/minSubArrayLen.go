package search

/*
给定一个含有 n 个正整数的数组和一个正整数 s ，找出该数组中满足其和 ≥ s 的长度最小的连续子数组，并返回其长度。如果不存在符合条件的连续子数组，返回 0。

示例:

输入: s = 7, nums = [2,3,1,2,4,3]
输出: 2
解释: 子数组 [4,3] 是该条件下的长度最小的连续子数组。
*/

// func minSubArrayLen(s int, nums []int) int {
//     n := len(nums)
//     if n == 0 {
//         return 0
//     }

//     sums := make([]int,n)
//     for k,v:= range nums{
//         if k == 0 {
//             sums[k] = v
//             continue
//         }
//         sums[k] = sums[k-1] + v
//     }

//     ret := n + 1
//     for i:=0;i<n;i++ {
//         for j:=i;j<n;j++{
//             sum := sums[j]-sums[i]+nums[i]
//             if sum < s {
//                 continue
//             }
//             ret = min(ret, j-i+1)
//             break //后面跨度只会更大
//         }
//     }

//     if ret == n+1 {
//         return 0
//     }
//     return ret
// }

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func minSubArrayLen(s int, nums []int) int {
	n := len(nums)
	if n == 0 {
		return n
	}

	i, j := 0, 0
	sum := 0
	ret := n + 1
	for j < n {
		sum += nums[j]
		for sum >= s && i < n {
			ret = min(ret, j-i+1)
			sum -= nums[i]
			i++
		}
		j++
	}

	if ret == n+1 {
		return 0
	}
	return ret
}
