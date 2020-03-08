package sort

import "sort"

func sortByBits(arr []int) []int {
	m := MyIntSlice(arr)
	sort.Sort(m)
	return arr
}

type MyIntSlice []int

func (p MyIntSlice) Len() int { return len(p) }

func (p MyIntSlice) Less(i, j int) bool {
	diffOne := countOne(p[i]) - countOne(p[j])
	switch {
	case diffOne < 0:
		return true
	case diffOne > 0:
		return false
	default:
		return p[i] < p[j]
	}
}

func (p MyIntSlice) Swap(i, j int) { p[i], p[j] = p[j], p[i] }

func countOne(x int) int {
	count := 0
	for x > 0 {
		count++
		x = x & (x - 1)
	}
	return count
}
