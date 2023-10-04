package cache


type node struct {
		key string
		val string
		next *node
		pre *node
}

	
type LRUCache struct {

	capacity int
	left *node
	right *node

	hash map[string]*node

}

func NewLRUCache(capacity int) LRUCache {
	
	ret := LRUCache{	
		capacity: capacity,
		left: &node{},
		right:&node{},
		hash: map[string]*node{},
	}

	ret.left.next=right
	ret.right.pre=left

	return ret
}




