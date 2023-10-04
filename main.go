package main

import (
	
	"fmt"
	"github.com/ayushman101/LRU_Cache_Go/LRU"	
)

func main(){

	Newcache:=cache.NewLRUCache(10)

	fmt.Println(Newcache.Capacity)

	Newcache.Put("hello ", "there")
	Newcache.Put("hey","there")

}
