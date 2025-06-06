package tree

import (
	"fmt"
	"math/rand"
	"reflect"
)

type Tree struct {
	val         int
	left, right *Tree
}

func TreeSort(values []int) {
	var root *Tree
	for _, v := range values {
		root = add(root, v)
	}
	values = appendValues(values[:0], root)
}

func add(t *Tree, val int) *Tree {
	if t == nil {
		return &Tree{val: val}
	}
	if val < t.val {
		t.left = add(t.left, val)
	} else {
		t.right = add(t.right, val)
	}
	return t
}

func appendValues(values []int, t *Tree) []int {
	if t != nil {
		values = appendValues(values, t.left)
		values = append(values, t.val)
		values = appendValues(values, t.right)
	}
	return values
}

func main() {
	arr := make([]int, 10)

	for i := range arr {
		arr[i] = rand.Intn(100)
	}

	for _, val := range arr {
		fmt.Printf("Value: %v\n", val)
		r := reflect.TypeOf(val)
		fmt.Printf("\n\rType: %v\n", r)
	}

	TreeSort(arr)
	fmt.Println("Sorted:", arr)
}
