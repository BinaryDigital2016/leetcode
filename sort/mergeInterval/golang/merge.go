/*
给出一个区间的集合，请合并所有重叠的区间。

示例 1:

输入: [[1,3],[2,6],[8,10],[15,18]]
输出: [[1,6],[8,10],[15,18]]
解释: 区间 [1,3] 和 [2,6] 重叠, 将它们合并为 [1,6].

示例 2:

输入: [[1,4],[4,5]]
输出: [[1,5]]
解释: 区间 [1,4] 和 [4,5] 可被视为重叠区间。
*/
func merge(intervals [][]int) [][]int {
    if len(intervals) <= 1{
        return intervals
    }

    interval := Int2Slice(intervals)
    sort.Sort(interval) //按起始值排序

    ret := make([][]int,0)
    for _,v:=range interval{
        n := len(ret)
        if n == 0 || ret[n-1][1] < v[0]{
            ret = append(ret, v)
        } else {
            ret[n-1][1] = max(ret[n-1][1], v[1])
        }
    }
    return ret
}

func max(a,b int) int{
    if a > b{
        return a
    }
    return b
}

type Int2Slice [][]int
func (p Int2Slice) Len() int           { return len(p) }
func (p Int2Slice) Less(i, j int) bool { return p[i][0] < p[j][0] }
func (p Int2Slice) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
