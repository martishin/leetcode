package largest_rectangle_in_histogram_84

import "fmt"

func largestRectangleAreaStack(heights []int) int {
	stack := []int{-1}
	maxArea := 0

	for i := range heights {
		for len(stack) > 1 && heights[stack[len(stack)-1]] > heights[i] {
			currentHeight := heights[stack[len(stack)-1]]
			currentWidth := i - stack[len(stack)-2] - 1
			currentArea := currentHeight * currentWidth
			if currentArea > maxArea {
				maxArea = currentArea
			}

			stack = stack[:len(stack)-1]
		}

		stack = append(stack, i)
	}

	for len(stack) > 1 {
		currentHeight := heights[stack[len(stack)-1]]
		currentWidth := len(heights) - stack[len(stack)-2] - 1
		currentArea := currentHeight * currentWidth
		if currentArea > maxArea {
			maxArea = currentArea
		}

		stack = stack[:len(stack)-1]
	}

	return maxArea
}

func TestStack() {
	fmt.Println(largestRectangleAreaStack([]int{2, 1, 5, 6, 2, 3}))
	fmt.Println(largestRectangleAreaStack([]int{2, 4}))
}
