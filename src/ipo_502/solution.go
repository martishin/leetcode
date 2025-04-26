package ipo_502

import (
	"container/heap"
	"fmt"
	"sort"
)

type Project struct {
	profit  int
	capital int
}

type MaxHeapProfit []Project

func (h MaxHeapProfit) Len() int           { return len(h) }
func (h MaxHeapProfit) Less(i, j int) bool { return h[i].profit > h[j].profit }
func (h MaxHeapProfit) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *MaxHeapProfit) Push(x any) {
	*h = append(*h, x.(Project))
}
func (h *MaxHeapProfit) Pop() any {
	n := h.Len()
	x := (*h)[n-1]
	*h = (*h)[:n-1]

	return x
}

func findMaximizedCapital(k int, w int, profits []int, capital []int) int {
	n := len(profits)
	projectsSortedByCapital := []Project{}

	for i := 0; i < n; i++ {
		projectsSortedByCapital = append(projectsSortedByCapital, Project{profits[i], capital[i]})
	}

	sort.Slice(projectsSortedByCapital, func(i, j int) bool {
		return projectsSortedByCapital[i].capital < projectsSortedByCapital[j].capital
	})

	profitHeap := &MaxHeapProfit{}
	heap.Init(profitHeap)
	projectsPtr := 0
	for i := 0; i < k; i++ {
		for projectsPtr < n && projectsSortedByCapital[projectsPtr].capital <= w {
			heap.Push(profitHeap, projectsSortedByCapital[projectsPtr])
			projectsPtr++
		}

		if profitHeap.Len() == 0 {
			break
		}

		w += heap.Pop(profitHeap).(Project).profit
	}

	return w
}

func Test() {
	fmt.Println(findMaximizedCapital(2, 0, []int{1, 2, 3}, []int{0, 1, 1}))
	fmt.Println(findMaximizedCapital(3, 0, []int{1, 2, 3}, []int{0, 1, 2}))
}
