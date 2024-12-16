package longest_continuous_subarray_with_absolute_diff_less_than_or_equal_to_limit_1438

import (
	"container/heap"
	"fmt"
)

type Pair struct {
	val int
	idx int
}

type MinHeap []Pair

func (h MinHeap) Len() int {
	return len(h)
}

func (h MinHeap) Less(i, j int) bool {
	return h[i].val < h[j].val
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
	*h = old[:n-1]
	return x
}

func (h *MinHeap) Peek() Pair {
	return (*h)[0]
}

func longestSubarrayTwoHeaps(nums []int, limit int) int {
	minHeap := &MinHeap{}
	maxHeap := &MinHeap{}

	left, longestSubarray := 0, 0
	for right, num := range nums {
		heap.Push(maxHeap, Pair{val: -num, idx: right})
		heap.Push(minHeap, Pair{val: num, idx: right})

		for -maxHeap.Peek().val-minHeap.Peek().val > limit {
			left = min(maxHeap.Peek().idx, minHeap.Peek().idx) + 1

			for maxHeap.Peek().idx < left {
				heap.Pop(maxHeap)
			}
			for minHeap.Peek().idx < left {
				heap.Pop(minHeap)
			}
		}

		longestSubarray = max(longestSubarray, right-left+1)
	}

	return longestSubarray
}

func TestTwoHeaps() {
	fmt.Println(longestSubarrayTwoHeaps([]int{8, 2, 4, 7}, 4))
	fmt.Println(longestSubarrayTwoHeaps([]int{10, 1, 2, 4, 7, 2}, 5))
	fmt.Println(longestSubarrayTwoHeaps([]int{4, 2, 2, 2, 4, 4, 2, 2}, 0))
}
