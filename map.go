package main

import (
	"fmt"
)

func main () {
	var m map[string] int32

	m = map[string] int32 {
		"key1": 20,
		"key2": 19,
	}

	fmt.Println("Raw map is", m)
	fmt.Println("Length of m:", len(m))

	m["key3"] = 54
	m["key4"] = 55
	fmt.Println("Updated map is: ", m)

	v,p := m["key5"]
	fmt.Println("v, p", v, p)

	for key, value := range m {
		fmt.Println("Inside loop", key, value)
	}

	// Initialize an empty map
	m1 := map[string]int32{}
	fmt.Println("Empty map:", m1)
}
