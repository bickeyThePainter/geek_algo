package Week01

// Design-circular-deque
// 1. array
// 2. linked-list
type MyArrayCircularDeque struct {
	capacity []int
	k int
	used int
	rear int
	front int
}


/** Initialize your data structure here. Set the size of the deque to be k. */
func ConstructorMyArrayCircularDeque(k int) MyArrayCircularDeque {
	return MyArrayCircularDeque{make([]int,k,k),k,0,0,k-1}
}


/** Adds an item at the front of Deque. Return true if the operation is successful. */
func (this *MyArrayCircularDeque) InsertFront(value int) bool {
	if this.IsFull(){
		return false
	}
	this.capacity[this.front]=value
	this.front=(this.front-1+this.k)%this.k
	this.used++
	return true
}


/** Adds an item at the rear of Deque. Return true if the operation is successful. */
func (this *MyArrayCircularDeque) InsertLast(value int) bool {
	if this.IsFull(){
		return false
	}
	this.capacity[this.rear]=value
	this.rear=(this.rear+1)%this.k
	this.used++
	return true
}


/** Deletes an item from the front of Deque. Return true if the operation is successful. */
func (this *MyArrayCircularDeque) DeleteFront() bool {
	if this.IsEmpty(){
		return false
	}
	this.front=(this.front+1)%this.k
	this.used--
	return true
}


/** Deletes an item from the rear of Deque. Return true if the operation is successful. */
func (this *MyArrayCircularDeque) DeleteLast() bool {
	if this.IsEmpty(){
		return false
	}
	this.rear=(this.rear-1+this.k)%this.k
	this.used--
	return true
}


/** Get the front item from the deque. */
func (this *MyArrayCircularDeque) GetFront() int {
	if this.IsEmpty(){
		return -1
	}
	return this.capacity[(this.front+1)%this.k]
}


/** Get the last item from the deque. */
func (this *MyArrayCircularDeque) GetRear() int {
	if this.IsEmpty(){
		return -1
	}
	return this.capacity[(this.rear-1+this.k)%this.k]
}


/** Checks whether the circular deque is empty or not. */
func (this *MyArrayCircularDeque) IsEmpty() bool {
	return this.used==0
}


/** Checks whether the circular deque is full or not. */
func (this *MyArrayCircularDeque) IsFull() bool {
	return this.used==this.k
}

