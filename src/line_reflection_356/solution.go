package line_reflection_356

import (
	"fmt"
	"math"
)

type Point struct {
	x int
	y int
}

func isReflected(points [][]int) bool {
	pointSet := make(map[Point]struct{})
	minX, maxX := math.MaxInt, math.MinInt

	for _, point := range points {
		pointSet[Point{point[0], point[1]}] = struct{}{}
		minX = min(minX, point[0])
		maxX = max(maxX, point[0])
	}

	candidate := float64(minX+maxX) / 2.0

	for point := range pointSet {
		pairPoint := Point{int(2*candidate) - point.x, point.y}
		if _, ok := pointSet[pairPoint]; !ok {
			return false
		}
	}

	return true
}

func Test() {
	fmt.Println(isReflected([][]int{{1, 1}, {-1, 1}}))
	fmt.Println(isReflected([][]int{{1, 1}, {-1, -1}}))
}
