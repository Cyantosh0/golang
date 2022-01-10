package go_beta

func SumIntsOrFloats[T int | float64](a, b T) T {
	return a+b
}