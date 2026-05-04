package main

import (
	"fmt"
	"strings"
)

// definition
type User struct {
	UserName string
	Email string
	Id uint32
}

func (u User) PrintUser () {
	u.UserName = "Sashi"
	fmt.Println("User name is", u.UserName)
}

func (u *User) UpdateUser (name string) {
	u.UserName = name
}

// alias for string and add a receiver
type Email string

func (e Email) IsValid () bool {
	return strings.Contains(string(e), "@")
}

func getName (uPtr *User) string {
	return (*uPtr).UserName
}

func main () {
	

	// Initialization
	user1 := User {
		UserName: "Surya",
		Email: "suryakant.hota2@gmail.com",
		Id: 23423423,
	}

	user1.PrintUser()
	fmt.Println(user1)
	user1.UpdateUser("Sashi")
	fmt.Println(user1)

	// Alias for string and validation
	var email = Email("suryakant.hota2@gmail.com")
	fmt.Println("Is Email valid", email.IsValid())

	user2 := User {
		UserName: "Surya",
		Email: "suryakant.hota2@gmail.com",
		Id: 12234,
	}

	fmt.Println("User name: ", getName(&user2))
}