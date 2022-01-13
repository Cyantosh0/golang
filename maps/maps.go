package maps

import "fmt"

func Maps() {
	// initialize map
	m := make(map[string]int)

	n := map[string]int{"n1": 1, "n2": 2}
	fmt.Println("n:", n)

	// Insert an element in map m
	m["m1"] = 42
	m["m2"] = 50
	displayMap(m)

	// Update an element in map m
	m["m1"] = 20
	m["m2"] = 30
	displayMap(m)

	// Delete an element
	delete(m, "m2")
	displayMap(m)

	// Test that a key is present
	m1, ok := m["m1"]
	fmt.Println(`The value in m["m1"]:`, m1, "   Present in map?", ok)

	m2, ok := m["m2"]
	fmt.Println(`The value in m["m2"]:`, m2, "   Present in map?", ok)
}

func displayMap(m map[string]int) {
	fmt.Println("m:", m, `   m["m1"]:`, m["m1"], `   m["m2"]:`, m["m2"])
}
