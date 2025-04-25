package subsets_78

import "fmt"

func subsetsMask(nums []int) [][]int {
	n := len(nums)
	total := 1 << n

	result := make([][]int, total)

	for mask := range total {
		subset := []int{}

		for i, num := range nums {
			if (1<<i)&mask != 0 {
				subset = append(subset, num)
			}
		}
		result[mask] = subset
	}

	return result
}

func TestMask() {
	fmt.Println(subsetsMask([]int{1, 2, 3}))
	fmt.Println(subsetsMask([]int{0}))
}
