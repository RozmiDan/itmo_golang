package lrucache

import (
	"iter"
	"project_1/pkg/linklist"
)

type LRU[K comparable, V any] interface{
	Put(key K, val V)
	Get(key K) V
	Size() int
	All() iter.Seq2[K,V]
}

type Node[K comparable, V any] struct{
	key K
	value V
}

type LruCache[K comparable, V any] struct{
	list *linklist.List[Node[K,V]]
	mp map[K]*linklist.Node[Node[K,V]]
	defaultValue V
	capacity int
}

func NewLruCache[K comparable, V any](cap int, defVal V) *LruCache[K,V] {
	return &LruCache[K, V] {
		list : linklist.NewList[Node[K,V]](),
		mp : make(map[K]*linklist.Node[Node[K,V]], cap),
		defaultValue: defVal,
		capacity : cap,
	}
}

func (lru *LruCache[K, V]) All() iter.Seq2[K, V]{
	return func(yield func(K, V) bool) {
		for it := range lru.list.All(){
			if !yield(it.key, it.value){
				return
			}
		}
	}
}

func (lru *LruCache[K, V]) Put(key K, val V) {
	if v, ok := lru.mp[key]; ok {
		lru.list.PutNewValue(v, Node[K, V]{key,val})
		lru.list.MoveToFront(v)
		return
	} 

	if lru.capacity <= lru.list.Size() {
		lru.list.PutNewValue(lru.list.Back(), Node[K, V]{key, val})
		lru.list.MoveToFront(lru.list.Back())
		lru.mp[key] = lru.list.Front()
		return
	} 

	lru.list.PushFront(Node[K, V]{key, val})
	lru.mp[key] = lru.list.Front()	
}

func (lru *LruCache[K, V]) Get(key K) V {
	if v, ok := lru.mp[key]; ok {
		lru.list.MoveToFront(v)
		return v.GetData().value
	} 
	return lru.defaultValue
}

func (lru *LruCache[K, V]) Size() int {
	return lru.list.Size()
}