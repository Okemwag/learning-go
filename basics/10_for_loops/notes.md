# Basics: `for`

## Purpose

This lesson covers:

- the four main ways to use `for`
- the complete `for` statement
- the condition-only `for` statement
- the infinite `for` statement
- `break` and `continue`
- the `for-range` statement
- labelling a `for` statement
- choosing the right `for` form

## Four Ways to Use `for`

Go has only one looping keyword, but it covers multiple patterns:

1. complete `for`
2. condition-only `for`
3. infinite `for`
4. `for-range`

That is one reason Go stays compact without losing expressiveness.

## The Complete `for` Statement

```go
for i := 0; i < 10; i++ {
    fmt.Println(i)
}
```

This form has:

- initialization
- condition
- post statement

Use it when you need a clear counter-based loop.

## The Condition-Only `for` Statement

```go
for count < 10 {
    count++
}
```

This is Go's equivalent of a traditional `while` loop.

Use it when:

- the loop is driven by a condition
- a counter exists but the `init; condition; post` form is less readable

## The Infinite `for` Statement

```go
for {
    // repeat forever until break/return/panic
}
```

Use it for:

- servers
- worker loops
- retry loops
- event processing

Always ensure there is a clear exit path when appropriate.

## `break` and `continue`

- `break` exits the current loop
- `continue` skips the current iteration and moves to the next one

These are useful, but overuse can make loops harder to read. Prefer clear conditions when possible.

## The `for-range` Statement

```go
for index, value := range items {
    fmt.Println(index, value)
}
```

This is the idiomatic way to iterate over:

- slices
- arrays
- strings
- maps
- channels

If you do not need one of the values, use `_`.

## Labelling Your `for` Statement

Nested loops sometimes need to break out of an outer loop, not just the inner one.

Use a label:

```go
Outer:
for ... {
    for ... {
        break Outer
    }
}
```

Labels are useful in specific cases, but should stay rare.

## Choosing the Right `for` Statement

Use:

- complete `for` for index-driven loops
- condition-only `for` for state-driven loops
- infinite `for` for ongoing processing loops
- `for-range` for collection traversal

Most of the time, `for-range` is the best default when iterating over existing data.

## Run It

```bash
go run ./basics/10_for_loops
```

## Deep Note

Good Go loops are boring in the best way: obvious start, obvious stop, obvious iteration behavior. Choose the simplest loop form that matches the job, and avoid writing loops that require the reader to simulate too much state mentally.
