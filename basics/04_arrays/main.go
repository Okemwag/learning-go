package main

import "fmt"

func main() {
	// Arrays have a fixed length, and that length is part of the type.
	var zeroed [3]int

	// Array literals create arrays with known contents.
	scores := [4]int{10, 20, 30, 40}

	// You can set specific indexes in a sparse array literal.
	checkpoints := [6]int{0: 100, 3: 400, 5: 600}

	fmt.Println("Zero-value array:", zeroed)
	fmt.Println("Scores:", scores)
	fmt.Println("Checkpoints:", checkpoints)
	fmt.Println("Length of scores:", len(scores))

	// Arrays are comparable when their element types are comparable.
	otherScores := [4]int{10, 20, 30, 40}
	fmt.Println("scores == otherScores:", scores == otherScores)
}
