package other

/*
给定一位研究者论文被引用次数的数组（被引用次数是非负整数）。编写一个方法，计算出研究者的 h 指数。

h 指数的定义: “h 代表“高引用次数”（high citations），一名科研人员的 h 指数是指他（她）的 （N 篇论文中）至多有 h 篇论文分别被引用了至少 h 次。（其余的 N - h 篇论文每篇被引用次数不多于 h 次。）”



示例:

输入: citations = [3,0,6,1,5]
输出: 3
解释: 给定数组表示研究者总共有 5 篇论文，每篇论文相应的被引用了 3, 0, 6, 1, 5 次。
     由于研究者有 3 篇论文每篇至少被引用了 3 次，其余两篇论文每篇被引用不多于 3 次，所以她的 h 指数是 3。



说明: 如果 h 有多种可能的值，h 指数是其中最大的那个。
*/

/*
方法一：排序

分析

我们想象一个直方图，其中 xxx 轴表示文章，yyy 轴表示每篇文章的引用次数。如果将这些文章按照引用次数降序排序并在直方图上进行表示，那么直方图上的最大的正方形的边长 hhh 就是我们所要求的 hhh。

算法

首先我们将引用次数降序排序，在排完序的数组 citations 中，如果 citations[i]>i，那么说明第 0 到 i 篇论文都有至少 i+1 次引用。
因此我们只要找到最大的 i 满足 citations[i]>i，那么 h 指数即为 i+1。例如：
i           	0 	1 	2 	3 	4 	5 	6
引用次数 	   10 	9 	5 	3 	3 	2 	1
citations[i]>i true 	true 	true 	false 	false 	false 	false

其中最大的满足 citations[i]>i的 i 值为 2，因此 h=i+1=3
*/

//func hIndex(citations []int) int {
//	sort.Sort(sort.IntSlice(citations))
//	i := 0
//	for i < len(citations) && citations[len(citations)-i-1] > i {
//		i++
//	}
//	return i
//}

/*
方法二：计数

分析

基于比较的排序算法存在时间复杂度下界 O(nlog⁡n)，如果要得到时间复杂度更低的算法，就必须考虑不基于比较的排序。

算法

方法一中，我们通过降序排序得到了 h 指数，然而，所有基于比较的排序算法，例如堆排序，合并排序和快速排序，都存在时间复杂度下界 O(nlog⁡n)。
要得到时间复杂度更低的算法. 可以考虑最常用的不基于比较的排序，计数排序。

然而，论文的引用次数可能会非常多，这个数值很可能会超过论文的总数 n，因此使用计数排序是非常不合算的（会超出空间限制）。
在这道题中，我们可以通过一个不难发现的结论来让计数排序变得有用，即：

    如果一篇文章的引用次数超过论文的总数 n，那么将它的引用次数降低为 n 也不会改变 h 指数的值。

由于 h 指数一定小于等于 n，因此这样做是正确的。在直方图中，将所有超过 y 轴值大于 n 的变为 n 等价于去掉 y>n 的整个区域。


从直方图中可以更明显地看出结论的正确性，将 y>n 的区域去除，并不会影响到最大的正方形，也就不会影响到 h 指数。

我们用一个例子来说明如何使用计数排序得到 hhh 指数。首先，引用次数如下所示：

citations=[1,3,2,3,100]

将所有大于 n=5 的引用次数变为 nnn，得到：

citations=[1,3,2,3,5]

计数排序得到的结果如下：
k   	0 	1 	2 	3 	4 	5
count 	0 	1 	1 	2 	0 	1
sk   	5 	5 	4 	3 	1 	1

其中 sk​ 表示至少有 k 次引用的论文数量，在表中即为在它之后的列（包括本身）的 count 一行的和。
根据定义，最大的满足 k≤s的 k 即为所求的 h。
在表中，这个 k 为 3，因此 h 指数为 3。
*/
func HIndex(citations []int) int {
	n := len(citations)
	s := make([]int, n+1)
	for _, v := range citations {
		s[min(n, v)]++
	}
	k := n
	for j := s[n]; k > j; j += s[k] {
		k--
	}
	return k
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
