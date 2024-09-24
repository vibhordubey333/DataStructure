package main

import (
	"container/list"
	"fmt"
)

type LRUCache struct {
	capacity     int
	cacheHashMap map[string]*list.Element
	list         *list.List
}

type KeyValue struct {
	key   int
	value interface{}
}

func Constructor(capacity int) LRUCache {

	return LRUCache{
		capacity:     capacity,
		cacheHashMap: make(map[string]*list.Element, 0),
		list:         list.New(),
	}
}

/*
Put adds the given cache item into LRU Cache.
If the cache capacity is full then LRU[Least Recently Used] item will be evicted before adding the new item
*/
func (this *LRUCache) Put(key int, cacheItem interface{}) {
	//Checking if item is in the cache, if yes then move it to front.
	if node, ok := this.cacheHashMap[string(key)]; ok {
		this.list.MoveToFront(node) //O(1) operation
		return
	}

	/*
		If items is not in cache then add it to the cache [Hashmap & DLL]
		check capacity if there is no capacity to store then delete the least recent used item.
		LRU item will be at the back of the DLL.
	*/
	if this.capacity == len(this.cacheHashMap) {
		//Deleting item from HashMap
		delete(this.cacheHashMap, string(this.list.Back().Value.(KeyValue).key))
		//Deleting item from DLL
		this.list.Remove(this.list.Back())
	}
	//Add the new item to the front of the DLL and to Hashmap
	node := this.list.PushFront(KeyValue{
		key:   key,
		value: cacheItem,
	})
	this.cacheHashMap[string(key)] = node
}

// Get method returns the cached element corresponding to the given key
// If the requested item is not present Get will add it to the cache.
func (this *LRUCache) Get(key int) interface{} {
	//Check if element is already present in cache
	if node, ok := this.cacheHashMap[string(key)]; ok {
		//Move requested item to front as it is recently used.
		this.list.MoveToFront(node)
		//Return cache value item.
		return node.Value.(KeyValue).value
	}
	//If the requested item is not already in cache, add it to the cache.
	value := getValueFromDatabase(key)
	this.Put(key, value)
	return value
}

func getValueFromDatabase(key int) interface{} {
	value := key
	return value
}

func main() {
	conObject := Constructor(2)
	fmt.Println(conObject.Get(1))
	fmt.Println(conObject.Get(2))

}
