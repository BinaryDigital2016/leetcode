package queue

type MaxQueue struct {
	q  []int
	sq []int //排序队列
	// n int //大小
}

func Constructor() MaxQueue {
	return MaxQueue{
		q:  make([]int, 0),
		sq: make([]int, 0),
	}
}

func (m *MaxQueue) Max_value() int {
	if len(m.sq) > 0 {
		return m.sq[0]
	}
	return -1
}

// 均摊O(1)
func (m *MaxQueue) Push_back(value int) {
	for len(m.sq) > 0 && m.sq[len(m.sq)-1] < value {
		m.sq = m.sq[:len(m.sq)-1]
	}
	m.sq = append(m.sq, value)
	m.q = append(m.q, value)
}

func (m *MaxQueue) Pop_front() int {
	if len(m.q) == 0 {
		return -1
	}
	if len(m.q) == 1 {
		t := m.q[0]
		m.q = make([]int, 0)
		m.sq = make([]int, 0)
		return t
	}
	t := m.q[0]
	m.q = m.q[1:]
	if t == m.sq[0] {
		m.sq = m.sq[1:]
	}
	return t
}

/**
 * Your MaxQueue object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Max_value();
 * obj.Push_back(value);
 * param_3 := obj.Pop_front();
 */
