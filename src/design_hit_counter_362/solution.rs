use std::collections::VecDeque;

struct HitCounter {
    hits: VecDeque<i32>,
}

/**
 * `&self` means the method takes an immutable reference.
 * If you need a mutable reference, change it to `&mut self` instead.
 */
impl HitCounter {
    fn new() -> Self {
        HitCounter {
            hits: VecDeque::new(),
        }
    }

    fn hit(&mut self, timestamp: i32) {
        self.hits.push_back(timestamp);
    }

    fn get_hits(&mut self, timestamp: i32) -> i32 {
        self.clean_old_hits(timestamp);
        self.hits.len() as i32
    }

    fn clean_old_hits(&mut self, current_time: i32) {
        let threshold = current_time - 5 * 60;
        while let Some(&oldest) = self.hits.front() {
            if oldest > threshold {
                break;
            }
            self.hits.pop_front();
        }
    }
}

/**
 * Your HitCounter object will be instantiated and called as such:
 * let obj = HitCounter::new();
 * obj.hit(timestamp);
 * let ret_2: i32 = obj.get_hits(timestamp);
 */
pub fn test() {
    let mut obj = HitCounter::new();
    obj.hit(1);
    obj.hit(2);
    obj.hit(3);
    println!("{}", obj.get_hits(4));
    obj.hit(300);
    println!("{}", obj.get_hits(300));
    println!("{}", obj.get_hits(301));
}
