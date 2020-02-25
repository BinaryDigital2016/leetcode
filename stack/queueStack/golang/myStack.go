type MyStack struct {
    e []int
}


/** Initialize your data structure here. */
func Constructor() MyStack {
    return MyStack{
        e: make([]int,0),
    }
}


/** Push element x onto stack. */
func (this *MyStack) Push(x int)  {
    this.e = append(this.e,x)
}


/** Removes the element on top of the stack and returns that element. */
func (this *MyStack) Pop() int {
    n := len(this.e)
    if n == 0{
        return -1
    }
    r := this.e[n-1]
    this.e = this.e[:n-1]
    return r
}


/** Get the top element. */
func (this *MyStack) Top() int {
    if len(this.e) == 0{
        return -1
    }
    return this.e[len(this.e)-1]
}


/** Returns whether the stack is empty. */
func (this *MyStack) Empty() bool {
    return len(this.e)==0
}


/**
 * Your MyStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Empty();
 */
