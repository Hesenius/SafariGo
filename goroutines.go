package main

import (
	"fmt"
)


func printThings(c chan bool) {
	for x := 0; x < 10_000; x++ {
		fmt.Println("> ", x)
	}
	c <- true
}

func main()  {
	c := make(chan bool)

	go printThings(c)
	fmt.Println("main still running....")
	b := <- c

	fmt.Println("got ", b)
}