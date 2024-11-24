package house_robber_ii_213

import "fmt"

func rob(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}

	if len(nums) == 2 {
		return max(nums[0], nums[1])
	}

	return max(robHelper(nums[1:]), robHelper(nums[:len(nums)-1]))
}

func robHelper(nums []int) int {
	prevPrev := nums[0]
	prev := max(nums[0], nums[1])

	for i := 2; i < len(nums); i++ {
		current := max(prev, prevPrev+nums[i])
		prevPrev = prev
		prev = current
	}

	return max(prev, prevPrev)
}

func Test() {
	fmt.Println(rob([]int{1, 2, 3, 1}))
	fmt.Println(rob([]int{2, 7, 9, 3, 1}))
}
