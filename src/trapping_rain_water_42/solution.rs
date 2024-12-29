pub struct Solution;

impl Solution {
    pub fn trap(height: Vec<i32>) -> i32 {
        let n = height.len();
        let mut all_trapped_water = 0;

        let mut current_trapped_water = 0;
        let mut current_height = 0;
        for i in 0..n {
            if height[i] >= current_height {
                all_trapped_water += current_trapped_water;
                current_height = height[i];
                current_trapped_water = 0;
            } else {
                current_trapped_water += current_height - height[i];
            }
        }

        current_trapped_water = 0;
        current_height = 0;
        for i in (0..n).rev() {
            if height[i] > current_height {
                all_trapped_water += current_trapped_water;
                current_height = height[i];
                current_trapped_water = 0;
            } else {
                current_trapped_water += current_height - height[i];
            }
        }

        return all_trapped_water;
    }
}

pub fn test() {
    println!(
        "{}",
        Solution::trap(vec![0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1])
    );
    println!("{}", Solution::trap(vec![4, 2, 0, 3, 2, 5]));
}
