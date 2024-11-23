// Definition for singly-linked list.
#[derive(PartialEq, Eq, Clone, Debug)]
pub struct ListNode {
    pub val: i32,
    pub next: Option<Box<ListNode>>,
}

impl ListNode {
    #[inline]
    fn new(val: i32) -> Self {
        ListNode { next: None, val }
    }
}

pub struct Solution;

impl Solution {
    pub fn reverse_list(head: Option<Box<ListNode>>) -> Option<Box<ListNode>> {
        let (mut prev, mut current) = (None, head);

        while let Some(mut node) = current {
            current = node.next;
            node.next = prev;
            prev = Some(node);
        }
        prev
    }
}

pub fn test() {
    fn build_list(values: Vec<i32>) -> Option<Box<ListNode>> {
        let mut head = None;
        for &val in values.iter().rev() {
            let mut new_node = Box::new(ListNode::new(val));
            new_node.next = head;
            head = Some(new_node);
        }
        head
    }

    fn print_list(head: Option<Box<ListNode>>) {
        let mut current = head;
        while let Some(node) = current {
            print!("{} -> ", node.val);
            current = node.next;
        }
        println!("Node");
    }

    let head = build_list(vec![1, 2, 3, 4, 5]);

    let reversed = Solution::reverse_list(head);
    print_list(reversed);
}