package lc146

type Node struct {
	key   int
	value int
	pre   *Node
	next  *Node
}

type LRUCache struct {
	Capacity   int
	Head, Tail *Node
	NodeMap    map[int]*Node
}

func Constructor(capacity int) LRUCache {
	head := &Node{}
	tail := &Node{}
	head.next = tail
	tail.pre = head
	return LRUCache{
		Capacity: capacity,
		Head:     head,
		Tail:     tail,
		NodeMap:  map[int]*Node{},
	}
}

func (c *LRUCache) addToHead(node *Node) {
	c.Head.next.pre = node
	node.next = c.Head.next
	node.pre = c.Head
	c.Head.next = node
}

func (c *LRUCache) moveToHead(node *Node) {
	c.deleteNode(node)
	c.addToHead(node)
}

func (c *LRUCache) removeFromTail() int {
	node := c.Tail.pre
	node.pre.next = c.Tail
	c.Tail.pre = node.pre
	c.deleteNode(node)
	return node.key
}

func (c *LRUCache) deleteNode(node *Node) {
	node.next.pre = node.pre
	node.pre.next = node.next
}

func (c *LRUCache) Get(key int) int {
	if node, ok := c.NodeMap[key]; ok {
		c.moveToHead(node)
		return node.value
	}
	return -1
}

func (c *LRUCache) Put(key int, value int) {
	if node, ok := c.NodeMap[key]; ok {
		node.value = value
		c.moveToHead(node)
		return
	}

	if len(c.NodeMap) == c.Capacity {
		key := c.removeFromTail()
		delete(c.NodeMap, key)
	}

	node := &Node{
		key:   key,
		value: value,
	}

	c.addToHead(node)
	c.NodeMap[key] = node
}
