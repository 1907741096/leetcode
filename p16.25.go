package main

type LRUNode struct {
	Key int
	Value int
	Next *LRUNode
	Prev *LRUNode
}

type LRUCache struct {
	hashMap map[int]*LRUNode
	head *LRUNode
	tail *LRUNode
	capacity int
	length int
}


func Constructor(capacity int) LRUCache {
	start := &LRUNode{
		Key: 0,
		Value: 0,
	}
	end := &LRUNode{
		Key: 0,
		Value: 0,
	}
	start.Next = end
	end.Prev = start
	return LRUCache{
		hashMap:  make(map[int]*LRUNode),
		length:	  0,
		head: start,
		tail: end,
		capacity: capacity,
	}
}

func (this *LRUCache) Get(key int) int {
	if node, ok := this.hashMap[key]; ok {
		this.moveHead(node)
		return node.Value
	}
	return -1
}

func (this *LRUCache) Put(key int, value int)  {
	if node, ok := this.hashMap[key]; ok {
		node.Value = value
		this.moveHead(node)
	} else {
		node = &LRUNode {
			Key: key,
			Value: value,
		}
		this.addHead(node)
		if this.length >= this.capacity {
			this.removeTail()
		} else {
			this.length ++
		}
		this.hashMap[key] = node
	}
}

func (this *LRUCache) moveHead (node *LRUNode) {
	node.remove()
	this.addHead(node)
}

func (node *LRUNode) remove () {
	node.Next.Prev = node.Prev
	node.Prev.Next = node.Next
}

func (this *LRUCache) addHead (node *LRUNode) {
	node.Next = this.head.Next
	node.Prev = this.head
	this.head.Next.Prev = node
	this.head.Next = node
}

func (this *LRUCache) removeTail () {
	node := this.tail.Prev
	delete(this.hashMap, node.Key)
	node.remove()
}


/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */