package main

import "fmt"

type User struct {
	ID   int
	Name string
}

type Account struct {
	ID   int
	Name string
}

func main() {
	// A named struct groups related fields into one value.
	user := User{
		ID:   1,
		Name: "Go Learner",
	}

	// Anonymous structs are useful for short-lived grouped data.
	session := struct {
		Token   string
		Expires bool
	}{
		Token:   "abc123",
		Expires: false,
	}

	// Structs are comparable when all fields are comparable.
	sameUser := User{ID: 1, Name: "Go Learner"}
	fmt.Println("user == sameUser:", user == sameUser)

	// Different named struct types can be converted if their fields are compatible.
	account := Account(user)

	fmt.Println("User:", user)
	fmt.Println("Anonymous struct:", session)
	fmt.Println("Converted Account:", account)
}
