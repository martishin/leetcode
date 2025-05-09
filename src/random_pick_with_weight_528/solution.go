package random_pick_with_weight_528

import (
	"fmt"
	"math/rand"
)

type Solution struct {
	prefixSum []int
	totalSum  int
}

func Constructor(w []int) Solution {
	prefixSum := make([]int, 0, len(w))
	sum := 0

	for _, val := range w {
		sum += val
		prefixSum = append(prefixSum, sum)
	}

	return Solution{prefixSum: prefixSum, totalSum: sum}
}

func (this *Solution) PickIndex() int {
	target := rand.Intn(this.totalSum)

	l, r := 0, len(this.prefixSum)

	for l < r {
		m := l + (r-l)/2
		if this.prefixSum[m] > target {
			r = m
		} else {
			l = m + 1
		}
	}

	return l
}

func Test() {
	obj := Constructor([]int{1})
	fmt.Println(obj.PickIndex())
}
