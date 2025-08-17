package min_stack

type Node struct {
    val int
    min int
}
type MinStack struct {
    stack []Node
}

func Constructor() MinStack {
    return MinStack{
        stack: make([]Node, 0),
    }
}


func (this *MinStack) Push(val int)  {
    minVal := val
    if len(this.stack) > 0 && this.GetMin() < val {
        minVal = this.GetMin()
    }
    this.stack = append(this.stack, Node{
        val: val,
        min: minVal,
    })
}

func (this *MinStack) Pop()  {
    if len(this.stack) == 0 {
        return
    }
    this.stack = this.stack[:len(this.stack)-1]
}


func (this *MinStack) Top() int {
    if len(this.stack) == 0 {
        return 0
    }
    
    return this.stack[len(this.stack)-1].val
}


func (this *MinStack) GetMin() int {
    if len(this.stack) == 0 {
        return 0
    }
    return this.stack[len(this.stack)-1].min
}


/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
