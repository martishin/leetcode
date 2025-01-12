package populating_next_right_pointers_in_each_node_ii_117

func connectLL(root *Node) *Node {
	if root == nil {
		return root
	}

	leftmost := root
	var prev *Node = nil
	var curr *Node = nil

	for leftmost != nil {
		prev = nil
		curr = leftmost

		leftmost = nil

		for curr != nil {
			prev, leftmost = processChild(curr.Left, prev, leftmost)
			prev, leftmost = processChild(curr.Right, prev, leftmost)

			curr = curr.Next
		}
	}
	return root
}

func processChild(childNode *Node, prev *Node, leftmost *Node) (*Node, *Node) {
	if childNode != nil {
		if prev != nil {
			prev.Next = childNode
		} else {
			leftmost = childNode
		}
		prev = childNode
	}
	return prev, leftmost
}
