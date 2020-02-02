package twoSum

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
