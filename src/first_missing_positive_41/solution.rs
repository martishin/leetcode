pub struct Solution;

impl Solution {
    pub fn first_missing_positive(mut nums: Vec<i32>) -> i32 {
        let n = nums.len();

        for i in 0..n {
            while nums[i] > 0 && nums[i] as usize <= n && nums[i] != nums[nums[i] as usize - 1] {
                let target_idx = nums[i] as usize - 1;
                nums.swap(i, target_idx);
            }
        }

        for i in 0..n {
            if nums[i] != (i as i32 + 1) {
                return i as i32 + 1;
            }
        }

        n as i32 + 1
    }
}

pub fn test() {
    println!("{}", Solution::first_missing_positive(vec![1, 2, 0]));
    println!("{}", Solution::first_missing_positive(vec![3, 4, -1, 1]));
    println!(
        "{}",
        Solution::first_missing_positive(vec![7, 8, 9, 11, 12])
    );
}
