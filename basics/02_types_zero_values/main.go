package main

import (
	"fmt"
)

func main() {
	// Predeclared types are built into the language.
	// Here we declare several values without assigning them.
	// Go gives each one its zero value automatically.
	var isReady bool
	var count int
	var price float64
	var name string
	var letter rune

	// Literals are values written directly in source code.
	boolLiteral := true
	intLiteral := 42
	floatLiteral := 19.95
	stringLiteral := "Go is explicit"
	runeLiteral := 'G'

	fmt.Println("Zero value bool:", isReady)
	fmt.Println("Zero value int:", count)
	fmt.Println("Zero value float64:", price)
	fmt.Println("Zero value string:", name)
	fmt.Println("Zero value rune:", letter)

	fmt.Println("Boolean literal:", boolLiteral)
	fmt.Println("Numeric literal (int):", intLiteral)
	fmt.Println("Numeric literal (float):", floatLiteral)
	fmt.Println("String literal:", stringLiteral)
	fmt.Println("Rune literal:", runeLiteral, "as character:", string(runeLiteral))
}
