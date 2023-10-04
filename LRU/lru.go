package cache


type node struct {
		key string
		val string
		next *node
		pre *node
}


	
type LRUCache struct {

	Capacity int
	left *node
	right *node

	hash map[string]*node

}




func NewLRUCache(capacity int) LRUCache {
	
	ret := LRUCache{	
		Capacity: capacity,
		left: &node{},
		right:&node{},
		hash: make(map[string]*node),
	}

	ret.left.next=ret.right
	ret.right.pre=ret.left

	return ret
}


func (lru *LRUCache) insert (key string, val string ) {

	lru.hash[key]= &node{ 
		key: key, 
		val:val, 
		next: lru.right,
		pre: lru.right.pre,
	} 

	lru.right.pre.next=lru.hash[key]
	lru.right.pre=lru.hash[key]
}


func (lru *LRUCache) remove (key string) {

	lru.hash[key].pre.next=lru.hash[key].next
	lru.hash[key].next.pre=lru.hash[key].pre

	delete(lru.hash,key)
}



func (lru *LRUCache) Put(key string, val string) {
	
	if _,ok := lru.hash[key] ; ok{
		
		lru.remove(key)
		lru.insert(key,val)
	} else{
		lru.insert(key,val)
        
	}	
	
	if len(lru.hash)>lru.Capacity {
		lru.remove(lru.left.next.key)
	}
	
}




