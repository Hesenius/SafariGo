package main

import "fmt"

func AddOne(x int) {
	x++ // these do not have "value"
	// ++x; // these don't exist
}

func AddOneViaPointer(xp *int) {
	// xp++ // Go prohibits arbitrary pointer manipulation (thank goodness!)
	(*xp)++
}

func main() {
	name := "Fred"
	otherName := name

	fmt.Printf("name is %s, stored at %p, type is %T\n", name, &name, name)
	fmt.Printf("othername is stored at %p\n", &otherName)

	x := 10
	fmt.Println(x)
	AddOne(x)
	fmt.Println(x)
	AddOneViaPointer(&x)
	fmt.Println(x)

}
