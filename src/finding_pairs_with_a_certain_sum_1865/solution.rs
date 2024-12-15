use std::collections::HashMap;

struct FindSumPairs {
    nums1: Vec<i32>,
    nums2: Vec<i32>,
    nums2_to_count: HashMap<i32, i32>,
}

impl FindSumPairs {
    fn new(mut nums1: Vec<i32>, nums2: Vec<i32>) -> Self {
        nums1.sort();

        let mut nums2_to_count = HashMap::new();
        for &num in &nums2 {
            *nums2_to_count.entry(num).or_insert(0) += 1;
        }

        FindSumPairs {
            nums1,
            nums2,
            nums2_to_count,
        }
    }

    fn add(&mut self, index: i32, val: i32) {
        let index = index as usize;
        self.nums2_to_count
            .entry(self.nums2[index])
            .and_modify(|e| *e -= 1);
        self.nums2[index] += val;
        *self.nums2_to_count.entry(self.nums2[index]).or_insert(0) += 1;
    }

    fn count(&self, tot: i32) -> i32 {
        let mut result = 0;

        for &num1 in &self.nums1 {
            if num1 >= tot {
                break;
            }

            if let Some(&count) = self.nums2_to_count.get(&(tot - num1)) {
                result += count;
            }
        }

        result
    }
}

pub fn test() {
    let mut obj = FindSumPairs::new(vec![1, 1, 2, 2, 2, 3], vec![1, 4, 5, 2, 5, 4]);
    println!("{}", obj.count(7));
    obj.add(3, 2);
    println!("{}", obj.count(8));
    println!("{}", obj.count(4));
    obj.add(0, 1);
    obj.add(1, 1);
    println!("{}", obj.count(7));
}
