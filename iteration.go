package main

import "fmt"

func main() {
	names := []string{"Fred", "Jim", "Sheila"}
	for i, v := range names {
		fmt.Println("index", i, "value", v)
	}

	numbers := []int{0, 1, 2, 3, 4, 5}
	// older Go created ONE variable i, and ONE variable v
	// and repeatedly overwrote it
	// after 1.18, before 1.22, this was changed :)
	// JavaScript closure "issue" using "var", instead of let
	for i, v := range numbers {
		fmt.Println(v)
		// v += 10
		numbers[i] += 10
		fmt.Println(v)
	}
	fmt.Println(numbers)
}
