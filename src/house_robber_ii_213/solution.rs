pub struct Solution;

impl Solution {
    pub fn rob(nums: Vec<i32>) -> i32 {
        let len = nums.len();
        if len == 1 {
            return nums[0];
        }

        if len == 2 {
            return nums[0].max(nums[1]);
        }

        Self::rob_helper(&nums[1..]).max(Self::rob_helper(&nums[..len - 1]))
    }

    fn rob_helper(nums: &[i32]) -> i32 {
        let mut second_last = nums[0];
        let mut last = nums[0].max(nums[1]);

        for &num in nums.iter().skip(2) {
            let current = last.max(second_last + num);
            second_last = last;
            last = current;
        }

        last
    }
}

pub fn test() {
    println!("{}", Solution::rob(vec![1, 2, 3, 1]));
    println!("{}", Solution::rob(vec![2, 7, 9, 3, 1]));
}
