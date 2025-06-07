package main

import (
	"fmt"
	"golang/stack"
)
//
// func adder(a, b float32) float32 {
// 	return a + b 
// }


func Reverse(a []int) []int {
	b := make([]int, len(a))
	copy(b, a)

	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}
	return b
}


func main() {
	// create a slice of integers and append some values
	var a []int
	a = append(a, 1, 2,3,4,5,6)
	a = append(a, 7,8,9,10)
	fmt.Printf("A: %v\n", a)

	// Loop over the slice 'a' with 2 variables - the iterator
	// and the value. Create a var 'r' to hold the result of
	// any whatever logic is performed. In this case, it's just 
	// adding the iterator value to each array value along
	// with the Sum.
	
	for i, val := range a {
    	r := i + val
    	fmt.Printf("i: %d | val: %d | i+val: %d\n", i, val, r)
	}

	// Stack is an integer implementation of &Stack 
	// Create a new instance of stack from the struct defined in 'stack/arrayStack.go'
	s := &stack.Stack{} 
	// Push an item onto the stack
	s.Push(1)
	s.Print()
	// Peek at the stack 
	s.Peek()
	// Push another int to the stack
	s.Push(2)
	s.Print()
	// Remove or "Pop" an item from the stack
	s.Pop()
	s.Push(33)
	s.Push(44)
	s.Push(77)
	// Print the stack
	s.Print()
	fmt.Println()

	// reverse the array 'a' [1,2,3,4...]
	fmt.Println(Reverse(a))
	// output: [10,9,8,7...]

	
	//
	// var x, y []int 
	// for i := range(12) {
	// 	y = append(x, i)
	// 	fmt.Printf("%d cap=%d\t%v\n", i, cap(y), y)
	// 	x = y
	// }
	//

}
