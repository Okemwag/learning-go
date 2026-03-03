package main

import "fmt"

func main() {
	// make allocates and initializes a writable map.
	scores := make(map[string]int)

	// Writing to a map uses assignment by key.
	scores["alice"] = 90
	scores["bob"] = 75

	// Reading from a map returns the value for the key.
	fmt.Println("alice score:", scores["alice"])

	// The comma-ok idiom distinguishes a missing key from a present zero value.
	score, ok := scores["charlie"]
	fmt.Println("charlie score:", score, "exists:", ok)

	// Delete removes a key if present. It is safe even if the key is absent.
	delete(scores, "bob")
	fmt.Println("After deleting bob:", scores)

	// Emptying a map can be done by deleting keys.
	for key := range scores {
		delete(scores, key)
	}
	fmt.Println("After emptying map:", scores)

	// Using a map as a set is common in Go.
	seen := map[string]bool{
		"go":   true,
		"rust": true,
	}
	fmt.Println("set contains go:", seen["go"])
}
