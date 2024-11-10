package closest_leaf_in_a_binary_tree_742

import (
	"fmt"
	"math"
)

// TreeNode defines a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Solution struct to hold memoization and path information.
type Solution struct {
	memo map[*TreeNode]closestLeafInfo
	path []*TreeNode
}

type closestLeafInfo struct {
	distance int
	leaf     *TreeNode
}

func Constructor() *Solution {
	return &Solution{
		memo: make(map[*TreeNode]closestLeafInfo),
		path: make([]*TreeNode, 0),
	}
}

func (s *Solution) findClosestLeaf(root *TreeNode, k int) int {
	s.dfs(root, k)

	minDistance := math.MaxInt32
	var closestLeaf *TreeNode

	for i, node := range s.path {
		info := s.closestLeaf(node)
		newDistance := info.distance + (len(s.path) - 1 - i)
		if newDistance < minDistance {
			minDistance = newDistance
			closestLeaf = info.leaf
		}
	}

	if closestLeaf != nil {
		return closestLeaf.Val
	}
	return -1
}

func (s *Solution) dfs(node *TreeNode, k int) bool {
	if node == nil {
		return false
	}

	s.path = append(s.path, node)
	if node.Val == k {
		return true
	}

	if s.dfs(node.Left, k) || s.dfs(node.Right, k) {
		return true
	}

	s.path = s.path[:len(s.path)-1]
	return false
}

func (s *Solution) closestLeaf(node *TreeNode) closestLeafInfo {
	if node == nil {
		return closestLeafInfo{distance: math.MaxInt32, leaf: nil}
	}

	if info, exists := s.memo[node]; exists {
		return info
	}

	if node.Left == nil && node.Right == nil {
		// Node is a leaf
		info := closestLeafInfo{distance: 0, leaf: node}
		s.memo[node] = info
		return info
	}

	leftInfo := s.closestLeaf(node.Left)
	rightInfo := s.closestLeaf(node.Right)

	var result closestLeafInfo
	if leftInfo.distance < rightInfo.distance {
		result = closestLeafInfo{distance: leftInfo.distance + 1, leaf: leftInfo.leaf}
	} else {
		result = closestLeafInfo{distance: rightInfo.distance + 1, leaf: rightInfo.leaf}
	}

	s.memo[node] = result
	return result
}

// Helper function to build a binary tree from a list of values.
func buildTreeFromList(values []interface{}) *TreeNode {
	if len(values) == 0 {
		return nil
	}

	rootValue, ok := values[0].(int)
	if !ok {
		return nil
	}
	root := &TreeNode{Val: rootValue}
	queue := []*TreeNode{root}
	values = values[1:]

	for len(values) > 0 && len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		if len(values) > 0 {
			if values[0] != nil {
				val := values[0].(int)
				current.Left = &TreeNode{Val: val}
				queue = append(queue, current.Left)
			}
			values = values[1:]
		}

		if len(values) > 0 {
			if values[0] != nil {
				val := values[0].(int)
				current.Right = &TreeNode{Val: val}
				queue = append(queue, current.Right)
			}
			values = values[1:]
		}
	}

	return root
}

func Test() {
	// Test case 1
	// Tree:
	//     1
	//    / \
	//   3   2
	root1 := &TreeNode{Val: 1}
	root1.Left = &TreeNode{Val: 3}
	root1.Right = &TreeNode{Val: 2}

	solution1 := Constructor()
	result1 := solution1.findClosestLeaf(root1, 1)
	fmt.Printf("Closest leaf to node with value 1: %d\n", result1) // Expected: 3 or 2

	// Test case 2
	// Tree:
	//         1
	//        / \
	//       2   3
	//      /
	//     4
	//    /
	//   5
	//  /
	// 6
	root2 := &TreeNode{Val: 1}
	root2.Left = &TreeNode{Val: 2}
	root2.Right = &TreeNode{Val: 3}
	root2.Left.Left = &TreeNode{Val: 4}
	root2.Left.Left.Left = &TreeNode{Val: 5}
	root2.Left.Left.Left.Left = &TreeNode{Val: 6}

	solution2 := Constructor()
	result2 := solution2.findClosestLeaf(root2, 2)
	fmt.Printf("Closest leaf to node with value 2: %d\n", result2) // Expected: 3
}
