package mysort

import "fmt"

// 数组快排
func QuickSort(arr []int) []int {
	return qSort(arr, 0, len(arr)-1)
}

func qSort(arr []int, l, r int) []int {
	if l < r {
		pivot := partition(arr, l, r)
		qSort(arr, 0, pivot-1)
		qSort(arr, pivot+1, r)
	}
	return arr
}

func partition(arr []int, l, r int) int {
	pivot := arr[l]
	i, j := l, r
	for i < j {
		for i < j && arr[j] >= pivot {
			j--
		}
		arr[i] = arr[j]
		for i < j && arr[i] <= pivot {
			i++
		}
		arr[j] = arr[i]
	}
	arr[i] = pivot

	return i
}

// 列表快排
const MIN = -0xFF

type LNode struct {
	Val  int
	Next *LNode
}

type LList struct {
	Head *LNode
	Len  int
}

func (l *LList) Append(val int) {
	new := LNode{val, nil}
	if l.Head == nil {
		l.Head = &LNode{MIN, &new}
		l.Len = 1
	} else {
		cur := l.Head
		for cur.Next != nil {
			cur = cur.Next
		}
		cur.Next = &new
		l.Len++
	}
}

func (l *LList) Show() {
	if l.Head == nil || l.Head.Next == nil {
		fmt.Print("empty")
		return
	}

	cur := l.Head.Next
	for cur != nil {
		fmt.Print(cur.Val)
		if cur.Next != nil {
			fmt.Print("->")
		} else {
			fmt.Print("\n")
		}
		cur = cur.Next
	}
}

func (l *LList) ShowReverse() {
	if l.Head == nil || l.Head.Next == nil {
		fmt.Print("empty")
		return
	}
	onShowReverse(l.Head.Next)
	fmt.Print("\n")
}

func (l *LList) Get(index int) int {
	if l.Head == nil || l.Head.Next == nil {
		return MIN
	}
	cur := l.Head
	for cur.Next != nil && index > 0 {
		cur = cur.Next
		index--
	}
	if index == 0 {
		return cur.Val
	}
	return MIN
}

func (l *LList) Set(index, val int) {
	if l.Head == nil || l.Head.Next == nil {
		return
	}
	cur := l.Head
	for cur.Next != nil && index > 0 {
		cur = cur.Next
		index--
	}
	if index == 0 {
		cur.Val = val
	}
}

func (l *LList) Swap(i, j int) {
	if i == j {
		return
	}
	iVal := l.Get(i)
	jVal := l.Get(j)
	l.Set(i, jVal)
	l.Set(j, iVal)
}

func onShowReverse(n *LNode) {
	if n.Next != nil {
		onShowReverse(n.Next)
		fmt.Print("<-")
	}

	fmt.Print(n.Val)
}

func QuickSort2(list *LList) *LList {
	qSort2(list, 1, list.Len)
	return list
}

func qSort2(list *LList, l, r int) {
	if l < r {
		pivot := partition2(list, l, r)
		qSort2(list, l, pivot-1)
		qSort2(list, pivot+1, r)
	}
	//return list
}

func partition2(list *LList, l, r int) int {
	pivot := list.Get(l)
	i := l
	j := i + 1
	for j <= r {
		for j <= r && list.Get(j) >= pivot {
			j++
		}
		if j <= r {
			i++
			list.Swap(i, j)
			j++
		}
	}

	list.Swap(i, l)
	return i
}
