package array

/*
给定一个排序数组，你需要在原地删除重复出现的元素，使得每个元素只出现一次，返回移除后数组的新长度。

不要使用额外的数组空间，你必须在原地修改输入数组并在使用 O(1) 额外空间的条件下完成。

示例 1:

给定数组 nums = [1,1,2],

函数应该返回新的长度 2, 并且原数组 nums 的前两个元素被修改为 1, 2。

你不需要考虑数组中超出新长度后面的元素。

示例 2:

给定 nums = [0,0,1,1,1,2,2,3,3,4],

函数应该返回新的长度 5, 并且原数组 nums 的前五个元素被修改为 0, 1, 2, 3, 4。

你不需要考虑数组中超出新长度后面的元素。
*/

// func RemoveDuplicates(nums []int) int {
//     size := len(nums)
//     if size <= 1{
//         return size
//     }

//     cnt := 0
//     for i := 1;i < size-cnt;{
//         if nums[i] == nums[i-1]{
//             for j := i; j < size-cnt; j++{
//                 nums[j-1]=nums[j]
//             }
//             cnt += 1
//         } else {
//             i++
//         }
//     }

//     return size - cnt
// }

func RemoveDuplicates(nums []int) int {
	i, j := 0, 1
	for ; j < len(nums); j++ {
		if nums[i] != nums[j] {
			i++
			nums[i] = nums[j]
		}
	}

	return i + 1
}
