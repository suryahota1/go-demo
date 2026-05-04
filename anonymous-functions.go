package main

import (
	"fmt"
)

func main () {
	func (n int32) {
		fmt.Println("Test", n)
	}(4)

	f := func(x string) {
		fmt.Println(x)
	}
	f("Hello")

	closure := func (context string) func () {
		x := 0
		return func () {
			x++;
			fmt.Println(context, x)
		}
	}

	h1 := closure("h1")
	h2 := closure("h2")

	h1()
	h2()

	h1()
	h2()
}