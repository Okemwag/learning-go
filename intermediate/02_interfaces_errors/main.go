package main

import (
	"fmt"
	"strings"
)

type Speaker interface {
	Speak() string
}

type Person struct {
	Name string
}

func (p Person) Speak() string {
	return "hello, I am " + p.Name
}

func normalizeName(input string) (string, error) {
	name := strings.TrimSpace(input)
	if name == "" {
		return "", fmt.Errorf("name cannot be empty")
	}

	return strings.ToUpper(name), nil
}

func main() {
	name, err := normalizeName("  gopher  ")
	if err != nil {
		fmt.Println("error:", err)
		return
	}

	var speaker Speaker = Person{Name: name}
	fmt.Println(speaker.Speak())
}
