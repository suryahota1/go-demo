package main

import (
	"fmt"
	"unsafe"
)

func main () {
	// declaration
	// var i int = 1 // declaration and initialization - long hand - type explicitly declared
	// j := 1 // declaration and initialization - short hand - type inferred
	// k := int64(1) // declaration and initialization - short hand - type explicitly declared

	// Constant
	const a int32 = 43

	const (
		year = 43
		leap = uint64(4)
	)

	var c = 90
	d := 91

	fmt.Println("Value of c is ", c, "Value of d is ", d)
	fmt.Printf("the type of c is %T and size of c is %d \n", c, unsafe.Sizeof(c))

	e := 5.67
	fmt.Printf("The type of e is %T and the size of e is %d", e, unsafe.Sizeof(e))

	// type conversion
	g := 5
	h := 9.45

	i := g + int(h)
	fmt.Println("\ni = ", i)
	fmt.Println("Addredd of a", &a);
}