package find_k_closest_elements_658

import (
	"math"
	"sort"
)

func findClosestElements(arr []int, k int, x int) []int {
	l := 0
	r := len(arr)

	for l < r {
		mid := (l + r) / 2
		if arr[mid] < x {
			l = mid + 1
		} else {
			r = mid
		}
	}

	l, r = l-1, l

	result := []int{}

	for i := 0; i < k; i++ {
		if l < 0 {
			result = append(result, arr[r])
			r++
		} else if r >= len(arr) {
			result = append(result, arr[l])
			l--
		} else if math.Abs(float64(arr[l]-x)) <= math.Abs(float64(arr[r]-x)) {
			result = append(result, arr[l])
			l--
		} else {
			result = append(result, arr[r])
			r++
		}
	}

	sort.Ints(result)

	return result
}
