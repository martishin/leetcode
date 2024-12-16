use std::collections::VecDeque;

pub struct Solution;

impl Solution {
    pub fn longest_subarray(nums: Vec<i32>, limit: i32) -> i32 {
        let mut min_deque: VecDeque<i32> = VecDeque::new();
        let mut max_deque: VecDeque<i32> = VecDeque::new();

        let mut left = 0;
        let mut longest_subarray = 0;

        for (right, &num) in nums.iter().enumerate() {
            while let Some(&back) = min_deque.back() {
                if back > num {
                    min_deque.pop_back();
                } else {
                    break;
                }
            }
            min_deque.push_back(num);

            while let Some(&back) = max_deque.back() {
                if back < num {
                    max_deque.pop_back();
                } else {
                    break;
                }
            }
            max_deque.push_back(num);

            while max_deque.front().unwrap() - min_deque.front().unwrap() > limit {
                if *max_deque.front().unwrap() == nums[left] {
                    max_deque.pop_front();
                }
                if *min_deque.front().unwrap() == nums[left] {
                    min_deque.pop_front();
                }
                left += 1;
            }

            longest_subarray = longest_subarray.max(right - left + 1);
        }

        longest_subarray as i32
    }
}

pub fn test() {
    let nums1 = vec![8, 2, 4, 7];
    let limit1 = 4;
    println!("{}", Solution::longest_subarray(nums1, limit1));

    let nums2 = vec![10, 1, 2, 4, 7, 2];
    let limit2 = 5;
    println!("{}", Solution::longest_subarray(nums2, limit2));

    let nums3 = vec![4, 2, 2, 2, 4, 4, 2, 2];
    let limit3 = 0;
    println!("{}", Solution::longest_subarray(nums3, limit3));
}
