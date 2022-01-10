package enums

import "fmt"

type Weekday int

const (
	Sun Weekday = iota + 1 // EnumIndex = 1
	Mon                    // EnumIndex = 2
	Tue                    // EnumIndex = 3
	Wed                    // EnumIndex = 4
	Thu                    // EnumIndex = 5
	Fri                    // EnumIndex = 6
	Sat                    // EnumIndex = 7
)

// String - Creating common behavior - give the type a String function
func (w Weekday) String() string {
	return [...]string{"Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"}[w-1]
}

// EnumIndex - Creating common behavior - give the type a EnumIndex function
func (w Weekday) EnumIndex() int {
	return int(w)
}

func (w Weekday) DisplayWeekDayDetails() {
	fmt.Println(w)             // Prints : Monday
	fmt.Println(w.String())    // Prints : Monday
	fmt.Println(w.EnumIndex()) // Prints : 2
}
