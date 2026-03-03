# Basics: Predeclared Types, Zero Values, and Literals

## Purpose

This lesson covers Go's foundational built-in data model:

- predeclared types
- zero values
- literals
- booleans
- numeric types
- strings
- runes

## Predeclared Types

Predeclared types are built into the language. You do not import anything to use them.

Common examples:

- `bool`
- `string`
- `int`, `int8`, `int16`, `int32`, `int64`
- `uint`, `uint8`, `uint16`, `uint32`, `uint64`
- `float32`, `float64`
- `complex64`, `complex128`
- `byte` (alias for `uint8`)
- `rune` (alias for `int32`)

## The Zero Value

Every declared variable in Go has a value, even if you do not assign one yourself.

That default is called the zero value.

Examples:

- `bool` -> `false`
- numeric types -> `0`
- `string` -> `""`
- pointers, maps, slices, functions, interfaces -> `nil`

This is one of Go's most useful design choices because it reduces uninitialized-memory mistakes and makes declarations safe by default.

## Literals

Literals are values written directly into your code.

Examples:

- boolean literal: `true`
- integer literal: `42`
- float literal: `19.95`
- string literal: `"hello"`
- rune literal: `'G'`

## Booleans

Booleans only have two values:

- `true`
- `false`

They are commonly used in:

- conditions
- flags
- state tracking

## Numeric Types

Go separates numeric types clearly.

- `int` is the default signed integer type for most general work
- `float64` is the default floating-point type for decimal calculations
- fixed-width types like `int32` or `uint64` are useful when size matters

Go does not silently mix numeric types. Conversions are explicit.

## Strings

A Go string is an immutable sequence of bytes, usually containing UTF-8 text.

Important implications:

- you cannot modify a string in place
- indexing a string gives you a byte, not necessarily a full character

## Runes

A rune represents a Unicode code point. It is an alias for `int32`.

Use runes when you care about characters rather than raw bytes.

This matters for:

- non-ASCII text
- iteration over Unicode content
- character-aware transformations

## Run It

```bash
go run ./basics/02_types_zero_values
```

## Deep Note

If you understand zero values and the difference between bytes, strings, and runes, you avoid a large class of beginner mistakes early. In Go, the language is deliberately small, but the exact meaning of each type matters.
