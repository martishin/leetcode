use std::collections::HashMap;

#[derive(Clone)]
struct ValueAtTime {
    value: String,
    timestamp: i32,
}

struct TimeMap {
    store: HashMap<String, Vec<ValueAtTime>>,
}

impl TimeMap {
    fn new() -> Self {
        TimeMap {
            store: HashMap::new(),
        }
    }

    fn set(&mut self, key: String, value: String, timestamp: i32) {
        let entry = self.store.entry(key).or_default();
        entry.push(ValueAtTime { value, timestamp });
    }

    fn get(&self, key: String, timestamp: i32) -> String {
        let Some(values) = self.store.get(&key) else {
            return "".to_string();
        };

        let mut left = 0;
        let mut right = values.len();

        while left < right {
            let mid = (left + right) / 2;
            if values[mid].timestamp <= timestamp {
                left = mid + 1;
            } else {
                right = mid;
            }
        }

        if right == 0 {
            return "".to_string();
        }

        values[right - 1].value.clone()
    }
}

/**
 * Your TimeMap object will be instantiated and called as such:
 * let obj = TimeMap::new();
 * obj.set(key, value, timestamp);
 * let ret_2: String = obj.get(key, timestamp);
 */
pub fn test() {
    let mut obj = TimeMap::new();
    obj.set("foo".to_string(), "bar".to_string(), 1);
    println!("{}", obj.get("foo".to_string(), 1));
    println!("{}", obj.get("foo".to_string(), 3));
    obj.set("foo".to_string(), "bar2".to_string(), 4);
    println!("{}", obj.get("foo".to_string(), 4));
    println!("{}", obj.get("foo".to_string(), 5));
}
