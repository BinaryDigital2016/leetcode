package others

/*
幂集。编写一种方法，返回某集合的所有子集。集合中不包含重复的元素。

说明：解集不能包含重复的子集。

示例:

 输入： nums = [1,2,3]
 输出：
[
  [3],
  [1],
  [2],
  [1,2,3],
  [1,3],
  [2,3],
  [1,2],
  []
]
*/

/*
先在result列表中将入一个空子集
然后遍历nums每一个元素，将元素与result中的每个元素结合加到result中
直到nums遍历完
*/
func subsets(nums []int) [][]int {
	ret := make([][]int, 0)
	ret = append(ret, []int{})
	for _, v := range nums {
		n := len(ret)
		for i := 0; i < n; i++ {
			a := copy(ret[i])
			ret = append(ret, append(a, v))
		}
	}
	return ret
}
func copy(a []int) []int {
	b := make([]int, len(a))
	for i := 0; i < len(a); i++ {
		b[i] = a[i]
	}
	return b
}
