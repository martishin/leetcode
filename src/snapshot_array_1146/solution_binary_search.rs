#[derive(Debug, Clone)]
struct Snapshot {
    version: i32,
    value: i32,
}

struct SnapshotArray {
    entries: Vec<Vec<Snapshot>>,
    version: i32,
}

impl SnapshotArray {
    fn new(length: i32) -> Self {
        SnapshotArray {
            entries: vec![
                vec![Snapshot {
                    version: 0,
                    value: 0
                }];
                length as usize
            ],
            version: 0,
        }
    }

    fn set(&mut self, index: i32, value: i32) {
        let index = index as usize;
        let last_snapshot = self.entries[index].last_mut().unwrap();
        if last_snapshot.version == self.version {
            last_snapshot.value = value
        } else {
            self.entries[index].push(Snapshot {
                version: self.version,
                value,
            });
        }
    }

    fn snap(&mut self) -> i32 {
        self.version += 1;
        self.version - 1
    }

    fn get(&self, index: i32, snap_id: i32) -> i32 {
        let index = index as usize;
        let snapshots = &self.entries[index];
        match snapshots.binary_search_by(|snapshot| snapshot.version.cmp(&snap_id)) {
            Ok(pos) => snapshots[pos].value,
            Err(pos) => {
                if pos == 0 {
                    0
                } else {
                    snapshots[pos - 1].value
                }
            }
        }
    }
}

/**
 * Your SnapshotArray object will be instantiated and called as such:
 * let obj = SnapshotArray::new(length);
 * obj.set(index, val);
 * let ret_2: i32 = obj.snap();
 * let ret_3: i32 = obj.get(index, snap_id);
 */
pub fn test() {
    let mut obj = SnapshotArray::new(3);
    obj.set(0, 5);
    print!("{},", obj.snap());
    obj.set(0, 6);
    println!("{}", obj.get(0, 0));

    let mut obj = SnapshotArray::new(2);
    print!("{},", obj.snap());
    print!("{},", obj.get(1, 0));
    print!("{},", obj.get(0, 0));
    obj.set(1, 8);
    print!("{},", obj.get(1, 0));
    obj.set(0, 20);
    println!("{}", obj.get(0, 0));
    obj.set(0, 7);
}
