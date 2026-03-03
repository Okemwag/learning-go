# Basics: Functions

## Purpose

This lesson covers:

- declaring and calling functions
- simulating named and optional parameters
- variadic input parameters and slices
- multiple return values
- why multiple return values are truly multiple values
- ignoring return values
- blank returns and why to avoid them
- functions as values
- function type declarations
- anonymous functions
- closures
- passing functions as parameters
- returning functions from functions
- `defer`
- Go as call by value

## Declaring and Calling Functions

A basic function looks like this:

```go
func add(a int, b int) int {
    return a + b
}
```

And you call it like this:

```go
result := add(2, 3)
```

Go requires explicit parameter and return types. That keeps function contracts clear.

## Simulating Named Parameters

Go does not support named parameters like some other languages.

A common workaround is a struct:

```go
type CreateUserParams struct {
    Name string
    Age  int
}
```

Then pass one value:

```go
createUser(CreateUserParams{Name: "Ada", Age: 30})
```

This gives you:

- clearer call sites
- easier extension later
- fewer mistakes from argument order

## Simulating Optional Parameters

Go also does not support optional parameters directly.

Common patterns:

- config structs
- helper wrapper functions
- zero-value defaults

A config struct is usually the cleanest choice for anything beyond a trivial case.

## Variadic Input Parameters and Slices

Variadic parameters let a function accept zero or more values:

```go
func sum(values ...int) int
```

Inside the function, `values` is a slice.

You can also pass an existing slice using `...`:

```go
nums := []int{1, 2, 3}
sum(nums...)
```

This is one of the simplest and most useful examples of the connection between function parameters and slices.

## Multiple Return Values

Go functions can return more than one value:

```go
func divide(a, b int) (int, int) {
    return a / b, a % b
}
```

This is heavily used in Go, especially for:

- `(value, error)`
- `(value, ok)`

## Multiple Return Values Are Multiple Values

Go does not wrap multiple returns into a hidden tuple object in the way people often imagine.

Instead, the language treats them as multiple distinct returned values that can be assigned directly:

```go
q, r := divide(10, 3)
```

That is why they integrate so naturally with assignment syntax.

## Ignoring Return Values

Use `_` when a returned value is intentionally ignored:

```go
q, _ := divide(10, 3)
```

This is common and idiomatic when one of the results is not needed.

## Blank Returns

Blank returns use named return values:

```go
func f() (result int) {
    result = 10
    return
}
```

This works, but you should almost never use it in normal code.

Why to avoid it:

- it hides what is being returned
- it makes functions harder to scan
- it becomes confusing in longer functions

It is only tolerable in very small, very obvious functions. Explicit return values are usually better.

## Functions as Values

Functions are first-class values in Go.

That means you can:

- assign them to variables
- pass them as arguments
- return them from functions

Example:

```go
op := add
```

## Function Type Declarations

You can name a function signature:

```go
type IntOperation func(int, int) int
```

This improves readability when:

- the signature appears repeatedly
- the function is passed around often
- you want to express intent clearly

## Anonymous Functions

An anonymous function has no name:

```go
func(a int) int {
    return a * 2
}
```

Use them for:

- short inline behavior
- callbacks
- closures

## Closures

A closure is a function value that captures variables from its surrounding scope.

Example:

```go
func makeCounter() func() int {
    count := 0
    return func() int {
        count++
        return count
    }
}
```

The returned function keeps access to `count` even after the outer function returns.

Closures are powerful, but keep them small and obvious.

## Passing Functions as Parameters

This is useful when behavior should be injected:

```go
func apply(a, b int, op IntOperation) int
```

That pattern is common in:

- callbacks
- reusable helpers
- middleware-like logic

## Returning Functions from Functions

This lets you generate specialized behavior:

```go
func makeMultiplier(factor int) func(int) int
```

This is useful for:

- factories
- adapters
- partial application patterns

## `defer`

`defer` schedules a function call to run when the surrounding function returns.

Example:

```go
defer file.Close()
```

Common uses:

- closing files
- unlocking mutexes
- cleanup logic

Important rule:

- deferred calls run in last-in, first-out order

## Go Is Call by Value

Go passes arguments by value. That means parameter variables receive copies.

For simple values like `int`, this is straightforward:

- changing the parameter does not change the caller's variable

For slices, maps, and some other values, the value being copied contains references to underlying runtime data. So the outer variable is still passed by value, but both copies may refer to shared underlying state.

That is why this is true at the same time:

- Go is call by value
- changing a slice element in a function can affect the caller

There is no contradiction once you distinguish:

- the copied slice header
- the shared backing array

## Run It

```bash
go run ./basics/12_functions
```

## Deep Note

Functions are where Go's design becomes especially visible: explicit inputs, explicit outputs, and very little hidden magic. If you keep function signatures clear, prefer explicit returns, and use higher-order functions sparingly and purposefully, your code stays easy to understand.
