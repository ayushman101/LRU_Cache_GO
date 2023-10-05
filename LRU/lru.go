package cache


import(
	"fmt"
	"errors"
	"golang.org/x/exp/constraints"
)

type node[K,T constraints.Ordered] struct {
		key K
		//val string
		val T
		next *node[K,T]
		pre *node[K,T]
}


	
type LRUCache[K,T constraints.Ordered] struct {

	Capacity int
	left *node[K,T]
	right *node[K,T]

	hash map[K]*node[K,T]

}




func NewLRUCache[K,T constraints.Ordered](capacity int) LRUCache[K,T] {
	
	ret := LRUCache[K,T]{	
		Capacity: capacity,
		left: &node[K,T]{},
		right:&node[K,T]{},
		hash: make(map[K]*node[K,T]),
	}

	ret.left.next=ret.right
	ret.right.pre=ret.left

	return ret
}


func (lru *LRUCache[K,T]) insert (key K, val T ) {

	lru.hash[key]= &node[K,T]{ 
		key: key, 
		val:val, 
		next: lru.right,
		pre: lru.right.pre,
	} 

	lru.right.pre.next=lru.hash[key]
	lru.right.pre=lru.hash[key]
}


func (lru *LRUCache[K,T]) remove (key K) {

	lru.hash[key].pre.next=lru.hash[key].next
	lru.hash[key].next.pre=lru.hash[key].pre

	delete(lru.hash,key)
}



func (lru *LRUCache[K,T]) Put(key K, val T) {
	
	if _,ok := lru.hash[key] ; ok{
		
		lru.remove(key)
	}

	lru.insert(key,val)
	
	if len(lru.hash)>lru.Capacity {
		lru.remove(lru.left.next.key)
	}
	
}


func (lru *LRUCache[K,T]) Get(key K) (T, error){

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
	var t T
	return t,fmt.Errorf("%w",err)
}



func (lru *LRUCache[K,T]) Display() {

	for key,Node:= range lru.hash {
		
		fmt.Printf("%s: %v\n",key,Node.val)
	}
}


