package main

import "fmt"

func main() {
	day := "Saturday"

	// A normal switch compares one value against multiple cases.
	switch day {
	case "Monday":
		fmt.Println("Start of the work week.")
	case "Saturday", "Sunday":
		fmt.Println("Weekend.")
	default:
		fmt.Println("Midweek.")
	}

	score := 81

	// A blank switch omits the switch expression.
	// Each case becomes a boolean condition.
	switch {
	case score >= 90:
		fmt.Println("Blank switch grade: A")
	case score >= 75:
		fmt.Println("Blank switch grade: B")
	default:
		fmt.Println("Blank switch grade: C or below")
	}

	// if is often better for a small number of simple checks.
	if score%2 == 0 {
		fmt.Println("if example: score is even")
	} else {
		fmt.Println("if example: score is odd")
	}

	// goto jumps to a label in the same function.
	// It should be used rarely and only when it makes flow clearer.
	count := 0

Start:
	if count < 2 {
		fmt.Println("goto loop iteration:", count)
		count++
		goto Start
	}
}
