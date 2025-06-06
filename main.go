package main

import (
	"fmt"
	"golang/channels"
	"golang/testMod"
	"golang/tree"
	"reflect"
	"slices"

)

var p = fmt.Println
var pf = fmt.Printf

func main() {


	// 'fmpf' is 'fmt.Printf()' shortcut


	fmt.Println(testmod.Greet("Keith"))

	msg := "Hello, gopher!"
	ch := make(chan string)
	channels.SendChanMsg(msg, ch)
	p("Message sent to channel:", <-ch)

	a := []int{1, 1, 2, 3, 5, 8, 13, 21, 34, 55, 89}
	s := slices.Clone(a)
	fmt.Printf("Original slice is: %v\n", s)
	tree.TreeSort(s)

	fmt.Printf("Sorted slice is: %v\n", s)
	fmt.Printf("Type of slice is: %v\n", reflect.TypeOf(s))
	fmt.Printf("Type of elements in slice is: %v\n", reflect.TypeOf(s[0]))
	fmt.Printf("Type of elements in slice is: %T\n", s[0])


	fmt.Println("\nTree Sort: ")
	var t tree.Tree 
	fmt.Println(t)


	fmt.Println("Slices: ")
	fmt.Println(s)
	i := 33
	w := append(s, i)
	fmt.Println(w)



}
