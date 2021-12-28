package main

import "fmt"

func main() {
	fmt.Printf("Generic Sums: %v and %v\n",
		SumIntsOrFloats(34,12),
		SumIntsOrFloats(35.98,26.99))
}

func SumIntsOrFloats[T int | float64](a, b T) T {
	return a+b
}

