package main

import "fmt"

func main() {
	Display(fmt.Sprintf("Generic Sums: %v and %v",
		SumIntsOrFloats(34,12),
		SumIntsOrFloats(35.98,26.99)))
}

func SumIntsOrFloats[T int | float64](a, b T) T {
	return a+b
}

func Display(data any) {
	fmt.Println(data)
}

