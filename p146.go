package main

type LRUCache struct {
	capacity int
	hashSet map[int]int
	list []int
	listIndex map[int]int
	count int
	length int
	index int
}


func Constructor(capacity int) LRUCache {
	l := LRUCache{
		capacity:capacity,
		hashSet:make(map[int]int),
		list: []int{},
		listIndex:make(map[int]int),
		count: 0,
		length: 0,
		index: 0,
	}

	return l
}


func (this *LRUCache) Get(key int) int {
	if value, ok := this.hashSet[key]; ok && value != -1 {
		index := this.listIndex[key]
		this.list[index] = -1
		this.list = append(this.list, key)
		this.length ++
		this.listIndex[key] = this.length - 1
		return value
	}
	return -1
}


func (this *LRUCache) Put(key int, value int)  {
	if value, ok := this.hashSet[key]; ok && value != -1 {
		index := this.listIndex[key]
		this.list[index] = -1
	} else {
		if this.count >= this.capacity {
			for ; this.index < this.length ; this.index ++ {
				if this.list[this.index] != -1 {
					this.hashSet[this.list[this.index]] = -1
					this.list[this.index] = -1
					break
				}
			}
		} else {
			this.count ++
		}
	}
	this.length ++
	this.list = append(this.list, key)
	this.listIndex[key] = this.length - 1
	this.hashSet[key] = value
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */