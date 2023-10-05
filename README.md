# LRU_CACHE 

<h2>A key-value cache with the key and value can be any of the following types: </h2>

<p>
	~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr | ~int | ~int8 | ~int16 | ~int32 | ~int64
	~float32 | ~float64 | ~string
<p>
<p>
Use the main.go file and play with it to understand the cache package.

</p>


<div>
<h1>Creating a new LRU Cache</h1>

	NewLRU:=cache.NewLRUCache[string,int](5)  

	//where the key is of type string and value is of type int. 

<p>
NewLRUCache takes an int for initialising the cache by specifying its max capacity. 
<p>
</div>


<div>
<h1>Put Function</h1>
	NewLRU.Put("hello",20)
	NewLRU.Put("hey",30)
</div>

<div>
<h1> Get Function</h1>

	val,err:=NewLRU.Get(key)

<p>
returns an error if there is no such key. 
</p>
</div>


<div>

<h1> Display</h1>

	NewLRU.Display()

<p> 
Use the display function to view the entire LRU cache.
</p>
</div>

