package main

import (
	"fmt"
	"golangtour/channels"
	"golangtour/tree"
	"reflect"
	"slices"
)

var p = fmt.Println
var pf = fmt.Printf

func main() {

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

}
