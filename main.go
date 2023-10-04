package main

import (
	
	"fmt"
	"github.com/ayushman101/LRU_Cache_Go/cache"

)

func main(){

	Newcache:=cache.NewLRUCache(10)

	fmt.Println(Newcache.capacity)
}
