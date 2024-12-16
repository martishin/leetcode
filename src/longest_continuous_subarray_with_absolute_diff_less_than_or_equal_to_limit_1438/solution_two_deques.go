package longest_continuous_subarray_with_absolute_diff_less_than_or_equal_to_limit_1438

import "fmt"

type Deque[T any] struct {
	data []T
}

func (d *Deque[T]) PushBack(value T) {
	d.data = append(d.data, value)
}

func (d *Deque[T]) PopBack() {
	if len(d.data) > 0 {
		d.data = d.data[:len(d.data)-1]
	}
}

func (d *Deque[T]) PopFront() {
	if len(d.data) > 0 {
		d.data = d.data[1:]
	}
}

func (d *Deque[T]) Front() T {
	if len(d.data) > 0 {
		return d.data[0]
	}
	var zero T
	return zero
}

func (d *Deque[T]) Back() T {
	if len(d.data) > 0 {
		return d.data[len(d.data)-1]
	}
	var zero T
	return zero
}

func (d *Deque[T]) IsEmpty() bool {
	return len(d.data) == 0
}

func longestSubarrayTwoDeques(nums []int, limit int) int {
	minDeque := &Deque[int]{}
	maxDeque := &Deque[int]{}

	left, longestSubarray := 0, 0
	for right, num := range nums {
		for !minDeque.IsEmpty() && minDeque.Back() > num {
			minDeque.PopBack()
		}
		minDeque.PushBack(num)

		for !maxDeque.IsEmpty() && maxDeque.Back() < num {
			maxDeque.PopBack()
		}
		maxDeque.PushBack(num)

		for maxDeque.Front()-minDeque.Front() > limit {
			if minDeque.Front() == nums[left] {
				minDeque.PopFront()
			}
			if maxDeque.Front() == nums[left] {
				maxDeque.PopFront()
			}
			left++
		}

		longestSubarray = max(longestSubarray, right-left+1)
	}

	return longestSubarray
}

// func longestSubarrayTwoDeques(nums []int, limit int) int {
// 	var maxDeque []int
// 	var minDeque []int
//
// 	left, longestSubarray := 0, 0
//
// 	for right, num := range nums {
// 		for len(maxDeque) > 0 && maxDeque[len(maxDeque)-1] < num {
// 			maxDeque = maxDeque[:len(maxDeque)-1]
// 		}
// 		maxDeque = append(maxDeque, num)
//
// 		for len(minDeque) > 0 && minDeque[len(minDeque)-1] > num {
// 			minDeque = minDeque[:len(minDeque)-1]
// 		}
// 		minDeque = append(minDeque, num)
//
// 		for maxDeque[0]-minDeque[0] > limit {
// 			if maxDeque[0] == nums[left] {
// 				maxDeque = maxDeque[1:]
// 			}
// 			if minDeque[0] == nums[left] {
// 				minDeque = minDeque[1:]
// 			}
// 			left++
// 		}
// 		longestSubarray = max(longestSubarray, right-left+1)
// 	}
//
// 	return longestSubarray
// }

func TestTwoDeques() {
	fmt.Println(longestSubarrayTwoDeques([]int{8, 2, 4, 7}, 4))
	fmt.Println(longestSubarrayTwoDeques([]int{10, 1, 2, 4, 7, 2}, 5))
	fmt.Println(longestSubarrayTwoDeques([]int{4, 2, 2, 2, 4, 4, 2, 2}, 0))
}
