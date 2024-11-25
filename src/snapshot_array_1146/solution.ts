interface Snapshot {
  version: number;
  val: number;
}

class SnapshotArray {
  private entries: Array<Array<Snapshot>>;
  private version: number;

  constructor(length: number) {
    this.entries = Array.from({ length }, () => [{ version: 0, val: 0 }]);
    this.version = 0;
  }

  set(index: number, val: number): void {
    const lastEntry = this.entries[index][this.entries[index].length - 1];
    if (lastEntry.version === this.version) {
      lastEntry.val = val;
    } else {
      this.entries[index].push({ version: this.version, val });
    }
  }

  snap(): number {
    this.version += 1;
    return this.version - 1;
  }

  get(index: number, snap_id: number): number {
    const snapshots = this.entries[index];

    let left = 0,
      right = snapshots.length;

    while (left < right) {
      const mid = Math.floor((left + right) / 2);
      if (snapshots[mid].version <= snap_id) {
        left = mid + 1;
      } else {
        right = mid;
      }
    }

    return snapshots[left - 1].val;

    // while (left < right) {
    //   const mid = Math.floor((left + right) / 2);
    //   if (snapshots[mid].version == snap_id) {
    //     return snapshots[mid].val;
    //   } else if (snapshots[mid].version < snap_id) {
    //     left = mid + 1;
    //   } else {
    //     right = mid;
    //   }
    // }
    //
    // if (right === 0) {
    //   return 0;
    // }
    //
    // return snapshots[right - 1].val;
  }
}

export function test(): void {
  const obj = new SnapshotArray(3);
  obj.set(0, 5);
  process.stdout.write(obj.snap() + ",");
  obj.set(0, 6);
  console.log(obj.get(0, 0));

  const obj2 = new SnapshotArray(2);
  process.stdout.write(obj2.snap() + ",");
  process.stdout.write(obj2.get(1, 0) + ",");
  process.stdout.write(obj2.get(0, 0) + ",");
  obj2.set(1, 8);
  process.stdout.write(obj2.get(1, 0) + ",");
  obj2.set(0, 20);
  console.log(obj2.get(0, 0));
  obj2.set(0, 7);
}
