package main

import "fmt"

func sum(vals ...int) (total int) {
	// get a slice of int
	fmt.Println("Got", len(vals))
	for _, v := range vals {
		total += v
	}
	return
}

func main() {
	fmt.Println(sum(1, 2, 3, 4, 5, 6, 7, 8, 9, 10))
	numArray := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	// fmt.Println(sum(numArray...)) // cannot "spread" an array
	fmt.Println(sum(numArray[2:7]...))
	fmt.Println(sum(numArray[:]...))
}
