package lru_cache_146

import "fmt"

type LNode struct {
	prev, next *LNode
	val, key   int
}

type LRUCache struct {
	capacity    int
	front, back *LNode
	mapping     map[int]*LNode
}

func Constructor(capacity int) LRUCache {
	front := &LNode{}
	back := &LNode{}

	front.next = back
	back.prev = front

	return LRUCache{
		capacity: capacity,
		front:    front,
		back:     back,
		mapping:  make(map[int]*LNode),
	}
}

func (c *LRUCache) pushFront(node *LNode) {
	c.mapping[node.key] = node

	c.front.next.prev = node
	node.next = c.front.next
	node.prev = c.front
	c.front.next = node
}

func (c *LRUCache) moveFront(node *LNode) {
	node.next.prev = node.prev
	node.prev.next = node.next

	c.pushFront(node)
}

func (c *LRUCache) removeBack() {
	delete(c.mapping, c.back.prev.key)
	c.back.prev.prev.next = c.back
	c.back.prev = c.back.prev.prev
}

func (c *LRUCache) Get(key int) int {
	node, ok := c.mapping[key]

	if ok {
		c.moveFront(node)
		return node.val
	}

	return -1
}

func (c *LRUCache) Put(key int, value int) {
	node, ok := c.mapping[key]

	if ok {
		node.val = value
		c.moveFront(node)
		return
	}

	newNode := &LNode{key: key, val: value}
	c.pushFront(newNode)
	if len(c.mapping) > c.capacity {
		c.removeBack()
	}
}

func Test() {
	obj := Constructor(2)
	obj.Put(1, 1)
	obj.Put(2, 2)
	fmt.Println(obj.Get(1))
	obj.Put(3, 3)
	fmt.Println(obj.Get(2))
	obj.Put(4, 4)
	fmt.Println(obj.Get(1))
	fmt.Println(obj.Get(3))
	fmt.Println(obj.Get(4))
}
