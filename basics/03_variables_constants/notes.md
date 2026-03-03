# Basics: `var`, `:=`, and Constants

## Purpose

This lesson focuses on:

- `var` vs `:=`
- using `const`
- typed constants
- untyped constants
- naming variables well

## `var` vs `:=`

Use `var` when:

- you want an explicit type
- you want the zero value first
- you are declaring package-level variables
- you are using grouped declarations

Use `:=` when:

- you are inside a function
- the value is obvious from the right-hand side
- you want concise local code

`:=` is not available at package scope.

## Using `const`

Constants are values that cannot change after declaration.

They are ideal for:

- fixed configuration values
- limits
- repeated string keys
- mathematically stable values

## Typed and Untyped Constants

Typed constant:

```go
const rate float64 = 0.15
```

Untyped constant:

```go
const maxUsers = 100
```

Untyped constants are more flexible because they can be used in different contexts until Go needs to choose a specific type.

## Naming Variables

Good naming in Go is:

- short
- precise
- context-aware

Prefer:

- `requestCount`
- `userID`
- `isReady`

Avoid names that hide meaning:

- `x`
- `tmpValue123`
- `dataThing`

## Run It

```bash
go run ./basics/03_variables_constants
```

## Deep Note

Go code is often judged by readability. The language removes many stylistic choices, which means naming becomes one of the main tools you have left to communicate intent.
