package dates

import (
	"errors"
	"fmt"
)

type Date struct {
	Day, Month, Year int
}

type Holiday struct {
	Date // embedded struct Holiday kinda "is-A"
	Name string
}

// methods on a type MUST be declared in the same package
func (d Date)GetYear() int {
	return d.Year
}

func (dp *Date) Tomorrow() {
	dp.Day++
	// should check for overflow beyond end of month...
}

func IsLeap(year int) bool {
	return year%4 == 0 && year%100 != 0 || year%400 == 0
}

func DaysInMonth(month, year int) (days int) {
	switch month {
	case 9, 4, 6, 11:
		days = 30
	// no fallthrough unless you use the keyword "fallthrough"
	// cases can be expresssions (like JavaScript) not only constants
	// also there's a mechanism for switch on type.
	case 1, 3, 5, 7, 8, 10, 12:
		days = 31
	case 2:
		// Go does not have a conditional operator (? :)
		if IsLeap(year) {
			days = 29
		} else {
			days = 28
		}
	default:
		panic(errors.New("illegal month number"))
	}
	return /* days */
}

func (d Date) String() string {
	return fmt.Sprintf("Date d=%d, m=%d, y=%d", d.Day, d.Month, d.Year)
}
