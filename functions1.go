package main

import (
	"errors"
	"fmt"
	"math"
)

// func Add(a int, b int) int {
//	func Add(a, b int) int {
//		// var sum int
//		// sum = a + b
//		// return sum
//		return a + b
//	}

func Add(a, b int) (sum int) {
	sum = a + b
	// automatically returns the declared "sum"
	//from the return type spec
	return
	// return sum
}

func SumAndDiff(a, b int) (sum int, diff int) {
	sum = a + b
	diff = a - b
	return
}

func SquareRoot(num float64) (float64, error) {
	if num < 0 {
		var err = errors.New("Square root of negative!")
		return 0, err
	} else {
		return math.Sqrt(num), nil
	}
}

func main() {
	a, b := 1, 2
	fmt.Println("sum of 1 and 2 is", Add(a, b))
	fmt.Println(SumAndDiff(a, b))
	// s := SumAndDiff(a, b) // NOT ALLOWED, two return values
	s, _ := SumAndDiff(a, b)
	// fmt.Println(s, d)
	fmt.Println(s)

	var err error
	// var sq float64
	if sq, err := SquareRoot(-1); err != nil {
		fmt.Println("Problem: err=", err)
	} else {
		fmt.Println("Root is ", sq)
	}
	fmt.Println("Error was", err) // NOPE,  out of scope
}
