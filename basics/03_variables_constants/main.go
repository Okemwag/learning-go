package main

import "fmt"

// Untyped constants can be used flexibly until a concrete type is required.
const maxUsers = 100

// Typed constants have a fixed type from the moment they are declared.
const discountRate float64 = 0.15

func main() {
	// The var keyword can be used at package scope or inside functions.
	// It is useful when you want the zero value or an explicit type.
	var accountName string = "Go Learner"

	// Short declaration with := is only allowed inside functions.
	// It is the most common style for local variables.
	points := 250

	// Grouped var declarations help when several values belong together.
	var (
		isActive bool   = true
		level    int    = 2
		label    string = "starter"
	)

	// Constant values never change after declaration.
	fmt.Println("Account:", accountName)
	fmt.Println("Points:", points)
	fmt.Println("Active:", isActive)
	fmt.Println("Level:", level)
	fmt.Println("Label:", label)
	fmt.Println("Untyped const maxUsers:", maxUsers)
	fmt.Println("Typed const discountRate:", discountRate)

	// Naming variables clearly matters more than clever abbreviations.
	// Prefer names that describe meaning in context.
	requestCount := 7
	fmt.Println("Well-named variable requestCount:", requestCount)
}
