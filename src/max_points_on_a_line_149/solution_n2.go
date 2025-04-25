package max_points_on_a_line_149

import (
	"fmt"
	"math"
)

func maxPointsN2(points [][]int) int {
	n := len(points)

	if n == 1 {
		return 1
	}

	result := 2
	for i := 0; i < n; i++ {
		atan2ToCount := make(map[float64]int)
		for j := 0; j < n; j++ {
			if i == j {
				continue
			}

			atan2 := math.Atan2(
				float64(points[i][1]-points[j][1]),
				float64(points[i][0]-points[j][0]),
			)
			atan2ToCount[atan2]++
		}

		for _, v := range atan2ToCount {
			result = max(result, v+1)
		}
	}

	return result
}

func TestN2() {
	fmt.Println(maxPointsN2([][]int{
		{1, 1},
		{2, 2},
		{3, 3},
	}))

	fmt.Println(maxPointsN2([][]int{
		{1, 1},
		{3, 2},
		{5, 3},
		{4, 1},
		{2, 3},
		{1, 4},
	}))
}
