package go_beta

import "fmt"

func WorkWithGoBeta() {
	Display(fmt.Sprintf("Generic Sums: %v and %v",
		SumIntsOrFloats(34, 12),
		SumIntsOrFloats(35.98, 26.99)))
}
