package interfaces

import "fmt"

type Mapper map[string]interface{}

func MapData(s string, i int, b bool) Mapper {
	return Mapper{
		"m1": s,
		"m2": i,
		"m3": b,
	}
}

func Interfaces() {
	var m interface{} = MapData("apple", 10, true)
	if v, ok := m.(Mapper); ok {
		fmt.Println(v)
		fmt.Println(v["m1"])
	}

}
