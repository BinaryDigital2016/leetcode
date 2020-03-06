package queue

type MyQueue struct {
	s1 []int
	s2 []int
}

/** Initialize your data structure here. */
func Constructor() MyQueue {
	return MyQueue{
		s1: make([]int, 0),
		s2: make([]int, 0),
	}
}

/** Push element x to the back of queue. */
func (this *MyQueue) Push(x int) {
	this.s1 = append(this.s1, x)
}

/** Removes the element from in front of queue and returns that element. */
func (this *MyQueue) Pop() int {
	if len(this.s2) == 0 {
		for i := len(this.s1) - 1; i >= 0; i-- {
			this.s2 = append(this.s2, this.s1[i])
		}
		this.s1 = make([]int, 0)
	}
	n := len(this.s2)
	if n == 0 {
		return -1
	}
	t := this.s2[n-1]
	this.s2 = this.s2[0 : n-1]
	return t
}

/** Get the front element. */
func (this *MyQueue) Peek() int {
	if len(this.s2) == 0 {
		for i := len(this.s1) - 1; i >= 0; i-- {
			this.s2 = append(this.s2, this.s1[i])
		}
		this.s1 = make([]int, 0)
	}
	n := len(this.s2)
	if n == 0 {
		return -1
	}
	return this.s2[n-1]
}

/** Returns whether the queue is empty. */
func (this *MyQueue) Empty() bool {
	return len(this.s1)+len(this.s2) == 0
}

/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */
