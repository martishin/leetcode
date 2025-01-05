package construct_binary_tree_from_inorder_and_postorder_traversal_106

import "fmt"

func buildTreeMap(inorder []int, postorder []int) *TreeNode {
	idxMap := map[int]int{}
	postIdx := len(postorder) - 1
	var helper func(inLeft int, inRight int) *TreeNode
	helper = func(inLeft int, inRight int) *TreeNode {
		if inLeft > inRight {
			return nil
		}
		rootVal := postorder[postIdx]
		root := &TreeNode{Val: rootVal}
		index := idxMap[rootVal]
		postIdx--
		root.Right = helper(index+1, inRight)
		root.Left = helper(inLeft, index-1)
		return root
	}
	for i := 0; i < len(inorder); i++ {
		idxMap[inorder[i]] = i
	}
	return helper(0, len(inorder)-1)
}

func TestMap() {
	fmt.Println(bfsTraversal(buildTreeMap([]int{9, 3, 15, 20, 7}, []int{9, 15, 7, 20, 3})))
	fmt.Println(bfsTraversal(buildTreeMap([]int{-1}, []int{-1})))
}
