package Week09

// 最简单方式 java->linkedHashMap(accessOrder=true)
// 双端队列维护 lru规则 map加快查找顺序 O(N)->O(1)
type LRUCache struct {
	head, tail *Node
	data       map[int]*Node
	capacity   int
}

type Node struct {
	key, value int
	next, prev *Node
}

func NewNode(key, value int) *Node {
	return &Node{
		key, value, nil, nil,
	}
}
func (n *Node) del() {
	n.prev.next = n.next
	n.next.prev = n.prev
	n.next, n.prev = nil, nil
}
func (n *Node) add(target *Node) {
	target.prev, target.next = n, n.next
	n.next, n.next.prev = target, target
}

func Constructor(capacity int) LRUCache {
	lru := LRUCache{
		head:     NewNode(0, 0),
		tail:     NewNode(0, 0),
		data:     make(map[int]*Node),
		capacity: capacity,
	}
	lru.head.next = lru.tail
	lru.tail.prev = lru.head
	return lru
}

func (this *LRUCache) Get(key int) int {
	cur, ok := this.data[key]
	if !ok {
		return -1
	} else {
		cur.del()
		this.head.add(cur)
		return cur.value
	}
}

func (this *LRUCache) Put(key int, value int) {
	if this.Get(key) == -1 {
		if len(this.data) == this.capacity {
			toDelete := this.tail.prev
			toDelete.del()
			delete(this.data, toDelete.key)
		}
		cur := NewNode(key, value)
		this.head.add(cur)
		this.data[key] = cur
	} else {
		cur := this.data[key]
		cur.value = value
	}
}
