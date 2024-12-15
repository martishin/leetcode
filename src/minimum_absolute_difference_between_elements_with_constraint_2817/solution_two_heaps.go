package minimum_absolute_difference_between_elements_with_constraint_2817

import (
	"container/heap"
	"fmt"
	"math"
	"sort"
)

type Pair struct {
	idx int
	val int
}

type MinHeap []Pair

func (h MinHeap) Len() int {
	return len(h)
}

func (h MinHeap) Less(i, j int) bool {
	return h[i].idx < h[j].idx
}

func (h MinHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *MinHeap) Push(x any) {
	*h = append(*h, x.(Pair))
}

func (h *MinHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func minAbsoluteDifferenceHeap(nums []int, x int) int {
	sortedList := make([]Pair, len(nums))
	for i, val := range nums {
		sortedList[i] = Pair{idx: i, val: val}
	}
	sort.Slice(sortedList, func(i, j int) bool {
		if sortedList[i].val == sortedList[j].val {
			return sortedList[i].idx < sortedList[j].idx
		}
		return sortedList[i].val < sortedList[j].val
	})

	minHeap := &MinHeap{}
	maxHeap := &MinHeap{}

	result := math.MaxInt
	for _, pair := range sortedList {
		val, idx := pair.val, pair.idx
		heap.Push(minHeap, Pair{idx: idx, val: val})
		heap.Push(maxHeap, Pair{idx: -idx, val: val})

		for minHeap.Len() > 0 && (*minHeap)[0].idx <= idx-x {
			top := heap.Pop(minHeap).(Pair)
			result = min(result, val-top.val)
		}

		for maxHeap.Len() > 0 && (*maxHeap)[0].idx <= -idx-x {
			top := heap.Pop(maxHeap).(Pair)
			result = min(result, val-top.val)
		}
	}

	return result
}

func TestHeap() {
	fmt.Println(minAbsoluteDifferenceHeap([]int{4, 3, 2, 4}, 2))
	fmt.Println(minAbsoluteDifferenceHeap([]int{5, 3, 2, 10, 15}, 1))
	fmt.Println(minAbsoluteDifferenceHeap([]int{1, 2, 3, 4}, 3))
	fmt.Println(minAbsoluteDifferenceHeap([]int{14, 111, 16}, 1))
}
