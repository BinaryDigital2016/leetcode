package array

import "sort"

// // map缓存
// func intersect(nums1 []int, nums2 []int) []int {
//     m := make(map[int]int)
//     for i:=0;i<len(nums1);i++{
//         m[nums1[i]]++
//     }

//     k := 0
//     for i:=0;i<len(nums2);i++{
//         if v,ok:=m[nums2[i]];ok&&v>0{
//             nums1[k]=nums2[i]
//             k++
//             m[nums2[i]]--
//         }
//     }

//     return nums1[:k]
// }

// 排序
func intersect(nums1 []int, nums2 []int) []int {
	sort.Sort(sort.IntSlice(nums1))
	sort.Sort(sort.IntSlice(nums2))

	i, j := 0, 0
	s := make([]int, 0)
	for i < len(nums1) && j < len(nums2) {
		if nums1[i] < nums2[j] {
			i++
		} else if nums1[i] > nums2[j] {
			j++
		} else {
			s = append(s, nums2[j])
			i++
			j++
		}
	}

	return s
}
