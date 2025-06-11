package main

import (
	"fmt"
	"golang/tree"
	"golang/utils"
	"strings"
)


const (
	name = "Keith"
	email = "camelbyte@proton.me"
)

var names []string

func main() {
	names = append(names, name)
	names = append(names, "Sophia", email)
	fmt.Println(strings.Join(names, ", "))

	a := utils.ReverseArrayCopy(names)
	fmt.Println(a)
	v := []int{1,2,3,4,54,4,5,6,67,78}
	flip := utils.Flip(v)
	fmt.Println(flip)

	// Treesort the slice of values
	fmt.Println(tree.TreeSort(v))


	// initialize a tree using our definition from tree/tree.go
	node := &tree.Tree{}
	// Tree is currently empty.
	// Functions of *Tree include [add, appendValues]
	fmt.Println(node)
	node.Add(4)
	node.AppendValues(v)
	fmt.Println(node)
	


}
