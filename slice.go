package main

import (
	"fmt"
)

func main () {
	// Slice literal
	s1 := []string{"v1", "v2", "v3", "v4"}
	fmt.Println("s1 is: ", s1)

	// Slice with make function
	s2 := make([]int, 5, 10)
	fmt.Println("s2 is: ", s2)

	// Slice from existing array
	arr := [5]int{5, 4, 3, 2, 1}
	s3 := arr[1:4]
	fmt.Println("s3 is: ", s3)

	s1 = append(s1, "v5", "v6", "v7")
	fmt.Println("Updated s1: ", s1)

	s2Copy := make([]int, len(s2))
	copy(s2Copy, s2)
	fmt.Println("S2 copy is: ", s2Copy)
}