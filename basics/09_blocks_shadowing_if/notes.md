# Basics: Blocks, Shadowing, and `if`

## Purpose

This lesson covers:

- blocks
- shadowing variables
- `if`

## Blocks

A block is any section of code enclosed in braces:

```go
{
    // block
}
```

Blocks define scope boundaries. Variables declared inside a block are only visible inside that block and nested inner blocks.

Common blocks in Go:

- function bodies
- `if` branches
- `for` bodies
- `switch` cases
- explicit inner `{ ... }` blocks

## Shadowing Variables

Shadowing happens when you declare a new variable with the same name in an inner scope:

```go
name := "outer"
{
    name := "inner"
}
```

The inner `name` does not replace the outer one. It temporarily hides it within the inner scope.

Why this matters:

- shadowing can be useful for short local values
- shadowing can also cause bugs when you think you are updating an outer variable but actually created a new one

Be especially careful with:

- `:=` inside `if`
- `:=` inside loops
- error variables like `err`

## `if`

Go's `if` statement does not require parentheses around the condition:

```go
if score > 50 {
    fmt.Println("pass")
}
```

You can also declare a short-lived variable before the condition:

```go
if grade := determineGrade(score); grade == "A" {
    // use grade
}
```

That variable exists only inside the `if`, `else if`, and `else` chain.

## Why This Matters

Understanding scope is foundational. Many beginner mistakes are not syntax problems but scope problems:

- using a value outside its block
- accidentally shadowing a variable
- assuming a variable survives longer than it does

## Run It

```bash
go run ./basics/09_blocks_shadowing_if
```

## Deep Note

Go keeps scope rules intentionally simple, but that simplicity is only useful if you stay disciplined. Prefer short scopes, but avoid reusing names so aggressively that shadowing becomes confusing.
