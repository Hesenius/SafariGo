package main

import "fmt"

func add(a, b int) int {
	return a + b
}

func makeIncrementer(inc int) func(int) int {
	return func(x int) int {
		return inc + x
	}
}

func main() {
	var sum func(int, int) int
	fmt.Printf("sum is %T\n", sum)
	sum = add
	fmt.Println(sum(1, 2))

	// res := func(a, b int) int {
	// 	return a + b
	// }(10, 20)
	// fmt.Println(res)

	// anonymous function == "lambda expression"
	sum = func(a, b int) int {
		return a + b
	}
	fmt.Println(sum(15, 15))

	addThree := makeIncrementer(3)
	addSeven := makeIncrementer(7)

	fmt.Println("5 addThree is", addThree(5))
	fmt.Println("5 addSeven is", addSeven(5))

}
