package main

import (
	"fmt"
	"container/list"
	"golang/functions"
	"golang/lists"
	"golang/tree"
	"golang/utils"
)

func main() {

	s := functions.CollectStuff()
	fmt.Println("\nHere is a collection of stuff: \n\n", s)
	
	var a []int
	a = []int{1, 43, 64, 22, 34, 78}

	fmt.Println(tree.TreeSort(a))
	fmt.Println(utils.Flip(a))

	var L lists.List[int]
	L.PushBack(10)
	L.PushFront(5)    
	L.PushBack(42)

	fmt.Println("Let's see what our list looked like as a slice:", L.MakeSlice()) 

	L.PushFront(33)
	L.PushBack(11)
	L.PushFront(22)
	fmt.Println(L.MakeSlice())


	// Instead of creating your own linked list you can use the
	// Go standard libary 1.18+ and import "container/list" then instantiate a list.New() 
	// Below is an example. Notice "container/list" methods provided are the same as the ones I created.

	ll := list.New()
	ll.PushFront(13333)
	ll.PushFront(44441)
	fmt.Println("Front:", ll.Back().Value)
	fmt.Println("Back:", ll.Front().Value)
	fmt.Println(ll.Back().Prev().Value)

}


