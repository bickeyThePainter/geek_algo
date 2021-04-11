package Week01

type MyCircularDequeLinkedList struct {
	head, tail *Node
	capacity   int
	used       int
}

type Node struct {
	prev, next *Node
	val        int
}

func NewNode(prev, next *Node, val int) *Node {
	return &Node{prev, next, val}
}

/** Initialize your data structure here. Set the size of the deque to be k. */
func ConstructorMyCircularDequeLinkedList(k int) MyCircularDequeLinkedList {
	res := MyCircularDequeLinkedList{NewNode(nil, nil, 0), NewNode(nil, nil, 0), k, 0}
	res.head.next, res.tail.prev = res.tail, res.head
	return res
}

/** Adds an item at the front of Deque. Return true if the operation is successful. */
func (this *MyCircularDequeLinkedList) InsertFront(value int) bool {
	if this.IsFull() {
		return false
	}
	cur := NewNode(nil, nil, value)
	this.head.next, this.head.next.prev, cur.next, cur.prev = cur, cur, this.head.next, this.head
	this.used++
	return true
}

/** Adds an item at the rear of Deque. Return true if the operation is successful. */
func (this *MyCircularDequeLinkedList) InsertLast(value int) bool {
	if this.IsFull() {
		return false
	}
	cur := NewNode(nil, nil, value)

	this.tail.prev, this.tail.prev.next, cur.next, cur.prev = cur, cur, this.tail, this.tail.prev
	this.used++
	return true
}

/** Deletes an item from the front of Deque. Return true if the operation is successful. */
func (this *MyCircularDequeLinkedList) DeleteFront() bool {
	if this.IsEmpty() {
		return false
	}
	this.head.next, this.head.next.next.prev = this.head.next.next, this.head
	this.used--
	return true
}

/** Deletes an item from the rear of Deque. Return true if the operation is successful. */
func (this *MyCircularDequeLinkedList) DeleteLast() bool {
	if this.IsEmpty() {
		return false
	}
	this.tail.prev, this.tail.prev.prev.next = this.tail.prev.prev, this.tail
	this.used--
	return true
}

/** Get the front item from the deque. */
func (this *MyCircularDequeLinkedList) GetFront() int {
	if this.IsEmpty() {
		return -1
	}
	return this.head.next.val
}

/** Get the last item from the deque. */
func (this *MyCircularDequeLinkedList) GetRear() int {
	if this.IsEmpty() {
		return -1
	}
	return this.tail.prev.val
}

/** Checks whether the circular deque is empty or not. */
func (this *MyCircularDequeLinkedList) IsEmpty() bool {
	return this.used == 0
}

/** Checks whether the circular deque is full or not. */
func (this *MyCircularDequeLinkedList) IsFull() bool {
	return this.used == this.capacity
}
