package others

/*
给定一个整数数组 nums 和一个目标值 target，请你在该数组中找出和为目标值的那 两个 整数，并返回他们的数组下标。
你可以假设每种输入只会对应一个答案。但是，你不能重复利用这个数组中同样的元素。
*/

// func twoSum(nums []int, target int) []int {
//     for k :=0;k<len(nums);k++{
//         for kk:=k+1;kk<len(nums);kk++{
//             if nums[k] + nums[kk] == target{
//                 return []int{k,kk}
//             }
//         }
//     }
//     return nil
// }

// 空间换时间
func twoSum(nums []int, target int) []int {
	mapItem := make(map[int]int, 0)
	for k, v := range nums {
		remain := target - v
		if kk, ok := mapItem[remain]; ok && kk != k {
			return []int{k, kk}
		}
		mapItem[v] = k
	}
	return nil
}
