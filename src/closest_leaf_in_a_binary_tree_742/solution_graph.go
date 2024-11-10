package closest_leaf_in_a_binary_tree_742

import (
	"fmt"
)

func findClosestLeaf(root *TreeNode, k int) int {
	// Map to represent the undirected graph
	graph := make(map[*TreeNode][]*TreeNode)
	var targetNode *TreeNode

	// Build the graph and find the target node
	var buildGraph func(node, parent *TreeNode)
	buildGraph = func(node, parent *TreeNode) {
		if node == nil {
			return
		}

		if parent != nil {
			graph[node] = append(graph[node], parent)
			graph[parent] = append(graph[parent], node)
		}

		if node.Val == k {
			targetNode = node
		}

		buildGraph(node.Left, node)
		buildGraph(node.Right, node)
	}

	buildGraph(root, nil)

	if targetNode == nil {
		return -1
	}

	// Perform BFS to find the closest leaf
	queue := []*TreeNode{targetNode}
	visited := make(map[*TreeNode]bool)
	visited[targetNode] = true

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]

		// Check if the node is a leaf node
		if node.Left == nil && node.Right == nil {
			return node.Val
		}

		for _, neighbor := range graph[node] {
			if !visited[neighbor] {
				visited[neighbor] = true
				queue = append(queue, neighbor)
			}
		}
	}

	// In case the root itself is the only node and is a leaf
	if root.Left == nil && root.Right == nil && root.Val == k {
		return root.Val
	}

	return -1
}

func TestGraph() {
	// Test case 1
	root1 := &TreeNode{Val: 1}
	root1.Left = &TreeNode{Val: 3}
	root1.Right = &TreeNode{Val: 2}

	result1 := findClosestLeaf(root1, 1)
	fmt.Printf("Closest leaf to node with value 1: %d\n", result1) // Expected: 3 or 2

	// Test case 2
	root2 := &TreeNode{Val: 1}
	root2.Left = &TreeNode{Val: 2}
	root2.Right = &TreeNode{Val: 3}
	root2.Left.Left = &TreeNode{Val: 4}
	root2.Left.Left.Left = &TreeNode{Val: 5}
	root2.Left.Left.Left.Left = &TreeNode{Val: 6}

	result2 := findClosestLeaf(root2, 2)
	fmt.Printf("Closest leaf to node with value 2: %d\n", result2) // Expected: 3
}
