package Week02

import (
	"fmt"
	"testing"
)

type Heap struct {
	data   []int
	offset int
}

func NewHeap(data []int) *Heap {
	heap := &Heap{data: data, offset: len(data)}
	heap.heapify()
	return heap
}

func (h *Heap) down(i int) {
	for h.firstBorn(i) < h.offset {
		maxChild := h.maxChild(i)
		if h.data[maxChild] > h.data[i] {
			h.data[maxChild], h.data[i] = h.data[i], h.data[maxChild]
			i = maxChild
		} else {
			break
		}
	}
}

func (h *Heap) heapify() {
	for i := (h.offset - 1) / 2; i >= 0; i-- {
		h.down(i)
	}
}

func (h *Heap) sort() []int {
	for h.offset > 0 {
		h.data[0], h.data[h.offset-1] = h.data[h.offset-1], h.data[0]
		h.offset--
		h.down(0)
	}
	return h.data
}

func (h *Heap) maxChild(i int) int {
	first := h.firstBorn(i)
	second := h.secondBorn(i)
	if first >= h.offset {
		return -1
	}
	if second >= h.offset {
		return first
	}
	if h.data[first] >= h.data[second] {
		return first
	} else {
		return second
	}
}

func (h *Heap) firstBorn(i int) int {
	return i*2 + 1
}

func (h *Heap) secondBorn(i int) int {
	return i*2 + 2
}

func TestHeap(t *testing.T) {
	h := NewHeap([]int{6, 1, 3, 2, 4, 7, 9, 1000, 88, 98, 2, 334})
	fmt.Println(h.sort())
}
