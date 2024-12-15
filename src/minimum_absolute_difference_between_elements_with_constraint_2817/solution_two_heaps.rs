use std::cmp::Reverse;
use std::collections::BinaryHeap;

pub struct Solution;

impl Solution {
    fn min_absolute_difference(nums: Vec<i32>, x: i32) -> i32 {
        let mut sorted_list: Vec<(i32, usize)> = nums
            .into_iter()
            .enumerate()
            .map(|(idx, val)| (val, idx))
            .collect();
        sorted_list.sort();

        let x = x as usize;
        let mut min_heap: BinaryHeap<Reverse<(usize, i32)>> = BinaryHeap::new();
        let mut max_heap: BinaryHeap<(usize, i32)> = BinaryHeap::new();

        let mut result = i32::MAX;

        for &(val, idx) in &sorted_list {
            min_heap.push(Reverse((idx, val)));
            max_heap.push((idx, val));

            while let Some(&Reverse((heap_idx, heap_val))) = min_heap.peek() {
                if idx >= x && heap_idx <= idx - x {
                    result = result.min((val - heap_val).abs());
                    min_heap.pop();
                } else {
                    break;
                }
            }

            while let Some(&(heap_idx, heap_val)) = max_heap.peek() {
                if heap_idx >= idx + x {
                    result = result.min((val - heap_val).abs());
                    max_heap.pop();
                } else {
                    break;
                }
            }
        }

        result
    }
}

pub fn test() {
    println!("{}", Solution::min_absolute_difference(vec![4, 3, 2, 4], 2));
    println!(
        "{}",
        Solution::min_absolute_difference(vec![5, 3, 2, 10, 15], 1)
    );
    println!("{}", Solution::min_absolute_difference(vec![1, 2, 3, 4], 3));
    println!(
        "{}",
        Solution::min_absolute_difference(vec![14, 111, 16], 1)
    );
    println!("{}", Solution::min_absolute_difference(vec![71, 4], 1));
}
