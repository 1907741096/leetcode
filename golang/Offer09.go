package main

type CQueue struct {
	stack1 Stack
	stack2 Stack
}

type Stack struct {
	list []int
}

func NewStack() Stack {
	return Stack{
		list: []int{},
	}
}

func (this *Stack) Push(value int) {
	this.list = append(this.list, value)
}

func (this *Stack) Pop() int {
	if len(this.list) == 0 {
		return -1
	}
	value := this.list[len(this.list) - 1]
	this.list = this.list[0: len(this.list) - 1]
	return value
}


func Constructor() CQueue {
	return CQueue{
		stack1: NewStack(),
		stack2: NewStack(),
	}
}


func (this *CQueue) AppendTail(value int)  {
	this.stack1.Push(value)
}


func (this *CQueue) DeleteHead() int {
	if len(this.stack1.list) == 0 {
		return -1
	}
	for len(this.stack1.list) > 1 {
		this.stack2.Push(this.stack1.Pop())
	}
	value := this.stack1.Pop()
	for len(this.stack2.list) > 0 {
		this.stack1.Push(this.stack2.Pop())
	}
	return value
}


/**
 * Your CQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AppendTail(value);
 * param_2 := obj.DeleteHead();
 */
