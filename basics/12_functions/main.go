package main

import "fmt"

// Function type declarations let you name a function signature.
// Any function with the same parameter and return types can use this type.
type IntOperation func(int, int) int

func main() {
	// Declaring and calling functions:
	// add is a normal named function declared below.
	total := add(10, 5)
	fmt.Println("add:", total)

	// Simulating named parameters:
	// Go has no named parameters, so structs are a common alternative.
	report := formatUser(FormatUserParams{
		FirstName: "Ada",
		LastName:  "Lovelace",
		Title:     "Engineer",
	})
	fmt.Println("simulated named params:", report)

	// Simulating optional parameters:
	// Go has no optional parameters, so a config struct works well.
	defaultGreeting := greet(GreetConfig{Name: "Gopher"})
	customGreeting := greet(GreetConfig{Name: "Gopher", Prefix: "Welcome"})
	fmt.Println("default greeting:", defaultGreeting)
	fmt.Println("custom greeting:", customGreeting)

	// Variadic parameters accept zero or more values.
	sum := sumAll(1, 2, 3, 4)
	fmt.Println("variadic sum:", sum)

	// A slice can be expanded into variadic arguments with ...
	numbers := []int{5, 6, 7}
	fmt.Println("variadic from slice:", sumAll(numbers...))

	// Multiple return values are truly multiple values, not a wrapped tuple object.
	quotient, remainder := divideAndRemainder(17, 5)
	fmt.Println("multiple returns:", quotient, remainder)

	// Ignoring return values with _ is common when a value is not needed.
	justQuotient, _ := divideAndRemainder(17, 5)
	fmt.Println("ignored remainder:", justQuotient)

	// Blank returns rely on named return values.
	// They work, but they make code harder to read in non-trivial functions.
	fmt.Println("blank return result:", blankReturnDemo(4))

	// Functions are values. You can assign them to variables.
	operation := add
	fmt.Println("function as value:", operation(3, 4))

	// Function type declarations make this clearer and easier to pass around.
	var op IntOperation = multiply
	fmt.Println("typed function value:", op(3, 4))

	// Anonymous functions are functions without a named identifier.
	inlineResult := func(a, b int) int {
		return a - b
	}(10, 3)
	fmt.Println("anonymous function:", inlineResult)

	// Closures capture variables from their surrounding scope.
	counter := makeCounter()
	fmt.Println("closure call 1:", counter())
	fmt.Println("closure call 2:", counter())

	// Passing functions as parameters enables reusable behavior.
	fmt.Println("passed function:", apply(8, 2, multiply))

	// Returning functions from functions enables factories and customization.
	doubler := makeMultiplier(2)
	fmt.Println("returned function:", doubler(9))

	// defer schedules a call to run when the surrounding function returns.
	defer fmt.Println("deferred: runs last in main")
	fmt.Println("before deferred call executes")

	// Go is call by value:
	// arguments are copied into parameters.
	original := 10
	tryToChange(original)
	fmt.Println("after tryToChange:", original)

	// Slices are still passed by value, but the slice header points to shared data.
	// That is why changing an element can still affect the caller's backing array.
	values := []int{1, 2, 3}
	changeFirst(values)
	fmt.Println("after changeFirst:", values)
}

func add(a int, b int) int {
	return a + b
}

func multiply(a int, b int) int {
	return a * b
}

type FormatUserParams struct {
	FirstName string
	LastName  string
	Title     string
}

func formatUser(params FormatUserParams) string {
	return params.Title + ": " + params.FirstName + " " + params.LastName
}

type GreetConfig struct {
	Name   string
	Prefix string
}

func greet(config GreetConfig) string {
	prefix := config.Prefix
	if prefix == "" {
		prefix = "Hello"
	}
	return prefix + ", " + config.Name
}

func sumAll(values ...int) int {
	total := 0
	for _, value := range values {
		total += value
	}
	return total
}

func divideAndRemainder(a int, b int) (int, int) {
	return a / b, a % b
}

func blankReturnDemo(value int) (doubled int) {
	doubled = value * 2

	// This works because doubled is a named return value.
	// Avoid this style in real code unless the function is tiny and obvious.
	return
}

func apply(a int, b int, op IntOperation) int {
	return op(a, b)
}

func makeCounter() func() int {
	count := 0

	return func() int {
		count++
		return count
	}
}

func makeMultiplier(factor int) func(int) int {
	return func(value int) int {
		return value * factor
	}
}

func tryToChange(value int) {
	value = 99
	fmt.Println("inside tryToChange:", value)
}

func changeFirst(values []int) {
	if len(values) > 0 {
		values[0] = 999
	}
}
