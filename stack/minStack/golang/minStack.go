type MinStack struct {
    elem []int
    min int
    len int
}


/** initialize your data structure here. */
func Constructor() MinStack {
    return MinStack{
        elem: make([]int,0),
        min: math.MaxInt64,
        len: 0,
    }
}


func (this *MinStack) Push(x int)  {
    this.elem = append(this.elem,x)
    this.len++
    if x < this.min {
        this.min = x
    }
}


func (this *MinStack) Pop()  {
    if len(this.elem) > 0 {
        this.len--
        t := this.elem[this.len]
        this.elem = this.elem[0:this.len]
        if t == this.min{
            this.min = math.MaxInt64
            for _,v := range this.elem{
                if v < this.min{
                    this.min = v
                }
            }
        }
    }
}


func (this *MinStack) Top() int {
    t := 0
    if len(this.elem) > 0 {
        t = this.elem[this.len-1]
    }
    return t
}


func (this *MinStack) GetMin() int {
    return this.min
}


/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
