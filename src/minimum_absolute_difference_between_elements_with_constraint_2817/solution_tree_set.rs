use std::collections::BTreeSet;
use std::collections::Bound::{Included, Unbounded};

pub struct Solution;

impl Solution {
    fn min_absolute_difference(nums: Vec<i32>, x: i32) -> i32 {
        let mut sorted_list = BTreeSet::new();
        let mut result = i32::MAX;
        let x = x as usize;

        for i in x..nums.len() {
            sorted_list.insert(nums[i - x]);

            // Find the largest element smaller than nums[i]
            // if let Some(&lower_bound) = sorted_list.range(..nums[i]).next_back() {
            //     result = result.min((nums[i] - lower_bound).abs());
            // }
            if let Some(&lower_bound) = sorted_list
                .range((Unbounded, Included(nums[i])))
                .next_back()
            {
                result = result.min((nums[i] - lower_bound).abs());
            }

            // Find the smallest element greater than nums[i]
            // if let Some(&upper_bound) = sorted_list.range(nums[i]..).next() {
            //     result = result.min((nums[i] - upper_bound).abs());
            // }
            if let Some(&upper_bound) = sorted_list.range((Included(nums[i]), Unbounded)).next() {
                result = result.min((nums[i] - upper_bound).abs());
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
}
