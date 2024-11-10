class TreeNode {
  val: number;
  left: TreeNode | null;
  right: TreeNode | null;

  constructor(val: number = 0, left: TreeNode | null = null, right: TreeNode | null = null) {
    this.val = val;
    this.left = left;
    this.right = right;
  }
}

class Solution {
  private graph: Map<TreeNode, TreeNode[]> = new Map();
  private targetNode: TreeNode | null = null;

  findClosestLeaf(root: TreeNode | null, k: number): number {
    this.buildGraph(root, null, k);

    if (!this.targetNode) {
      return -1;
    }

    const queue: (TreeNode | null)[] = [];
    const seen: Set<TreeNode> = new Set();

    queue.push(this.targetNode);
    seen.add(this.targetNode);

    while (queue.length > 0) {
      const node = queue.shift();
      if (node) {
        const neighbors = this.graph.get(node) || [];
        if (neighbors.length <= 1 && node !== this.targetNode && !node.left && !node.right) {
          return node.val;
        }
        for (const neighbor of neighbors) {
          if (!seen.has(neighbor)) {
            seen.add(neighbor);
            queue.push(neighbor);
          }
        }
      }
    }

    return -1;
  }

  private buildGraph(node: TreeNode | null, parent: TreeNode | null, k: number): void {
    if (!node) {
      return;
    }

    if (node.val === k) {
      this.targetNode = node;
    }

    if (!this.graph.has(node)) {
      this.graph.set(node, []);
    }

    if (parent) {
      this.graph.get(node)!.push(parent);
      if (!this.graph.has(parent)) {
        this.graph.set(parent, []);
      }
      this.graph.get(parent)!.push(node);
    }

    this.buildGraph(node.left, node, k);
    this.buildGraph(node.right, node, k);
  }
}

export function test() {
  // Test case 1
  const root1 = new TreeNode(1);
  root1.left = new TreeNode(3);
  root1.right = new TreeNode(2);

  const solution1 = new Solution();
  console.log(solution1.findClosestLeaf(root1, 1)); // Expected output: 3 or 2

  // Test case 2
  const root2 = new TreeNode(1);
  root2.left = new TreeNode(2);
  root2.right = new TreeNode(3);
  root2.left.left = new TreeNode(4);
  root2.left.left.left = new TreeNode(5);
  root2.left.left.left.left = new TreeNode(6);

  const solution2 = new Solution();
  console.log(solution2.findClosestLeaf(root2, 2)); // Expected output: 3
}
