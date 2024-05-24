package main

import (
	"fmt"
	"myprog/dates"
)

// interfaces specify a requirement, or "contract", for
// certain *methods* (method, not functions)
type HasYear interface {
	GetYear() int
}

type Vehicle struct {
	Year int
}

func (v Vehicle) GetYear() int {
	return v.Year
}

func main() {
	today := dates.Date{Day: 24, Month: 5, Year: 2024}
	fmt.Println(today)
	today.Tomorrow()
	fmt.Println(today)

	newYearsDay := dates.Holiday{dates.Date{1, 1, 2025}, "New Year's Day"}
	fmt.Println(newYearsDay)
	newYearsDay.Tomorrow()
	fmt.Println(newYearsDay)

	var yearThing HasYear
	fmt.Printf("yearThing is %T, %v\n", yearThing, yearThing)
	yearThing = today
	fmt.Printf("yearThing is %T, %v\n", yearThing, yearThing)
	fmt.Println(yearThing.GetYear())

	yearThing = Vehicle{1995}
	fmt.Println(yearThing.GetYear())
}
