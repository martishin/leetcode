package design_hit_counter_362

import "fmt"

type HitCounter struct {
	hits []int
}

func Constructor() HitCounter {
	return HitCounter{
		hits: []int{},
	}
}

func (this *HitCounter) Hit(timestamp int) {
	this.hits = append(this.hits, timestamp)
}

func (this *HitCounter) GetHits(timestamp int) int {
	this.cleanOldHits(timestamp)
	return len(this.hits)
}

func (this *HitCounter) cleanOldHits(currentTime int) {
	threshold := currentTime - 300
	i := 0
	for i < len(this.hits) && this.hits[i] <= threshold {
		i++
	}

	this.hits = this.hits[i:]
}

func Solution() {
	counter := Constructor()
	counter.Hit(1)
	counter.Hit(2)
	counter.Hit(3)
	fmt.Println(counter.GetHits(4)) // Should return 3
	counter.Hit(300)
	fmt.Println(counter.GetHits(300)) // Should return 4
	fmt.Println(counter.GetHits(301)) // Should return 3
}
