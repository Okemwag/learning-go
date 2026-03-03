# Basics: `switch` and `goto`

## Purpose

This lesson covers:

- `switch`
- blank switches
- choosing between `if` and `switch`
- `goto`

## `switch`

A `switch` is useful when one value or one decision point has multiple possible branches.

Example:

```go
switch day {
case "Monday":
    fmt.Println("start")
case "Friday":
    fmt.Println("end")
default:
    fmt.Println("middle")
}
```

Why `switch` is often clearer than long `if/else if` chains:

- the decision branches line up vertically
- the shared decision point is obvious
- adding cases stays readable

Go's `switch` also does not fall through by default. That prevents a common class of bugs from C-like languages.

## Blank Switches

A blank switch omits the expression:

```go
switch {
case score >= 90:
    fmt.Println("A")
case score >= 75:
    fmt.Println("B")
}
```

Each `case` is treated as a boolean expression.

This is a clean alternative to long ordered `if/else if` chains when:

- multiple conditions are checked in sequence
- a value range is being classified

## Choosing Between `if` and `switch`

Prefer `if` when:

- there are only one or two simple conditions
- the logic is highly specific
- you need a short, direct branch

Prefer `switch` when:

- many branches depend on one decision
- the conditions form a category list
- readability improves from vertically aligned cases

A blank `switch` is especially good when several related boolean tests belong to one conceptual decision.

## `goto`

`goto` jumps to a label within the same function.

Example:

```go
goto Retry
```

`goto` is legal Go, but it should be rare.

It can be reasonable when:

- escaping deeply nested logic
- handling a tightly scoped low-level control-flow pattern

It is usually a bad choice when:

- it makes flow harder to follow
- it replaces clearer loop or function structure

Most code should prefer:

- `for`
- `break`
- `continue`
- helper functions

## Run It

```bash
go run ./basics/11_switch_goto
```

## Deep Note

`switch` is a readability tool. `goto` is a last-resort control-flow tool. If both are available, `switch` usually makes code more maintainable, while `goto` often does the opposite unless used with discipline.
