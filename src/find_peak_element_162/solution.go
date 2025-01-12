package find_peak_element_162

import (
	"fmt"
	"math"
)

func findPeakElement(nums []int) int {
	n := len(nums)

	if n == 0 {
		return math.MaxInt32
	} else if n == 1 {
		return 0
	} else if nums[0] > nums[1] {
		return 0
	} else if nums[n-1] > nums[n-2] {
		return n - 1
	}

	l := 0
	r := n
	for l < r {
		m := (l + r) / 2
		if nums[m] < nums[m+1] {
			l = m
		} else if nums[m-1] > nums[m] {
			r = m
		} else {
			return m
		}
	}

	return r - 1
}

func Test() {
	fmt.Println(findPeakElement([]int{1, 2, 3, 1}))
	fmt.Println(findPeakElement([]int{1, 2, 1, 3, 5, 6, 4}))
}
