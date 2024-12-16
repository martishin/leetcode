use std::cmp::Reverse;
use std::collections::BinaryHeap;

pub struct Solution;

impl Solution {
    pub fn longest_subarray(nums: Vec<i32>, limit: i32) -> i32 {
        let mut min_heap: BinaryHeap<Reverse<(i32, usize)>> = BinaryHeap::new();
        let mut max_heap: BinaryHeap<(i32, usize)> = BinaryHeap::new(); // BinaryHeap is max-heap by default in Rust

        let mut left = 0;
        let mut longest_subarray = 0;

        for (right, &num) in nums.iter().enumerate() {
            min_heap.push(Reverse((num, right)));
            max_heap.push((num, right));

            while max_heap.peek().unwrap().0 - min_heap.peek().unwrap().0 .0 > limit {
                left = std::cmp::min(max_heap.peek().unwrap().1, min_heap.peek().unwrap().0 .1) + 1;

                while max_heap.peek().unwrap().1 < left {
                    max_heap.pop();
                }
                while min_heap.peek().unwrap().0 .1 < left {
                    min_heap.pop();
                }
            }

            longest_subarray = longest_subarray.max(right - left + 1);
        }

        longest_subarray as i32
    }
}

pub fn test() {
    println!("{}", Solution::longest_subarray(vec![8, 2, 4, 7], 4));
    println!("{}", Solution::longest_subarray(vec![10, 1, 2, 4, 7, 2], 5));
    println!(
        "{}",
        Solution::longest_subarray(vec![4, 2, 2, 2, 4, 4, 2, 2], 0)
    );
}
