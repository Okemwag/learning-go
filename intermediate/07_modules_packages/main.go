package main

import (
	"fmt"
	stdmath "math"

	hello "github.com/Okemwwag/learning-go/intermediate/07_modules_packages/greeting"
	"github.com/Okemwwag/learning-go/intermediate/07_modules_packages/internal/secret"
)

func main() {
	// Repository, module, and package are different levels:
	// repository: the overall source control project
	// module: the versioned unit defined by go.mod
	// package: a directory of Go files compiled together

	// Creating and accessing a package:
	// greeting is a local package inside this module.
	msg := hello.Build("gopher")

	// Importing and exporting:
	// Build, Message, and Text are exported from greeting.
	fmt.Println(msg.Text())
	fmt.Println(hello.DebugText("gopher"))

	// Overriding a package's name:
	// The math package is imported with the alias stdmath.
	fmt.Println("Pi from aliased import:", stdmath.Pi)

	// Using the internal package:
	// This works because this main package sits inside the parent tree
	// that owns the internal directory.
	fmt.Println(secret.Hint())
}
