package main

import "fmt"

func main() {
	// A string is immutable byte data, usually UTF-8 text.
	word := "Go語"

	// Indexing a string returns a byte, not necessarily a full character.
	firstByte := word[0]

	// Converting to []byte exposes raw encoded bytes.
	rawBytes := []byte(word)

	// Converting to []rune exposes Unicode code points.
	runes := []rune(word)

	fmt.Println("String:", word)
	fmt.Println("Byte at index 0:", firstByte)
	fmt.Println("Bytes:", rawBytes)
	fmt.Println("Runes:", runes)
	fmt.Println("Rune count:", len(runes))

	// Iterating with range over a string yields rune indexes and rune values.
	for index, r := range word {
		fmt.Println("range index:", index, "rune:", r, "char:", string(r))
	}
}
