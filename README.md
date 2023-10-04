# LRU_CACHE 

<h2>A key-value cache with key and value being strings</h2>
<p>
Use the main.go file and play with it to understand the cache package.
</p>


<div>
<h1>Creating a new LRU Cache</h1>

	NewLRU:=cache.NewLRUCache(5)

<p>
NewLRUCache takes an int for initialising the cache by specifying its max capacity. 
<p>
</div>


<div>
<h1>Put Function</h1>
	NewLRU.Put("hello","there")
	NewLRU.Put("hey","there")
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

