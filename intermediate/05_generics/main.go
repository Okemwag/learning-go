package main

import "fmt"

// Generics reduce repetitive code and increase type safety.
// Before generics, you would often duplicate this kind of logic for []int, []string, etc.

// SumNumbers uses type terms to specify the operators we need.
// The + operator only works because the constraint allows numeric types.
type Number interface {
	~int | ~int64 | ~float64
}

func SumNumbers[T Number](values []T) T {
	var total T
	for _, value := range values {
		total += value
	}
	return total
}

// Generic functions abstract algorithms:
// Contains works for any slice whose element type is comparable.
func Contains[T comparable](values []T, target T) bool {
	for _, value := range values {
		if value == target {
			return true
		}
	}
	return false
}

// Type inference usually lets callers omit explicit type arguments.
func First[T any](values []T) (T, bool) {
	var zero T
	if len(values) == 0 {
		return zero, false
	}
	return values[0], true
}

// Generic data structures combine well with generic functions.
type Stack[T any] struct {
	items []T
}

func (s *Stack[T]) Push(value T) {
	s.items = append(s.items, value)
}

func (s *Stack[T]) Pop() (T, bool) {
	var zero T
	if len(s.items) == 0 {
		return zero, false
	}

	last := len(s.items) - 1
	value := s.items[last]
	s.items = s.items[:last]
	return value, true
}

func (s Stack[T]) Len() int {
	return len(s.items)
}

// Generics and interfaces work together.
// The constraint says T can be any type that has String() string.
type Stringer interface {
	String() string
}

func JoinStrings[T Stringer](values []T) []string {
	out := make([]string, 0, len(values))
	for _, value := range values {
		out = append(out, value.String())
	}
	return out
}

type Label string

func (l Label) String() string {
	return "label:" + string(l)
}

// Type elements can limit constants and operations.
// Because T is constrained to ~int | ~int64, the untyped constant 10 can be converted into T.
func AddTen[T ~int | ~int64](value T) T {
	return value + 10
}

// Things left out:
// Go generics do not support specialization, generic methods on non-generic interfaces,
// or arbitrary operator overloading. Constraints deliberately stay simple.

func main() {
	// Generics reduce repetitive code and increase type safety:
	// one function, multiple concrete types, compile-time checked.
	ints := []int{1, 2, 3, 4}
	floats := []float64{1.5, 2.5, 3.0}

	fmt.Println("sum ints:", SumNumbers(ints))
	fmt.Println("sum floats:", SumNumbers(floats))

	// comparable is useful for equality-based generic algorithms.
	fmt.Println("contains int:", Contains([]int{1, 2, 3}, 2))
	fmt.Println("contains string:", Contains([]string{"go", "rust"}, "go"))

	// Type inference means you usually do not write First[int](...).
	first, ok := First([]string{"generics", "in", "go"})
	fmt.Println("first string:", first, ok)

	// You can still be explicit if needed.
	explicit, explicitOK := First[int]([]int{9, 8, 7})
	fmt.Println("explicit type args:", explicit, explicitOK)

	// Generic functions and generic data structures work well together.
	var stack Stack[string]
	stack.Push("a")
	stack.Push("b")
	fmt.Println("stack len:", stack.Len())

	popped, poppedOK := stack.Pop()
	fmt.Println("stack pop:", popped, poppedOK)

	// Generics and interfaces:
	labels := []Label{"one", "two"}
	fmt.Println("join stringers:", JoinStrings(labels))

	// Type elements limit constants:
	// the constant 10 must fit the permitted types.
	fmt.Println("add ten:", AddTen(32))

	// More on comparable:
	// slices are not comparable, so this would not compile:
	// _ = Contains([][]int{{1}}, []int{1})
	//
	// comparable allows == and !=, but only for types the language marks as comparable.
}
