package main

import (
	"fmt"
)

func Greet(name string) string {
	fmt.Printf("Hello, %s!!\n", name)
	return name
}

func main() {
	Greet("World")

	data := TodoPageData{
		PageTitle: "My TODO list",
		Todos: []Todo{
			{Title: "Task 1", Done: false},
			{Title: "Task 2", Done: true},
			{Title: "Task 3", Done: true},
		},
	}


	fmt.Printf("%+v\n", data)
}

