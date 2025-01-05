package construct_binary_tree_from_inorder_and_postorder_traversal_106

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(inorder []int, postorder []int) *TreeNode {
	if len(postorder) == 0 {
		return nil
	}

	return buildTreeHelper(inorder, postorder)
}

func buildTreeHelper(inorder []int, postorder []int) *TreeNode {
	if len(postorder) == 0 {
		return nil
	}

	rootVal := postorder[len(postorder)-1]
	rootIdx := findRootNode(inorder, rootVal)
	root := &TreeNode{Val: rootVal}

	root.Left = buildTreeHelper(inorder[:rootIdx], postorder[:rootIdx])
	root.Right = buildTreeHelper(inorder[rootIdx+1:], postorder[rootIdx:len(postorder)-1])

	return root
}

func findRootNode(inorder []int, root int) int {
	for i, val := range inorder {
		if val == root {
			return i
		}
	}

	return -1
}

func bfsTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	var result []int

	var queue []*TreeNode
	queue = append(queue, root)

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]

		if node == nil {
			result = append(result, 0)
		} else {
			result = append(result, node.Val)
			queue = append(queue, node.Left)
			queue = append(queue, node.Right)
		}
	}

	for len(result) > 0 && result[len(result)-1] == 0 {
		result = result[:len(result)-1]
	}

	return result
}

func Test() {
	fmt.Println(bfsTraversal(buildTree([]int{9, 3, 15, 20, 7}, []int{9, 15, 7, 20, 3})))
	fmt.Println(bfsTraversal(buildTree([]int{-1}, []int{-1})))
}
