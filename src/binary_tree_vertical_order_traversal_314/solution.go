package binary_tree_vertical_order_traversal_314

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type NodeAndCol struct {
	Node *TreeNode
	Col  int
}

func verticalOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}

	columnToVals := make(map[int][]int)
	minCol := 0
	maxCol := 0

	queue := []NodeAndCol{{root, 0}}

	for len(queue) > 0 {
		nodeAndCol := queue[0]
		queue = queue[1:]

		node, col := nodeAndCol.Node, nodeAndCol.Col

		minCol = min(minCol, col)
		maxCol = max(maxCol, col)

		columnToVals[col] = append(columnToVals[col], node.Val)

		if node.Left != nil {
			queue = append(queue, NodeAndCol{node.Left, col - 1})
		}

		if node.Right != nil {
			queue = append(queue, NodeAndCol{node.Right, col + 1})
		}
	}

	result := [][]int{}
	for i := minCol; i <= maxCol; i++ {
		result = append(result, columnToVals[i])
	}

	return result
}

func Test() {
	head := &TreeNode{3,
		&TreeNode{9, nil, nil},
		&TreeNode{20,
			&TreeNode{15, nil, nil},
			&TreeNode{7, nil, nil},
		},
	}
	fmt.Println(verticalOrder(head))
}
