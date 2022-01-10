package main

import (
	"fmt"

	"github.com/Cyantosh0/golang/enums"
)

func main() {
	displayWeekDayDetails(enums.Mon)

	// workWithGoBeta()
}

func displayWeekDayDetails(weekday enums.Weekday) {
	fmt.Println(weekday)             // Prints : Monday
	fmt.Println(weekday.String())    // Prints : Monday
	fmt.Println(weekday.EnumIndex()) // Prints : 2
}

// func workWithGoBeta() {
// 	go_beta.Display(fmt.Sprintf("Generic Sums: %v and %v",
// 		go_beta.SumIntsOrFloats(34, 12),
// 		go_beta.SumIntsOrFloats(35.98, 26.99)))
// }
