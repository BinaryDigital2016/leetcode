/*
给定一个整数数组  nums，求出数组从索引 i 到 j  (i ≤ j) 范围内元素的总和，包含 i,  j 两点。

示例：

给定 nums = [-2, 0, 3, -5, 2, -1]，求和函数为 sumRange()

sumRange(0, 2) -> 1
sumRange(2, 5) -> -1
sumRange(0, 5) -> -3

说明:

    你可以假设数组不可变。
    会多次调用 sumRange 方法。
*/
// 暴力法
// type NumArray struct {
//     s []int
// }


// func Constructor(nums []int) NumArray {
//     return NumArray{
//         s : nums[:],
//     }
// }


// func (this *NumArray) SumRange(i int, j int) int {
//     sum := 0
//     for m:=i;m<=j;m++{
//         sum += this.s[m]
//     }

//     return sum
// }

//带缓存
type NumArray struct {
    s []int
    sum []int
}


func Constructor(nums []int) NumArray {
    num := NumArray{
        s : nums[:],
        sum : make([]int,len(nums)+1),
    }

    n:=len(nums)
    for i:=0;i<n;i++{
        num.sum[i+1] = num.sum[i] + nums[i]
    }
    return num
}


func (this *NumArray) SumRange(i int, j int) int {
    if i > len(this.sum) || j > len(this.sum){
        return 0
    }
    if i > j {
        i,j=j,i
    }
    return this.sum[j+1] - this.sum[i]
}


/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.SumRange(i,j);
 */
