package search_in_rotated_sorted_array_33

import "fmt"

func search(nums []int, target int) int {
	n := len(nums)
	if n == 0 {
		return -1
	}

	l := 0
	r := n
	for l < r {
		m := l + (r-l)/2
		if nums[m] == target {
			return m
		} else if nums[l] <= nums[m] {
			if nums[l] <= target && target < nums[m] {
				r = m
			} else {
				l = m + 1
			}
		} else {
			if nums[m] < target && target <= nums[r-1] {
				l = m + 1
			} else {
				r = m
			}
		}
	}

	return -1
}

func Test() {
	fmt.Println(search([]int{4, 5, 6, 7, 0, 1, 2}, 0))
	fmt.Println(search([]int{4, 5, 6, 7, 0, 1, 2}, 3))
	fmt.Println(search([]int{1}, 0))
}
