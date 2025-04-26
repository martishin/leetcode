package insert_delete_get_random_380

import (
	"fmt"
	"math/rand"
)

type RandomizedSet struct {
	values  []int
	mapping map[int]int
}

func Constructor() RandomizedSet {
	return RandomizedSet{
		values:  []int{},
		mapping: make(map[int]int),
	}
}

func (s *RandomizedSet) Insert(val int) bool {
	if _, ok := s.mapping[val]; ok {
		return false
	}

	s.values = append(s.values, val)
	s.mapping[val] = len(s.values) - 1
	return true
}

func (s *RandomizedSet) Remove(val int) bool {
	idx, ok := s.mapping[val]
	if !ok {
		return false
	}

	n := len(s.values)
	s.values[idx], s.values[n-1] = s.values[n-1], s.values[idx]
	s.mapping[s.values[idx]] = idx
	s.values = s.values[:n-1]
	delete(s.mapping, val)
	return true
}

func (s *RandomizedSet) GetRandom() int {
	randIdx := rand.Intn(len(s.values))
	return s.values[randIdx]
}

func Test() {
	obj := Constructor()
	fmt.Println(obj.Insert(1))
	fmt.Println(obj.Remove(2))
	fmt.Println(obj.Insert(2))
	fmt.Println(obj.GetRandom())
	fmt.Println(obj.Remove(1))
	fmt.Println(obj.Insert(2))
	fmt.Println(obj.GetRandom())
}
