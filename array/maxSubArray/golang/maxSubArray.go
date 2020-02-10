func maxSubArray(nums []int) int {
    n := len(nums)
    if n<=0 {
        return 0
    }
    sum := nums[0]
    cur := nums[0]

    for i:=1;i<n;i++{
        if cur > 0{
            cur += nums[i]
        } else {
            cur = nums[i]
        }
        sum = max(cur, sum)
    }
    
    return sum
}

func max(x,y int) int{
    if x < y{
        return y
    }
    return x
}
