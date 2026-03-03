package main

import "fmt"

func main() {
	// 1. The complete for statement: init; condition; post
	for i := 0; i < 3; i++ {
		fmt.Println("Complete for:", i)
	}

	// 2. The condition-only for statement: works like a while loop.
	counter := 0
	for counter < 2 {
		fmt.Println("Condition-only for:", counter)
		counter++
	}

	// 3. The infinite for statement: loop until break is used.
	steps := 0
	for {
		fmt.Println("Infinite for iteration:", steps)
		steps++
		if steps == 2 {
			break
		}
	}

	// 4. The for-range statement: iterate over collections.
	values := []string{"go", "is", "clear"}
	for index, value := range values {
		if index == 1 {
			// continue skips the rest of this iteration.
			continue
		}
		fmt.Println("Range loop:", index, value)
	}

	// Labels let break or continue target an outer loop.
	rows := []string{"A", "B"}
	cols := []int{1, 2, 3}

Outer:
	for _, row := range rows {
		for _, col := range cols {
			if row == "B" && col == 2 {
				fmt.Println("Breaking outer loop at", row, col)
				break Outer
			}
			fmt.Println("Cell:", row, col)
		}
	}
}
