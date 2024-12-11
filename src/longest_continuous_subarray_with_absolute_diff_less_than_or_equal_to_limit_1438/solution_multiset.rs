use std::collections::BTreeMap;

pub struct Solution;

impl Solution {
    pub fn longest_subarray(nums: Vec<i32>, limit: i32) -> i32 {
        let mut result = 1;
        let mut set_values = BTreeMap::new();
        let mut set_len = 0;
        let mut l = 0;

        for &num in &nums {
            set_len += 1;
            *set_values.entry(num).or_insert(0) += 1;

            while set_values.keys().last().unwrap() - set_values.keys().next().unwrap() > limit {
                if let Some(count) = set_values.get_mut(&nums[l]) {
                    *count -= 1;
                    set_len -= 1;
                    if *count == 0 {
                        set_values.remove(&nums[l]);
                    }
                }
                l += 1;
            }
            result = result.max(set_len);
        }

        result
    }
}

pub fn test() {
    println!("{:?}", Solution::longest_subarray(vec![8, 2, 4, 7], 4));
    println!(
        "{:?}",
        Solution::longest_subarray(vec![10, 1, 2, 4, 7, 2], 5)
    );
    println!(
        "{:?}",
        Solution::longest_subarray(vec![4, 2, 2, 2, 4, 4, 2, 2], 0)
    );
}
