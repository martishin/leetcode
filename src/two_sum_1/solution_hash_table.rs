use std::collections::HashMap;

pub struct Solution;

impl Solution {
    pub fn two_sum(nums: Vec<i32>, target: i32) -> Vec<i32> {
        let mut lookup_table: HashMap<i32, i32> = HashMap::new();

        for (idx, &num) in nums.iter().enumerate() {
            let idx = idx as i32;
            let find = target - num;

            if let Some(&idx_find) = lookup_table.get(&find) {
                return vec![idx_find, idx];
            }

            lookup_table.insert(num, idx);
        }

        vec![-1, -1]
    }
}

pub fn test() {
    println!("{:?}", Solution::two_sum(vec![3, 2, 4], 6))
}
