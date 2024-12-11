package main

import (
	"fmt"
	"project_1/pkg/linklist"
	"project_1/pkg/lrucache"
)


func main(){
	b := linklist.NewList[int]()
	c := lrucache.NewLruCache[int, int](10, 0)
	c.Put(23, 14)
	// b.PushBack("1")
	// b.PushFront("2")
	// b.PushBack("3")
	// b.PushFront("4")
	// b.PushBack("5")
	// b.PushFront("6")
	// b.PushBack("7")
	// b.PushFront("8")

	b.PushBack(1)
	res := b.Remove(b.Back())
	fmt.Println(res)
	b.PushBack(2)
	res = b.Remove(b.Front())
	fmt.Println(res)
	fmt.Println(b.Size())
	fmt.Println()

	b.PushBack(3)
	b.PushFront(4)
	b.MoveToFront(b.Back())
	b.PushBack(5)
	b.MoveToFront(b.Back())
	b.PushFront(6)
	b.MoveToBack(b.Front())
	b.PushBack(7)
	b.PushBack(8)
	// 5 3 4 6 7 8 

	// b.Remove(b.Front())
	// b.Remove(b.Front())
	// b.Remove(b.Front())
	//6 4 1 2 3 5 7 8
	// res = b.Remove(b.Back())
	// fmt.Println(res)
	// res = b.Remove(b.Front())
	// fmt.Println(res)
	// res = b.Remove(b.Back())
	// fmt.Println(res)
	// fmt.Println()
	
	for it := range b.All(){
		fmt.Println(it)
	}

	// b.All(func(n *linklist.Node[int]){
	// 	n.Print()
	// })
}