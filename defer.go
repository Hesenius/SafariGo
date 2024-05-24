package main

import (
	"errors"
	"fmt"
	"math/rand"
)

func df(msg string) {
	// can use "recover" here (not considered normal)
	fmt.Println("Deferred message", msg)
}

func getMsg() string {
	fmt.Println("getting message")
	return "My message"
}

func mightBreak() {
	defer df(getMsg())

	fmt.Println("Trying something")
	if rand.Float32() > 0.5 {
		panic(errors.New("oops!!!"))
	}
	fmt.Println("Succeeded")
}

func main() {
	mightBreak()
	fmt.Println()
}
