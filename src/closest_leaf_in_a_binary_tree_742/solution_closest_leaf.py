from typing import Optional, Tuple, Dict, List


# Definition for a binary tree node.
class TreeNode:
    def __init__(
        self,
        val: int = 0,
        left: Optional["TreeNode"] = None,
        right: Optional["TreeNode"] = None,
    ):
        self.val = val
        self.left = left
        self.right = right


class Solution:
    def __init__(self):
        self._memo: Dict[TreeNode, Tuple[int, Optional[TreeNode]]] = {}
        self._path: List[TreeNode] = []

    def closest_leaf(self, node: Optional[TreeNode]) -> Tuple[int, Optional[TreeNode]]:
        if node in self._memo:
            return self._memo[node]

        if not node:
            ans = float("inf"), None
        elif not node.left and not node.right:
            ans = 0, node
        else:
            d1, leaf1 = self.closest_leaf(node.left)
            d2, leaf2 = self.closest_leaf(node.right)
            print(
                f"node: {node.val}, d1: {d1}, leaf1: {leaf1}, d2: {d2}, leaf2: {leaf2}"
            )

            ans = min(d1, d2) + 1, (leaf1 if d1 < d2 else leaf2)

        self._memo[node] = ans
        return ans

    def dfs(self, node: Optional[TreeNode], k: int) -> bool:
        if not node:
            return False

        if node.val == k:
            self._path.append(node)
            return True

        self._path.append(node)
        if self.dfs(node.left, k) or self.dfs(node.right, k):
            return True

        self._path.pop()
        return False

    def findClosestLeaf(self, root: Optional[TreeNode], k: int) -> int:
        self.dfs(root, k)

        dist, leaf = float("inf"), None
        for i, node in enumerate(self._path):
            new_dist, new_leaf = self.closest_leaf(node)
            new_dist += len(self._path) - 1 - i
            if new_dist < dist:
                dist = new_dist
                leaf = new_leaf

        return leaf.val if leaf else -1


def test():
    # case 1
    root = TreeNode(1)
    root.right = TreeNode(2)
    root.left = TreeNode(3)
    print(Solution().findClosestLeaf(root, 1))

    # case 3
    root = TreeNode(1)
    root.right = TreeNode(3)
    root.left = TreeNode(2)
    root.left.left = TreeNode(4)
    root.left.left.left = TreeNode(5)
    root.left.left.left.left = TreeNode(6)
    print(Solution().findClosestLeaf(root, 2))
