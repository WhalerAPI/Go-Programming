package main

import (
	"fmt"
	"golang/utils"
	"strings"
	"golang/tree"
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

	s := strings.Join(names, email)
	fmt.Println(s)

	v := []int{1,2,3,4,54,4,5,6,67,78}
	fmt.Println(len(v)-1-len(v))
	flip := utils.Flip(v)
	fmt.Println(flip)

	fmt.Println(tree.TreeSort(v))
}
