package cache


import(
	"fmt"
	"errors"
	"golang.org/x/exp/constraints"
)

type node[T constraints.Ordered] struct {
		key string
		//val string
		val T
		next *node[T]
		pre *node[T]
}


	
type LRUCache[T constraints.Ordered] struct {

	Capacity int
	left *node[T]
	right *node[T]

	hash map[string]*node[T]

}




func NewLRUCache[T constraints.Ordered](capacity int) LRUCache[T] {
	
	ret := LRUCache[T]{	
		Capacity: capacity,
		left: &node[T]{},
		right:&node[T]{},
		hash: make(map[string]*node[T]),
	}

	ret.left.next=ret.right
	ret.right.pre=ret.left

	return ret
}


func (lru *LRUCache[T]) insert (key string, val T ) {

	lru.hash[key]= &node[T]{ 
		key: key, 
		val:val, 
		next: lru.right,
		pre: lru.right.pre,
	} 

	lru.right.pre.next=lru.hash[key]
	lru.right.pre=lru.hash[key]
}


func (lru *LRUCache[T]) remove (key string) {

	lru.hash[key].pre.next=lru.hash[key].next
	lru.hash[key].next.pre=lru.hash[key].pre

	delete(lru.hash,key)
}



func (lru *LRUCache[T]) Put(key string, val T) {
	
	if _,ok := lru.hash[key] ; ok{
		
		lru.remove(key)
	}

	lru.insert(key,val)
	
	if len(lru.hash)>lru.Capacity {
		lru.remove(lru.left.next.key)
	}
	
}


func (lru *LRUCache[T]) Get(key string) (T, error){

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



func (lru *LRUCache[T]) Display() {

	for key,Node:= range lru.hash {
		
		fmt.Printf("%s: %v\n",key,Node.val)
	}
}


