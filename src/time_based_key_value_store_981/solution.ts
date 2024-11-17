interface TimestampValue {
  timestamp: number;
  value: string;
}

class TimeMap {
  private keyToValues: Map<string, TimestampValue[]>;

  constructor() {
    this.keyToValues = new Map();
  }

  set(key: string, value: string, timestamp: number): void {
    if (!this.keyToValues.has(key)) {
      this.keyToValues.set(key, []);
    }
    this.keyToValues.get(key)!.push({ timestamp, value });
  }

  get(key: string, timestamp: number): string {
    const values = this.keyToValues.get(key);

    if (!values) {
      return "";
    }

    let left = 0;
    let right = values.length;

    while (left < right) {
      const mid = Math.floor((left + right) / 2);
      if (values[mid].timestamp <= timestamp) {
        left = mid + 1;
      } else {
        right = mid;
      }
    }

    if (right === 0) {
      return "";
    }

    return values[right - 1].value;
  }
}

/**
 * Your TimeMap object will be instantiated and called as such:
 * var obj = new TimeMap()
 * obj.set(key,value,timestamp)
 * var param_2 = obj.get(key,timestamp)
 */
export function test() {
  const obj = new TimeMap();
  obj.set("foo", "bar", 1);
  console.log(obj.get("foo", 1));
  console.log(obj.get("foo", 3));
  obj.set("foo", "bar2", 4);
  console.log(obj.get("foo", 4));
  console.log(obj.get("foo", 5));
}
