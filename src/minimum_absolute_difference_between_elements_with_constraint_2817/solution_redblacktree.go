package minimum_absolute_difference_between_elements_with_constraint_2817

import (
	"fmt"

	"github.com/emirpasic/gods/trees/redblacktree"
)

func minAbsoluteDifferenceRbTree(nums []int, x int) int {
	rbt := redblacktree.NewWithIntComparator()
	ans := 1 << 30
	for i := x; i < len(nums); i++ {
		rbt.Put(nums[i-x], nil)
		c, _ := rbt.Ceiling(nums[i])
		f, _ := rbt.Floor(nums[i])
		if c != nil {
			ans = min(ans, c.Key.(int)-nums[i])
		}
		if f != nil {
			ans = min(ans, nums[i]-f.Key.(int))
		}
	}
	return ans
}

func TestRbTree() {
	fmt.Println(minAbsoluteDifferenceRbTree([]int{4, 3, 2, 4}, 2))
	fmt.Println(minAbsoluteDifferenceRbTree([]int{5, 3, 2, 10, 15}, 1))
	fmt.Println(minAbsoluteDifferenceRbTree([]int{1, 2, 3, 4}, 3))
	fmt.Println(minAbsoluteDifferenceRbTree([]int{14, 111, 16}, 1))
}
