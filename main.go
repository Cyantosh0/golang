package main

import (
	"fmt"

	go_beta "github.com/Cyantosh0/golang/go_1.18"
)

func main() {

	workWithGoBeta()
}

func workWithGoBeta() {
	go_beta.Display(fmt.Sprintf("Generic Sums: %v and %v",
		go_beta.SumIntsOrFloats(34, 12),
		go_beta.SumIntsOrFloats(35.98, 26.99)))
}
