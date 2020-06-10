package main

type MinStack struct {
	list []int
	minNum int
	len int
}


/** initialize your data structure here. */
//func Constructor() MinStack {
//	MinStack := *new(MinStack)
//	MinStack.len = 0
//	MinStack.list = []int{}
//	return MinStack
//}


func (this *MinStack) Push(x int)  {
	this.list = append(this.list, x)
	this.len ++
	if this.len == 1 || x < this.minNum{
		this.minNum = x
	}
}


func (this *MinStack) Pop()  {
	if this.list[this.len - 1] == this.minNum {
		minN := this.list[0]
		for i := 1 ; i < this.len - 1 ; i ++ {
			if minN > this.list[i] {
				minN = this.list[i]
			}
		}
		this.minNum = minN
	}
	this.len --
	this.list = this.list[0: this.len]
}


func (this *MinStack) Top() int {
	return this.list[this.len - 1]
}


func (this *MinStack) GetMin() int {
	return this.minNum
}


/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */