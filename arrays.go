package main

import "fmt"
import "sort"

func main () {
	// define an array
	var arr [5]int

	arr[0] = 4
	arr[3] = 4

	fmt.Println(arr)

	// declare and initialize in one statement
	arr1 := [3]int{6, 4, 5}

	fmt.Println(arr1)

	// Iterate over an array
	for i:= 0; i<len(arr); i++ {
		fmt.Println(arr[i])
	}

	// Iterate using range operator
	for _, val := range arr1 {
		fmt.Println(val)
	}

	// Array slicing
	fmt.Println(arr1[:2])
	fmt.Println(arr1[1:3])
	fmt.Println(arr1[:3])

	// Array comparison
	arr2 := [4]int{1, 2, 3, 4}
	arr3 := [4]int{1, 2, 3, 5}
	arr4 := [4]int{1, 2, 3, 4}

	fmt.Println(arr2 == arr3)
	fmt.Println(arr2 == arr4)

	// Sort array using sort package
	sort.Ints(arr1[:])
	fmt.Println(arr1)

	// Copy array
	arr5 := [4]int {3, 7, 1, 2}
	arr6 := [4]int {2, 1, 3, 7}

	copy(arr5[1:4], arr6[0:2])
	fmt.Println(arr5)

	// Array of structs to help indicate if its an actual zero or not existing index

	type Slot struct {
		val int;
		exists bool;
	}

	arr7 := [4]Slot{}
	arr7[0] = Slot{val: 0, exists: true}

	fmt.Println(arr7)

	// Find element in array

	i1 := sort.Search(len(arr), func(i int) bool { return arr[i] >= 4})

	fmt.Println(i1)

	// 2D array

	arr8 := [2][3]int {{1, 2, 3}}
	fmt.Println(arr8)

	// Array of maps
	var arr9 [2] map[string]int
	arr9[0] = map[string]int{"a": 1, "b": 2}
	arr9[1] = map[string]int{"c": 3, "d": 4}
	fmt.Println(arr9)
}
