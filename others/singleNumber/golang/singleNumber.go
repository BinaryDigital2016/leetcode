package others

/*
给定一个非空整数数组，除了某个元素只出现一次以外，其余每个元素均出现两次。找出那个只出现了一次的元素。

说明：

你的算法应该具有线性时间复杂度。 你可以不使用额外空间来实现吗？

示例 1:

输入: [2,2,1]
输出: 1

示例 2:

输入: [4,1,2,1,2]
输出: 4
*/

// //使用异或运算：相同值异或为自身，不同值异或为0，0异或任何数等于0
// func SingleNumber(nums []int) int {
//     ret := 0
//     for _,v:=range nums{
//         ret ^= v
//     }
//     return ret
// }

// 使用map暴力解
func SingleNumber(nums []int) int {
	m := make(map[int]struct{})
	for _, v := range nums {
		if _, ok := m[v]; ok {
			delete(m, v)
		} else {
			m[v] = struct{}{}
		}
	}
	for k, _ := range m {
		return k
	}
	return 0
}
