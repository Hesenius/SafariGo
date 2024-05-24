package main

import "fmt"

func main() {
	// var names [3]string
	// var names [3]string =	[3]string{"", "", ""}

	// names := [3]string{"", "", ""}
	// names := [3]string{"Albert"}
	names := [3]string{"Albert", "Alice", "Algernon"}

	fmt.Printf("names is %#v, contains %d elements\n", names, len(names))

	// cannot assign or == compare dissimilar (size or base type) arrays
	// var moreNames [4]string
	var moreNames [3]string

	moreNames = names
	fmt.Printf("morenames is %#v, contains %d elements\n", moreNames, len(names))
	fmt.Println("equal?", names == moreNames)

	people := []string{"Fred", "Jim", "Sheila", "Alice", "Bob", "Susan", "Bob", "Colin"}
	fmt.Println("first person is %s\n", people[0])

	somePeople := people[3:5] // people is a slice, so is somePeople, cannot compare with ==
	// var moreStill []string
	// moreStill = somePeople
	fmt.Println(somePeople)
	somePeople[0] = "Frank"
	fmt.Println(somePeople)
	fmt.Println(people)

	somePeople = append(somePeople, "Bill", "Larry", "John", "Jim", "Jane")
	fmt.Println(somePeople)
	fmt.Println(people)
}
