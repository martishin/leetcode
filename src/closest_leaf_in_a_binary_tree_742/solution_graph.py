from collections import defaultdict, deque
from typing import Optional, Dict, List, Set, Deque


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
        self._graph: Dict[TreeNode, List[TreeNode]] = defaultdict(list)

    def findClosestLeaf(self, root: Optional[TreeNode], k: int) -> int:
        self.build_graph(root)

        queue: Deque[TreeNode] = deque(
            node for node in self._graph if node and node.val == k
        )
        seen: Set[TreeNode] = set(queue)

        while queue:
            node = queue.popleft()
            if node:
                if len(self._graph[node]) <= 1:
                    return node.val
                for nbi in self._graph[node]:
                    if nbi not in seen:
                        seen.add(nbi)
                        queue.append(nbi)

        return -1

    def build_graph(self, node: Optional[TreeNode], parent: Optional[TreeNode] = None):
        if node is None:
            return

        self._graph[node].append(parent)
        self._graph[parent].append(node)
        self.build_graph(node.left, node)
        self.build_graph(node.right, node)


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
