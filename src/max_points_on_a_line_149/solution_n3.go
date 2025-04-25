package max_points_on_a_line_149

import "fmt"

func checkPointsOnLine(p1, p2, p3 []int) bool {
	return (p2[1]-p1[1])*(p2[0]-p3[0]) == (p2[1]-p3[1])*(p2[0]-p1[0])
}

func maxPointsN3(points [][]int) int {
	n := len(points)

	if n == 0 {
		return 0
	}

	resultCount := 1
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			currentCount := 2
			for k := j + 1; k < n; k++ {
				if checkPointsOnLine(points[i], points[j], points[k]) {
					currentCount++
				}
			}
			resultCount = max(resultCount, currentCount)
		}
	}

	return resultCount
}

func TestN3() {
	fmt.Println(maxPointsN3([][]int{
		{1, 1},
		{2, 2},
		{3, 3},
	}))

	fmt.Println(maxPointsN3([][]int{
		{1, 1},
		{3, 2},
		{5, 3},
		{4, 1},
		{2, 3},
		{1, 4},
	}))
}
