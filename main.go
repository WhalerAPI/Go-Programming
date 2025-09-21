package main

import (
	"Go-Data-Structures/functions"
	"fmt"
	"time"
)

func sort(n []int) []int {
	for i := range n {
		for j := i + 1; j < len(n); j++ {
			if n[i] > n[j] {
				n[i], n[j] = n[j], n[i]
			}
		}
	}
	return n
}

func main() {
	start := time.Now()
	numbers := []int{3, 2, 38, 61, 4, 11, 9, 01, 44, 43, 55, 23, 67, 89, 12, 34, 56, 78}
	maxC := functions.Max(numbers)
	fmt.Println("Elapsed time after compute: ", time.Since(start))
	fmt.Println("Max found:", maxC)
	fmt.Println("Execution time:", time.Since(start))

	fmt.Println("duration from start")
	fmt.Println(sort(numbers))
	fmt.Println("Execution time:", time.Since(start))

}
