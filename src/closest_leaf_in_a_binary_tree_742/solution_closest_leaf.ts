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
  private memo: Map<TreeNode | null, [number, TreeNode | null]> = new Map();
  private path: TreeNode[] = [];

  findClosestLeaf(root: TreeNode | null, k: number): number {
    this.memo.clear();
    this.path = [];

    this.dfs(root, k);

    let minDistance = Number.POSITIVE_INFINITY;
    let closestLeaf: TreeNode | null = null;

    for (let i = 0; i < this.path.length; i++) {
      const node = this.path[i];
      let [distance, leaf] = this.closestLeaf(node);
      distance += this.path.length - 1 - i;
      if (distance < minDistance) {
        minDistance = distance;
        closestLeaf = leaf;
      }
    }

    return closestLeaf ? closestLeaf.val : -1;
  }

  private dfs(node: TreeNode | null, k: number): boolean {
    if (!node) {
      return false;
    }

    if (node.val === k) {
      this.path.push(node);
      return true;
    }

    this.path.push(node);
    if (this.dfs(node.left, k) || this.dfs(node.right, k)) {
      return true;
    }

    this.path.pop();
    return false;
  }

  private closestLeaf(node: TreeNode | null): [number, TreeNode | null] {
    if (this.memo.has(node)) {
      return this.memo.get(node)!;
    }

    let result: [number, TreeNode | null];

    if (!node) {
      result = [Number.POSITIVE_INFINITY, null];
    } else if (!node.left && !node.right) {
      result = [0, node];
    } else {
      const [leftDistance, leftLeaf] = this.closestLeaf(node.left);
      const [rightDistance, rightLeaf] = this.closestLeaf(node.right);
      console.log(
        `node: ${node.val}, leftDistance: ${leftDistance}, leftLeaf: ${leftLeaf?.val}, rightDistance: ${rightDistance}, rightLeaf: ${rightLeaf?.val}`,
      );

      if (leftDistance < rightDistance) {
        result = [leftDistance + 1, leftLeaf];
      } else {
        result = [rightDistance + 1, rightLeaf];
      }
    }

    this.memo.set(node, result);
    return result;
  }
}

export function test() {
  // Test case 1
  // Tree:
  //     1
  //    / \
  //   3   2
  const root1 = new TreeNode(1);
  root1.left = new TreeNode(3);
  root1.right = new TreeNode(2);
  console.log(new Solution().findClosestLeaf(root1, 1)); // Expected output: 3 or 2

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
  const root2 = new TreeNode(1);
  root2.left = new TreeNode(2);
  root2.right = new TreeNode(3);
  root2.left.left = new TreeNode(4);
  root2.left.left.left = new TreeNode(5);
  root2.left.left.left.left = new TreeNode(6);

  console.log(new Solution().findClosestLeaf(root2, 2)); // Expected output: 3
}
