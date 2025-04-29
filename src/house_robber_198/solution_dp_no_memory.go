package house_robber_198

import "fmt"

func robDpNM(nums []int) int {
	n := len(nums)

	if n == 0 {
		return 0
	} else if n == 1 {
		return nums[0]
	}

	prevPrev := nums[0]
	prev := max(nums[0], nums[1])

	for i := 2; i < n; i++ {
		tmp := max(prevPrev+nums[i], prev)
		prevPrev = prev
		prev = tmp
	}

	return prev
}

func TestDpNM() {
	fmt.Println(robDp([]int{1, 2, 3, 1}))
	fmt.Println(robDp([]int{2, 7, 9, 3, 1}))
}
