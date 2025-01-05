package populating_next_right_pointers_in_each_node_ii_117

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

func connect(root *Node) *Node {
	if root == nil {
		return nil
	}

	queue := []*Node{}

	queue = append(queue, root)
	queue = append(queue, nil)

	for len(queue) > 0 {
		n := len(queue)

		for range n {
			node := queue[0]
			queue = queue[1:]

			if node == nil {
				continue
			}

			if len(queue) > 0 {
				node.Next = queue[0]
			}

			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}

		if len(queue) > 0 {
			queue = append(queue, nil)
		}
	}

	return root
}
