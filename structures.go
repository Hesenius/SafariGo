package main

import "fmt"

type Person struct {
	Name, Address string
	CreditLimit   int
}

func (p *Person) IncreaseCredit() {
	fmt.Printf("address of p is %p\n", p)
	// p->Cre...?? Not needed here
	p.CreditLimit += 1000
}

func (p Person) ShowPerson() {
	fmt.Printf("Person: %s, %s, %d\n", p.Name, p.Address, p.CreditLimit)
}

// func IncreaseCredit(p *Person) {
// 	fmt.Printf("address of p is %p\n", p)
// 	// p->Cre...?? Not needed here
// 	p.CreditLimit += 1000
// }

// func ShowPerson(p Person) {
// 	fmt.Printf("Person: %s, %s, %d\n", p.Name, p.Address, p.CreditLimit)
// }

func main() {
	// var p1 Person = Person{"Fred", "Over the Hill", 1000}
	// var p1 Person = Person{
	// 	"Fred",
	// 	"Over the Hill",
	// 	1000, // trailing comma required on multi-line form
	// }
	var p1 Person = Person{
		Address: "Over the Hill", // named form likely preferred?
		Name:    "Fred",
	}
	fmt.Println(p1)
	fmt.Println(p1.CreditLimit)
	// ShowPerson(p1)
	p1.ShowPerson()
	fmt.Printf("address of p1 is %p\n", &p1)
	p1.IncreaseCredit() // Pointer taken implicitly!!!
	p1.ShowPerson()
	// IncreaseCredit(&p1)
	// ShowPerson(p1)
}
