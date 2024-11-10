use std::cell::RefCell;
use std::collections::HashMap;
use std::hash::{Hash, Hasher};
use std::rc::Rc;

#[derive(Debug)]
pub struct TreeNode {
    pub val: i32,
    pub left: Option<Node>,
    pub right: Option<Node>,
}

impl TreeNode {
    #[inline]
    pub fn new(val: i32) -> Self {
        Self {
            val,
            left: None,
            right: None,
        }
    }
}

type Node = Rc<RefCell<TreeNode>>;

// Wrapper for Node to implement custom Hash and Eq based on pointer identity
#[derive(Clone)]
struct NodePtr(Node);

impl PartialEq for NodePtr {
    fn eq(&self, other: &Self) -> bool {
        Rc::ptr_eq(&self.0, &other.0)
    }
}

impl Eq for NodePtr {}

impl Hash for NodePtr {
    fn hash<H: Hasher>(&self, state: &mut H) {
        Rc::as_ptr(&self.0).hash(state);
    }
}

pub struct Solution {
    memo: HashMap<NodePtr, (i32, Node)>,
    path: Vec<Node>,
}

impl Solution {
    pub fn new() -> Self {
        Self {
            memo: HashMap::new(),
            path: Vec::new(),
        }
    }

    pub fn find_closest_leaf(&mut self, root: Option<Node>, k: i32) -> i32 {
        self.find_path_to_target(root.as_ref(), k);

        let mut min_distance = i32::MAX;
        let mut closest_leaf_val = -1;

        for (i, node) in self.path.clone().iter().enumerate() {
            let (leaf_distance, leaf_node) = self.calculate_closest_leaf(node.clone());
            let adjusted_distance = leaf_distance + (self.path.len() - 1 - i) as i32;

            if adjusted_distance < min_distance {
                min_distance = adjusted_distance;
                closest_leaf_val = leaf_node.borrow().val;
            }
        }

        closest_leaf_val
    }

    fn find_path_to_target(&mut self, node: Option<&Node>, k: i32) -> bool {
        if let Some(n) = node {
            self.path.push(Rc::clone(n));

            if n.borrow().val == k {
                return true;
            }

            if self.find_path_to_target(n.borrow().left.as_ref(), k)
                || self.find_path_to_target(n.borrow().right.as_ref(), k)
            {
                return true;
            }

            self.path.pop();
        }
        false
    }

    fn calculate_closest_leaf(&mut self, node: Node) -> (i32, Node) {
        let node_ptr = NodePtr(node.clone());

        if let Some(result) = self.memo.get(&node_ptr) {
            return result.clone();
        }

        let result = {
            let n = node.borrow();
            if n.left.is_none() && n.right.is_none() {
                (0, node.clone())
            } else {
                let left_result = n
                    .left
                    .as_ref()
                    .map(|left| self.calculate_closest_leaf(left.clone()))
                    .unwrap_or((i32::MAX, node.clone()));
                let right_result = n
                    .right
                    .as_ref()
                    .map(|right| self.calculate_closest_leaf(right.clone()))
                    .unwrap_or((i32::MAX, node.clone()));

                if left_result.0 < right_result.0 {
                    (left_result.0 + 1, left_result.1)
                } else {
                    (right_result.0 + 1, right_result.1)
                }
            }
        };

        self.memo.insert(node_ptr, result.clone());
        result
    }
}

pub fn test() {
    let mut solution = Solution::new();

    // Test case 1
    let root = Rc::new(RefCell::new(TreeNode::new(1)));
    let left = Rc::new(RefCell::new(TreeNode::new(2)));
    let right = Rc::new(RefCell::new(TreeNode::new(3)));
    root.borrow_mut().left = Some(left.clone());
    root.borrow_mut().right = Some(right.clone());

    let result = solution.find_closest_leaf(Some(root.clone()), 1);
    println!("Closest leaf to node with value 1: {}", result); // Expected: 2 or 3

    let mut solution = Solution::new();

    // Test case 3
    let root = Rc::new(RefCell::new(TreeNode::new(1)));
    let node2 = Rc::new(RefCell::new(TreeNode::new(2)));
    let node3 = Rc::new(RefCell::new(TreeNode::new(3)));
    let node4 = Rc::new(RefCell::new(TreeNode::new(4)));
    let node5 = Rc::new(RefCell::new(TreeNode::new(5)));
    let node6 = Rc::new(RefCell::new(TreeNode::new(6)));

    root.borrow_mut().left = Some(node2.clone());
    root.borrow_mut().right = Some(node3.clone());
    node2.borrow_mut().left = Some(node4.clone());
    node4.borrow_mut().left = Some(node5.clone());
    node5.borrow_mut().left = Some(node6.clone());

    let result = solution.find_closest_leaf(Some(root), 2);
    println!("Closest leaf to node with value 2: {}", result); // Expected: 3
}
