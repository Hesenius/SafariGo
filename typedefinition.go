package main

import "fmt"

type DayOfMonth int

// constants are handled by the compiler
// arithmetic is VERY VERY high precision and range
// also assignable to things of various appropriate target types
const THIRTEEN = 13

func main()  {
	var thirteenth DayOfMonth
	thirteenth = 13 // allowed ONLY BECAUSE 13 is a constant
	var fourteen int = 14
	// thirteenth = fourteen // requires cast call!
	fmt.Println(thirteenth, fourteen)
}