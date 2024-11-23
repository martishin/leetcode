package two_sum_1

import (
	"fmt"
	"sort"
)

func twoSumIteration(nums []int, target int) []int {
	numsWithIndices := make([]Pair, len(nums))

	for idx, num := range nums {
		numsWithIndices[idx] = Pair{
			num: num,
			idx: idx,
		}
	}

	sort.Slice(numsWithIndices, func(i, j int) bool {
		return numsWithIndices[i].num < numsWithIndices[j].num
	})

	l := 0
	r := len(numsWithIndices) - 1

	for l < r {
		sum := numsWithIndices[l].num + numsWithIndices[r].num
		if sum < target {
			l++
		} else if sum > target {
			r--
		} else {
			return []int{numsWithIndices[l].idx, numsWithIndices[r].idx}
		}
	}

	return []int{-1, -1}
}

func TestIteration() {
	fmt.Println(twoSumIteration([]int{3, 2, 4}, 6))
}
