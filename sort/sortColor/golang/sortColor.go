/*
给定一个包含红色、白色和蓝色，一共 n 个元素的数组，原地对它们进行排序，使得相同颜色的元素相邻，并按照红色、白色、蓝色顺序排列。

此题中，我们使用整数 0、 1 和 2 分别表示红色、白色和蓝色。

注意:
不能使用代码库中的排序函数来解决这道题。

示例:

输入: [2,0,2,1,1,0]
输出: [0,0,1,1,2,2]

进阶：

    一个直观的解决方案是使用计数排序的两趟扫描算法。
    首先，迭代计算出0、1 和 2 元素的个数，然后按照0、1、2的排序，重写当前数组。
    你能想出一个仅使用常数空间的一趟扫描算法吗？
*/
// // 冒泡排序
// func sortColors(nums []int)  {
//     n:=len(nums)
//     for i:=0;i<n;i++{
//         for j:=0;j<n-i-1;j++{
//             if nums[j]>nums[j+1]{
//                 nums[j],nums[j+1]=nums[j+1],nums[j]
//             }
//         }
//     } 
// }

// // 插入排序
// func sortColors(nums []int)  {
//     n:=len(nums)
//     for i:=1;i<n;i++{
//         if nums[i]<nums[i-1]{
//             t := nums[i]
//             for j:=i;j>=0;j--{
//                 if j>0&&t<nums[j-1]{
//                     nums[j]=nums[j-1]
//                 } else {
//                     nums[j] = t
//                     break
//                 }
//             }
//         }
//     } 
// }

// 选择排序
// func sortColors(nums []int)  {
//     n:=len(nums)
//     for i:=0;i<n;i++{
//         minIndex := i
//         for j:=i+1;j<n;j++{
//             if nums[j]<nums[minIndex]{
//                 minIndex = j
//             }
//         }
//         if i != minIndex{
//             nums[i],nums[minIndex]=nums[minIndex],nums[i]
//         }
//     } 
// }

// 三指针法
func sortColors(nums []int)  {
    p0,p2,cur:=0,len(nums)-1,0
    for ;cur<=p2; {
        switch nums[cur]{
        case 0:
            nums[cur],nums[p0]=nums[p0],nums[cur]
            cur++ //因为curr左边的值已经扫描过了，所以curr要++继续扫描下一位，而与p2交换的值，curr未扫描，要停下来扫描一下，所以curr不用++
            p0++
        case 2:
            nums[cur],nums[p2]=nums[p2],nums[cur]
            p2--
        default:
            cur++
        }
    }
}
