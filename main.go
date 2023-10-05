package main

import (
	
	"fmt"
	"github.com/ayushman101/LRU_Cache_Go/LRU"	
)

func main(){

	Newcache:=cache.NewLRUCache[string,float64](10)

	fmt.Println(Newcache.Capacity)

	Newcache.Put("hello",10.01)
	Newcache.Put("hey",20.222)
	_,err:= Newcache.Get("hey")

	if err!=nil{
		panic(err)
	}
	
	fmt.Println(Newcache)
	Newcache.Display()
}
