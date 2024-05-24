package main

import (
	"fmt"
	"math/rand"
	"strconv"

	// redundantly (in this case!) specifying the package name
	// this might be useful if package name does NOT match directory name
	utils "myprog/utils"
)

var Random = rand.Float32()

func main() {
	// var random float32
	// random = rand.Float32()

	// var random float32 = rand.Float32()
	// var random = rand.Float32()
	random := rand.Float32()
	fmt.Println("Hello Go World!")
	fmt.Println(random)
	fmt.Println(Random)

	fmt.Println(utils.Thingy)
	// fmt.Println(utils.name) // Not exported, lower case first letter

	// var num int64
	// string is actually a "slice" of bytes -- intended to be UTF-8

	var num64 int64
	var num16 int16 = 32
	// num64 = num16 // NOT VALID assignment
	num64 = int64(num16)
	fmt.Println(num64)
	num16 = 65

	var strNum string = string(num16)
	fmt.Println(strNum)
	sixtyFive, err := strconv.Atoi("xx65")
	fmt.Println("sixtyFive is ", sixtyFive, "x is", err)

	someNum := fmt.Sprintf("The number is %d", 12345)
	fmt.Println(someNum)
}
