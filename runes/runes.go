package runes

import (
	"fmt"
	"reflect"
)

func Runes() {
	a1 := '🦍'
	var a2 = 'k'
	var a3 rune = '🦡'

	fmt.Printf("%c - %s\n", a1, reflect.TypeOf(a1))
	fmt.Printf("%c - %s\n", a2, reflect.TypeOf(a2))
	fmt.Printf("%c - %s\n", a3, reflect.TypeOf(a3))

	a4 := '\U0001F3A8'
	fmt.Printf("%c\n", a4)

	s1 := "falcon"
	r1 := []rune(s1)
	fmt.Printf("%U\n", r1)
}
