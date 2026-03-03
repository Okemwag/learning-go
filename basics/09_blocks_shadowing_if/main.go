package main

import "fmt"

func main() {
	// A block is any region enclosed in braces: { ... }.
	// Blocks create scope boundaries for variables.
	status := "outer"
	fmt.Println("Before inner block:", status)

	{
		// This status shadows the outer variable.
		// It exists only inside this inner block.
		status := "inner"
		fmt.Println("Inside inner block:", status)
	}

	// Outside the inner block, the outer variable is still visible.
	fmt.Println("After inner block:", status)

	score := 82

	// if can declare a short-lived variable before the condition.
	// That variable exists only for the if/else chain.
	if grade := determineGrade(score); grade == "A" {
		fmt.Println("Excellent. Grade:", grade)
	} else if grade == "B" {
		fmt.Println("Good work. Grade:", grade)
	} else {
		fmt.Println("Keep improving. Grade:", grade)
	}

	// This is a plain if with a boolean condition.
	isPassing := score >= 50
	if isPassing {
		fmt.Println("The score is passing.")
	}
}

func determineGrade(score int) string {
	if score >= 90 {
		return "A"
	}
	if score >= 75 {
		return "B"
	}
	if score >= 60 {
		return "C"
	}
	return "D"
}
