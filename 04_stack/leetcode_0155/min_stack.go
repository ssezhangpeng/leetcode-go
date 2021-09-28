package leetcode_0155

type MinStack struct {
	MinStack []int
	Stack []int
}

func Constructor() MinStack {
	return MinStack{
		MinStack: make([]int, 0),
		Stack: make([]int, 0),
	}
}


func (this *MinStack) Push(val int)  {
	this.Stack = append(this.Stack, val)

	if len(this.MinStack) == 0 {
		this.MinStack = append(this.MinStack, val)
	} else {
		topVal := this.MinStack[len(this.MinStack)-1]
		if topVal >= val {
			this.MinStack = append(this.MinStack, val)
		} else {
			this.MinStack = append(this.MinStack, topVal)
		}
	}
}


func (this *MinStack) Pop()  {
	this.Stack = this.Stack[:len(this.Stack)-1]
	this.MinStack = this.MinStack[:len(this.MinStack)-1]
}


func (this *MinStack) Top() int {
	return this.Stack[len(this.Stack)-1]
}


func (this *MinStack) GetMin() int {
	return this.MinStack[len(this.MinStack)-1]
}


/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
