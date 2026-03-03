# Intermediate: Generics

## Purpose

This lesson covers:

- how generics reduce repetitive code and increase type safety
- introducing generics in Go
- generic functions as algorithm abstractions
- generics and interfaces
- using type terms to specify operators
- type inference
- how type elements limit constants
- combining generic functions with generic data structures
- more on `comparable`
- things left out of Go generics
- idiomatic Go and generics
- adding generics to the standard library

## Introducing Generics in Go

Generics let you write functions and types that work across multiple concrete types while preserving compile-time type safety.

The basic shape is:

```go
func First[T any](values []T) (T, bool)
```

Here:

- `T` is a type parameter
- `any` is the constraint
- the function works for any concrete type that satisfies the constraint

Go generics were added to reduce duplication without giving up clarity.

## Generics Reduce Repetitive Code and Increase Type Safety

Without generics, you often wrote near-duplicate functions:

```go
func SumInts(values []int) int
func SumFloat64s(values []float64) float64
```

With generics:

```go
func SumNumbers[T Number](values []T) T
```

This reduces repetitive code because:

- one algorithm can serve multiple types
- you do not need copy-pasted variants

It increases type safety because:

- the compiler still knows the concrete type
- results stay strongly typed
- there is no need to fall back to `any` and type assertions for common cases

That is the key value of generics in Go: less duplication without sacrificing compile-time checks.

## Generic Functions Abstract Algorithms

Generics are best when the algorithm is the same across types.

Examples:

- searching
- mapping
- filtering
- min/max logic
- stack/queue operations

This is where generics shine:

- the control flow stays identical
- only the element type changes

When the behavior itself changes significantly by type, interfaces or separate functions are often a better fit.

## Generics and Interfaces

Constraints are often interfaces.

For example:

```go
func JoinStrings[T Stringer](values []T) []string
```

That means generics and interfaces are not rivals. They solve different parts of the problem:

- interfaces express required behavior
- generics abstract over concrete type parameters

Together, they let you write code that is both reusable and strongly typed.

## Use Type Terms to Specify Operators

If a generic function uses operators like `+`, `<`, or `==`, the constraint must permit them.

For example:

```go
type Number interface {
    ~int | ~int64 | ~float64
}
```

This uses type terms and unions to say:

- `T` can be types whose underlying type is `int`, `int64`, or `float64`

The `~` means "any type with this underlying type."

This is important because operators in generics are only available when the constraint guarantees they are valid.

## Type Inference and Generics

Go often infers type arguments automatically:

```go
first, ok := First([]string{"a", "b"})
```

You rarely need to write:

```go
First[string](...)
```

This keeps generic call sites readable.

Use explicit type arguments only when inference is unclear or when you want to be specific.

## Type Elements Limit Constants

Untyped constants interact with generic code through the constraint's permitted types.

Example:

```go
func AddTen[T ~int | ~int64](value T) T {
    return value + 10
}
```

The constant `10` works because it can be represented in the allowed types.

Constraints effectively limit which constants and operations are valid in generic code.

This is another way the type system keeps behavior explicit.

## Combining Generic Functions with Generic Data Structures

Generics are especially useful when both algorithms and containers need to stay type-safe.

A generic stack:

```go
type Stack[T any] struct {
    items []T
}
```

This lets you write:

- one stack implementation
- no `any`
- no type assertions

Then generic helper functions can work naturally with that type.

This combination is one of the clearest practical wins of generics in Go.

## More on `comparable`

`comparable` is a predeclared constraint for types that support `==` and `!=`.

It is useful for:

- equality checks
- map keys
- membership tests like `Contains`

But not every type is comparable.

Examples of non-comparable types:

- slices
- maps
- functions

If your generic code relies on equality, `comparable` is often the right constraint.

## Things That Are Left Out

Go's generics are intentionally limited.

Things deliberately left out include:

- template metaprogramming complexity
- arbitrary specialization
- operator overloading
- highly advanced type-level programming
- complex generic features that make code harder to read

This is intentional. Go wants generics to solve common reuse problems, not become a second language inside the language.

## Idiomatic Go and Generics

Idiomatic Go still prefers simplicity first.

Use generics when:

- they remove obvious duplication
- they preserve readability
- the algorithm is genuinely type-agnostic

Do not use generics when:

- a concrete type is simpler
- the abstraction is used only once
- an interface expresses the design more clearly

The best Go generic code is usually boring:

- small constraints
- clear names
- obvious behavior

## Adding Generics to the Standard Library

Generics influenced the standard library by enabling more reusable typed helpers without forcing everything through `interface{}`.

A clear example is the newer `slices` and `maps` packages in the standard library ecosystem, which provide generic utilities for common operations.

That matters because it shows the intended direction:

- use generics for practical collection helpers
- avoid overcomplicating core APIs

Go did not retrofit every package with generics. The standard library has adopted them where they provide clear value.

## Run It

```bash
go run ./intermediate/05_generics
```

## Deep Note

Generics in Go are best viewed as a precision tool, not a default style. They are excellent for reusable algorithms and containers, especially where `any` would otherwise weaken type safety. But they work best when the abstraction stays obvious. If a generic abstraction makes the code harder to understand than the duplication it removed, it is usually not idiomatic Go.
