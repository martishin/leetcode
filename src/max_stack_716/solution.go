package max_stack_716

import (
	"container/heap"
	"container/list"
	"fmt"
)

type MaxHeap []*list.Element

func (h MaxHeap) Len() int           { return len(h) }
func (h MaxHeap) Less(i, j int) bool { return h[i].Value.(int) >= h[j].Value.(int) }
func (h MaxHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *MaxHeap) Push(x any) {
	*h = append(*h, x.(*list.Element))
}
func (h *MaxHeap) Pop() any {
	n := h.Len()
	el := (*h)[n-1]
	*h = (*h)[:n-1]

	return el
}

type MaxStack struct {
	heap  *MaxHeap
	stack *list.List
}

func Constructor() MaxStack {
	h := &MaxHeap{}
	heap.Init(h)

	s := list.New()
	s.PushFront(1)

	return MaxStack{
		heap:  h,
		stack: s,
	}
}

// O(log n)
func (s *MaxStack) Push(x int) {
	el := s.stack.PushFront(x)
	heap.Push(s.heap, el)
}

// O(1)
func (s *MaxStack) Pop() int {
	el := s.stack.Front()
	s.stack.Remove(el)

	return el.Value.(int)
}

// O(1)
func (s *MaxStack) Top() int {
	el := s.stack.Front()

	return el.Value.(int)
}

// O(log n)
func (s *MaxStack) PeekMax() int {
	el := heap.Pop(s.heap).(*list.Element)
	for el.Next() == nil {
		el = heap.Pop(s.heap).(*list.Element)
	}

	heap.Push(s.heap, el)
	return el.Value.(int)
}

// O(log n)
func (s *MaxStack) PopMax() int {
	el := heap.Pop(s.heap).(*list.Element)
	for el.Next() == nil {
		el = heap.Pop(s.heap).(*list.Element)
	}

	s.stack.Remove(el)
	return el.Value.(int)
}

func Test() {
	obj := Constructor()
	obj.Push(5)
	obj.Push(1)
	obj.Push(5)
	fmt.Println(obj.Top())
	fmt.Println(obj.PopMax())
	fmt.Println(obj.Top())
	fmt.Println(obj.PeekMax())
	fmt.Println(obj.Pop())
	fmt.Println(obj.Top())
}
