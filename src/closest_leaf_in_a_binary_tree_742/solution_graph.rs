use std::cell::RefCell;
use std::collections::{HashMap, HashSet, VecDeque};
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
    graph: HashMap<NodePtr, Vec<Node>>,
}

impl Solution {
    pub fn new() -> Self {
        Self {
            graph: HashMap::new(),
        }
    }

    pub fn find_closest_leaf(&mut self, root: Option<Node>, k: i32) -> i32 {
        let mut target_node: Option<Node> = None;
        self.build_graph(&root, None, &mut target_node, k);

        let start_node = match target_node {
            Some(node) => node,
            None => return -1,
        };

        let mut queue = VecDeque::new();
        let mut visited = HashSet::new();

        queue.push_back(start_node.clone());
        visited.insert(NodePtr(start_node.clone()));

        while let Some(node) = queue.pop_front() {
            let node_ptr = NodePtr(node.clone());
            let neighbors = self.graph.get(&node_ptr).unwrap();
            let neighbor_count = neighbors.len();

            // Check if the node is a leaf node
            if neighbor_count <= 1 && node.borrow().left.is_none() && node.borrow().right.is_none()
            {
                return node.borrow().val;
            }

            for neighbor in neighbors {
                let neighbor_ptr = NodePtr(neighbor.clone());
                if !visited.contains(&neighbor_ptr) {
                    visited.insert(neighbor_ptr.clone());
                    queue.push_back(neighbor.clone());
                }
            }
        }

        -1
    }

    fn build_graph(
        &mut self,
        node: &Option<Node>,
        parent: Option<Node>,
        target_node: &mut Option<Node>,
        k: i32,
    ) {
        if let Some(current) = node {
            let current_ptr = NodePtr(current.clone());

            if current.borrow().val == k {
                *target_node = Some(current.clone());
            }

            self.graph
                .entry(current_ptr.clone())
                .or_insert_with(Vec::new);

            if let Some(parent_node) = parent {
                self.graph
                    .get_mut(&current_ptr)
                    .unwrap()
                    .push(parent_node.clone());

                let parent_ptr = NodePtr(parent_node.clone());
                self.graph
                    .entry(parent_ptr.clone())
                    .or_insert_with(Vec::new);
                self.graph
                    .get_mut(&parent_ptr)
                    .unwrap()
                    .push(current.clone());
            }

            self.build_graph(
                &current.borrow().left,
                Some(current.clone()),
                target_node,
                k,
            );
            self.build_graph(
                &current.borrow().right,
                Some(current.clone()),
                target_node,
                k,
            );
        }
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

    let result = solution.find_closest_leaf(Some(root.clone()), 2);
    println!("Closest leaf to node with value 2: {}", result); // Expected: 3
}
