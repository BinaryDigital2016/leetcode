package sort

/*
给你两个数组，arr1 和 arr2，

    arr2 中的元素各不相同
    arr2 中的每个元素都出现在 arr1 中

对 arr1 中的元素进行排序，使 arr1 中项的相对顺序和 arr2 中的相对顺序相同。未在 arr2 中出现过的元素需要按照升序放在 arr1 的末尾。



示例：

输入：arr1 = [2,3,1,3,2,4,6,7,9,2,19], arr2 = [2,1,4,3,9,6]
输出：[2,2,2,1,4,3,3,9,6,7,19]



提示：

    arr1.length, arr2.length <= 1000
    0 <= arr1[i], arr2[i] <= 1000
    arr2 中的元素 arr2[i] 各不相同
    arr2 中的每个元素 arr2[i] 都出现在 arr1 中
*/
// 暴力法
// func relativeSortArray(arr1 []int, arr2 []int) []int {
//     t := 0
//     for i:=0;i<len(arr2);i++{
//         for j:=0;j<len(arr1);j++{
//             if arr1[j] == arr2[i]{
//                 arr1[j],arr1[t]=arr1[t],arr1[j]
//                 t++
//             }
//         }
//     }

//     tmp := sort.IntSlice(arr1[t:])
//     sort.Sort(tmp)

//     return arr1
// }

// 计数排序
func relativeSortArray(arr1 []int, arr2 []int) []int {
	tmp := make([]int, 1001) //arr1.length, arr2.length <= 1000
	ret := make([]int, 0, len(arr1))
	for _, v := range arr1 {
		tmp[v] += 1
	}

	for _, v := range arr2 {
		for tmp[v] > 0 {
			ret = append(ret, v)
			tmp[v]--
		}
	}

	for i := 0; i < 1001; i++ {
		for tmp[i] > 0 {
			ret = append(ret, i)
			tmp[i]--
		}
	}

	return ret
}
