package time_based_key_value_store_981

import "fmt"

type ValueAtTime struct {
	value     string
	timestamp int
}

type TimeMap struct {
	store map[string][]ValueAtTime
}

func Constructor() TimeMap {
	return TimeMap{
		store: make(map[string][]ValueAtTime),
	}
}

func (this *TimeMap) Set(key string, value string, timestamp int) {
	newValueAtTime := ValueAtTime{value: value, timestamp: timestamp}
	this.store[key] = append(this.store[key], newValueAtTime)
}

func (this *TimeMap) Get(key string, timestamp int) string {
	values, found := this.store[key]
	if !found {
		return ""
	}

	l := 0
	r := len(values)

	for l < r {
		mid := (l + r) / 2

		if values[mid].timestamp <= timestamp {
			l = mid + 1
		} else {
			r = mid
		}
	}

	if r == 0 {
		return ""
	}

	return values[r-1].value
}

// Test
/**
 * Your TimeMap object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Set(key,value,timestamp);
 * param_2 := obj.Get(key,timestamp);
 */
func Test() {
	obj := Constructor()
	obj.Set("foo", "bar", 1)
	fmt.Println(obj.Get("foo", 1))
	fmt.Println(obj.Get("foo", 3))
	obj.Set("foo", "bar2", 4)
	fmt.Println(obj.Get("foo", 4))
	fmt.Println(obj.Get("foo", 5))
}
