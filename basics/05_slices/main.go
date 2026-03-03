package main

import "fmt"

func main() {
	// Declaring a nil slice gives you a zero-value slice.
	// It has len 0, cap 0, and no backing array yet.
	var nilSlice []int

	// Slice literals create a slice with initial values.
	numbers := []int{10, 20, 30}

	// make creates a slice with a chosen length and capacity.
	buffer := make([]int, 2, 5)

	fmt.Println("nilSlice len/cap:", len(nilSlice), cap(nilSlice))
	fmt.Println("numbers len/cap:", len(numbers), cap(numbers))
	fmt.Println("buffer len/cap:", len(buffer), cap(buffer))

	// append adds elements and may allocate a new backing array.
	numbers = append(numbers, 40, 50)
	fmt.Println("After append:", numbers, "len/cap:", len(numbers), cap(numbers))

	// Slicing a slice creates a new slice view into the same backing array.
	window := numbers[1:4]
	fmt.Println("window:", window)

	// Emptying a slice while keeping the backing array is common.
	numbers = numbers[:0]
	fmt.Println("Emptied numbers:", numbers, "len/cap:", len(numbers), cap(numbers))

	// copy duplicates values from source into destination.
	source := []int{1, 2, 3, 4}
	destination := make([]int, len(source))
	copied := copy(destination, source)
	fmt.Println("Copied count:", copied, "destination:", destination)

	// Arrays can be sliced to produce a slice.
	array := [4]int{7, 8, 9, 10}
	fromArray := array[:]
	fmt.Println("Slice from array:", fromArray)

	// A practical way to turn slice data into an array value is to copy
	// the needed elements into a new array.
	firstThree := numbersOrFallback(source)
	fmt.Println("Array value copied from slice data:", firstThree)
}

func numbersOrFallback(values []int) [3]int {
	var result [3]int
	copy(result[:], values)
	return result
}
