package cache


import(
	"fmt"
	"errors"
)

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
	}

	lru.insert(key,val)
	
	if len(lru.hash)>lru.Capacity {
		lru.remove(lru.left.next.key)
	}
	
}


func (lru *LRUCache) Get(key string) (string, error){

	Node,ok:=lru.hash[key]

	if ok {	
		//Instead of removing and inserting, We just reassign the 
		//Node to the head of the double linked list.

		Node.pre.next=Node.next
		Node.next.pre=Node.pre

		Node.next=lru.right
		Node.pre=lru.right.pre

		lru.right.pre.next=Node

		lru.right.pre=Node

		return Node.val,nil
	}

	err:= errors.New("Key Not Found")
	return "",fmt.Errorf("%w",err)
}



func (lru *LRUCache) Display() {

	for key,Node:= range lru.hash {
		
		fmt.Printf("%s: %s\n",key,Node.val)
	}
}


