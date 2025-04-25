package largest_rectangle_in_histogram_84

import "fmt"

// SegTreeNode represents a node in the segment tree
type SegTreeNode struct {
	start, end  int
	minIndex    int
	left, right *SegTreeNode
}

// largestRectangleAreaSegTree is the main entry function
func largestRectangleAreaSegTree(heights []int) int {
	if len(heights) == 0 {
		return 0
	}
	root := buildSegmentTree(heights, 0, len(heights)-1)
	return calculateMax(heights, root, 0, len(heights)-1)
}

// buildSegmentTree builds a segment tree storing the index of the minimum element
func buildSegmentTree(heights []int, start, end int) *SegTreeNode {
	if start > end {
		return nil
	}
	node := &SegTreeNode{start: start, end: end}
	if start == end {
		node.minIndex = start
		return node
	}
	mid := (start + end) / 2
	node.left = buildSegmentTree(heights, start, mid)
	node.right = buildSegmentTree(heights, mid+1, end)
	leftMin := node.left.minIndex
	rightMin := node.right.minIndex
	if heights[leftMin] < heights[rightMin] {
		node.minIndex = leftMin
	} else {
		node.minIndex = rightMin
	}
	return node
}

// calculateMax recursively computes the max area
func calculateMax(heights []int, root *SegTreeNode, start, end int) int {
	if start > end {
		return 0
	}
	if start == end {
		return heights[start]
	}
	minIndex := query(root, heights, start, end)
	leftMax := calculateMax(heights, root, start, minIndex-1)
	rightMax := calculateMax(heights, root, minIndex+1, end)
	minMax := heights[minIndex] * (end - start + 1)
	return max(leftMax, rightMax, minMax)
}

// query returns the index of the minimum height in a range
func query(root *SegTreeNode, heights []int, start, end int) int {
	if root == nil || end < root.start || start > root.end {
		return -1
	}
	if start <= root.start && end >= root.end {
		return root.minIndex
	}
	leftMin := query(root.left, heights, start, end)
	rightMin := query(root.right, heights, start, end)
	if leftMin == -1 {
		return rightMin
	}
	if rightMin == -1 {
		return leftMin
	}
	if heights[leftMin] < heights[rightMin] {
		return leftMin
	}
	return rightMin
}

func TestSegTree() {
	fmt.Println(largestRectangleAreaSegTree([]int{2, 1, 5, 6, 2, 3}))
	fmt.Println(largestRectangleAreaSegTree([]int{2, 4}))
}
