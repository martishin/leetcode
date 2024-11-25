package snapshot_array_1146

import "fmt"

type Snapshot struct {
	version int
	val     int
}

type SnapshotArray struct {
	entries [][]*Snapshot
	version int
}

func Constructor(length int) SnapshotArray {
	return SnapshotArray{
		entries: make([][]*Snapshot, length),
		version: 0,
	}
}

func (this *SnapshotArray) Set(index int, val int) {
	if len(this.entries[index]) > 0 {
		lastEntry := this.entries[index][len(this.entries[index])-1]

		if lastEntry.version == this.version {
			lastEntry.val = val
			return
		}
	}

	this.entries[index] = append(this.entries[index], &Snapshot{version: this.version, val: val})
}

func (this *SnapshotArray) Snap() int {
	this.version++
	return this.version - 1
}

func (this *SnapshotArray) Get(index int, snap_id int) int {
	if len(this.entries[index]) == 0 {
		return 0
	}

	idxEntries := this.entries[index]

	l := 0
	r := len(idxEntries)

	for l < r {
		m := (l + r) / 2
		if idxEntries[m].version == snap_id {
			return idxEntries[m].val
		} else if idxEntries[m].version < snap_id {
			l = m + 1
		} else {
			r = m
		}
	}

	if r == 0 {
		return 0
	}

	return idxEntries[r-1].val
}

// Test
/**
 * Your SnapshotArray object will be instantiated and called as such:
 * obj := Constructor(length);
 * obj.Set(index,val);
 * param_2 := obj.Snap();
 * param_3 := obj.Get(index,snap_id);
 */
func Test() {
	obj := Constructor(3)
	obj.Set(0, 5)
	fmt.Print(obj.Snap(), ",")
	obj.Set(0, 6)
	fmt.Println(obj.Get(0, 0))

	obj = Constructor(2)
	fmt.Print(obj.Snap(), ",")
	fmt.Print(obj.Get(1, 0), ",")
	fmt.Print(obj.Get(0, 0), ",")
	obj.Set(1, 8)
	fmt.Print(obj.Get(1, 0), ",")
	obj.Set(0, 20)
	fmt.Println(obj.Get(0, 0))
	obj.Set(0, 7)
}
