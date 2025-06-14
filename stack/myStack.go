package stack 

import (
	"fmt"
)


type sNode struct {
	data string 
	next *Node
}


type sStack struct {
	top *Node 
	length int
}


func (s *Stack) Push(data string) {
	newNode := &Node {
		data: data,
		next: s.top,
	}
	s.top = newNode
	s.length++
	fmt.Println(newNode.data)
}






