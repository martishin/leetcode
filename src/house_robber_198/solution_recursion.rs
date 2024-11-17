use std::cmp::max;
use std::collections::HashMap;

pub struct SolutionRecursion;

impl SolutionRecursion {
    fn max_rob(nums: &[i32], cache: &mut HashMap<usize, i32>, idx: usize) -> i32 {
        if idx >= nums.len() {
            return 0;
        }

        if let Some(&value) = cache.get(&idx) {
            return value;
        }

        let result = max(
            nums[idx] + Self::max_rob(nums, cache, idx + 2),
            Self::max_rob(nums, cache, idx + 1),
        );

        cache.insert(idx, result);
        result
    }

    pub fn rob(nums: Vec<i32>) -> i32 {
        let mut cache = HashMap::new();

        Self::max_rob(&nums, &mut cache, 0)
    }
}

pub fn test() {
    println!("{}", SolutionRecursion::rob(vec![1, 2, 3, 1]));
    println!("{}", SolutionRecursion::rob(vec![2, 7, 9, 3, 1]));
}
